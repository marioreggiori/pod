# pod

## Install
### Linux
- todo
### Windows
- todo
### Mac
- todo

## Usage

```bash
# simple command
$ pod go mod init
# command with arguments
$ pod node -e "console.log('hello world')"
$ pod npx create-react-app --template typescript
# map port to localhost
$ pod -p 3000 npm start
$ pod -p 8080:80 npm start
# enable verbose mode (prints infos like container pull)
$ pod --verbose node --version
```

### Setup aliases for commands
```bash
# ~/.bashrc
alias node='pod node'
alias npm='pod npm'
...
```