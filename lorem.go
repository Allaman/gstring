package main

import (
	"strings"
)

var loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ultrices sagittis orci a scelerisque purus. At volutpat diam ut venenatis tellus in. Mattis rhoncus urna neque viverra justo nec ultrices dui sapien. Tellus orci ac auctor augue mauris. Eu scelerisque felis imperdiet proin fermentum. Tortor id aliquet lectus proin nibh nisl condimentum id. Id donec ultrices tincidunt arcu non sodales. Ultrices dui sapien eget mi proin. Bibendum neque egestas congue quisque egestas diam. Sem fringilla ut morbi tincidunt augue interdum. Vel risus commodo viverra maecenas accumsan lacus vel facilisis volutpat. Morbi blandit cursus risus at ultrices mi tempus. Adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus. Sapien pellentesque habitant morbi tristique senectus et netus et malesuada.`

var loremIpsumWords = strings.Fields(loremIpsum)
var loremIpsumSentences = strings.Split(loremIpsum, ".")

func generateLoremIpsum(lengthType string, count int) string {
	switch lengthType {
	case "words":
		return generateWords(count)
	case "sentences":
		return generateSentences(count)
	case "paragraphs":
		return generateParagraphs(count)
	default:
		return "Invalid length type. Please specify 'words', 'sentences', or 'paragraphs'."
	}
}

func generateWords(count int) string {
	var words strings.Builder
	for i := range count {
		index := i % len(loremIpsumWords)
		if i > 0 {
			words.WriteString(" ")
		}
		words.WriteString(loremIpsumWords[index])
	}
	return words.String()
}

func generateSentences(count int) string {
	var sentences strings.Builder
	for i := range count {
		index := i % len(loremIpsumSentences)
		sentences.WriteString(loremIpsumSentences[index])
		sentences.WriteString(".")
	}
	return sentences.String()
}

func generateParagraphs(count int) string {
	paragraphs := []string{}
	for range count {
		paragraphs = append(paragraphs, loremIpsum)
	}
	return strings.Join(paragraphs, "\n\n")
}
