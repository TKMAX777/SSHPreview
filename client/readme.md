# SSH Preview (Client)
## 概要
このプログラムはSSHの接続先に配置するプログラムです。このプログラムは指定したプレビューしたいファイルを指定したポートにHTTPで飛ばします。

## 目次
<!-- TOC -->

- [SSH Preview (Client)](#ssh-preview-client)
    - [概要](#概要)
    - [目次](#目次)
    - [Install](#install)
    - [設定方法](#設定方法)
        - [UnixDomainSocketを使う場合](#unixdomainsocketを使う場合)
        - [ポート転送を用いる場合](#ポート転送を用いる場合)
    - [利用方法](#利用方法)

<!-- /TOC -->

## Install

```sh
go install github.com/TKMAX777/SSHPreview/client/cmd/open@latest
```

## 設定方法
### UnixDomainSocketを使う場合
接続元で指定したsocketディレクトリを次のように指定します。

```sh
export PreviewListenSock=/path/to/dir
```

### ポート転送を用いる場合
接続元で指定した、接続先の転送ポートを次のように環境変数で指定します。

```sh
export PreviewListenPort=ポート番号
```

## 利用方法

```sh
open /path/to/pictures
```