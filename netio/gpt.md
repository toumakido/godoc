3. 並行ダウンローダー (Goroutine & HTTP Client)
Go の net/http と sync.WaitGroup を使い、複数の URL から並行してファイルをダウンロードするプログラムを作成します。

要件:

複数の URL を受け取り、並行処理 でダウンロード
各ファイルのサイズを出力
sync.WaitGroup を使ってすべてのダウンロードが完了するまで待つ
io.Copy を使ってデータを /tmp/ に保存