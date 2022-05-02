learn_gql
====

自分の学習用。下記記事と同じ。

【Go言語】はじめてのGraphQLサーバ実装 | gqlgen - Qiita
https://qiita.com/hiroyky/items/4d7764172e73ff54f18b

以下は学習メモ。

# コード生成の注意

元のチュートリアルにある生成方法だとエラーになる。以下を実行すること。

```
go run -mod=mod github.com/99designs/gqlgen generate
```

# Go側のモデル実装

Go側のモデル実装はゼロから書いて苦労するより、自動生成されたモデルに手を入れた方が良いという気付きを得た。

## 自動生成コードの重複定義

https://qiita.com/hiroyky/items/4d7764172e73ff54f18b#%E3%82%B3%E3%83%BC%E3%83%89%E7%94%9F%E6%88%902-%E6%A7%8B%E9%80%A0%E4%BD%93%E3%82%92%E7%B7%A8%E9%9B%86%E3%81%97%E3%81%A6%E5%86%8D%E7%94%9F%E6%88%90

上記のように自動生成されたモデルを再定義しても、再生成の都度にmodels_gen.goにまた同じモデルが生成されてしまう。

これは毎回手作業で削除するってことなんだろうか。

### 重複定義の解決

再定義したモデルが自動生成で重複定義されないためには、```gqlgen.yml```で以下を定義する必要がある。

```
autobind:
#  - "github.com/eyasuyuki/learn_gql/graph/model"
```

自分のリポジトリへのautbindのコメントを解除する。

```
autobind:
  - "github.com/eyasuyuki/learn_gql/graph/model"
```

# ```mutation.resolver.go```が生成されない?

原因はtypoだった。Mutationの綴りが間違っていた。訂正したら生成された。

# テストを書く

とりあえず ```graph/model/models_test.go```は書いたが、resolverのテストをどう書けば良いのか分からない。

https://github.com/99designs/gqlgen/blob/master/TESTING.md

これはgqlgenそのもののテストのことで、gqlgenで生成したサーバーのテストではないし。

## GraphQLのクエリを自動的にテストする - Qiita

https://qiita.com/pocke/items/bfe120f07bd8d94724a7

その後これを見つけたがRuby書きたくない。

## timqian/gql-generator: Generate queries from graphql schema, used for writing api test.

https://github.com/timqian/gql-generator

これやってみよう。

### 単一のスキーマファイルにしか対応していない

なので生成できない....

## ```time.Now()```を文字列化してparseするとタイムゾーンが違うのでテストが失敗する

```go
const TIMESTAMP_PATTERN = "2006-01-02 15:04:05.999999999"
now := time.Now() // +0900
nowStr := now.Format(TIMESTAMP_PATTERN) // タイムゾーンつかない
parsedTime := time.Parse(TIMESTAMP_PATTERN, nowStr) // UTCとしてparseされる
now == parsedTime // false
```

## ```Format```の前にUTCにしておく

```go
utc,_ := time.LoadLocation("UTC")
nowUtc := time.Now().In(utc)
nowStr := nowUtc.Format(TIMESTAMP_PATTERN) // UTCに変更されてフォーマットされる
parsedTime := time.Parse(TIMESTAMP_PATTERN, nowStr)
nowUtc == parsedTime // true
```

### MariaDB(MySQL)への接続で```loc=UTC```を指定する

これがないとUTCの値を保存してもデフォルトタイムゾーンの値が返ってくる。

## テストを書くために実装を```service```パッケージに移動した

実装を```service```パッケージに移動してユニットテストを書いた。

# Mutationにリレーションについての定義がない

```departments```テーブルは```company_id```列、```employees```テーブルは```department_id```列、```company_id````列を持っているが、誰がこれをupdateするのか分からない。一般的にどう解決しているのだろうか? gqlgenのドキュメントを読む方が早いか?

## Mutationの```xxxInput```定義にリレーションがないのが原因?

結局```mutation.graphqls```の```xxxInput```の定義にリレーションが一切ないのでリレーションに関する定義が生成されないのではなかろうか。

## Mutationでリレーションを張るディレクティブを増やす

```graphql
addCompanyToDepartment(id: ID!, companyID: ID!): Department!
addDepartmentToEmployee(id: ID!, departmentID: ID!): Employee!
addCompanyToEmployee(id: ID!, companyID: ID!): Employee!
```

結局こういうのを増やしたけどこれで良かったんだろうか。