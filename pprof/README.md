# net/http/pprof Demo

`net/http/pprof` パッケージの動作を確認するためのサンプルプログラム。

## 起動方法

```bash
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

## プロファイリング用エンドポイント

### 負荷を生成するエンドポイント

- `GET /cpu` - CPU集約的な処理
- `GET /memory` - メモリ割り当て（100MB）
- `GET /goroutine` - 100個のgoroutineを生成
- `GET /block` - ブロッキング操作

### プロファイリングエンドポイント（net/http/pprof提供）

- `GET /debug/pprof/` - 利用可能なプロファイルの一覧
- `GET /debug/pprof/profile?seconds=30` - 30秒間のCPUプロファイル
- `GET /debug/pprof/heap` - ヒープメモリプロファイル
- `GET /debug/pprof/goroutine` - goroutineスタックトレース
- `GET /debug/pprof/block` - ブロックプロファイル
- `GET /debug/pprof/mutex` - ミューテックスプロファイル
- `GET /debug/pprof/allocs` - メモリアロケーションプロファイル

## 使用例

### 1. CPUプロファイルの取得と解析

```bash
# ターミナル1: サーバー起動
go run main.go

# ターミナル2: 負荷をかける
for i in {1..10}; do curl http://localhost:8080/cpu & done

# ターミナル3: CPUプロファイルを取得して解析
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
# インタラクティブモードで top, list, web などのコマンドを実行可能
```

### 2. ヒープメモリプロファイル

```bash
# メモリを割り当てる
curl http://localhost:8080/memory
curl http://localhost:8080/memory
curl http://localhost:8080/memory

# ヒーププロファイルを解析
go tool pprof http://localhost:8080/debug/pprof/heap
```

### 3. Goroutineプロファイル

```bash
# Goroutineを生成
curl http://localhost:8080/goroutine

# テキスト形式で確認
curl http://localhost:8080/debug/pprof/goroutine?debug=1

# pprofツールで解析
go tool pprof http://localhost:8080/debug/pprof/goroutine
```

### 4. ブラウザでの確認

```bash
# ブラウザで http://localhost:8080/debug/pprof/ にアクセス
# 各プロファイルのリンクをクリックして確認可能
```

### 5. プロファイルの可視化

```bash
# CPUプロファイルをグラフ化（graphvizが必要）
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/profile?seconds=30
# ブラウザで http://localhost:8081 にアクセスしてフレームグラフを表示
```

## pprofコマンド（インタラクティブモード）

プロファイルを取得後、以下のコマンドが使用可能：

- `top` - リソース使用量の多い関数トップ10
- `top -cum` - 累積リソース使用量でソート
- `list <function>` - 指定関数のソースコード表示
- `web` - グラフをブラウザで表示（graphviz必要）
- `pdf` - PDFファイルに出力
- `help` - ヘルプ表示

## 注意事項

- 本番環境では `/debug/pprof/` を公開しないこと（認証・アクセス制限推奨）
- メモリプロファイリングは `-memprofile` フラグでも可能
- ブロック・ミューテックスプロファイルは `runtime.SetBlockProfileRate()` や `runtime.SetMutexProfileFraction()` で有効化が必要な場合がある
