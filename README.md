# 準備

## go

go-app ディレクトリで実行ファイル(main.go)を用意しておく。名前を変更する場合は Docker/go/Dockerfile を編集。

## react

Dockerfile では node のみが使えて react-app ディレクトリ以下を丸々コピペする内容になっているので事前に react-app にて任意の React ディレクトリを入れておく。

注意：Dockerfile 上では node は 16.15.0 になっているので必要に応じて変更が必要  
(Docker 上で create-react-app して詰まったのと Next や GatsbyJS なども入れられるようにコピペで対応させるようにした)

例：$react-app npx crate-react-app . --typescript
