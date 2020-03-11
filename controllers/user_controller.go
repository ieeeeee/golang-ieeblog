package controllers

import(
	"strconv"
	"ieeblog/services"
	"fmt"
	"html/template"
	"os"
	"io"
	"time"
	"crypto/md5"
	"path"
	"net/http"
	//"strings"
	//"ieeblog/ieecom/loghelper"
)
//UserController struct
type UserController struct{
	Controller
}

var viewsPath string="./views"

//GetUser a user
func(userCtrl *UserController) GetUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("user Get")
	/*
	params:=make(map[string]string)
	for k,v:=range r.Form{
		params[k]=strings.Join(v,"")
		fmt.Println("key",k)
		fmt.Println("val",strings.Join(v,""))
	}
	*/
	//userCtrl.Ct=&userCtrl.Ct{Request:r,ResponseWriter:w}
	
	userCtrl.Init(w,r,"UserController")
	//userCtrl.Ct.Request.ParseForm() 	//解析参数
	//id,_:= strconv.ParseInt(params["id"],0,64)
	id,_:= strconv.ParseInt(userCtrl.Request.Form["id"][0],0,64)
	userCtrl.Data["User"] = services.UserService.GetUserByID(id)
	fmt.Println("user Get1")
	userCtrl.Layout=append(userCtrl.Layout,"master/header.tpl") 
	userCtrl.Layout=append(userCtrl.Layout,"master/footer.tpl") 
	userCtrl.Layout=append(userCtrl.Layout,"user/detail.tpl") 
	userCtrl.TplNames="content"
	//userCtrl.TplNames = "user/detail.tpl"
	err:=userCtrl.Render()
	if err!=nil{
		fmt.Println("render error",err)
	}
	/*
	params:=make(map[string]string)
	r.ParseForm() 	//解析参数
	fmt.Println(r.Form) //输出到服务器端打印信息
	for k,v:=range r.Form{
		params[k]=strings.Join(v,"")
		fmt.Println("key",k)
		fmt.Println("val",strings.Join(v,""))
	}
	id,_:= strconv.ParseInt(params["id"],0,64)
	fmt.Println(id)
	data:=make(map[interface{}]interface{})
	data["User"] = services.UserService.GetUserByID(id)
	fmt.Println("data user",data["User"])
	var layout []string
	layout=append(layout,"layout.gtpl") 
	layout=append(layout,"user/detail.gtpl") 
	tplNames:="layoutContent"
	//tplNames:="user/detail.gtpl"
	fmt.Println("User Render")
	if len(layout)>0{
		var filenames []string
		for _,file:=range layout{
			filenames=append(filenames,path.Join(viewsPath,file)) //ViewsPath 配置文件配置
		}
		fmt.Println("filenames",filenames)
		t,err:=template.ParseFiles(filenames...)
		if err!=nil{
			fmt.Println("template ParseFiles err:", err)
			//loghelper.Trace("template ParseFiles err:", err)
		}
		err=t.ExecuteTemplate(w,tplNames,data)//path.Join(viewsPath, tplNames)
		if err!=nil{
			fmt.Println("template ExecuteTemplate err:", err)
			//loghelper.Trace("template Execute err:", err)
		}
	}else{
		
		if c.TplNames==""{
			c.TplNames=c.ChildName+"/"+c.Ct.Request.Method+"."+c.TplExt
		}
		
		t, err := template.ParseFiles(path.Join(viewsPath, tplNames))
        if err != nil {
			fmt.Println("template ParseFiles err:", err)
           // loghelper.Trace("template ParseFiles err:", err)
        }
        err = t.Execute(w, data)
        if err != nil {
			fmt.Println("template Execute err:", err)
            //loghelper.Trace("template Execute err:", err)
        }
	}
	*/
	
}

//GetFollowerListByUserID article list Get()
func(userCtrl *UserController) GetFollowerListByUserID(w http.ResponseWriter, r *http.Request){
    userCtrl.Init(w,r,"userController")
	userCtrl.Data["Followers"] = services.UserService.GetFollowerListByUserID()
	userCtrl.Layout=append(userCtrl.Layout,"master/header.tpl") 
	userCtrl.Layout=append(userCtrl.Layout,"master/footer.tpl") 
	userCtrl.Layout=append(userCtrl.Layout,"user/follower.tpl") 
	userCtrl.TplNames="content"
	err:=userCtrl.Render()
	if err!=nil{
		fmt.Println("render user follower list error",err)
	}
}

//GetFollowingListByUserID article list Get()
func(userCtrl *UserController) GetFollowingListByUserID(w http.ResponseWriter, r *http.Request){
    userCtrl.Init(w,r,"userController")
	userCtrl.Data["Followings"] = services.UserService.GetFollowingListByUserID()
	userCtrl.Layout=append(userCtrl.Layout,"master/header.tpl") 
	userCtrl.Layout=append(userCtrl.Layout,"master/footer.tpl") 
	userCtrl.Layout=append(userCtrl.Layout,"user/following.tpl") 
	userCtrl.TplNames="content"
	err:=userCtrl.Render()
	if err!=nil{
		fmt.Println("render user following list error",err)
	}
}
//var viewsPath string="./views/user"

//Upload test
func(userCtrl *UserController) Upload(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method=="GET"{
		crutime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))
		fmt.Println("upload parsefilename:",path.Join(viewsPath, "user/upload.gtpl"))
		t,err:=template.ParseFiles(path.Join(viewsPath, "user/upload.gtpl")) //template.ParseFiles(path.Join(ViewsPath, c.TplNames)
		if err!=nil{
			fmt.Println("Parse file error",err)
			return
		}
		t.Execute(w,token)
	}else{
		r.ParseMultipartForm(32<<20) //把上传的文件存储在内存和临时文件中
		file,handler,err:=r.FormFile("uploadfile") //获取文件句柄
		if err!=nil{
			fmt.Println("FormFile errot",err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header) //map[Content-Disposition:[form-data; name="uploadfile"; filename="11.txt"] Content-Type:[text/plain]]
		f,err:=os.OpenFile("./file/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err!=nil{
			fmt.Println("Open file error",err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}
