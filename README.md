# janken_go_grpc
goとgRPCの勉強

## メモ
- google/protobuf/timestampがimportできなかった 
  - [このページ](https://qiita.com/revenue-hack/items/7221f8e015d47d894854)を参考にした
  - 結局stringで対応させた
- `go run cmd/api`とするとインターフェイスの一覧のようなものが流れてきてサーバを立ち上げられなかった
  - `go run ./cmd/api`が正しい
- grpcurlをつかってサーバ側の動作確認をした（11/12）
- ほかのプロジェクトでも使えるものを用意しておきたかったのでgprc_cliをwsl2内でビルドして使えるようにした（11/13）
  - `~/grpc/cmake/build/grpc_cli`

※実際の流れ（正しいかはわからない）
```bash
cd /janken_go_grpc
go get github.com/fullstorydev/grpcurl
cd $GOPATH
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

![two-wins-result](img/two-wins-in-a-row.png)
![two-wins-history](img/two-wins-in-a-row_history.png)
## 参考
- [Go言語で簡単なgRPCサーバーを作成](https://dev.classmethod.jp/articles/golang-grpc-sample-project/)
- [gRPCサーバーの動作確認をgrpcurlでやってみた](https://qiita.com/yukina-ge/items/a84693f01f3f0edba482)
- [gRPC command line tool](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)
