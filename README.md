# 楽天QuizAPI

## 概要
楽天APIからランダムに問題のためのデータを排出するAPI
.envファイルに設定したApplicationIDとAffiliateId(任意)からリクエストができる

## Quick start
- 楽天APIからアプリケーションIDを取得する


```bash
#Golang1.18以上の環境がない人
apt install golang
```

```bash
#create directory
mkdir /serve
cd /serve
#リポジトリからコードをクローンしてくる
git clone github.com/GiikuCAMPvol2/QuizAPI.git
#envファイルを作る
nano /.env
# APPLICATION_ID=your_application_id
# AFFILIATE_ID= your_affiliate_id
#を追記する

go build -o serve
./serve
```
## API reference

GET  /quiz
リアルタイムでRakutenから取得する
### パラメータ

|Query|Detail|
|----|----|
|keyword|キーワードで絞って検索する|
|genreid|ジャンルIDで絞って検索する|
|page|楽天検索のページを設定|
|hits|取得件数|
|sort|以下省略|


### sortのメソッド
|query|detail|
|----|----|
|0|楽天標準ソート <- default|
|1|発売日順(降順)|
|2|発売日順(昇順)|
|3|売上順(降順)|
|4|売上順(昇順)|
|5|満足順（降順|
|6|満足順（昇順）|

GET /quizlake
Mongoに貯めているデータを取得する

### パラメータ

|Query|Detail|
|----|----|
|hits|取得件数|

