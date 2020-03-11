package controllers

import (
    "net/http"
    "fmt"
	"ieeblog/services"
	//"html/template"
)

//LoginController not inherit base controller
type LoginController struct {
	Controller
}

//Login session
func(loginCtrl *LoginController) Login(w http.ResponseWriter, r *http.Request){
	fmt.Println("metgod:", r.Method) //获取请求的方法
	//sess:=loginCtrl.CurrentSession.SessionStart(w,r) //session test add
	r.ParseForm()  //session test add
	loginCtrl.Init(w,r,"loginController")
	/*
	if r.Method=="GET"{
		t,_:=template.ParseFiles("./views/login/login.gtpl")
		//log.Println(t.Execute(w,nil)) //session test comment
		w.Header().Set("Content-Type","text/html") //session test add
		t.Execute(w,nil) //session test add sess.Get("username")
	}else{
		//r.ParseForm() //才会解析 否则下面的输出为空白
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:",r.Form["password"])
		//validation
		ok:=services.UserService.LoginValidate(r.FormValue("username"),r.FormValue("password"))
		if ok{
			//sess.Set("username",r.Form["username"])
			(&ArticleController{}).GetArticleList(w,r)
		}
	}*/
	if r.Method=="GET"{
		loginCtrl.Layout=append(loginCtrl.Layout,"login/login.tpl") 
		loginCtrl.TplNames="login"
		err:=loginCtrl.Render()
		if err!=nil{
			fmt.Println("render user following list error",err)
		}
	}else{
		//validation
		fmt.Println("login username:", r.FormValue("username"))
		fmt.Println("login password:",r.FormValue("password"))
		ok:=services.UserService.LoginValidate(r.FormValue("username"),r.FormValue("password"))
		fmt.Println("ok:",ok)
		if ok{
			//sess.Set("username",r.Form["username"])
			loginCtrl.Redirect("/article/list",304)
			(&ArticleController{}).GetArticleList(w,r)
		}
	}
}


//PostLogin sign in
func(loginCtrl *LoginController) PostLogin(username string,password string)bool{

	ok:=services.UserService.LoginValidate(username,password)
	if ok{
		return true
	}
	return false
	
}

//SignUp username pwd
func(loginCtrl *LoginController) SignUp(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	loginCtrl.Init(w,r,"loginController")
	if r.Method=="GET"{
		loginCtrl.Layout=append(loginCtrl.Layout,"login/signup.tpl") 
		loginCtrl.TplNames="signup"
		err:=loginCtrl.Render()
		if err!=nil{
			fmt.Println("render sign up get error",err)
		}
	}else{
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:",r.Form["password"])
		row,_:= services.UserService.AddUser(r.Form["username"][0],r.Form["email"][0],r.Form["password"][0])
		if row>0{ //sign up successfully
			
			loginCtrl.Redirect("/login",304)
			/*
			loginCtrl.Layout=append(loginCtrl.Layout,"login/login.tpl") 
			loginCtrl.TplNames="login"
			err:=loginCtrl.Render()
			if err!=nil{
				fmt.Println("render login in get error",err)
			}
			*/
		}
	}
}

//LoginOut 退出时销毁session
func(loginCtrl *LoginController) LoginOut(w http.ResponseWriter, r *http.Request){
	loginCtrl.CurrentSession.SessionDestroy(w,r)
	fmt.Println("session destroy")
}

