{{define "content"}}
{{template "header"}}
<div class="ui main container">
    <div class="box wrap-left">
        <div class="ui items ">
        {{range .Articles}}
        <div class="item">
            <div class="ui image middle article-list-cover ">
                <a href="/article/detail?id={{.ArticleID}}"><img src="/static/images/cat-walk.png"></a>
            </div>
            <div class="content">
                <a class="header" href="/article/detail?id={{.ArticleID}}">{{.Title}}</a>
                <div class="description" id="article-description">
                    {{.Content}}
                </div>
                <div class="extra">
                    <div class="ui label"><i class="eye icon"></i>10</div>
                    <div class="ui label"><i class="clock icon"></i>{{.LastChangedDateTime}}</div>
                </div>
            </div>
        </div>
        <div class="ui fitted divider"></div>
        
        {{end}} 
        </div>
    </div>
</div>
<script type="text/javascript">
    document.getElementById("article-description").innerHTML=$("#article-description").text()
</script>
{{template "footer"}}
{{end}}

