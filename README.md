# go-websocket-health-request
goでwebsocketサーバへconnectionしてjsonrpcプロトコルにそってhealthリクエストを投げ続ける
``gotest.go``でリクエスト終了時刻やリクエスト間隔を設定する

```sh
docker run -it --rm -v /vagrant_shared/gotest:/opt golang:websocket-client go run gotest.go
```