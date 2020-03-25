# Overview 

- Twitter Clone App with Golang1.12
- Only server side and endpoints are below
  - get timeline
  - create tweet
  - delete tweet
  - favorite
  - delete favorite
  - reply

# 初めてのGolangアプリで苦労した点、難しかった点

- 設計
  - ネット上を参考にしてもなかなか色んなパターン（DDD, Clean Architecture）がある
  - MVCはそんなに見かけない
  - ただ未経験のDDDなどを採用するハードルが高かったため（時間的制約など）MVCライクな設計に落ち着いた
  - 参考にしたリポジトリ
    - https://github.com/chilledoj/realworld-starter-kit
  
- Router
  - 最終的に採用ライブラリが https://github.com/gorilla/mux
  - デファクトスタンダードっぽく結構使用例もあったのが採用理由
  - ただ、最終的に6つのエンドポイントにも関わらず大きくなってしまったのが問題かなと感じている
    - 修正し、リクエストをControllerに投げる形で解決
    - Routerはシンプルになった
  - どこまでをrouterの責務にして処理を書くかが悩んだが、エンドポイント自体は参照できるし、一つのメソッドがめちゃくちゃ長いわけではない

- Timeline
  - Timelineはusersテーブルとtweetテーブルのデータをいい感じに取得したものたち
  - それらの構造体をどこで定義するのが良いかに迷った
  - 最終的にはtimeline_controllerにて構造体を定義した
  - また、usersモデルからrowsを返してcontoroller側でscanしたがそれがいい使い方なのかは分かっていない
  - 何かベストな方法が知りたかった
  - 意図としては以下
    - モデルには最低限の責務しか負わせたくなかったから
    - なるべくSQLへのアクセスは減らしたい（Joinで一回で取得したい）

- エラーハンドリング
  - 今回はなるべくハンドリングしたがtry catchがない大変さを感じた
  - これも色んな例が出てくるのでスタンダードを知りたい

- リプライ機能
  - どういう構成にするか悩んだ
  - 候補は以下
    - commentテーブル
    - tweetsテーブルに内包
  - 最終的に採用したのはtweetsテーブルに内包
  - 理由としては以下
    - Twitterウェブ版のURIと振る舞いを見てそれに則った
    - commentテーブルにするとリプライへのリプライが難しくなる
