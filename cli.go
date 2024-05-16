package main

import (
	"fmt"
)

type CLI struct {
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
		Random randomCaseCmd `cmd:"" help:"ForMat rANdom CaSE"`
		Snake  snakeCaseCmd  `cmd:"" help:"format_snake_case"`
	} `cmd:"" help:"Format various cases"`
	HtPassWD HtPassWDCmd `cmd:"" name:"htpasswd" help:"Create a htpasswd string"`
	Reverse  reverseCmd  `cmd:"" help:"Reverse the input"`
	SHA      struct {
		SHA256 sha256Cmd `cmd:"" name:"256" help:"Calculate SHA 256"`
		SHA512 sha512Cmd `cmd:"" name:"512" help:"Calculate SHA 512"`
	} `cmd:"" help:"Calculate SHA Hashsums"`
	Version versionCmd `cmd:"" help:"Show version information"`
}

type HtPassWDCmd struct {
	User string `required:"" short:"u"`
	Pass string `required:"" short:"p"`
}

type versionCmd struct {
	Version string
}

func (c *versionCmd) Run() error {
	fmt.Println(Version)
	return nil
}

type countCharsCmd struct {
	Char string `optional:"" short:"c" help:"Count only a specific character"`
}

type countWordsCmd struct {
	Word string `optional:"" short:"w" help:"Count only a specific word"`
}

func (c *countCharsCmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	if c.Char != "" {
		runes := []rune(c.Char)
		if len(runes) > 1 {
			return fmt.Errorf("provide only one character")
		}
		fmt.Println(countChars(in, rune(runes[0])))
	} else {
		fmt.Println(countChars(in))
	}
	return nil
}

func (c *countWordsCmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	if c.Word != "" {
		fmt.Println(countWords(in, c.Word))
	} else {
		fmt.Println(countWords(in))
	}
	return nil
}

func (c *HtPassWDCmd) Run() error {
	entry, err := generateHtpasswdEntry(c.User, c.Pass)
	if err != nil {
		return err
	}
	fmt.Println(entry)
	return nil
}

type encodeBase64Cmd struct{}
type decodeBase64Cmd struct{}

func (c *encodeBase64Cmd) Run() error {
	return func() error {
		in, err := readFromSTDIN()
		if err != nil {
			return err
		}
		fmt.Println(encodeBase64(in))
		return nil
	}()
}
func (c *decodeBase64Cmd) Run() error {
	return func() error {
		in, err := readFromSTDIN()
		if err != nil {
			return err
		}
		decoded, err := decodeBase64(in)
		if err != nil {
			return err
		}
		fmt.Println(decoded)
		return nil
	}()
}

type camelCaseCmd struct{}

func (c *camelCaseCmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	fmt.Println(toCamelCase(in))
	return nil
}

type randomCaseCmd struct{}

func (c *randomCaseCmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	fmt.Println(toRandomCase(in))
	return nil
}

type snakeCaseCmd struct{}

func (c *snakeCaseCmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	fmt.Println(toSnakeCase(in))
	return nil
}

type reverseCmd struct{}

func (c *reverseCmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	fmt.Println(reverseString(in))
	return nil
}

type sha256Cmd struct{}
type sha512Cmd struct{}

func (c *sha256Cmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	fmt.Println(calculateSHA256(in))
	return nil
}
func (c *sha512Cmd) Run() error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	fmt.Println(calculateSHA512(in))
	return nil
}
