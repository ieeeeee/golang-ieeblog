{{define "content"}}
{{template "header"}}
<div class="ui main container">
    <div class="box wrap-left">
        <div class="ui items ">
        {{range .Followers}}
        <div class="item">
            <div class="ui image middle article-list-cover ">
                <a href="/article/detail?id={{.FollowerUserID}}"><img src="http://logo.pizza/img/cat-walk/cat-walk.png"></a>
            </div>
            <div class="content">
                <a class="header" href="/user/follower?id={{.FollowerUserID}}">{{.FollowerName}}</a>
                <div class="description" id="article-description">
                    {{.FollowerIntroduction}}
                </div>
                <div class="extra">
                    <div class="ui label"><i class="eye icon"></i>{{.ReadCount}}</div>
                    <div class="ui label"><i class="eye icon"></i>{{.StarCount}}</div>
                    <div class="ui label"><i class="eye icon"></i>{{.LikeCount}}</div>
                    <div class="ui label"><i class="clock icon"></i>{{.LastChangedDateTime}}</div>
                    <button class="ui button">Follow</button>
                </div>
            </div>
        </div>
        <div class="ui fitted divider"></div>
        
        {{end}} 
        </div>
    </div>
</div>
{{template "footer"}}
{{end}}

