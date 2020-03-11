{{define "content"}}
{{template "header"}}
<div id="main" class="ui main container">
    <div class="box wrap-left">
        <div id="main-doc" class="ui text">
            <div id="article_{{.Article.ArticleID}}" class="ui content">
                <h1 class="ui header">{{.Article.Title}}</h1>
                <div class="ui left detail-article-extra">
                    <div class="ui label"><i></i>Author</div>
                    <div class="ui label"><i></i>PostTime</div>
                    <div class="ui label"><i class="eye icon"></i>ReadCount</div>
                </div>
                <div class="ui message">
                    Article description need
                </div>
                <div class="description detail-article-content">
                    <div class="detail-article-content-view" id="article-content-view">
                            {{.Article.Content}}                 
                    </div>
                    
                </div>
                <div class="ui bottom detail-article-action">
                    <div class="ui  labeled button" tabindex="0">
                        <div class="ui basic red button">
                            <i class="heart icon"></i> Star
                          </div>
                        <a class="ui basic red left pointing label">
                          2,048
                        </a>
                        
                    </div>
                    <div class="ui left labeled button" tabindex="0">
                        <div class="ui basic orange button">
                            <i class="heart icon"></i> Star
                          </div>
                        <a class="ui basic orange left pointing label">
                          2,048
                        </a>
                        
                    </div>
                    <div class="ui left labeled button" tabindex="0">
                        <div class="ui basic blue button">
                            <i class="share icon"></i>Share
                        </div>
                        <a class="ui basic left pointing blue label">
                            1,048
                        </a>
                        
                    </div>
                </div>
                <div class="ui message">Notice Message need</div>
            </div>
        </div>
    </div>
</div>
<script type="text/javascript">
    document.getElementById("article-content-view").innerHTML=$("#article-content-view").text()
</script>
{{template "footer"}}
{{end}}

