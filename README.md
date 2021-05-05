# CodeChallenge3 - Search

CLI search written in GO.  Currently only reads a remote URL and prints to CLI.

## Usage

```
go run challenge3_search.go -source=https://appdev-code-challenge.s3.amazonaws.com/challenge3_search/SW_EpisodeIV.txt

go run challenge3_search.go -h
```

## Things I Learned

- `flags.String(name, default, help msg)`, usefulf or CLI parameters, good [link](https://gobyexample.com/command-line-flags)
- `go mod init (package name)`, 
  - for this project `go mod init github.com/yamori/challenge3_search`
  - where `challenge3_search` is the name of the entry file, convention
- `go install`
  - executed locally will install to `$GOPATH/bin/`, and if this is on `PATH` it can then be used immediately
- VSCode - go extensions, auto complete etc.