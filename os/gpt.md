# OS package
## 課題
### chatGPTに質問
goの練習用にosパッケージを使って、100行程度で実装できる課題を出して

### 回答
課題: 簡易ファイルマネージャー
概要
コマンドラインで動作する簡易的なファイルマネージャーを作成してください。このプログラムは以下の機能を提供します：

指定されたディレクトリ内のファイル・フォルダ一覧を表示
ファイルを作成
ファイルの内容を表示
環境変数を利用してデフォルトの作業ディレクトリを指定可能
要件
ディレクトリの一覧表示:
指定されたディレクトリ内の全てのファイルとフォルダをリストアップする。
ディレクトリはos.Argsで渡すか、デフォルトの環境変数DEFAULT_DIRから取得する。
ファイル作成:
指定した名前のファイルを作成する。
ファイル名はコマンドライン引数で指定する。
ファイル内容の表示:
指定したファイルの内容を表示する。
エラーハンドリング:
入力が不正な場合やファイル・ディレクトリが存在しない場合に適切なエラーメッセージを表示する。

