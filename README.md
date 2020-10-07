go_simple_blog

標準パッケージを利用してシンプルなブログAppを作成



フォルダ構成
go-simple-blog (MVC)
 -handler
  - handler.go (MVCのC的な役割：リクエストを受け取りHTMLを返す) 
 -model
  - model.go (MVCのM的な役割：DBとの接続)
 -templates
  - (index,new,show).html
 -src
  - js 
  - css 
 -main.go
  - 
 
 
URL構成 

GET  /             => topページ(一覧ページ)
GET  /articles/new => 記事投稿ページ   
GET  /articles/id  => 記事詳細ページ
POST /articles/    => 記事投稿(一覧ページに飛ぶ)