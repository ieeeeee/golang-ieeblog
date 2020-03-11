{{define "login"}}
<html>
<head>
<title>SignIn/SignUp</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta charset="utf-8" />
<link type="text/css" href="/static/css/bootstrap.css" rel="stylesheet">
<link type="text/css" href="/static/other/semantic/semantic.min.css" rel="stylesheet" />
<link type="text/css" href="/static/css/ieestyle.css" rel="stylesheet" />
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
<script type="text/javascript" src="/static/other/semantic/semantic.min.js"></script> 
<style>
    body > .grid {
      height: 100%;
    }
    .image {
      margin-top: -100px;
    }
  .column {
      max-width: 450px;
    }
</style>
</head>
<body>
    <div class="ui middle aligned center aligned grid">
        <div class="column">
            <h2 class="ui teal image header">
                <img class="image" src="/static/images/cat-walk.png">
                <div class="content">Log-in to your account</div>
            </h2>
            <form class="ui large form" method="post">
                <div class="ui stacked segment">
                  <div class="field">
                    <div class="ui left icon input">
                      <i class="user icon"></i>
                      <input type="text" name="username" placeholder="Name">
                    </div>
                  </div>
                  <div class="field">
                    <div class="ui left icon input">
                      <i class="lock icon"></i>
                      <input type="password" name="password" placeholder="Password">
                    </div>
                  </div>
                  <div class="ui fluid large teal submit button">Login</div>
                </div>
          
                <div class="ui error message"></div>
          
            </form>
            <div class="ui message">
                New to us?<a href="/login/signup">Sign Up</a>
            </div>
        </div>
    </div>
    <script>
        $(document)
          .ready(function() {
            $('.ui.form')
              .form({
                fields: {
                  email: {
                    identifier  : 'email',
                    rules: [
                      {
                        type   : 'empty',
                        prompt : 'Please enter your account name'
                      }
                    ]
                  },
                  password: {
                    identifier  : 'password',
                    rules: [
                      {
                        type   : 'empty',
                        prompt : 'Please enter your password'
                      },
                      {
                        type   : 'length[6]',
                        prompt : 'Your password must be at least 6 characters'
                      }
                    ]
                  }
                }
              })
            ;
          })
        ;
        </script>
</body>
</html>
{{end}}