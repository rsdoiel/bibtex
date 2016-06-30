//
// Package bibtex is a quick and dirty plain text parser for generating
// a Bibtex citation
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of mkbib nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package bibtex

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"

	// My packages
	"github.com/rsdoiel/tok"
)

const (
	// Version of BibTeX package
	Version = "0.0.0"

	// A template for printing an element
	ElementTmplSrc = `
@{{- .Type -}}{
    {{-range .Keys}}
	{{ . -}},
	{{end}}
	{{-range $key, $val := .Tags}}
		{{- $key -}} = {{- $val -}},
	{{end}}
}
`
)

// Generic Element
type Element struct {
	XMLName xml.Name          `json:"-"`
	Type    string            `xml:"type" json:"type"`
	Keys    []string          `xml:"keys" json:"keys"`
	Tags    map[string]string `xml:"tags" json:"tags"`
}
type Elements []*Element

type TagTypes struct {
	Required []string
	Optional []string
}

// Entry types
var (
	elementTypes = &map[string]*TagTypes{
		"article": &TagTypes{
			Required: []string{"author", "title", "journal", "year", "volume"},
			Optional: []string{"number", "pages", "month", "note"},
		},
		"book": &TagTypes{
			Required: []string{"author", "editor", "title", "publisher", "year"},
			Optional: []string{"volume", "number", "series", "address", "edition", "month", "note"},
		},
		"booklet": &TagTypes{
			Required: []string{"Title"},
			Optional: []string{"author", "howpublished", "address", "month", "year", "note"},
		},
		"inbook": &TagTypes{
			Required: []string{"author", "editor", "title", "chapter", "pages", "publisher", "year"},
			Optional: []string{"volume", "number", "series", "type", "address", "edition", "month", "note"},
		},
		"incollection": &TagTypes{
			Required: []string{"author", "title", "booktitle", "publisher", "year"},
			Optional: []string{"editor", "volume", "number", "series", "type", "chapter", "pages", "address", "edition", "month", "note"},
		},
		"inproceedings": &TagTypes{
			Required: []string{"author", "title", "booktitle", "year"},
			Optional: []string{"editor", "volume", "number", "series", "pages", "address", "month", "organization", "publisher", "note"},
		},
		"conference": &TagTypes{
			Required: []string{"author", "title", "booktitle", "year"},
			Optional: []string{"editor", "volume", "number", "series", "pages", "address", "month", "organization", "publisher", "note"},
		},
		"manual": &TagTypes{
			Required: []string{"title"},
			Optional: []string{"author", "organization", "address", "edition", "month", "year", "note"},
		},
		"masterthesis": &TagTypes{
			Required: []string{"author", "title", "school", "year"},
			Optional: []string{"type", "address", "month", "note"},
		},
		"misc": &TagTypes{
			Required: []string{},
			Optional: []string{"author", "title", "howpublished", "month", "year", "note"},
		},
		"phdthesis": &TagTypes{
			Required: []string{"author", "title", "school", "year"},
			Optional: []string{"type", "address", "month", "note"},
		},
		"proceedings": &TagTypes{
			Required: []string{"title", "year"},
			Optional: []string{"editor", "volume", "series", "address", "month", "publisher", "organization", "note"},
		},
		"techreport": &TagTypes{
			Required: []string{"author", "title", "institution", "year"},
			Optional: []string{"type", "number", "address", "month", "note"},
		},
		"unpublished": &TagTypes{
			Required: []string{"author", "title", "note"},
			Optional: []string{"month", "year"},
		},
	}
)

// Render a single BibTeX element
func (element *Element) String() string {
	var out []string

	out = append(out, fmt.Sprintf("@%s{", element.Type))
	for _, ky := range element.Keys {
		out = append(out, fmt.Sprintf("    %s,", ky))
	}
	for ky, val := range element.Tags {
		out = append(out, fmt.Sprintf("    %s = %s,", ky, val))
	}

	out = append(out, fmt.Sprintf("}"))
	return strings.Join(out, "\n")
}

//
// Parser related structures
//

func advanceTo(targetType string, lineNo int, buf []byte) (int, *tok.Token, []byte, error) {
	var (
		token *tok.Token
	)
	startLine := lineNo
	for {
		if len(buf) == 0 {
			return lineNo, token, buf, fmt.Errorf("%d expected %s after %d", lineNo, targetType, startLine)
		}
		token, buf = tok.Tok2(buf, tok.Bib)
		if token.Type == tok.Space {
			// Keep track of what line we're on from source buffer
			lineNo = lineNo + bytes.Count(token.Value, []byte("\n"))
		}
		//fmt.Printf("DEBUG %d token: %s\n", lineNo, token)
		if token.Type == targetType {
			return lineNo, token, buf, nil
		}
	}
}

func skipSpaces(lineNo int, buf []byte) (int, *tok.Token, []byte) {
	var (
		token *tok.Token
	)
	for {
		if len(buf) == 0 {
			return lineNo, token, buf
		}
		token, buf = tok.Tok2(buf, tok.Bib)
		if token.Type == tok.Space {
			// Keep track of what line we're on from source buffer
			lineNo = lineNo + bytes.Count(token.Value, []byte("\n"))
		} else {
			//fmt.Printf("DEBUG %d token: %s\n", lineNo, token)
			return lineNo, token, buf
		}
	}
}

