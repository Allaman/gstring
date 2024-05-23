package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Globals struct {
	Trim bool `default:"false" short:"t" help:"Trim output (for piping)"`
}

type CLI struct {
	Globals
	Base64 struct {
		Encode encodeBase64Cmd `cmd:"" help:"Encodes with base64"`
		Decode decodeBase64Cmd `cmd:"" help:"Decodes with base64"`
	} `cmd:"" name:"base64" help:"Base64 Encoding and Decoding"`
	Count struct {
		Bytes countBytesCmd `cmd:"" help:"Counts bytes"`
		Chars countCharsCmd `cmd:"" help:"Counts characters"`
		Words countWordsCmd `cmd:"" help:"Counts words"`
	} `cmd:"" help:"Count various things"`
	Case struct {
		Camel  camelCaseCmd  `cmd:"" help:"formatsCamelCase"`
		Lower  lowerCaseCmd  `cmd:"" help:"formats lower"`
		Random randomCaseCmd `cmd:"" help:"ForMats rANdom CaSE"`
		Snake  snakeCaseCmd  `cmd:"" help:"formats_snake_case"`
		Upper  upperCaseCmd  `cmd:"" help:"FORMATS UPPER"`
	} `cmd:"" help:"Format various cases"`
	Hash struct {
		SHA256 sha256Cmd `cmd:"" name:"sha256" help:"Calculates Sha256"`
		SHA512 sha512Cmd `cmd:"" name:"sha512" help:"Calculates Sha512"`
		MD5    md5Cmd    `cmd:"" name:"md5" help:"Calculates MD5"`
	} `cmd:"" help:"Calculate SHA Hashsums"`
	Head headCmd `cmd:"" help:"Returns the first n lines"`
	Hex  struct {
		ToHex   toHexCmd   `cmd:"" name:"to" help:"Converts to hexadecimal"`
		FromHex fromHexCmd `cmd:"" name:"from" help:"Converts hexadecimal back"`
	} `cmd:"" name:"hex" help:"Hexadeicmal converting"`
	Lorem struct {
		Bytes      loremBytesCmd      `cmd:"" help:"Returns Bytes"`
		Paragraphs loremParagraphsCmd `cmd:"" help:"Returns Paragraphs"`
		Sentences  loremSentencesCmd  `cmd:"" help:"Returns Sentences"`
		Words      loremWordsCmd      `cmd:"" help:"Returns Words"`
	} `cmd:"" help:"Generate Lorem Ipsum"`
	HtPassWD         HtPassWDCmd         `cmd:"" name:"htpasswd" help:"Create a htpasswd string"`
	Permissions      permissionsCmd      `cmd:"" help:"Parses Unix permissions"`
	PwGen            pwGenCmd            `cmd:"" name:"pwgen" help:"Password generator (letters and numbers)"`
	RemoveWhitespace removeWhitespaceCmd `cmd:"" help:"Removes whitespace"`
	Reverse          reverseCmd          `cmd:"" help:"Reverses the input"`
	Rng              rngCmd              `cmd:"" help:"Random number generator"`
	Sort             sortCmd             `cmd:"" help:"Sorts the input by line"`
	Split            splitCmd            `cmd:"" help:"Splits a string"`
	Tail             tailCmd             `cmd:"" help:"Returns the last n lines"`
	Time             struct {
		FromUnixTime fromUnixTimeCmd `cmd:"" name:"from-unix" help:"Converts from Unix time to normal time"`
		GetUnixTime  getUnixTimeCmd  `cmd:"" name:"get-unix" help:"Returns Unix time"`
	} `cmd:"" help:"Time conversions"`
	URL struct {
		Encode encodeURLCmd `cmd:"" help:"Encodes string to valid URL"`
		Decode decodeURLCmd `cmd:"" help:"Decodes URL to string"`
	} `cmd:"" help:"URL Encoding and Decoding"`
	Version versionCmd `cmd:"" help:"Shows version information"`
}

type HtPassWDCmd struct {
	User string `required:"" short:"u" help:"The username"`
	Pass string `required:"" short:"p" help:"The password (pay attention to shell history)"`
}

type versionCmd struct{}

func (c *versionCmd) Run() error {
	printOutput(version, false)
	return nil
}

type countBytesCmd struct {
	Format bool `optional:"" default:"false" short:"f" help:"Human readable format"`
}

type countCharsCmd struct {
	Char string `optional:"" short:"c" help:"Count only a specific character"`
}

type countWordsCmd struct {
	Word string `optional:"" short:"w" help:"Count only a specific word"`
}

