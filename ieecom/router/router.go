/*
而 router 可以理解为一个容器，或者说一种机制，它管理了一组 route。
简单来说，route 只是进行了URL和函数的映射
，而在当接收到一个URL之后，去路由映射表中查找相应的函数，
这个过程是由 router 来处理的。
The router routes you to a route
*/

//路由要自动匹配 {{controller}/{method}/{param}}


package router
import(
	"strings"
	"regexp"
	//"path"
	"reflect"
	"net/http"
	"net/url"
	"fmt"
)
/*
1, save router
2, forward router
*/
/*
//how to use
beego.BeeApp.RegisterController("/",&contrpllers,MainController{}
)

//param register
beego.BeeApp.RegisterController("/:param", &controllers.UserController{})

//regex
beego.BeeApp.RegisterController("/users/:uid([0-9]+)", &controllers.UserController{})
*/
//动态请求和静态资源分离
//save path
type controllerInfo struct{
	regex *regexp.Regexp   //带正则表达式
	params map[int]string  //带参数
	controllerType  reflect.Type //request.Method Get,Post,Put,Delete  用户访问一个handler，可以用GET、POST、DELETE、HEAD等方式访问，所以会在handler中写if method==get
}

//ControllerRegister struct
type ControllerRegister struct { 
	routers []*controllerInfo  
	//Application *App
}
//Context struct
type Context struct {
	Request        *http.Request
	ResponseWriter  http.ResponseWriter//*http.Response//http.ResponseWriter
	Params   map[string]string
}
//ControllerInterface interface define  Define the methods that the controller should have
type ControllerInterface interface{
	Init(ct *Context, cn string)    //初始化上下文和子类名称
    Prepare()                       //开始执行之前的一些处理
    Get()                           //method=GET的处理
    Post()                          //method=POST的处理
    Delete()                        //method=DELETE的处理
    Put()                           //method=PUT的处理
    Head()                          //method=HEAD的处理
    Patch()                         //method=PATCH的处理
    Options()                       //method=OPTIONS的处理
    Finish()                        //执行完成之后的处理        
    Render() error                  //执行完method对应的方法之后渲染页面
}

// NewControllerRegister returns a new ControllerRegister.
func NewControllerRegister() *ControllerRegister {
	fmt.Println("router NewControllerRegister")
	return &ControllerRegister{
		routers:  make([]*controllerInfo,0),	
	}
}
//Dynamic routing implementation

//Add a dynamic route
//对外的接口函数
func (p *ControllerRegister) Add(pattern string,c ControllerInterface){
	fmt.Println("router add",pattern)
	parts:=strings.Split(pattern,"/")

	j:=0
	params:=make(map[int]string)
	for i,part:=range parts{
		if strings.HasPrefix(part,":"){
			expr:="([^/])"
			//a user may choose to override the defult expression
			// similar to expressjs: ‘/user/:id([0-9]+)’
			
			if index:=strings.Index(part,"(");index!=-1{
				expr=part[index:]
				part=part[:index]
			}
			params[j]=part
			parts[i]=expr
			j++			
		}
	}
	//recreate the url pattern, with parameters replaced
	//by regular expressions. then compile the regex
	pattern=strings.Join(parts,"/")
	regex,regexErr:=regexp.Compile(pattern)
	if regexErr!=nil{
		//TODO add error handling here to avoid panic
        panic(regexErr)
        return
	}

	//now create the Route
	t:=reflect.Indirect(reflect.ValueOf(c)).Type()
	route:=&controllerInfo{}
	route.regex=regex
	route.params=params
	route.controllerType=t
	fmt.Println("router add,controllerType",t)
	fmt.Println("router add,controllerType",route)
	p.routers=append(p.routers,route)
}
/*
//SetStaticPath static routing implementation
func(app *App) SetStaticPath(url string,path string) *App{
	StaticDir[url]=path
	return app
}
//beego.SetStaticPath("/img","/static/img")
*/

//StaticDir map[string]string
var StaticDir map[string]string

