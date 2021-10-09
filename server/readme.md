# SSH Preview (Server)
## 概要
このプログラムはSSHの接続元に配置するプログラムです。このプログラムはClientからのリクエストを受けとり、Chromeを立ち上げ、プレビューを転送します。

## 目次
<!-- TOC -->

- [SSH Preview (Server)](#ssh-preview-server)
    - [概要](#概要)
    - [目次](#目次)
    - [Install](#install)
    - [設定方法](#設定方法)
        - [Unix Domain Socketを用いる場合](#unix-domain-socketを用いる場合)
        - [ポート転送を行う場合](#ポート転送を行う場合)

<!-- /TOC -->

## Install

```sh
git clone github.com/TKMAX777/SSHPreview
cd SSHPreview/server
```


## 設定方法 
接続先で用いるソケットとして、通常のポート転送に加え、unix domain socketを用いることができます。

### Unix Domain Socketを用いる場合
接続元の任意のリッスンポートを次のように環境変数で指定します。

```sh
export PreviewListenPort=ポート番号
```

SSH時にポートをソケットに転送するように設定します。

```config
# ~/.ssh/config
Host 設定したいホスト
HostName 設定したいホストアドレス
RemoteForward /path/to/socket_dir 127.0.0.1:先程のポート番号
```

`/path/to/socket_dir` には、接続先の適当なディレクトリを指定してください。

この後、常態化してください。


### ポート転送を行う場合
接続元の任意のリッスンポートを次のように環境変数で指定します。

```sh
export PreviewListenPort=ポート番号
```

SSH時にポートを接続先のポートへと転送するように設定します。

```config
# ~/.ssh/config
Host 設定したいホスト
HostName 設定したいホストアドレス
RemoteForward :転送先のポート 127.0.0.1:先程のポート番号
```

この後、常態化してください。



