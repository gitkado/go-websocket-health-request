# go-websocket-health-request
goでwebsocketサーバへconnectionしてjsonrpcプロトコルにそってhealthリクエストを投げ続ける  
``gotest.go``でリクエスト終了時刻やリクエスト間隔を設定する

```sh
# 実行コマンド
docker run -it --rm -v $PWD:/opt golang:websocket-client go run gotest.go
```