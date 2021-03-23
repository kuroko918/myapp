# 本番環境 URL
## Web
- https://nuxt-dtlmvofcca-an.a.run.app

## Mobile

# 使用技術
### サーバーサイド
- cmd/app ディレクトリ
  - go
  - gin（webフレームワーク）
  - gorm（ORマッパー）
- cmd/grpc-app ディレクトリ
  - go
  - gRPC（grpc-gateway）
  - Cloud Firestore
  - [本番環境]Google App Engine

### フロントエンド
- Nuxt.js（SSR）
- TypeScript（Vuexの型付け：vuex-smart-module）
- Jest
- [本番環境]Cloud Run

### ネイティブアプリ
- Flutter(flutter_hooks)
  - 状態管理: Riverpod(hooks_riverpod)
  - 状態変更通知: StateNotifier
  - HTTP通信: Retrofit(Dio)
  - Json解析： freezed

### 共通
- 認証系：firebase Authorization

# アーキテクチャ
### サーバーサイド
  - Clean Architecture

# 環境構築

### サーバーサイド
```bash
$ cd cmd/app

# serve at localhost:50051
$ go run main.go

# reverse proxy serve at localhost:8080
$ cd reverse-proxy & go run main.go

# How to deploy
## Create mod vendor
$ go mod vendor
## Deploy gRPC server to GAE
$ cd ~/src/github.com/kuroko918/myapp/cmd/grpc-app
$ gcloud app deploy
```

- gRPC server 及び gRPC reverse proxy server の生成
```bash
$ protoc \
    -I$GOPATH/src/github.com/kuroko918/myapp/cmd/grpc-app/proto \
    -I$GOPATH/src/github.com/kuroko918/myapp/pkg/envoyproxy/protoc-gen-validate \
    -I$GOPATH/src/github.com/kuroko918/myapp/pkg/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out==plugins=grpc:$GOPATH/src \
    --go-grpc_out==plugins=grpc:$GOPATH/src \
    --validate_out=lang=go:$GOPATH/src \
    --grpc-gateway_out=logtostderr=true:$GOPATH/src \
    $GOPATH/src/github.com/kuroko918/myapp/cmd/grpc-app/proto/*.proto
```

### フロントエンド
```bash
$ cd web

# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev

# build for production and launch server
$ yarn build
$ yarn start

# How to deploy
### Remove .env in web/.gitignore file

### Change URL ENV in .env into 'https://myapp-kuroko918.an.r.appspot.com'

### Create docker image & Upload to Container Registry
$ gcloud builds submit --tag asia.gcr.io/myapp-kuroko918/nuxt:<tag> .

### Deploy into Cloud Run
$ gcloud run deploy --image=asia.gcr.io/myapp-kuroko918/nuxt:<tag> --platform managed --port 3000 --region asia-northeast1
```

### ネイティブアプリ
```bash
$ cd native

# open server
$ flutter run
```
