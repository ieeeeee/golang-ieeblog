package models

import(
	"fmt"
	"database/sql"
)


func GetAll()(rows *sql.Rows){
	db:=ieecom.GetDB()
	err:=db.Open()
	if err!=nil{
		fmt.Println("sql open:", err)
        return
	}

	defer db.Close()
	rows,err:=db.Query("select * from dbo.Article")
	if err != nil {
        fmt.Println("query: ", err)
        return
	}
	return rows
}