func (c *countBytesCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(countBytes(in, c.Format), globals.Trim)
	return nil
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
		printOutput(countChars(in, rune(runes[0])), globals.Trim)
	} else {
		printOutput(countChars(in), globals.Trim)
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
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(encodeBase64(in), globals.Trim)
	return nil
}
func (c *decodeBase64Cmd) Run(globals *Globals) error {
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

type toHexCmd struct {
	Formated bool `short:"f" default:"false" help:"Format hexadecimal result with whitespaces"`
}

func (c *toHexCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(stringToHex(in, c.Formated), globals.Trim)
	return nil
}

type fromHexCmd struct{}

func (c *fromHexCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	out, err := hexToString(in)
	if err != nil {
		return err
	}
	printOutput(out, globals.Trim)
	return nil
}

type encodeURLCmd struct{}
type decodeURLCmd struct{}

func (c *encodeURLCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(encodeURL(in), globals.Trim)
	return nil
}
func (c *decodeURLCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	decoded, err := decodeURL(in)
	if err != nil {
		return err
	}
	printOutput(decoded, globals.Trim)
	return nil
}

type loremBytesCmd struct {
	Count int `default:"5" short:"c" help:"Number of bytes"`
}
type loremWordsCmd struct {
	Count int `default:"5" short:"c" help:"Number of words"`
}
type loremSentencesCmd struct {
	Count int `default:"1" short:"c" help:"Number of sentences"`
}
type loremParagraphsCmd struct {
	Count int `default:"1" short:"c" help:"Number of Paragraphs"`
}

func (c *loremBytesCmd) Run(globals *Globals) error {
	lorem := generateLoremIpsum("bytes", c.Count)
	printOutput(lorem, globals.Trim)
	return nil
}

func (c *loremWordsCmd) Run(globals *Globals) error {
	lorem := generateLoremIpsum("words", c.Count)
	printOutput(lorem, globals.Trim)
	return nil
}

func (c *loremSentencesCmd) Run(globals *Globals) error {
	lorem := generateLoremIpsum("sentences", c.Count)
	printOutput(lorem, globals.Trim)
	return nil
}

func (c *loremParagraphsCmd) Run(globals *Globals) error {
	lorem := generateLoremIpsum("paragraphs", c.Count)
	printOutput(lorem, globals.Trim)
	return nil
}

type splitCmd struct {
	Sep string `default:" " short:"s" help:"Separator to split"`
}

func (c *splitCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(formatSplittedString(splitString(in, c.Sep)), globals.Trim)
	return nil
}

type headCmd struct {
	Num int `default:"1" short:"n" help:"Number of lines"`
}

func (c *headCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(head(in, c.Num), globals.Trim)
	return nil
}

type tailCmd struct {
	Num int `default:"1" short:"n" help:"Number of lines"`
}

func (c *tailCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(tail(in, c.Num), globals.Trim)
	return nil
}

type rngCmd struct {
	Start     int    `default:"0" short:"s" help:"Min RNG"`
	End       int    `default:"1" short:"e" help:"Max RNG"`
	Count     int    `default:"1" short:"c" help:"Number of generated RNGs"`
	Delimiter string `default:" " short:"d" help:"Output delimiter"`
}

func (c *rngCmd) Run(globals *Globals) error {
	rngs, err := pseudoRandomGenerator(c.Start, c.End, c.Count, c.Delimiter)
	if err != nil {
		return err
	}
	printOutput(rngs, globals.Trim)
	return nil
}

type pwGenCmd struct {
	Length    int    `default:"10" short:"n" help:"Length of password"`
	Count     int    `default:"1" short:"c" help:"Number of passwords to generate"`
	Delimiter string `default:"\n" short:"d" help:"Output delimiter"`
}

func (c *pwGenCmd) Run(globals *Globals) error {
	passwords := strings.Join(generateRandomPasswords(c.Length, c.Count), c.Delimiter)
	printOutput(passwords, globals.Trim)
	return nil
}

type sortCmd struct {
	Desc             bool `default:"false" short:"d" help:"Descending"`
	IgnoreEmptyLines bool `default:"true" help:"Ignore empty lines"`
	Unique           bool `default:"false" short:"u" help:"Only unique lines"`
}

func (c *sortCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	printOutput(sortLines(in, c.Desc, c.IgnoreEmptyLines, c.Unique), globals.Trim)
	return nil
}

type fromUnixTimeCmd struct {
	Format string `default:"2006-01-02 15:04:05.000000000" short:"f" help:"Time format as Go reference time"`
}

func (c *fromUnixTimeCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	timestamp, err := strconv.ParseInt(strings.TrimSpace(in), 10, 64)
	if err != nil {
		return err
	}
	t, err := convertUnixTimestamp(timestamp, c.Format)
	if err != nil {
		return err
	}
	printOutput(t, globals.Trim)
	return nil
}

type getUnixTimeCmd struct {
}

func (c *getUnixTimeCmd) Run(globals *Globals) error {
	t := unixTimestamp()
	printOutput(t, globals.Trim)
	return nil
}

type removeWhitespaceCmd struct {
	Spaces     bool `default:"false" short:"s" help:"Remove spaces"`
	Tabs       bool `default:"false" short:"b" help:"Remove tabs"`
	CR         bool `default:"false" short:"c" help:"Remove CRs"`
	LE         bool `default:"false" short:"l" help:"Remove line endings"`
	EmptyLines bool `default:"false" short:"e" help:"Remove empty lines"`
}

func (c *removeWhitespaceCmd) Run(globals *Globals) error {
	in, err := readFromSTDIN()
	if err != nil {
		return err
	}
	t := removeWhitespace(in, c.Spaces, c.Tabs, c.CR, c.LE, c.EmptyLines)
	printOutput(t, globals.Trim)
	return nil
}

type permissionsCmd struct {
	Permission string `required:"" short:"p" help:"Permission in octal format (both 3 and 4 digits)"`
}

func (c *permissionsCmd) Run(globals *Globals) error {
	t, err := parseUnixPermissions(c.Permission)
	if err != nil {
		return err
	}
	printOutput(t, globals.Trim)
	return nil
}
