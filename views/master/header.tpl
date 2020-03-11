{{define "header"}}
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta charset="utf-8" />
    <!--<meta name="viewport" content="width=device-width, initial-scale=1.0">-->
    <link type="text/css" href="/static/css/bootstrap.css" rel="stylesheet">
    <link type="text/css" href="/static/css/font-awesome.css" rel="stylesheet" />
    <link type="text/css" href="/static/css/icon-font.min.css" rel="stylesheet" />
    <link type="text/css" href="/static/css/wangEditor.css" rel="stylesheet" />
    <!--link type="text/css" href="/static/layui/css/layui.css" rel="stylesheet" /-->
    <link type="text/css" href="/static/other/semantic/semantic.min.css" rel="stylesheet" />
    <link type="text/css" href="/static/css/ieestyle.css" rel="stylesheet" />
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <script type="text/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-transition.js"></script>
    <script type="text/javascript" src="/static/other/semantic/semantic.min.js"></script>
    <script type="text/javascript" src="/static/js/layer.js"></script>   
</head>

<body>
  <div id="pageDiv">
    <div class="ui container">
      <div class="ui secondary pointing menu">
          <a class="active item" href="/article/list">Blog</a>
          <a class="item" href="/article/edit">New</a>
          <a class="item" href="/article/user">Mine</a>
          <a class="item" href="/user/following">Following</a>
          <a class="item" href="/user/follower">Follower</a>
          <a class="item" href="/download">Download</a>
        <div class="right menu">
          <div class="item">
            <div class="ui transparent icon input">
              <input type="text" placeholder="Search...">
              <i class="search link icon"></i>
            </div>
          </div>
          <a class="ui item"href="/login">Sign in/Sign up</a>
        </div>
      </div>
      
      <!--
        <div class="ui inverted segment">
        <div class="ui inverted secondary pointing menu">
          
          <a  class="active item" href="/article/list">Blog</a>
          <a class="item" href="/article/edit">New</a>
          <a class="item" href="/article/user">Mine</a>
          <a class="item" href="/user/following">Following</a>
          <a class="item" href="/user/follower">Follower</a>
          <a class="item" href="/download">Download</a>
          <a class="item" href="/login">Sign in/Sign up</a>
        </div>
      </div>
      
              <div class="logo">
        <img id="header-image" src="http://logo.pizza/img/cat-walk/cat-walk.png" alt="Cloud Cat LOGO"/>
      </div>
        <nav id="nav-bar">
        <ul class="layui-nav" lay-filter="">
          <li class="layui-nav-item layui-this"><a href="/article/list">Blog</a></li>
          <li class="layui-nav-item"><a href="/article/edit">New</a></li>
            
          <li class="layui-nav-item"><a href="/article/user">Mine</a></li>
          <li class="layui-nav-item"><a href="/user/following">Following</a></li>
          <li class="layui-nav-item"><a href="/user/follower">Follower</a></li>
          <li class="layui-nav-item"><a href="/download">Download</a></li>  
          <li class="layui-nav-item"><a href="/login">Sign in/Sign up</a></li>
        </ul>
      </nav>
        <nav id="nav-bar">
        <ul>
          <li class="nav-link"><a href="/article/list">Blog</a></li>
          <li class="nav-link"><a href="/article/edit">New</a></li>
          <li class="nav-link"><a href="/article/user">Mine</a></li>
          <li class="nav-link"><a href="/user/following">Following</a></li>
          <li class="nav-link"><a href="/user/follower">Follower</a></li>
          <li class="nav-link"><a href="/download">Download</a></li>
          <li class="nav-link"><a href="/login">Sign in/Sign up</a></li>
        </ul>
      </nav>
      -->
      
    </div>

{{end}}
        <!--{{.LayoutContent}}-->
