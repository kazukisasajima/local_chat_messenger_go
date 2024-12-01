# local_chat_messenger_go

## 概要

Goのソケット通信とfakerパッケージを使用して、クライアントサーバ間で情報をやりとりするシンプルなアプリケーションです。このプロジェクトはコンピュータサイエンス学習サービス[Recursion](https://recursion.example.com)の課題でpythonで作成したものをGoでも作成してみました。

## インストール

- スクリプトを実行するにはGoのfakerパッケージをインストールする必要があります。

```sh
go get github.com/bxcodec/faker/v3
```


## デモ

### 1. ターミナルを2つ開きます。
- 1つ目のターミナルでサーバーを起動します
```sh
go run main.go server
```
- 2つ目のターミナルでクライアントを起動します
```sh
go run main.go client
```
  
### 2. クライアントでメッセージを入力:
- クライアント側のターミナルで以下のメッセージを入力し、Enterを押します。
  - "name": サーバーがランダムな名前を生成してクライアントに送信します。
  - "email": サーバーがランダムなメールアドレスを生成してクライアントに送信します。
  - その他の文字列: "Unknown request"というメッセージが返されます。
  - "exit": クライアント側の接続を終了します。

<img width="1200" alt="local_chat_messenger_go" src="https://github.com/user-attachments/assets/149c5a06-9235-42a9-8923-13674bb333c0">
<img width="1200" alt="local_chat_messenger_go3" src="https://github.com/user-attachments/assets/67ce97c8-2696-467f-a086-264caa30b39b">

### 注意
- 名前付きパイプ（UNIXソケット）を使用しているため、Windowsの標準的なコマンドプロンプトやPowerShellでは動作しません。
- このアプリケーションはLinuxまたはWSL（Windows Subsystem for Linux）環境で動作します。
