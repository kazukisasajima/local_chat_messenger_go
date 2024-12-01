# local_chat_messenger_go

## 概要

Goのソケット通信とfakerパッケージを使用して、クライアントサーバ間で情報をやりとりするシンプルなアプリケーションです。このプロジェクトはコンピュータサイエンス学習サービス[Recursion](https://recursion.example.com)の課題でpythonで作成したものをGoでも作成してみました。

## インストール

- スクリプトを実行するにはGoのfakerパッケージをインストールする必要があります。

```sh
go get github.com/bxcodec/faker/v3
```


## 実行方法

- 以下のコマンドを使用して、実行する。

```sh
go run main.go server
```
```sh
go run main.go client
```
クライアント側のターミナルで任意の値を入力

### 注意
- 名前付きパイプ（UNIXソケット）を使用しているため、Windowsの標準的なコマンドプロンプトやPowerShellでは動作しません。
- このアプリケーションはLinuxまたはWSL（Windows Subsystem for Linux）環境で動作します。


## デモ
