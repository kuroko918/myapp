# 使用技術
### サーバーサイド
- go
- gin（webフレームワーク）
- gorm（ORマッパー）

### フロントエンド
- Nuxt.js
- TypeScript（Vuexの型付け：vuex-smart-module）
- Jest

### ネイティブアプリ
- Flutter

# アーキテクチャ
`Clean Architecture`を採用

# 環境構築

### サーバーサイド
```bash
$ cd cmd/app

# build
$ go build

# serve at localhost:8080
$ ./myapp
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
