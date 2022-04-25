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
git clone git@github.com:TKMAX777/SSHPreview.git
cd SSHPreview/server
go build
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
RemoteForward /path/to/socket/dir/http.sock 127.0.0.1:先程のポート番号
StreamLocalBindUnlink yes
```

`/path/to/socket/dir` には、接続先の適当なディレクトリを指定してください。

この後、常態化してください。

#### Hint
Unix Domain Socketのソケットファイルの権限を絞ることで、接続ユーザを絞ることができます。
例えば、接続元を自分のみに制限する場合は、次のように実行することで生成されるソケットファイルの権限を変えられます。

```sh
chmod 700 /path/to/socket/dir
```

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

## <参考> 常態化方法について

### Windows

タスクスケジューラを用いると簡単にこれを行うことができます。

1. バイナリファイルと同じディレクトリに次のVBSファイルを作成します。

```vbs
CreateObject("WScript.Shell").Run "C:\path\to\SSHPreview\server\server.exe", 0
```

2. タスクスケジューラで適当な新たなタスクの作成をし、起動時に実行するようにします。

3. 実行時の開始(オプション)に、 `C:\path\to\SSHPreview\server` と記述します。

### macOS

1. 次の設定ファイルを `/Library/LaunchAgents/ssh_preview.service.plist` に保存します。

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>ssh_preview.service.plist</string>
 
  <key>ProgramArguments</key>
  <array>
    <string>/path/to/ssh_preview/server/server</string>
  </array>

  <key>WorkingDirectory</key>
  <string>/path/to/ssh_preview/server</string>

  <key>EnvironmentVariables</key>
  <dict>
    <key>PreviewListenPort</key>
    <string> PORT_NUMBER_LIKE_8000 </string>
  </dict>

  <key>RunAtLoad</key>
  <true/>
  
  <key>UserName</key>
  <string>USER_NAME</string>

</dict>
</plist>
```

2. 登録する

```sh
$ launchctl load /Library/LaunchAgents/ssh_preview.service.plist
```

