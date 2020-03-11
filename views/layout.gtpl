{{define "masterLayout"}}
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta charset="utf-8" />
    <!--<meta name="viewport" content="width=device-width, initial-scale=1.0">-->
    <link type="text/css" href="/static/css/bootstrap.css" rel="stylesheet">
    <link type="text/css" href="/static/css/font-awesome.css" rel="stylesheet" />
    <link type="text/css" href="/static/css/icon-font.min.css" rel="stylesheet" />
    <link type="text/css" href="/static/css/mainlayout.css" rel="stylesheet" />
    <script type="application/javascript" src="/static/js/bootstrap.min.js"></script>
    <script type="application/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
    <script type="application/javascript" src="/static/js/bootstrap-transition.js"></script>    
</head>

<body>
  <div id="pageDiv">
    <header id="header">
      <div class="logo">
        <img id="header-image" src="http://logo.pizza/img/cat-walk/cat-walk.png" alt="Cloud Cat LOGO"/>
      </div>
      <nav id="nav-bar">
        <ul>
          <li class="nav-link" href="/">Blog</li>
          <li class="nav-link" href="/">Mine</li>
          <li class="nav-link" href="/following">Following</li>
          <li class="nav-link" href="/follower">Follower</li>
          <li class="nav-link" href="/download">Download</li>
          <li class="nav-link" href="/login">Sign in/Sign up</li>
        </ul>
      </nav>
    </header>
    <div id="main" class="container">
      {{end}}
        <!--{{.LayoutContent}}-->

{{define "footerLayout"}}
    </div>
    <div class="footer">
      <div class="footer-bg">
        <div class="footer-left">
          <div class="share-link-list">
            <a href="#">Facebook</a>
            <a href="#">Twitter</a>
          </div>
        </div>
        <div class="footer-right">
          <ul>
            <li>
              <a href="#">Privacy</a>
            </li>
            <li>
              <a href="#">Terms</a>
            </li>
            <li>
              <a href="#">Contact</a>
            </li>
          </ul>
          <!--<span>Copyright 2020, All Rights Reserved.</span>-->
        </div>
      </div>
    </div>
  </div> 
</body>
</html>
{{end}}