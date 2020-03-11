package main

import(
	"flag"
	"fmt"
	"strconv"
	"html/template"
	"os"
	"io"
	"time"
	"crypto/md5"
	"path"
	"net/http"
	"ieeblog/ieecom/config"
	"ieeblog/ieecom/loghelper"
	"ieeblog/ieecom/session"
	//"ieeblog/ieecom/router"
	//"ieeblog/services"
	"ieeblog/controllers"
)
/*
session
1, var globalSession 
2, init()-> newSessionManager(),go globalSession.GC()
router
3, main()-> http.HandleFunc(),http.ListenAndServe

*/



//配置文件
var configFile=flag.String("config", "./web.config", "配置文件路径")
//GlobalSessions 全局session
var globalSessions *session.Manager
func init(){
	//初始化配置文件
	config.InitConfig(*configFile)

	//设置日志
	loghelper.InitFileLogger() 

	//设置session
	var err error
	globalSessions,err=session.NewManager("memory","gosessionid",3600)
	(&controllers.Controller{}).CurrentSession=globalSessions
	if err!=nil{
		fmt.Println("new session manager error: ",err)
		return
	}

	go globalSessions.GC()
}


func main(){

	/*添加路由*/
	/*
	blogRouter:=router.NewControllerRegister()
	blogRouter.Add("/views/user", &controllers.UserController{})
	fmt.Println("main blogRouter",blogRouter)
	blogRouter.ServeHTTP(Response{},&http.Request{})
	*/
	//启用静态服务
	staticServe:=http.FileServer(http.Dir("static"))
	http.Handle("/static/",http.StripPrefix("/static/",staticServe))//启动静态服务
	http.HandleFunc("/user/detail", (&controllers.UserController{}).GetUser)
	http.HandleFunc("/user/upload", (&controllers.UserController{}).Upload)
	http.HandleFunc("/user/follower", (&controllers.UserController{}).GetFollowerListByUserID)
	http.HandleFunc("/user/following", (&controllers.UserController{}).GetFollowingListByUserID)

	http.HandleFunc("/article/list", (&controllers.ArticleController{}).GetArticleList)
	http.HandleFunc("/article/detail", (&controllers.ArticleController{}).GetArticleByID)
	http.HandleFunc("/article/edit", (&controllers.ArticleController{}).GetArticleNew)
	http.HandleFunc("/article/post", (&controllers.ArticleController{}).PostArticleNew)
	http.HandleFunc("/article/user", (&controllers.ArticleController{}).GetArticleListByUserID)

	http.HandleFunc("/login", (&controllers.LoginController{}).Login)
	http.HandleFunc("/login/signup",  (&controllers.LoginController{}).SignUp)
	err:=http.ListenAndServe(":9090",nil) //设置监听的端口 nil 默认获取handler=DefaultServerMux
	if err!=nil{
		loghelper.Critical("ListenAndServe:",err)
	}
	
	

/* Test
	loghelper.Info("log test")
	fmt.Println("config test:%v",config.Conf.Mssql.DataSource)
	fmt.Println("config test:%v",config.Conf.Mssql.Database)
	fmt.Println("config test:%v",config.Conf.Mssql.Windows)
	fmt.Println("config test:%v",config.Conf.Mssql.User)
	fmt.Println("config test:%v",config.Conf.Mssql.Password)
	
	
	//插入一个user
	userID:=services.UserService.AddUser()
	fmt.Println("sql test:%v",userID)
	user:=services.UserService.GetUserByID(1)
	fmt.Println("sql test:%v",user.UserNo)
*/
/*
	//显示博客首页
	beego.Router("/", &controllers.IndexController{})
	//查看博客详细信息
	beego.Router("/view/:id([0-9]+)", &controllers.ViewController{})
	//新建博客博文
	beego.Router("/new", &controllers.NewController{})
	//删除博文
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})
	//编辑博文
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})
*/
}


/*
func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("metgod:", r.Method) //获取请求的方法
	sess:=globalSession.SessionStart(w,r) //session test add
	r.ParseForm()  //session test add
	if r.Method=="GET"{
		t,_:=template.ParseFiles("./views/login/login.gtpl")
		//log.Println(t.Execute(w,nil)) //session test comment
		w.Header().Set("Content-Type","text/html") //session test add
		t.Execute(w,sess.Get("username")) //session test add
	}else{
		//r.ParseForm() //才会解析 否则下面的输出为空白
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:",r.Form["password"])
		//validation
		ok:=(&controllers.LoginController{}).PostLogin(r.FormValue("username"),r.FormValue("password"))
		if ok{

			sess.Set("username",r.Form["username"])
		}
		
	}
}
*/

/*
func test_session_valid(w http.ResponseWriter, r *http.Request) {  
    var sessionID = sessionMgr.CheckCookieValid(w, r)  
  
    if sessionID == "" {  
        http.Redirect(w, r, "/login", http.StatusFound)  
        return  
    }  
}
*/

var viewsPath string="./views/user"
func upload(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method=="GET"{
		crutime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))
		t,err:=template.ParseFiles(path.Join(viewsPath, "upload.gtpl")) //template.ParseFiles(path.Join(ViewsPath, c.TplNames)
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

