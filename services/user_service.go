package services
import(
	"fmt"
	"ieeblog/ieecom/dbhelper"
	"ieeblog/models"
	"github.com/weimingjue/json"
	"time"
)

type userService struct{

}

func newUserService() *userService{
	return &userService{}
}

//UserService public
var UserService=newUserService()


func (userS *userService) GetUserByID(id int64) *models.User{
	db:=dbhelper.Mssql{}
	//获取userMap
	userMap,err:=db.DoQuery("select * from dbo.UserInfo where UserID=?",id)
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//userMap转userStruct
	userModel:=models.User{}
	err=json.Map2Struct(userMap[0],&userModel)
	if err != nil {
        fmt.Println("map2struct error: ", err)
        return nil
	}
	fmt.Println(userModel)
	fmt.Println(userModel.UserNo)
	return &userModel
}

func (userS *userService) GetUserByNamePwd(username string,password string) *models.User{
	db:=dbhelper.Mssql{}

	//获取userMap
	userMap,err:=db.DoQuery("select * from dbo.UserInfo where IsDeleted=0 and UserName=? and Password=?",username,password)
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//userMap转userStruct
	userModel:=models.User{}
	if len(userMap)>0{
		err=json.Map2Struct(userMap[0],&userModel)
		if err != nil {
			fmt.Println("map2struct error: ", err)
			return nil
		}
	}
	fmt.Println(userModel)
	return &userModel
}

//verify the login
func  (userS *userService) LoginValidate(username string,password string) bool{
	fmt.Println("LoginValidate")
	userModel:=userS.GetUserByNamePwd(username,password)
	if userModel==nil{
		return false
	}
	fmt.Println("LoginValidate userModel",userModel)
	return true
}

func(userS *userService) AddUser(username string,email string,password string) (row int64,errMsg string){
	fmt.Println("add user:")
	db:=dbhelper.Mssql{}
	/*
	err:=db.Open()
	if err!=nil{
		fmt.Println("sql open:", err)
        return -1,fmt.Sprintf("%s", err)
	}
	defer db.Close()
	*/

	//if the user is exist
	if isExistUser:=userS.IsExistUserByName(username);isExistUser!=nil{
		return -1,"The username already exists, please enter a new name."
	}
	
	//if the email is exist
	if isExistUser:=userS.IsExistUserByEmail(email);isExistUser!=nil{
		return -2,"The email already exists, please enter a new email."
	}
	//password needs to be encrypted.

	sqlInfo:="insert into dbo.UserInfo(UserName,NickName,Email,Password,DataOwner,CreateDateTime)values(?,?,?,?,?,?)"
	result,err:=db.DoExec(sqlInfo,username,username,email,password,"000000",time.Now().Format("2006-01-02 15:04:05"))
	if err!=nil{
		fmt.Println("sql exec:", err)
        return -3,fmt.Sprintf("%s", err)
	}
	/*
	stmt,err:=db.Prepare("insert into dbo.UserInfo(UserName,Email,Password,DataOwner,CreateDateTime)values(?,?,?,?,?)")
	if err!=nil{
		fmt.Println("sql prepare:", err)
        return -3,fmt.Sprintf("%s", err)
	}
	result,err:=stmt.Exec(username,email,password,"000000",time.Now().Format("2006-01-02 15:04:05"))
	if err!=nil{
		fmt.Println("sql exec:", err)
        return -3,fmt.Sprintf("%s", err)
	}
	*/
	row,_=result.RowsAffected()
	fmt.Println("RowsAffected:",row)
	return row,""
}

func (userS *userService) IsExistUserByName(username string) *models.User{
	db:=dbhelper.Mssql{}

	//获取userMap
	userMap,err:=db.DoQuery("select * from dbo.UserInfo where IsDeleted=0 and UserName=?",username)
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//userMap转userStruct
	userModel:=models.User{}
	if len(userMap)>0{
		err=json.Map2Struct(userMap[0],&userModel)
		if err != nil {
			fmt.Println("map2struct error: ", err)
			return nil
		}
	}else{
		return nil
	}
	return &userModel
}
func (userS *userService) IsExistUserByEmail(email string) *models.User{
	db:=dbhelper.Mssql{}

	//获取userMap
	userMap,err:=db.DoQuery("select * from dbo.UserInfo where IsDeleted=0 and Email=?",email)
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//userMap转userStruct
	userModel:=models.User{}
	if len(userMap)>0{
		err=json.Map2Struct(userMap[0],&userModel)
		if err != nil {
			fmt.Println("map2struct error: ", err)
			return nil
		}
	}else{
		return nil
	}
	return &userModel
}
//GetFollowerListByUserID list
func (userS *userService) GetFollowerListByUserID() *[]models.Follower{
	db:=dbhelper.Mssql{}
	/*分页

	*/
	//获取Map
	followerMap,err:=db.DoQuery("select * from dbo.Follower where IsDeleted=0")//分页
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//Map转Struct
	followerModel:=make([]models.Follower,len(followerMap))
	for i:=0;i<len(followerMap);i++{
		err=json.Map2Struct(followerMap[i],&followerModel[i])
		if err != nil {
			fmt.Println("map2struct error: ", err)
			return nil
		}
	}
	return &followerModel
}

//GetFollowingListByUserID list
func (userS *userService) GetFollowingListByUserID() *[]models.Follower{
	db:=dbhelper.Mssql{}
	/*分页

	*/
	//获取Map
	followingMap,err:=db.DoQuery("select * from dbo.Follower where IsDeleted=0")//分页
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//Map转Struct
	followingModel:=make([]models.Follower,len(followingMap))
	for i:=0;i<len(followingMap);i++{
		err=json.Map2Struct(followingMap[i],&followingModel[i])
		if err != nil {
			fmt.Println("map2struct error: ", err)
			return nil
		}
	}
	return &followingModel
}

/*
stmt,err:=db.Prepare("insert into dbo.UserInfo(UserNo,UserName,NickName,Password,DataOwner,CreateDateTime)values(?,?,?,?,?,?)")
	if err!=nil{
		fmt.Println("sql prepare:", err)
        return -1
	}
	result,err:=stmt.Exec("000002","iee","ieee","123456","000000",time.Now().Format("2006-01-02 15:04:05"))
*/