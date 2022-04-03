# learn_gql

自分の学習用。下記記事と同じ。

【Go言語】はじめてのGraphQLサーバ実装 | gqlgen - Qiita
https://qiita.com/hiroyky/items/4d7764172e73ff54f18b

以下は学習メモ。

## コード生成の注意

元のチュートリアルにある生成方法だとエラーになる。以下を実行すること。

```
go run -mod=mod github.com/99designs/gqlgen generate
```

## Go側のモデル実装

Go側のモデル実装はゼロから書いて苦労するより、自動生成されたモデルに手を入れた方が良いという気付きを得た。

### 自動生成コードの重複定義

https://qiita.com/hiroyky/items/4d7764172e73ff54f18b#%E3%82%B3%E3%83%BC%E3%83%89%E7%94%9F%E6%88%902-%E6%A7%8B%E9%80%A0%E4%BD%93%E3%82%92%E7%B7%A8%E9%9B%86%E3%81%97%E3%81%A6%E5%86%8D%E7%94%9F%E6%88%90

上記のように自動生成されたモデルを再定義しても、再生成の都度にmodels_gen.goにまた同じモデルが生成されてしまう。

これは毎回手作業で削除するってことなんだろうか。
