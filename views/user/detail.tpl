{{define "content"}}
{{template "header"}}
<div id="main" class="container">
  <div class="right-side">
      <main>
          <h1>Personal Information</h1>
          <div id="form-outer">
            <p id="description">Hi, I am designing a personal accounting management application</p>
            <form id="user-form">
              <div class="rowTab">
                <div class="labels">
                  <label id="name-label" for="name">Name:</label>
                </div>
                <div class="rightTab">
                  <input type="text" id="name" name="name" class="input-field" value="{{.User.NickName}}"/>
                </div>
              </div>
              <div class="rowTab">
                <div class="labels">
                  <label id="email-label" for="email">Email:</label>
                </div>
                <div class="rightTab">
                  <input type="email" id="email" name="email" class="input-field" value="{{.User.Email}}"/>
                </div>
              </div>
              <div class="rowTab">
                <div class="labels">
                  <label id="number" for="age">Age:</label>
                </div>
                <div class="rightTab">
                  <input type="text"  id="number" name="age" class="input-field" value="{{.User.Age}}"/>
                </div>
              </div> 
              <div class="rowTab">
                <div class="labels">
                  <label for="comments">Personal Introduction</label>
                </div>
                <div class="rightTab">
                  <textarea id="comments" name="comments" class="input-field" style="height:50px; resize:vertical;" value="{{.User.Introduction}}"></textarea>
                </div>
              </div>
              <div class="RowTab">
                <button type="submit" id="submit">Submit</button>
              </div>
            </form>
          </div>
        </main>
    </div>
</div>
  {{template "footer"}}
{{end}}

