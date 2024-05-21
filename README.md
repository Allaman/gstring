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
  base64 encode        Encodes with base64
  base64 decode        Decodes with base64
  count bytes          Counts bytes
  count chars          Counts characters
  count words          Counts words
  case camel           formatsCamelCase
  case lower           formats lower
  case random          ForMats rANdom CaSE
  case snake           formats_snake_case
  case upper           FORMATS UPPER
  hash sha256          Calculates Sha256
  hash sha512          Calculates Sha512
  hash md5             Calculates MD5
  head                 Returns the first n lines
  hex to               Converts to hexadecimal
  hex from             Converts hexadecimal back
  lorem bytes          Returns Bytes
  lorem paragraphs     Returns Paragraphs
  lorem sentences      Returns Sentences
  lorem words          Returns Words
  htpasswd             Create a htpasswd string
  pwgen                Password generator (letters and numbers)
  remove-whitespace    Removes whitespace
  reverse              Reverses the input
  rng                  Random number generator
  sort                 Sorts the input by line
  split                Splits a string
  tail                 Returns the last n lines
  time from-unix       Converts from Unix time to normal time
  time get-unix        Returns Unix time
  url encode           Encodes string to valid URL
  url decode           Decodes URL to string
  version              Shows version information

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
date +%s | gstring time from-unix -f "2006-01-02 15:04"
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
