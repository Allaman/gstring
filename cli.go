package main

import (
	"fmt"
)

type Globals struct {
	Trim bool `default:"false" short:"t" help:"Trim output (for piping)"`
}

type CLI struct {
	Globals
	Base64 struct {
		Encode encodeBase64Cmd `cmd:"" help:"Encode with base64"`
		Decode decodeBase64Cmd `cmd:"" help:"Decode with base64"`
	} `cmd:"" name:"base64" help:"Base64 Encoding and Decoding"`
	Count struct {
		Chars countCharsCmd `cmd:"" help:"Count characters"`
		Words countWordsCmd `cmd:"" help:"Count words"`
	} `cmd:"" help:"Count various things"`
	Case struct {
		Camel  camelCaseCmd  `cmd:"" help:"formatCamelCase"`
		Lower  lowerCaseCmd  `cmd:"" help:"format lower"`
		Random randomCaseCmd `cmd:"" help:"ForMat rANdom CaSE"`
		Snake  snakeCaseCmd  `cmd:"" help:"format_snake_case"`
		Upper  upperCaseCmd  `cmd:"" help:"FORMAT UPPER"`
	} `cmd:"" help:"Format various cases"`
	Hash struct {
		SHA256 sha256Cmd `cmd:"" name:"sha256" help:"Calculate Sha256"`
		SHA512 sha512Cmd `cmd:"" name:"sha512" help:"Calculate Sha512"`
		MD5    md5Cmd    `cmd:"" name:"md5" help:"Calculate MD5"`
	} `cmd:"" help:"Calculate SHA Hashsums"`
	HtPassWD HtPassWDCmd `cmd:"" name:"htpasswd" help:"Create a htpasswd string"`
	Reverse  reverseCmd  `cmd:"" help:"Reverse the input"`
	Version  versionCmd  `cmd:"" help:"Show version information"`
}

type HtPassWDCmd struct {
	User string `required:"" short:"u"`
	Pass string `required:"" short:"p"`
}

type versionCmd struct{}

func (c *versionCmd) Run() error {
	printOutput(version, false)
	return nil
}

type countCharsCmd struct {
	Char string `optional:"" short:"c" help:"Count only a specific character"`
}

type countWordsCmd struct {
	Word string `optional:"" short:"w" help:"Count only a specific word"`
}

func (c *countCharsCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	if c.Char != "" {
		runes := []rune(c.Char)
		if len(runes) > 1 {
			return fmt.Errorf("provide only one character")
		}
		printOutput((countChars(in, rune(runes[0]))), globals.Trim)
	} else {
		printOutput((countChars(in)), globals.Trim)
	}
	return nil
}

func (c *countWordsCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	if c.Word != "" {
		printOutput(countWords(in, c.Word), globals.Trim)
	} else {
		printOutput(countWords(in), globals.Trim)
	}
	return nil
}

func (c *HtPassWDCmd) Run(globals *Globals) error {
	entry, err := generateHtpasswdEntry(c.User, c.Pass)
	if err != nil {
		return err
	}
	printOutput(entry, globals.Trim)
	return nil
}

type encodeBase64Cmd struct{}
type decodeBase64Cmd struct{}

func (c *encodeBase64Cmd) Run(globals *Globals) error {
	return func() error {
		in, err := readFromSTDIN()
		if err != nil {
			return err
		}
		printOutput(encodeBase64(in), globals.Trim)
		return nil
	}()
}
func (c *decodeBase64Cmd) Run(globals *Globals) error {
	return func() error {
		in, err := readFromSTDIN()
		if err != nil {
			return err
		}
		decoded, err := decodeBase64(in)
		if err != nil {
			return err
		}
		printOutput(decoded, globals.Trim)
		return nil
	}()
}

type camelCaseCmd struct{}

func (c *camelCaseCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(toCamelCase(in), globals.Trim)
	return nil
}

type randomCaseCmd struct{}

func (c *randomCaseCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(toRandomCase(in), globals.Trim)
	return nil
}

type snakeCaseCmd struct{}

func (c *snakeCaseCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(toSnakeCase(in), globals.Trim)
	return nil
}

type upperCaseCmd struct{}

func (c *upperCaseCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(toUpper(in), globals.Trim)
	return nil
}

type lowerCaseCmd struct{}

func (c *lowerCaseCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(toLower(in), globals.Trim)
	return nil
}

type reverseCmd struct{}

func (c *reverseCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(reverseString(in), globals.Trim)
	return nil
}

type sha256Cmd struct{}
type sha512Cmd struct{}

func (c *sha256Cmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(calculateSHA256(in), globals.Trim)
	return nil
}

func (c *sha512Cmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(calculateSHA512(in), globals.Trim)
	return nil
}

type md5Cmd struct{}

func (c *md5Cmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(calculateMD5(in), globals.Trim)
	return nil
}
