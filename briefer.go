// The MIT License (MIT)
//
// Copyright (c) 2018 Keigo Motosugi
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/JesusIslam/tldr"
	"github.com/ikawaha/kagome/tokenizer"
)

var (
	langOpt          = flag.String("l", "en", "language setting. support \"en\" and \"ja\". default is \"en\"")
	sentenceCountOpt = flag.Int("c", 3, "display sentence count.")
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("input file isn't specifyed.")
		os.Exit(1)
	}

	inputFile := flag.Args()[0]

	intoSentences := *sentenceCountOpt
	iText, _ := ioutil.ReadFile(inputFile)
	t := tokenizer.New()
	inText := ""
	if *langOpt == "ja" {
		tokens := t.Tokenize(string(iText))
		for _, token := range tokens {
			if token.Class == tokenizer.DUMMY {
				continue
			}
			addText := ""
			if token.Surface == "。" {
				addText = "."
			} else if token.Surface == "？" {
				addText = "?"
			} else {
				addText = token.Surface
			}
			inText = inText + " " + addText
		}
	} else {
		inText = string(iText)
	}
	bag := tldr.New()
	result, _ := bag.Summarize(inText, intoSentences)
	fmt.Println(strings.Replace(result, " ", "", -1))
}
