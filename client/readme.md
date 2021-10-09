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
            - [Sudu権限がある場合](#sudu権限がある場合)
            - [ユーザ権限で行う場合](#ユーザ権限で行う場合)
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
export PreviewListenSock=/path/to/socket/dir
```


#### Sudo権限がある場合
SSH接続時に、前回の接続ソケットを上書きするため、次の設定をsshdにする必要があります。

```sh
sudo sh -c 'echo "StreamLocalBindUnlink yes" >> /etc/ssh/sshd_config'.
```

但し、この設定は他のSSHのUnixソケット転送にも影響を及ぼすことに注意してください。

#### ユーザ権限で行う場合
ソケットファイルの有効状態を監視し、無効な場合削除するデーモンを建てることができます。

```sh
go install github.com/TKMAX777/SSHPreview/client/cmd/PreviewCheckUp@latest
```

このプログラムは指定した時間(初期値では5秒ごと)にソケットファイルの存在を確認し、存在すればサーバにリクエストを行うことで、有効状態を確認するものです。

次の環境変数で確認時間(秒)を変更することができます。

```
export PreviewCheckInterval=5
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
