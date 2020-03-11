//session provider is memory

package session
import(
	"container/list"
	"sync"
	"time"
)

/*
1,define a sessionStore,as a receiver to implement session interface
2,define a FromMemory, to implement the provider
*/

//provider is memory


var pder=&FromMemory{list:list.New()}

func init(){
	pder.sessions=make(map[string]*list.Element,0)
	Register("memory",pder)//注册 memory 
}

//SessionStore struct Session 实现其接口？
type SessionStore struct{
	sid string
	LastAccessedTime time.Time
	value map[interface{}]interface{}
}
//Set implement
func (st *SessionStore) Set(key, value interface{}) error{
	st.value[key]=value
	pder.SessionUpdate(st.sid)
	return nil
}
//Get implement
func(st *SessionStore) Get(key interface{}) interface{}{
	pder.SessionUpdate(st.sid)
	if v,ok:=st.value[key];ok{
		return v
	}
	return nil
}
//Delete implement
func(st *SessionStore) Delete(key interface{})error{
	delete(st.value,key)
	pder.SessionUpdate(st.sid)
	return nil
}
//SessionID implement
func(st *SessionStore) SessionID() string{
	return st.sid
}

//FromMemory 来自于内存的provider实现
type FromMemory struct{
	lock sync.Mutex
	sessions map[string]*list.Element //用来存储在内存
	list *list.List
}
//SessionInit Provider interface implement
func(frommemory *FromMemory) SessionInit(sid string)(Session,error){
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	v:=make(map[interface{}]interface{},0)
	newsess:=&SessionStore{
		sid:sid,
		LastAccessedTime:time.Now(),
		value:v,
	}
	element:=frommemory.list.PushBack(newsess)
	frommemory.sessions[sid]=element
	return newsess,nil
}
//SessionRead Provider interface implement
func(frommemory *FromMemory) SessionRead(sid string)(Session,error){
	if element,ok:=frommemory.sessions[sid];ok{
		return element.Value.(*SessionStore),nil
	}else{
		sess,err:=frommemory.SessionInit(sid)
		return sess,err
	}
	return nil,nil
}
//SessionDestroy Provider interface implement
func(frommemory *FromMemory) SessionDestroy(sid string) error{
	if element,ok:=frommemory.sessions[sid]; ok{
		delete(frommemory.sessions,sid)
		frommemory.list.Remove(element)
		return nil
	}
	return nil
}
//SessionGC Provider interface implement
func(frommemory *FromMemory) SessionGC(maxLifeTime int64){
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	for{
		element:=frommemory.list.Back()
		if element==nil{
			break
		}
		if(element.Value.(*SessionStore).LastAccessedTime.Unix()+maxLifeTime)<time.Now().Unix(){
			frommemory.list.Remove(element)
			delete(frommemory.sessions,element.Value.(*SessionStore).sid)
			
		}else{
			break
		}
	}
}

//SessionUpdate Provider interface implement
func(frommemory *FromMemory) SessionUpdate(sid string) error{
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	if element,ok:=frommemory.sessions[sid];ok{
		element.Value.(*SessionStore).LastAccessedTime=time.Now()
		frommemory.list.MoveToFront(element)
		return nil
	}
	return nil
}