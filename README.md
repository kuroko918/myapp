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

### フロントエンド
- Nuxt.js（SSR）
- TypeScript（Vuexの型付け：vuex-smart-module）
- Jest
- Cloud Run(動的コンテンツの配信用)
- firebase Hosting(静的コンテンツの配信用)

### ネイティブアプリ
- Flutter

### 共通
- 認証系： firebase Authorization

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

# generate static project
$ yarn generate
```

### ネイティブアプリ
```bash
$ cd native

# open server
$ flutter run
```
