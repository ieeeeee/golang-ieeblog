package dbhelper

import(
	"database/sql"
	"strings"
	_"github.com/mattn/go-adodb" //mssql 驱动包
	"ieeblog/ieecom/config" //sql配置从配置文件取
	"fmt"
)

/*
sql.DB manages the open and close operations of the underlying database connection 
for us through the database driver.
Open() creates a DB
调用 db.Query 执行 SQL 语句, 此方法会返回一个 Rows 作为查询的结果,A row is not a hash map, but an abstraction of a cursor
通过 rows.Next() 迭代查询数据.
通过 rows.Scan() 读取每一行的值
调用 db.Close() 关闭查询
QueryRow() 查询行
stmt:=db.Prepare() 准备要执行操作的语句
stmt.Exec() 执行操作 update,insert,delete
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
//插入的id
id, err := result.LastInsertId()
//可以获得影响行数
affect, err := result.RowsAffected()

tx:=db.Beigin() 开启事务
tx.Rollback()
tx.Commit()
设置连接池
db.SetMaxIdleConns(n)
db.SetMaxOpenConns(n)
*/

//Mssql struct
type Mssql struct{
	*sql.DB	
}

//Open Conn
func(m *Mssql) Open()(err error){
	fmt.Println("Open in")
	var conn []string
	conn=append(conn,"Provider=SQLOLEDB")
	conn=append(conn,"Data Source="+config.Conf.Mssql.DataSource)
	if config.Conf.Mssql.Windows{		// windwos: true 为windows身份验证，false 必须设置sa账号和密码
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，
		// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
		conn=append(conn,"integrated security=SSPI")
	}
	conn=append(conn,"Initial Catalog="+config.Conf.Mssql.Database)
	conn=append(conn,"user id="+config.Conf.Mssql.User)
	conn=append(conn,"password="+config.Conf.Mssql.Password)
	m.DB,err=sql.Open("adodb",strings.Join(conn,";"))
	if err!=nil{
		return err
	}
	return nil
}

//重点是如何将结果和struct匹配
//rows转成map map再和model中的struct匹配

//DoQuery query  Rows to Map
func(m *Mssql)DoQuery(sqlInfo string,args ...interface{})([]map[string]interface{},error){
	fmt.Println("sqlhelper doquery")
	err:=m.Open()
	if err!=nil{
		fmt.Println("sql open:", err)
        return nil,err
	}

	defer m.DB.Close()
	
	rows,err:=m.DB.Query(sqlInfo,args...)
	if err!=nil{
		return nil,err
	}
	cols,_:=rows.Columns()
	colsLen:=len(cols)
	cache:=make([]interface{},colsLen) //临时存储每行数据
	for index:=range cache{ //为每一列初始化一个指针
		var a interface{}
		cache[index]=&a
	}
	var listRet []map[string]interface{}
	for rows.Next(){ //每行
		_=rows.Scan(cache...)  
		item:=make(map[string]interface{}) //每列
		for i,data:=range cache{
			item[cols[i]]=*data.(*interface{})
		}
		listRet=append(listRet,item)
	}
	_=rows.Close()
	return listRet,nil
}

//DoExec insert
func(m *Mssql)DoExec(sqlInfo string,vals ...interface{})(sql.Result,error){
	fmt.Println("sqlhelper doexec")
	err:=m.Open()
	if err!=nil{
		fmt.Println("sql open:", err)
        return nil,err
	}

	defer m.DB.Close()
	
	
	stmt,err:=m.DB.Prepare(sqlInfo)
	if err!=nil{
		fmt.Println("sql prepare:", err)
        return nil,err
	}
	result,err:=stmt.Exec(vals...)
	if err!=nil{
		fmt.Println("sql exec:", err)
        return nil,err
	}

	return result,nil
}

/*
type DB struct {
	*sync.RWMutex
	DB    *sql.DB
	stmts map[string]*sql.Stmt
}

func (d *DB) Begin() (*sql.Tx, error) {
	return d.DB.Begin()
}
*/

/*
//Mssql struct
type Mssql struct{
	*sql.DB
	dataSource string
	database string
	windows bool
	sa SA
}

//SA struct
type SA struct{
	user string
	password string
}
//DB global //注意逗号,
var DB=Mssql{
	dataSource:config.Conf.Mssql.DataSource,  
	database:config.Conf.Mssql.Database,
	windows:config.Conf.Mssql.Windows,
	sa:SA{
		user:config.Conf.Mssql.User,
		password:config.Conf.Mssql.Password,
	},
}

//Open Conn
func(m *Mssql) Open()(err error){
	fmt.Println("Open in")
	var conn []string
	conn=append(conn,"Provider=SQLOLEDB")
	conn=append(conn,"Data Source="+m.dataSource)
	if m.windows{		// windwos: true 为windows身份验证，false 必须设置sa账号和密码
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，
		// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
		conn=append(conn,"integrated security=SSPI")
	}
	conn=append(conn,"Initial Catalog="+m.database)
	conn=append(conn,"user id="+m.sa.user)
	conn=append(conn,"password="+m.sa.password)
	m.DB,err=sql.Open("adodb",strings.Join(conn,";"))
	if err!=nil{
		return err
	}
	return nil
}
*/
//应从配置文件取,配置文件写成XMl
/*
func GetDB()(db *sql.DB){
	db:=Mssql{
		dataSource:"." //127.0.0.1\\SQLEXPRESS
		database:"gowebblog"
		windows:false 
		sa:SA{
			user:"sa",
			password:"zld607608",
		},
	}
}
*/

//GetDB config