func parseTagValue(lineNo int, buf []byte) (int, []byte, []byte, error) {
	var (
		quoteValue []byte
		i          int
		token      *tok.Token
		val        []byte
		err        error
	)
	lineNo, token, buf = skipSpaces(lineNo, buf)
	if token.Type == "AlphaNumeric" {
		lineNo, _, buf, err = advanceTo("Comma", lineNo, buf)
		//fmt.Printf("DEBUG %d, val: %s\n", lineNo, token.Value)
		return lineNo, token.Value, buf, err
	} else {
		quoteValue = token.Value
		for i = 0; i < len(buf); i++ {
			//FIXME: Need to handle concatenation # when quotaVale is single/double quote
			//FIXME: Need to handle escaping the quoteValue...
			if bytes.Equal(buf[i:i+1], quoteValue) {
				break
			}
		}
	}
	// Copy out the value from buffer and advance
	val, buf = buf[0:i], buf[i:]
	//fmt.Printf("DEBUG %d quoteValue: %s, val: %s\n", lineNo, quoteValue, val)
	return lineNo, val, buf, nil
}

func consumeBibString(quoteType string, val []byte, lineNo int, buf []byte) (int, []byte, *tok.Token, []byte) {
	var (
		token      *tok.Token
		quoteCount int
	)
	// When consumeBibString has happened we're already inside a quote
	if quoteType == tok.DoubleQuote || quoteType == tok.OpenCurlyBracket {
		quoteCount++
	}

	for {
		if len(buf) == 0 {
			break
		}
		token, buf = tok.Tok2(buf, tok.Bib)
		switch {
		case quoteType == tok.OpenCurlyBracket && token.Type == tok.OpenQurlyBracket:
			quoteCount++
		case quoteType == tok.OpenCurlyBracket && token.Type == tok.CloseQurlyBracket:
			quoteCount--
			// NOTE: This handles the case of a missing trailing comma in a tag or key entry
			if quoteQount < 1 {
				break
			}
		case quoteType != tok.OpenCurlyBracket && token.Type == tok.DoubleQuote:
			// NOTE: If we're not using curlies to quote a string then we may have concatenated strings
			// so want to toggle our quote count.
			if quoteCount == 1 {
				quoteCount = 0
			} else {
				quoteCount = 1
			}
		case quoteCount == 0 && token.Type == "Comma":
			// NOTE: tag/key ends normally
			break
		case quoteCount == 0 && token.Type == tok.CloseCurlyBracket:
			// NOTE: tag/key is missing trailing comma in a tag or key entry
			break
		case quoteCount == 0 && token.Type == tok.EqualSign:
			// NOTE: Handle the case where we have a key for a tag.
			break
		}
		val = append(val, token.Value)
	}
	return lineNo, val, token, buf
}

func parseKeysAndTags(lineNo int, buf []byte) (int, []string, map[string]string, []byte, error) {
	var (
		keys       []string
		tags       map[string]string
		ky         string
		val        []byte
		token      *tok.Token
		err        error
		isKeyValue bool
	)
	tags = make(map[string]string)
	// Skip leading spaces
	lineNo, token, buf = skipSpaces(lineNo, buf)
	for {
		if len(buf) == 0 {
			break
		}
		// Now we need to evaluate the entry
		// Do we have a key, a tag or comment string?
		isKeyValue = true
		switch {
		case token.Type == "AlphaNumeric":
			ky = fmt.Sprintf("%s", token.Value)
			lineNo, token, buf = skipSpaces(lineNo, buf)
		default:
			// We have some sort of ID, string or key
			lineNo, val, token, buf = consumeBibString(tok.Type, token.Value, lineNo, buf)
			ky = fmt.Sprintf("%q", val)
		}

		switch token.Type {
		case tok.EqualSign:
			isKeyValue = false
			lineNo, token, buf = skipSpaces(lineNo, buf)
			lineNo, val, buf, err = consumeBibString(token.Type, lineNo, buf)
			if err != nil {
				return lineNo, keys, tags, buf, err
			}
			if val != nil {
				tags[ky] = fmt.Sprintf("%s", val)
			}
		case "Comma":
			if isKeyValue == true {
				keys = append(keys, ky)
			}
		default:
			break
		}
	}
	fmt.Printf("DEBUG ParseKeysAndTags() ended with Curly! %d, %d\n", lineNo, len(buf))
	return lineNo, keys, tags, buf, nil
}

// Parse a BibTeX file into appropriate structures
func Parse(buf []byte) ([]*Element, error) {
	var (
		lineNo   int
		token    *tok.Token
		elements []*Element
		err      error
		keys     []string
		tags     map[string]string
	)
	lineNo = 1
	for {
		if len(buf) == 0 {
			break
		}
		fmt.Printf("DEBUG finding next AtSign: %d\n", len(buf))
		lineNo, token, buf, err = advanceTo(tok.AtSign, lineNo, buf)
		if err != nil && len(elements) == 0 {
			return elements, err
		}
		fmt.Printf("DEBUG buf len: %d\n", len(buf))
		if token.Type == tok.AtSign {
			lineNo, token, buf, err = advanceTo("AlphaNumeric", lineNo, buf)
			if err != nil {
				return elements, fmt.Errorf("line %d, %s", lineNo, err)
			}
			elementType := fmt.Sprintf("%s", token.Value)
			lineNo, token, buf, err = advanceTo(tok.OpenCurlyBracket, lineNo, buf)
			if err != nil {
				return elements, fmt.Errorf("line %d, %s", lineNo, err)
			}
			lineNo, keys, tags, buf, err = parseKeysAndTags(lineNo, buf)
			if err != nil {
				return elements, fmt.Errorf("line %d, %s", lineNo, err)
			}
			// Add element to the list
			elements = append(elements, &Element{
				Type: elementType,
				Keys: keys,
				Tags: tags,
			})
			fmt.Printf("DEBUG last element:\n%s\n", elements[len(elements)-1])
			fmt.Printf("DEBUG elements:\n%s\n", elements)
		}
		fmt.Printf("DEBUG buf len: %d\n", len(buf))
	}
	return elements, nil
}