//ServerHTTP Forward route AutoRoute
func (p *ControllerRegister) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Println("router ServeHTTP")
/*	defer func(){
		if err:=recover();err!=nil{
			if !RecoverPanic{
				//go back to panic
				panic(err)
			}else{
				Critical("Handler crashed with error",err)
				for i:=1; ;i++{
					_,file,line,ok:=runtime.Caller(i)
					if !ok{
						break
					}
					Critical(file,line)
				}
			}
		}
	}()*/
	var started bool
	//静态资源？？？
	StaticDir["/static"] = "static"
	for prefix,staticDir:=range StaticDir{ //StaticDir从配置取
		if strings.HasPrefix(r.URL.Path,prefix){  //判断r.URL.Path是否有前缀字符串prefix。
			file:=staticDir+r.URL.Path[len(prefix):]
			http.ServeFile(w,r,file)  //ServeFile回复请求filename指定的文件或者目录的内容。
			started=true
			return
		}
	}
	requestPath:=r.URL.Path

	//find a matching route
	for _,route:=range p.routers{
		//check if Route pattern matches url
		if !route.regex.MatchString(requestPath){
			continue
		}

		//get submatches(params)
		matches:=route.regex.FindStringSubmatch(requestPath) //Find返回一个保管正则表达式re在b中的最左侧的一个匹配结果以及（可能有的）分组匹配的结果的[]string切片。如果没有匹配到，会返回nil。

		//double check that the Route matches the URL pattern.
		if len(matches[0])!=len(requestPath){
			continue
		}

		params:=make(map[string]string)
		if len(route.params)>0{
			//add url parameters to the qyery param map
			values:=r.URL.Query()  //Query方法解析RawQuery字段并返回其表示的Values类型键值对。
			for i,match:=range matches[1:]{
				values.Add(route.params[i],match)
				params[route.params[i]]=match
			}

			//reassemble query params and add to RawQuery
			r.URL.RawQuery=url.Values(values).Encode()+"&"+r.URL.RawQuery //Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
		}

		//Invoke the request handler
		vc:=reflect.New(route.controllerType) //返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针，返回值的Type为PtrTo(typ)。
		//返回v的名为name的方法的已绑定（到v的持有值的）状态的函数形式的Value封装。
		//返回值调用Call方法时不应包含接收者；返回值持有的函数总是使用v的持有者作为接收者（即第一个参数）。
		//如果未找到该方法，会返回一个Value零值。
		init:=vc.MethodByName("Init")

		in:=make([]reflect.Value,2)
		ct:=&Context{ResponseWriter:w,Request:r,Params:params}
		in[0]=reflect.ValueOf(ct) //ValueOf返回一个初始化为i接口保管的具体值的Value，ValueOf(nil)返回Value零值。
		in[1]=reflect.ValueOf(route.controllerType.Name()) // Refleft.Type.Name()返回该类型在自身包内的类型名，如果是未命名类型会返回""
		/*Call方法使用输入的参数in调用v持有的函数。例如，如果len(in) == 3，v.Call(in)代表调用v(in[0], in[1], in[2])（其中Value值表示其持有值）。
		如果v的Kind不是Func会panic。它返回函数所有输出结果的Value封装的切片。
		和go代码一样，每一个输入实参的持有值都必须可以直接赋值给函数对应输入参数的类型。
		如果v持有值是可变参数函数，Call方法会自行创建一个代表可变参数的切片，将对应可变参数的值都拷贝到里面。*/
		method := vc.MethodByName("Prepare")
		init.Call(in)
		if r.Method=="GET"{
			method = vc.MethodByName("Get") //返回名字叫Get方法
            method.Call(in) //执行Controller中的方法名为Get()
        } else if r.Method == "POST" {
            method = vc.MethodByName("Post")
            method.Call(in)
        } else if r.Method == "HEAD" {
            method = vc.MethodByName("Head")
            method.Call(in)
        } else if r.Method == "DELETE" {
            method = vc.MethodByName("Delete")
            method.Call(in)
        } else if r.Method == "PUT" {
            method = vc.MethodByName("Put")
            method.Call(in)
        } else if r.Method == "PATCH" {
            method = vc.MethodByName("Patch")
            method.Call(in)
        } else if r.Method == "OPTIONS" {
            method = vc.MethodByName("Options")
            method.Call(in)
        }
		
		//if AutoRender{
			method=vc.MethodByName("Render")
			method.Call(in)
        //}
        method = vc.MethodByName("Finish")
        method.Call(in)
		started=true
		break
	}

	//if no matches to url,throw a not found exception
	if started==false{
		http.NotFound(w,r)

	}
}

