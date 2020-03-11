package session

import (
	"fmt"
	"net/http"
	"sync"
	"crypto/rand"
	"io"
	"net/url"
	"time"
	"encoding/base64"
)
/*
1,define a session manager
2,new sessionManager func
3,define a provider
4,provider interface{Init(),Read(),Destory(),GC()}
5,session interface{Set(),Get(),Delete(),SessionID}
5,register, makes a session provide available by the provided name
6,new a sessionID
7,cookie->session,return(session.Read()) or new (session.Init())
*/


//Manager  define a global session manager
type Manager struct{
	cookieName string //private cookiename
	lock sync.Mutex //protects session  互斥锁
	provider Provider  //session保存的底层存储结构
	maxlifetime int64 
}

//NewManager new a session manager
func NewManager(provideName,cookieName string, maxlifetime int64)(*Manager,error){
	provider,ok:=provides[provideName]
	if!ok{
		return nil,fmt.Errorf("session:unkown provide %q(forgotten import?)",provideName)
	}
	return &Manager{provider:provider,cookieName:cookieName,maxlifetime:maxlifetime},nil

}

//Provider session data storage structure 存储方式 服务器上的内存？文件？数据库？
type Provider interface{
	SessionInit(sid string)(Session,error) 
	SessionRead(sid string)(Session,error)
	SessionDestroy(sid string) error
	SessionGC(maxlifeTime int64)
}
//Session 操作 设置 读取 删除 获取sesssionID
type Session interface{
	Set(key,value interface{}) error //set session value
	Get(key interface{}) interface{} //get session value
	Delete(key interface{}) error  //delete session value
	SessionID() string //back current sessionID
}
var provides=make(map[string] Provider)

//Register makes a session provide available by the provided name
//由实现Provider接口的结构体调用
func Register(name string,provider Provider){
	if provider==nil{
		panic("session:Register provide is nil")
	}
	if _,dup:=provides[name];dup{
		panic("session:Register called twice for provide "+name)
	}
	provides[name]=provider
}

//SessionId New a sessionID
func (manager *Manager) sessionID() string{
	b:=make([]byte,32)
	if _,err:=io.ReadFull(rand.Reader,b);err!=nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//SessionStart 判断当前请求的cookie中是否存在有效的session,存在返回，否则创建
func(manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request)(session Session){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err:=r.Cookie(manager.cookieName)
	if err!=nil||cookie.Value==""{
		sid:=manager.sessionID()
		session,_=manager.provider.SessionInit(sid)
		cookie:=http.Cookie{
			Name:manager.cookieName,
			Value:url.QueryEscape(sid), //url.QueryEscape()转码，才能安全的在url中使用
			Path:"/",
			HttpOnly:true,
			MaxAge:int(manager.maxlifetime),
			Expires:time.Now().Add(time.Duration(manager.maxlifetime)),
		}
		http.SetCookie(w,&cookie)
	}else{
		sid,_:=url.QueryUnescape(cookie.Value) //将转码后的sid还原
		session,_=manager.provider.SessionRead(sid)
	}
	return session
}

//SessionDestroy 重置？
func(manager *Manager) SessionDestroy(w http.ResponseWriter,r *http.Request){
	cookie,err:=r.Cookie(manager.cookieName)
	if err!=nil||cookie.Value==""{
		return
	}else{
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration:=time.Now()
		cookie:=http.Cookie{
			Name:manager.cookieName,
			Path:"/",
			HttpOnly:true,
			Expires:expiration,
			MaxAge:-1,
		}
		http.SetCookie(w,&cookie)
	}
}

//只有在main中的init函数进行启用
func(manager *Manager) GC(){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime),func(){manager.GC()})
}