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
<<<<<<< HEAD
|hits|取得問題数を指定する 1~10|