/*
// AddAuto router to ControllerRegister.
// example beego.AddAuto(&MainContorlller{}),
// MainController has method List and Page.
// visit the url /main/list to execute List function
// /main/page to execute Page function.
func (p *ControllerRegister) AddAuto(c ControllerInterface) {
	p.AddAutoPrefix("/", c)
}

// AddAutoPrefix Add auto router to ControllerRegister with prefix.
// example beego.AddAutoPrefix("/admin",&MainContorlller{}),
// MainController has method List and Page.
// visit the url /admin/main/list to execute List function
// /admin/main/page to execute Page function.
func (p *ControllerRegister) AddAutoPrefix(prefix string, c ControllerInterface) {
	reflectVal := reflect.ValueOf(c)
	rt := reflectVal.Type()
	ct := reflect.Indirect(reflectVal).Type() //返回持有v持有的指针指向的值的Value。如果v持有nil指针，会返回Value零值；如果v不持有指针，会返回v。
	controllerName := strings.TrimSuffix(ct.Name(), "Controller") //返回去除s可能的后缀suffix的字符串。
	for i := 0; i < rt.NumMethod(); i++ {
		//if !utils.InSlice(rt.Method(i).Name, exceptMethod) { //方法名不在exceptMethod里
			route := &ControllerInfo{}
			//route.routerType = routerTypeBeego
			route.methods = map[string]string{"*": rt.Method(i).Name}
			route.controllerType = ct
			pattern := path.Join(prefix, strings.ToLower(controllerName), strings.ToLower(rt.Method(i).Name), "*") //path包Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。结果是经过简化的，所有的空字符串元素会被忽略。
			//patternInit := path.Join(prefix, controllerName, rt.Method(i).Name, "*")
			patternFix := path.Join(prefix, strings.ToLower(controllerName), strings.ToLower(rt.Method(i).Name))
			patternFixInit := path.Join(prefix, controllerName, rt.Method(i).Name)
			route.pattern = pattern
			for m := range HTTPMETHOD {
				p.addToRouter(m, pattern, route)
				p.addToRouter(m, patternInit, route)
				p.addToRouter(m, patternFix, route)
				p.addToRouter(m, patternFixInit, route)
			}
		//}
	}
}
*/

// InSlice checks given string in string slice or not.
func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

/*
type ResponseWriter interface {
    // Header返回一个Header类型值，该值会被WriteHeader方法发送。
    // 在调用WriteHeader或Write方法后再改变该对象是没有意义的。
    Header() Header
    // WriteHeader该方法发送HTTP回复的头域和状态码。
    // 如果没有被显式调用，第一次调用Write时会触发隐式调用WriteHeader(http.StatusOK)
    // WriterHeader的显式调用主要用于发送错误码。
    WriteHeader(int)
    // Write向连接中写入作为HTTP的一部分回复的数据。
    // 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
    // 如果Header中没有"Content-Type"键，
    // 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
    Write([]byte) (int, error)
}

type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string    // path (relative paths may omit leading slash)
    RawPath    string    // encoded path hint (see EscapedPath method)
    ForceQuery bool      // append a query ('?') even if RawQuery is empty
    RawQuery   string    // encoded query values, without '?'
    Fragment   string    // fragment for references, without '#'
}

type Values map[string][]string

v := url.Values{}
v.Set("name", "Ava")
v.Add("friend", "Jess")
v.Add("friend", "Sarah")
v.Add("friend", "Zoe")
// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
fmt.Println(v.Get("name"))
fmt.Println(v.Get("friend"))
fmt.Println(v["friend"])

Output:

Ava
Jess
[Jess Sarah Zoe]


u, err := url.Parse("http://bing.com/search?q=dotnet")
if err != nil {
    log.Fatal(err)
}
u.Scheme = "https"
u.Host = "google.com"
q := u.Query()
q.Set("q", "golang")
u.RawQuery = q.Encode()
fmt.Println(u)

Output:

https://google.com/search?q=golang


*/