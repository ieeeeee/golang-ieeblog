<div class="right-side">
    <main id="main-doc" class="doc">   
{{range .articles}}
      <section id="article_{{.id}}" class="main-section">
        <header><a href="/blog/article/{{.ArticleID}}">{{.Title}}</a></header>
        <article>
          {{.Content}}
        </article>
        <dl class="author-bar">
            <dt></dt>
            <dd class="author-name">
                <a href="/user/bloglist{{.UserID}}"></a>
            </dd>
            <div class="active-count">
                
            </div>
        </dl>
      </section>
{{end}}
    </main>
  </div>

