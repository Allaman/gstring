<h1 align="center">Gstring</h1>

<div align="center">
    <img alt="GoReleaser Logo" src="https://github.com/Allaman/gstring/assets/12184268/b302769b-4cfe-4ef4-83c4-f01c9f505bb9?v=3&s=200" height="140" />
  <p>
    <img src="https://github.com/Allaman/gstring/actions/workflows/release.yaml/badge.svg" alt="Release"/>
    <img src="https://img.shields.io/github/repo-size/Allaman/gstring" alt="size"/>
    <img src="https://img.shields.io/github/issues/Allaman/gstring" alt="issues"/>
    <img src="https://img.shields.io/github/last-commit/Allaman/gstring" alt="last commit"/>
    <img src="https://img.shields.io/github/license/Allaman/gstring" alt="license"/>
    <img src="https://img.shields.io/github/v/release/Allaman/gstring?sort=semver" alt="last release"/>
  </p>
  <p>
    Swiss army knife for working with strings
  </p>
</div>

## Use cases

```
❯ gstring --help
Usage: gstring <command> [flags]

A tool to work with strings

Flags:
  -h, --help    Show context-sensitive help.
  -t, --trim    Trim output (for piping)

Commands:
  base64 encode       Encode with base64
  base64 decode       Decode with base64
  count bytes         Count bytes
  count chars         Count characters
  count words         Count words
  case camel          formatCamelCase
  case lower          format lower
  case random         ForMat rANdom CaSE
  case snake          format_snake_case
  case upper          FORMAT UPPER
  hash sha256         Calculate Sha256
  hash sha512         Calculate Sha512
  hash md5            Calculate MD5
  head                Returns the first n lines
  hex to              Convert to hexadecimal
  hex from            Convert hexadecimal back
  lorem bytes         Bytes
  lorem paragraphs    Paragraphs
  lorem sentences     Sentences
  lorem words         Words
  htpasswd            Create a htpasswd string
  pwgen               Password generator (letters and numbers)
  reverse             Reverse the input
  rng                 Random number generator
  sort                Sort the input by line
  split               Split a string
  tail                Returns the last n lines
  url encode          Encode string to valid URL
  url decode          Decode URL to string
  version             Show version information

Run "gstring <command> --help" for more information on a command.
```

## Get gstring

with curl

```sh
VERSION=$(curl -s https://api.github.com/repos/allaman/gstring/releases/latest | grep tag_name | cut -d '"' -f 4)
curl -sLo gstring https://github.com/Allaman/gstring/releases/download/${VERSION}/gstring_${VERSION}_$(uname -s)_$(uname -m)
```

with wget

```sh
VERSION=$(wget -qO - https://api.github.com/repos/allaman/gstring/releases/latest | grep tag_name | cut -d '"' -f 4)
wget -qO gstring https://github.com/Allaman/gstring/releases/download/${VERSION}/gstring_${VERSION}_$(uname -s)_$(uname -m)
```

```sh
chmod +x gstring
./gstring
```

## Usage

Most of the commands work with stdin so you have to pipe content, for example:

```sh
echo "hello world" | gstring case camel
gstring hash sha256 < main.go # equivalent to sha256sum main.go
```

However, some commands read from named arguments, for example:

```sh
gstring htpasswd -u bar -p test
gstring lorem sentences -c 20
gstring rng -s 1 -e 100 -c 10 -d ";"
gstring pwgen -n 20 -c 10 -d ";"
```

## What about the name?

Well, of course, from **G**olang and **string**. What are you thinking?
