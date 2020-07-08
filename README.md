# CodeRemote

This is CLI tool for opening remote directories with Vscode: remote-ssh plugin.

# Install

## Mac OS

```
VERSION=v0.1.0
curl -L "https://github.com/Raita876/coderemote/releases/download/${VERSION}/release-mac64.zip" -o ./release-mac64.zip
unzip ./release-mac64.zip
chmod +rx ./coderemote
mv ./coderemote /usr/local/bin/
```

## Linux

```
VERSION=v0.1.0
curl -L "https://github.com/Raita876/coderemote/releases/download/${VERSION}/release-lin64.zip" -o ./release-lin64.zip
unzip ./release-lin64.zip
chmod +rx ./coderemote
mv ./coderemote /usr/local/bin/
```

# How to use

```
$ coderemote --workdir /home/vagrant/Workdir --remote-host ubuntu repository
code --folder-uri vscode-remote://ssh-remote+ubuntu/home/vagrant/Workdir/repository

# use Environment
$ export CODE_WORKDIR=/home/vagrant/Workdir
$ export CODE_HOST=ubuntu
$ coderemote repository
code --folder-uri vscode-remote://ssh-remote+ubuntu/home/vagrant/Workdir/repository
```
