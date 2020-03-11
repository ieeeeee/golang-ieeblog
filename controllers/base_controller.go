/*Controller 基类*/

package controllers
import(
	"net/http"
	"html/template"
	//"ieeblog/ieecom/loghelper"
	"path"
	//"ieeblog/ieecom/router"
	"fmt"
	"strings"
	"ieeblog/ieecom/session"
)

//Controller define a base controller struct
type Controller struct{
	//Ct *router.Context
	Tpl *template.Template
	Data map[interface{}]interface{}
	ChildName string
	TplNames string
	Layout []string
	TplExt string
	ResponseWriter http.ResponseWriter
	Request *http.Request
	Params map[string]string
	// session
	CurrentSession *session.Manager
}


//implement this interface

//Init a baseController
func (c *Controller) Init(w http.ResponseWriter, r *http.Request,cn string){ //ct *router.Context,
	params:=make(map[string]string)
	r.ParseForm()
	fmt.Println(r.Form) 
	for k,v:=range r.Form{
		params[k]=strings.Join(v,"")
		fmt.Println("key",k)
		fmt.Println("val",strings.Join(v,""))
	}
	c.Data=make(map[interface{}]interface{})
	c.Layout = make([]string, 0)
    c.TplNames = ""
	c.ChildName = cn
	c.ResponseWriter=w
	c.Request=r
	c.Params=params
   // c.Ct = &router.Context{Request:r,ResponseWriter:w,Params:params}
	c.TplExt = "tpl"
	fmt.Println("base controller init")
}

//Prepare implement
func(c *Controller) Prepare(){

}
//Finish implement
func (c *Controller) Finish() {

}
//Get implement
func (c *Controller) Get() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}
//Post implement
func (c *Controller) Post() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}
//Delete implement
func (c *Controller) Delete() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}
//Put implement
func (c *Controller) Put() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}
//Head implement
func (c *Controller) Head() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}
//Patch implement
func (c *Controller) Patch() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}
//Options implement
func (c *Controller) Options() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}
//ViewsPath config
var ViewsPath string="./views" //ViewsPath 应从配置文件配置

//Render the page after executing the method corresponding to the method
func (c *Controller) Render() error{
	fmt.Println("Base Render1")
	if len(c.Layout)>0{
		var filenames []string
		for _,file:=range c.Layout{
			filenames=append(filenames,path.Join(ViewsPath,file)) 
		}
		fmt.Println("base filenames",filenames)
		t,err:=template.ParseFiles(filenames...)
		if err!=nil{
			fmt.Println("template ParseFiles err:", err)
			//loghelper.Trace("template ParseFiles err:", err)
		}
		err=t.ExecuteTemplate(c.ResponseWriter,c.TplNames,c.Data)
		if err!=nil{
			fmt.Println("template ParseFiles err:", err)
			//loghelper.Trace("template Execute err:", err)
		}
	}else{
		if c.TplNames==""{
			c.TplNames=c.ChildName+"/"+c.Request.Method+"."+c.TplExt
		}
		t, err := template.ParseFiles(path.Join(ViewsPath, c.TplNames))
        if err != nil {
			fmt.Println("template ParseFiles err:", err)
            //loghelper.Trace("template ParseFiles err:", err)
        }
        err = t.Execute(c.ResponseWriter, c.Data)
        if err != nil {
			fmt.Println("template ParseFiles err:", err)
            //loghelper.Trace("template Execute err:", err)
        }
	}
	return nil
}

//Redirect url
func (c *Controller) Redirect(url string, code int) {
    http.Redirect(c.ResponseWriter,c.Request,url,code)
}

/*
// SetSession puts value into session.
func (c *Controller) SetSession(name interface{}, value interface{}) {
	c.CurrentSession.Set(name, value)
}

// GetSession gets value from session.
func (c *Controller) GetSession(name interface{}) interface{} {
	return c.CurrentSession.Get(name)
}

// DelSession removes value from session.
func (c *Controller) DelSession(name interface{}) {
	c.CurrentSession.Delete(name)
}



通过路由根据url执行相应的controller的原则，会依次执行如下：
Init()      初始化
Prepare()   执行之前的初始化，每个继承的子类可以来实现该函数
method()    根据不同的method执行不同的函数：GET、POST、PUT、HEAD等，
			子类来实现这些函数，如果没实现，那么默认都是403
Render()    可选，根据全局变量AutoRender来判断是否执行
Finish()    执行完之后执行的操作，每个继承的子类可以来实现该函数

*/


/*
//MainController struct
type MainController struct {
    ieeblog.Controller
}

//Get implement
func (mc *MainController) Get() {
    mc.Data["Username"] = "astaxie"
    mc.Data["Email"] = "astaxie@gmail.com"
    mc.TplNames = "index.tpl"
}



//GetJson get json data of remote URL address
func GetJson() {
    resp, err := http.Get(beego.AppConfig.String("url"))
    if err != nil {
        logger.Critical("http get info error")
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body, &AllInfo)
    if err != nil {
        logger.Critical("error:", err)
    }
}
*/