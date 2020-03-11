{{define "content"}}
{{template "header"}}
<div id="main" class="ui main container">
  <div class="box wrap-left">
      <div id="main-doc" class="ui text">
            <div class="ui form" action="">
                <div class="field">
                    <input type="text" id="article-title" name="title" placeholder="Please enter the title" autocomplete="off">
                </div>
                <div class="field">
                    <div class="article-editor" id="article-editor">
                      
                    </div>
                </div>
                <div class="field">
                  <label >Article Tag</label>
                  <select multiple="" class="ui dropdown">
                    <option value="">Select Tag</option>
                    <option value="0">Golang</option>
                    <option value="1">Redis</option>
                    <option value="2">SQL</option>
                    <option value="3">JS</option>
                    <option value="4">Note</option>
                  </select>
                </div>
                <div class="inline field">
                  <label>Article Column</label>
                  <div class="ui checkbox">
                    <input type="checkbox" tabindex="0" name="like[write]" title="写作">
                    <input type="checkbox" tabindex="0"  name="like[read]" title="阅读" checked>
                    <input type="checkbox" tabindex="0" name="like[dai]" title="发呆">
                  </div>
                </div>
                <div class="inline fields">
                  <label for="releaseForm">Release Form</label>
                  <div class="field">
                    <div class="ui radio checkbox">
                      <input type="radio" name="isPrivate" value="0" title="Public" class="hidden">
                      <label>Public</label>
                    </div>
                    <div class="ui radio checkbox">
                      <input type="radio" name="isPrivate" value="1" title="Private" checked>
                      <label>Private</label>
                    </div>
                    
                  </div>
                </div>
               <div>
                <button class="ui secondary button" id="postArticle">Post Article</button>
                <button type="reset" class="ui button">Save Draft</button>
               </div>
              </div>
      <!--
<div class="form-content">
                    <div class="article-input-title">
                        <input type="text" id="article-title" name="titles" placeholder="请输入文章标题"/>
                    </div>
                    <div class="article-editor" id="article-editor" name="content">

                    </div>
                    <div class="article-extra-set">
                        <div class="extra-item"></div>
                        <div class="extra-item">
                            <input type="button" id="postArticle" value="Post Article"/>
                            <button type="button" >Save Draft</button>
                        </div>
                    </div>
                </div>
      -->
                

      </div>
    </div>
</div>
  <script type="text/javascript" src="/static/js/wangEditor.min.js"></script>
    <script type="text/javascript">
    //Init a editor
        var E = window.wangEditor
        var editor = new E('#article-editor')
        // 或者 var editor = new E( document.getElementById('editor') )
        editor.create()

        //Post Article Button Event
        $("#postArticle").click(function(){
            //editor.txt.html()
           // alert(editor.txt.html())
           // alert($("#article-title").val())
            $.post("/article/post",{"title":$("#article-title").val(),"content":editor.txt.html()},function(result){
                console.log(result);
                if(result.ok=="true"){
                  layer.msg("Posted");
                }
            },"json")
            
        })
    </script>
  {{template "footer"}}
{{end}}

