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
)

// Version of mkbib library
const Version = "0.0.0"

type Citation struct {
	XMLName   xml.Name `json:"-"`
	Type      string   `xml:"type,omitempty", json:"type,omitempty"`
	Authors   []string `xml:"author,omitempty" json:"author,omitempty"`
	Title     string   `xml:"title,omitempty" json:"title,omitemtpy"`
	Publisher string   `xml:"publisher,omitempty" json:"publisher,omitempty"`
	Year      string   `xml:"year,omitempty" json:"year,omitempt"`
}

func (c *Citation) String() string {
	return fmt.Sprintf(`
 @%s{
	author = %q,
	title = %q,
	publisher = %q,
	year = %q
 }
 `, c.Type, strings.Join(c.Authors, ", "), c.Title, c.Publisher, c.Year)
}

func mkCitation(s string) *Citation {
	c := new(Citation)
	fmt.Printf("DEBUG FIXME mkCitation: %s\n", s)
	c.Authors = []string{"R. S. Doiel"}
	c.Title = "This is a big test"
	c.Publisher = "The BFG"
	c.Year = "2016"
	return c
}

func Parse(buf []byte) ([]byte, error) {
	out := []string{}
	in := bytes.NewBuffer(buf)
	i := 0
	var err error
	for err == nil {
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}
		i++
		cite := mkCitation(line)
		fmt.Printf("DEBUG line: %d %s, %s\n", i, line, cite)
		out = append(out, cite.String())
	}
	return []byte(strings.Join(out, "\n")), nil
}
