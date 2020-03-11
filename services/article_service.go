package services
import(
	"fmt"
	"ieeblog/ieecom/dbhelper"
	"ieeblog/models"
	"github.com/weimingjue/json"
	"time"
	//"html"
)

type articleService struct{

}

func newArticleService() *articleService{
	return &articleService{}
}
//ArticleService public
var ArticleService=newArticleService()

func (articleS *articleService) GetArticleList() *[]models.Article{
	//db:=dbhelper.GetDB()
	fmt.Println("article service GetList:")
	db:=dbhelper.Mssql{}
	/*分页
	select * from (
select distinct row_number() over (order by a.articleID)as rowID,* from 
Article a where IsDeleted=0
) as Result 
where rowID between ((pageNum-1)*pageSize)+1 and (pageNum*pageSize)

	*/
	//获取Map
	articleMap,err:=db.DoQuery("select * from dbo.Article where IsDeleted=0")//分页
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//Map转Struct
	articleModel:=make([]models.Article,len(articleMap))
	for i:=0;i<len(articleMap);i++{
		//fmt.Println("articlemap :",articleMap[i])
		err=json.Map2Struct(articleMap[i],&articleModel[i])
		//fmt.Println("articleModel :", &articleModel[i])
		if err != nil {
			fmt.Println("map2struct error: ", err)
			return nil
		}
	}
	//fmt.Println(articleModel)
	return &articleModel
}

func (articleS *articleService) GetArticleByID(id int64) *models.Article{
	db:=dbhelper.Mssql{}
	//获取Map
	articleMap,err:=db.DoQuery("select * from dbo.Article where ArticleID=?",id)
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//Map转Struct
	articleModel:=models.Article{}
	err=json.Map2Struct(articleMap[0],&articleModel)
	if err != nil {
        fmt.Println("map2struct error: ", err)
        return nil
	}
	fmt.Println(articleModel)
	return &articleModel
}

func (articleS *articleService) GetArticleListByUserID(id int64) *[]models.Article{
	db:=dbhelper.Mssql{}
	/*分页*/
	
	//获取Map
	articleMap,err:=db.DoQuery("select * from dbo.Article where IsDeleted=0 and UserID=?",id)//分页
	if err != nil {
        fmt.Println("query: ", err)
        return nil
	}
	//Map转Struct
	articleModel:=make([]models.Article,len(articleMap))
	for i:=0;i<len(articleMap);i++{
		//fmt.Println("articlemap :",articleMap[i])
		err=json.Map2Struct(articleMap[i],&articleModel[i])
		//fmt.Println("articleModel :", &articleModel[i])
		if err != nil {
			fmt.Println("map2struct error: ", err)
			return nil
		}
	}
	return &articleModel
}

//AddArticle return ok
func(articleS *articleService) AddArticle(params map[string]string) bool{
	fmt.Println("add article:")
	db:=dbhelper.Mssql{}
/*
	err:=db.Open()
	if err!=nil{
		fmt.Println("sql open:", err)
        return false
	}

	defer db.Close()
	*/

	sqlInfo:="insert into dbo.Article(UserID,Title,Content,IsPrivate,DataOwner,CreateDateTime,LastChangedBy,LastChangedDateTime)values(?,?,?,?,?,?,?,?)"
	_,err:=db.DoExec(sqlInfo,"1",params["title"],params["content"],"1","000000",time.Now().Format("2006-01-02 15:04:05"),"000000",time.Now().Format("2006-01-02 15:04:05"))
	//_,err=stmt.Exec("1","Article Test Title2","Article Test Content2","1","000000",time.Now().Format("2006-01-02 15:04:05"),"000000",time.Now().Format("2006-01-02 15:04:05"))
	if err!=nil{
		fmt.Println("sql exec:", err)
        return false
	}
/*
	stmt,err:=db.Prepare("insert into dbo.Article(UserID,Title,Content,IsPrivate,DataOwner,CreateDateTime,LastChangedBy,LastChangedDateTime)values(?,?,?,?,?,?,?,?)")
	if err!=nil{
		fmt.Println("sql prepare:", err)
        return false
	}
	_,err=stmt.Exec("1",params["title"],params["content"],"1","000000",time.Now().Format("2006-01-02 15:04:05"),"000000",time.Now().Format("2006-01-02 15:04:05"))
	//_,err=stmt.Exec("1","Article Test Title2","Article Test Content2","1","000000",time.Now().Format("2006-01-02 15:04:05"),"000000",time.Now().Format("2006-01-02 15:04:05"))
	if err!=nil{
		fmt.Println("sql exec:", err)
        return false
	}
	*/
	//userID,_:=result.LastInsertId()
	//fmt.Println("add user LastInsertId:", userID)
	return true
}

/*
select * from (
select distinct row_number() over (order by a.articleID)as rowID,* from 
Article a where IsDeleted=0
) as Result 
where rowID between ((pageNum-1)*pageSize)+1 and (pageNum*pageSize)

--insert into Article(,,,)

--insert into dbo.Article(UserID,Title,Content,IsPrivate,DataOwner,CreateDateTime,LastChangedBy,LastChangedDateTime)
values(1,'Article test title1','Article test content1',1,'000000',CONVERT(varchar,GETDATE(),20),'000000',CONVERT(varchar,GETDATE(),20))

*/