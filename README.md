# learn_gql

自分の学習用。下記記事と同じ。

【Go言語】はじめてのGraphQLサーバ実装 | gqlgen - Qiita
https://qiita.com/hiroyky/items/4d7764172e73ff54f18b

## コード生成の注意

元のチュートリアルにある生成方法だとエラーになる。以下を実行すること。

```
go run -mod=mod github.com/99designs/gqlgen generate
```