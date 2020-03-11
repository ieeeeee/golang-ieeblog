package controllers

import (
    "net/http"
    "fmt"
    "ieeblog/services"
	"strconv"
	"io"
)

//ArticleController inherit base controller
type ArticleController struct {
    Controller
}
//GetArticleList article list Get()
func(articleCtrl *ArticleController) GetArticleList(w http.ResponseWriter, r *http.Request){
    articleCtrl.Init(w,r,"ArticleController")
	articleCtrl.Data["Articles"] = services.ArticleService.GetArticleList()
	articleCtrl.Layout=append(articleCtrl.Layout,"master/header.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"master/footer.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"article/list.tpl") 
	articleCtrl.TplNames="content"
	err:=articleCtrl.Render()
	if err!=nil{
		fmt.Println("render article list error",err)
	}
}
//GetArticleByID article detail Get()
func(articleCtrl *ArticleController) GetArticleByID(w http.ResponseWriter, r *http.Request){
    articleCtrl.Init(w,r,"ArticleController")
	id,_:= strconv.ParseInt(articleCtrl.Request.Form["id"][0],0,64)
	articleCtrl.Data["Article"] = services.ArticleService.GetArticleByID(id)
	articleCtrl.Layout=append(articleCtrl.Layout,"master/header.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"master/footer.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"article/detail.tpl") 
	articleCtrl.TplNames="content"
	err:=articleCtrl.Render()
	if err!=nil{
		fmt.Println("render article detail error",err)
	}
}

//GetArticleListByUserID article detail Get()
func(articleCtrl *ArticleController) GetArticleListByUserID(w http.ResponseWriter, r *http.Request){
    articleCtrl.Init(w,r,"ArticleController")
	id,_:= strconv.ParseInt(articleCtrl.Request.Form["id"][0],0,64)
	articleCtrl.Data["Articles"] = services.ArticleService.GetArticleListByUserID(id)
	articleCtrl.Layout=append(articleCtrl.Layout,"master/header.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"master/footer.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"article/user.tpl") 
	articleCtrl.TplNames="content"
	err:=articleCtrl.Render()
	if err!=nil{
		fmt.Println("render article list by user error",err)
	}
}


//GetArticleNew load article edit page
func(articleCtrl *ArticleController) GetArticleNew(w http.ResponseWriter, r *http.Request){
    articleCtrl.Init(w,r,"ArticleController")
	//id,_:= strconv.ParseInt(articleCtrl.Request.Form["id"][0],0,64)
	//articleCtrl.Data["Article"] = services.ArticleService.GetArticleByID(id)
	articleCtrl.Layout=append(articleCtrl.Layout,"master/header.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"master/footer.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"article/edit.tpl") 
	articleCtrl.TplNames="content"
	err:=articleCtrl.Render()
	if err!=nil{
		fmt.Println("render article edit error",err)
	}
}

//PostArticleNew article edit post button
func(articleCtrl *ArticleController) PostArticleNew(w http.ResponseWriter, r *http.Request){
	articleCtrl.Init(w,r,"ArticleController")
	fmt.Println("PostArticleNew")
	//data:=articleCtrl.Params
	//id,_:= strconv.ParseInt(articleCtrl.Request.Form["id"][0],0,64)
	ok:= services.ArticleService.AddArticle(articleCtrl.Params)
	io.WriteString(w,fmt.Sprintf("%v",ok))
	/*
	articleCtrl.Layout=append(articleCtrl.Layout,"master/header.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"master/footer.tpl") 
	articleCtrl.Layout=append(articleCtrl.Layout,"article/list.tpl") 
	articleCtrl.TplNames="content"
	err:=articleCtrl.Render()
	if err!=nil{
		fmt.Println("render article post error",err)
	}
	*/
}

