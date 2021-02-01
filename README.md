# pod
todo: intro

## Usage

```bash
# print available commands
$ pod
# simple command
$ pod go mod init
# command with arguments
$ pod node -e "console.log('hello world')"
$ pod npx create-react-app --template typescript
# map port to localhost
$ pod -p 3000 npm start
$ pod -p 8080:80 npm start
# map volume to container
$ pod -v /some/local/dir:/some/container/dir <command> --arg 42
# enable verbose mode (prints infos like container pull)
$ pod --verbose node --version
```

### Setup aliases for commands
```bash
# ~/.bashrc
alias node='pod node'
alias npm='pod npm'
alias python='pod python'
alias pip='pod pip'
alias python2='pod -t 2.7 python'
alias pip2='pod -t 2.7 pip'
...

# /bin/bash
$ npm init
$ python2 --version
```

## Install
### Linux & Mac
```
sudo curl -L "https://github.com/marioreggiori/pod/releases/latest/download/pod-$(uname -s)-$(uname -m)" -o /usr/local/bin/pod
```
### Windows
todo

## Documentation
Browse all available commands [here](docs/pod.md)