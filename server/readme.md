# SSH Preview (Client)
## 概要
このプログラムはSSHの接続元に配置するプログラムです。このプログラムはClientからのリクエストを受けとり、Chromeを立ち上げ、プレビューを転送します。

## 目次
<!-- TOC -->

- [SSH Preview (Client)](#ssh-preview-client)
    - [概要](#概要)
    - [目次](#目次)
    - [設定方法](#設定方法)

<!-- /TOC -->

## 設定方法

1. ビルドします

```sh
go build
```

2. 接続元の任意のリッスンポートを次のように環境変数で指定します。

```sh
export PreviewListenPort=ポート番号
```

3. SSH時にポートを転送するように設定します。

```config
# ~/.ssh/config
Host 設定したいホスト
HostName 設定したいホスト
RemoteForward :接続先での転送ポート 127.0.0.1:先程のポート番号
```

4. 何かしらの手段で常態化してください。