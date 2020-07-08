# CodeRemote

This is CLI tool for opening remote directories with Vscode: remote-ssh plugin.

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
