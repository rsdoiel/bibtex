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

// Version of mkbib library
const (
	Version = "0.0.0"
)

// Generic Element
type Element struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:"type" json:"type"`
	Value   interface{}
}
type Elements []*Element

// Entry types

// Article type Bib Element
type Article struct {
	// Required fields: author, title, journal, year, volume
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"`
	Title       string   `xml:"title" json:"title"`
	Journal     string   `xml:"journal" json:"journal"`
	Year        string   `xml:"year", json:"year"`
	Volume      string   `xml:"volume" json:"volume"`

	// Optional fields: number, pages, month, note, key
	Number string `xml:"number,omitempty" json:"number,omitempty"`
	Pages  string `xml:"pages,omitempty" json:"pages,omitempty"`
	Month  string `xml:"month,omitempty" json:"month,omitempty"`
	Note   string `xml:"note,omitempty" json:"note,omitempty"`
	Key    string `xml:"key,omitempty" json:"key,omitempty"`
}

type Book struct {
	// Required fields: author/editor, title, publisher, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"` // You need at least one Author or Editor, can also have both
	Editor      string   `xml:"editor" json:"editor"`
	Title       string   `xml:"title" json:"title"`
	Publisher   string   `xml:"publisher" json:"publisher"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: volume/number, series, address, edition, month, note, key
	Volume  string `xml:"volume,omitempty" json:"volume,omitempty"`
	Number  string `xml:"number,omitempty" json:"number,omitempty"`
	Series  string `xml:"series,omitempty" json:"series,omitempty"`
	Address string `xml:"address,omitempty" json:"address,omitempty"`
	Edition string `xml:"edition,omitempty" json:"edition,omitempty"`
	Month   string `xml:"month,omitempty" json:"month,omitempty"`
	Note    string `xml:"note,omitempty" json:"note,omitempty"`
	Key     string `xml:"key,omitempty" json:"key,omitempty"`
}

type Booklet struct {
	// Required fields: title
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Title       string   `xml:"title" json:"title"`

	// Optional fields: author, howpublished, address, month, year, note, key
	Author       string `xml:"author,omitempty" json:"author,omitempty"`
	HowPublished string `xml:"howpublished,omitempty" json:"howpublished,omitempty"`
	Address      string `xml:"address,omitempty" json:"address,omitempty"`
	Month        string `xml:"month,omitempty" json:"month,omitempty"`
	Year         string `xml:"year,omitempty" json:"year,omitempty"`
	Note         string `xml:"note,omitempty" json:"note,omitempty"`
	Key          string `xml:"key,omitempty" json:"key,omitempty"`
}

type InBook struct {
	// Reuqired fields: author/editor, title, chapter/pages, publisher, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"` // You need at least one Author or Editor, can also have both
	Editor      string   `xml:"editor" json:"editor"`
	Title       string   `xml:"title" json:"title"`
	Chapter     string   `xml:"chapter" json:"chapter"` // You need at least Chapter or Pages, can also have both
	Pages       string   `xml:"pages" json:"pages"`
	Publisher   string   `xml:"publisher" json:"publisher"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: volume/number, series, type, address, edition, month, note, key
	Volume  string `xml:"volume,omitempty" json:"volume,omitempty"` // You may have Volune, Number or both
	Number  string `xml:"number,omitempty" json:"number,omitempty"`
	Series  string `xml:"series,omitempty" json:"series,omitempty"`
	Type    string `xml:"type,omitempty" json:"type,omitempty"`
	Address string `xml:"address,omitempty" json:"address,omitempty"`
	Edition string `xml:"edition,omitempty" json:"edition,omitempty"`
	Month   string `xml:"month,omitempty" json:"month,omitempty"`
	Note    string `xml:"note,omitempty" json:"note,omitempty"`
	Key     string `xml:"key,omitempty" json:"key,omitempty"`
}

type InCollection struct {
	// Reuqired fields: author, title, booktitle, publisher, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"`
	Title       string   `xml:"title" json:"title"`
	BookTitle   string   `xml:"booktitle" json:"booktitle"`
	Publisher   string   `xml:"publisher" json:"publisher"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: editor, volume/number, series, type, chapter, pages, address, edition, month, note, key
	Editor  string `xml:"editor,omitempty" json:"editor,omitempty"`
	Volume  string `xml:"volume,omitempty" json:"volume,omitempty"`
	Number  string `xml:"number,omitempty" json:"number,omitempty"`
	Series  string `xml:"series,omitempty" json:"series,omitempty"`
	Type    string `xml:"type,omitempty" json:"type,omitempty"`
	Chapter string `xml:"chapter,omitempty" json:"chapter,omitempty"`
	Pages   string `xml:"pages,omitempty" json:"pages,omitempty"`
	Address string `xml:"address,omitempty" json:"address,omitempty"`
	Edition string `xml:"edition,omitempty" json:"edition,omitempty"`
	Month   string `xml:"month,omitempty" json:"month,omitempty"`
	Note    string `xml:"note,omitempty" json:"note,omitempty"`
	Key     string `xml:"key,omitempty" json:"key,omitempty"`
}

type InProceedings struct {
	// Required fields: author, title, booktitle, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"`
	Title       string   `xml:"title" json:"title"`
	BookTitle   string   `xml:"booktitle" json:"booktitle"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: editor, volume/number, series, pages, address, month, organization, publisher, note, key
	Editor       string `xml:"editor,omitempty" json:"editor,omitempty"`
	Volume       string `xml:"volume,omitempty" json:"volume,omitempty"`
	Number       string `xml:"number,omitempty" json:"number,omitempty"`
	Series       string `xml:"series,ommitempty" json:"series,omitempty"`
	Pages        string `xml:"pages,omitempty" json:"pages,omitempty"`
	Address      string `xml:"address,omitempty" json:"address,omitempty"`
	Month        string `xml:"month,omitempty" json:"month,omitempty"`
	Organization string `xml:"organization,omitempty" json:"organization,omitempty"`
	Publisher    string `xml:"publisher,omitempty" json:"publisher,omitempty"`
	Note         string `xml:"note,omitempty" json:"note,omitempty"`
	//FIXME: Make sure that key and citation key are not the name.
	Key string `xml:"key,omitempty" json:"key,omitempty"`
}

type Conference InProceedings

type Manual struct {
	// Required fields: title
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Title       string   `xml:"title" json:"title"`

	// Optional fields: author, organization, address, edition, month, year, note, key
	Author       string `xml:"author,omitempty" json:"author,omitempty"`
	Organization string `xml:"organization,omitempty" json:"organization,omitempty"`
	Address      string `xml:"address,omitempty" json:"address,omitempty"`
	Edition      string `xml:"edition,omitempty" json:"edition,omitempty"`
	Month        string `xml:"month,omitempty" json:"month,omitempty"`
	Year         string `xml:"year,omitempty" json:"year,omitempty"`
	Note         string `xml:"note,omitempty" json:"note,omitempty"`
	Key          string `xml:"key,omitempty" json:"key,omitempty"`
}

type MastersThesis struct {
	// Required fields: author, title, school, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"`
	Title       string   `xml:"title" json:"title"`
	School      string   `xml:"school" json:"school"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: type, address, month, note, key
	Type    string `xml:"type,omitempty" json:"type,omitempty"`
	Address string `xml:"address,omitempty" json:"address,omitempty"`
	Month   string `xml:"month,omitempty" json:"month,omitempty"`
	Note    string `xml:"note,omitempty" json:"note,omitempty"`
	Key     string `xml:"key,omitempty" json:"key,omitempty"`
}

type Misc struct {
	// Required fields: none
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`

	// Optional fields: author, title, howpublished, month, year, note, key
	Author       string `xml:"author,omitempty" json:"author,omitempty"`
	Title        string `xml:"title,omitempty" json:"title,omitempty"`
	HowPublished string `xml:"how_published,omitempty" json:"how_published,omitempty"`
	Month        string `xml:"month,omitempty" json:"month,omitempty"`
	Year         string `xml:"year" json:"year"`
	Note         string `xml:"note,omitempty" json:"note,omitempty"`
	Key          string `xml:"key,omitempty" json:"key,omitempty"`
}

type PhDThesis struct {
	// Required fields: author, title, school, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"`
	Title       string   `xml:"title" json:"title"`
	School      string   `xml:"school" json:"school"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: type, address, month, note, key
	Type    string `xml:"type,omitempty" json:"type,omitempty"`
	Address string `xml:"address,omitempty" json:"address,omitempty"`
	Month   string `xml:"month,omitempty" json:"month,omitempty"`
	Note    string `xml:"note,omitempty" json:"note,omitempty"`
	Key     string `xml:"key,omitempty" json:"key,omitempty"`
}

type Proceedings struct {
	// Required fields: title, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Title       string   `xml:"title" json:"title"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: editor, volume/number, series, address, month, publisher, organization, note, key
	Editor       string `xml:"editor" json:"editor"`
	Volume       string `xml:"volume,omitempty" json:"volume,omitempty"`
	Number       string `xml:"number,omitempty" json:"number,omitempty"`
	Series       string `xml:"series,omitempty" json:"series,omitempty"`
	Address      string `xml:"address,omitempty" json:"address,omitempty"`
	Month        string `xml:"month,omitempty" json:"month,omitempty"`
	Publisher    string `xml:"publisher,omitempty" json:"publisher,omitempty"`
	Organization string `xml:"organization,omitempty" json:"organization,omitempty"`
	Note         string `xml:"note,omitempty" json:"note,omitempty"`
	Key          string `xml:"key,omitempty" json:"key,omitempty"`
}

type TechReport struct {
	// Required fields: author, title, institution, year
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"`
	Title       string   `xml:"title" json:"title"`
	Institution string   `xml:"institution" json:"institution"`
	Year        string   `xml:"year" json:"year"`

	// Optional fields: type, number, address, month, note, key
	Type    string `xml:"type,omitempty" json:"type,omitempty"`
	Number  string `xml:"number,omitempty" json:"number,omitempty"`
	Address string `xml:"address,omitempty" json:"address,omitempty"`
	Month   string `xml:"month,omitempty" json:"month,omitempty"`
	Note    string `xml:"note,omitempty" json:"note,omitempty"`
	Key     string `xml:"key,omitempty" json:"key,omitempty"`
}

type Unpublished struct {
	// Required fields: author, title, note
	XMLName     xml.Name `json:"-"`
	CitationKey string   `xml:"citation_key", json:"citation_key"`
	Author      string   `xml:"author" json:"author"`
	Title       string   `xml:"title" json:"title"`
	Note        string   `xml:"note" json:"note"`

	// Optional fields: month, year, key
	// month, year, key
	Month string `xml:"month,omitempty" json:"month,omitempty"`
	Year  string `xml:"year,omitempty" json:"year,omitempty"`
	Key   string `xml:"key,omitempty" json:"key,omitempty"`
}

// String conversions render in BibText format
func (a *Article) String() string {
	var (
		kv []string
	)
	if a.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", a.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "author", a.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", a.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "journal", a.Journal))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", a.Year))
	kv = append(kv, fmt.Sprintf("%s = %q", "volume", a.Volume))

	// number, pages, month, note, key
	if a.Number != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", a.Number))
	}
	if a.Pages != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", a.Pages))
	}
	if a.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", a.Month))
	}
	if a.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", a.Note))
	}
	if a.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", a.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@article { %s }`, data)
}

func (b *Book) String() string {
	var (
		kv []string
	)
	// author/editor, title, publisher, year
	if b.Author != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", b.Author))
	}
	if b.Editor != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", b.Editor))
	}
	if b.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", b.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "title", b.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "publisher", b.Publisher))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", b.Year))

	//volume/number, series, address, edition, month, note, key
	if b.Volume != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", b.Volume))
	}
	if b.Number != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", b.Number))
	}
	if b.Series != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", b.Series))
	}
	if b.Edition != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", b.Edition))
	}
	if b.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", b.Month))
	}
	if b.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", b.Note))
	}
	if b.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", b.Key))
	}

	data := strings.Join(kv, ", ")
	//author/editor, title, publisher, year
	return fmt.Sprintf(`@book{ %s }`, data)
}

func (bl *Booklet) String() string {
	var (
		kv []string
	)

	// Required fields: Title
	if bl.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", bl.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "title", bl.Title))
	// Optional fields: author, howpublished, address, month, year, note, key
	if bl.Author != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", bl.Author))
	}
	if bl.HowPublished != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "howpublished", bl.HowPublished))
	}
	if bl.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", bl.Address))
	}
	if bl.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", bl.Month))
	}
	if bl.Year != "" {
		kv = append(kv, fmt.Sprintf("%s = %s", "year", bl.Year))
	}
	if bl.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", bl.Note))
	}
	if bl.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", bl.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@booklet{ %s }`, data)
}

func (ib *InBook) String() string {
	var (
		kv []string
	)
	// Required fields: author/editor, title, chapter/pages, publisher, year
	if ib.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", ib.CitationKey))
	}
	if ib.Author != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", ib.Author))
	}
	if ib.Editor != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", ib.Editor))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "title", ib.Title))
	if ib.Chapter != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "chapter", ib.Chapter))
	}
	if ib.Pages != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", ib.Chapter))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "publisher", ib.Publisher))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", ib.Year))

	// Optional fields: volume/number, series, type, address, edition, month, note, key
	if ib.Volume != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", ib.Volume))
	}
	if ib.Number != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", ib.Number))
	}
	if ib.Series != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", ib.Series))
	}
	if ib.Type != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", ib.Type))
	}
	if ib.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", ib.Address))
	}
	if ib.Edition != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", ib.Edition))
	}
	if ib.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", ib.Month))
	}
	if ib.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", ib.Note))
	}
	if ib.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", ib.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@inbook{ %s }`, data)
}

func (ic *InCollection) String() string {
	var (
		kv []string
	)
	// Required fields: author, title, booktitle, publisher, yeic.
	if ic.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", ic.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "author", ic.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", ic.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "booktitle", ic.BookTitle))
	kv = append(kv, fmt.Sprintf("%s = %q", "publisher", ic.Publisher))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", ic.Year))

	// Optional fields: editor, volume/number, series, type, chapter, pages, address, edition, month, note, key
	if ic.Editor != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", ic.Editor))
	}
	if ic.Volume != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", ic.Volume))
	}
	if ic.Number != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", ic.Number))
	}
	if ic.Series != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", ic.Series))
	}
	if ic.Type != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", ic.Type))
	}
	if ic.Chapter != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "chapter", ic.Chapter))
	}
	if ic.Pages != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", ic.Pages))
	}
	if ic.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", ic.Address))
	}
	if ic.Edition != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", ic.Edition))
	}
	if ic.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", ic.Month))
	}
	if ic.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", ic.Note))
	}
	if ic.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", ic.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@incollection{ %s }`, data)
}

func (ip *InProceedings) String() string {
	var (
		kv []string
	)
	// Required fields: author, title, booktitle, year
	if ip.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", ip.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "author", ip.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", ip.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "booktitle", ip.BookTitle))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", ip.Year))

	// Optional fields: editor, volume/number, series, pages, address, month, organization, publisher, note, key
	if ip.Editor != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", ip.Editor))
	}
	if ip.Volume != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", ip.Volume))
	}
	if ip.Number != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", ip.Number))
	}
	if ip.Series != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", ip.Series))
	}
	if ip.Pages != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", ip.Pages))
	}
	if ip.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", ip.Address))
	}
	if ip.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", ip.Month))
	}
	if ip.Organization != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "organization", ip.Organization))
	}
	if ip.Publisher != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "publisher", ip.Publisher))
	}
	if ip.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", ip.Note))
	}
	if ip.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", ip.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@inproceedings{ %s }`, data)
}

func (m *Manual) String() string {
	var (
		kv []string
	)
	// Required fields: title
	if m.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", m.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "title", m.Title))
	// Optional fields: author, organization, address, edition, month, year, note, key
	if m.Author != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", m.Author))
	}
	if m.Organization != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "organization", m.Organization))
	}
	if m.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", m.Address))
	}
	if m.Edition != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", m.Edition))
	}
	if m.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", m.Month))
	}
	if m.Year != "" {
		kv = append(kv, fmt.Sprintf("%s = %s", "year", m.Year))
	}
	if m.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", m.Note))
	}
	if m.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", m.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@manual{ %s }`, data)
}

func (mt *MastersThesis) String() string {
	var (
		kv []string
	)
	// Required fields: author, title, school, year
	if mt.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", mt.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "author", mt.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", mt.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "school", mt.School))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", mt.Year))

	// Optional fields: type, address, month, note, key
	if mt.Type != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", mt.Type))
	}
	if mt.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", mt.Address))
	}
	if mt.School != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "school", mt.School))
	}
	if mt.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", mt.Month))
	}
	if mt.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", mt.Note))
	}
	if mt.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", mt.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@mastersthesis{ %s }`, data)
}

func (misc *Misc) String() string {
	var (
		kv []string
	)
	// Required fields: none
	if misc.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", misc.CitationKey))
	}
	// Optional fields: author, title, howpublished, month, year, note, key
	if misc.Author != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", misc.Author))
	}
	if misc.Title != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "title", misc.Title))
	}
	if misc.HowPublished != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "howpublished", misc.HowPublished))
	}
	if misc.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", misc.Month))
	}
	if misc.Year != "" {
		kv = append(kv, fmt.Sprintf("%s = %s", "year", misc.Year))
	}
	if misc.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", misc.Note))
	}
	if misc.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", misc.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@misc{ %s }`, data)
}

func (phd *PhDThesis) String() string {
	var (
		kv []string
	)
	// Required fields: author, title, school, year
	if phd.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", phd.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "author", phd.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", phd.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "school", phd.School))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", phd.Year))

	// Optional fields: type, address, month, note, key
	if phd.Type != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", phd.Type))
	}
	if phd.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", phd.Address))
	}
	if phd.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", phd.Month))
	}
	if phd.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", phd.Note))
	}
	if phd.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", phd.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@phdthesis{ %s }`, data)
}

func (p *Proceedings) String() string {
	var (
		kv []string
	)
	// Required fields: title, year
	if p.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", p.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "title", p.Title))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", p.Year))

	// Optional fields: editor, volume/number, series, address, month, publisher, organization, note, key
	if p.Editor != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", p.Editor))
	}
	if p.Volume != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", p.Volume))
	}
	if p.Number != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", p.Number))
	}
	if p.Series != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", p.Series))
	}
	if p.Address != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", p.Address))
	}
	if p.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", p.Month))
	}
	if p.Publisher != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "publisher", p.Publisher))
	}
	if p.Organization != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "organization", p.Organization))
	}
	if p.Note != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", p.Note))
	}
	if p.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", p.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@proceedings{ %s }`, data)
}

func (t *TechReport) String() string {
	var (
		kv []string
	)
	// Required fields: author, title, institution, year
	if t.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", t.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "author", t.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", t.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "institution", t.Institution))
	kv = append(kv, fmt.Sprintf("%s = %s", "year", t.Year))

	// Optional fields: type, number, address, month, note, key
	if t.Type != "" {
		kv = append(kv, fmt.Sprintf("%s =%q", "type", t.Type))
	}
	if t.Number != "" {
		kv = append(kv, fmt.Sprintf("%s =%q", "number", t.Number))
	}
	if t.Address != "" {
		kv = append(kv, fmt.Sprintf("%s =%q", "address", t.Address))
	}
	if t.Month != "" {
		kv = append(kv, fmt.Sprintf("%s =%q", "month", t.Month))
	}
	if t.Note != "" {
		kv = append(kv, fmt.Sprintf("%s =%q", "note", t.Note))
	}
	if t.Key != "" {
		kv = append(kv, fmt.Sprintf("%s =%q", "key", t.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@techreport{ %s }`, data)
}

func (u *Unpublished) String() string {
	var (
		kv []string
	)

	// Required fields: author, title, note
	if u.CitationKey != "" {
		kv = append(kv, fmt.Sprintf("%s", u.CitationKey))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "author", u.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", u.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "note", u.Note))

	// Optional fields: month, year, key
	if u.Month != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", u.Month))
	}
	if u.Year != "" {
		kv = append(kv, fmt.Sprintf("%s = %s", "year", u.Year))
	}
	if u.Key != "" {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", u.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@unpublished{ %s }`, data)
}

func (elements Elements) String() string {
	var out []string

	for _, element := range elements {
		switch element.Type {
		default:
			// FIXME: Cast the element.Value to appropriate type and render with String()
			out = append(out, fmt.Sprintf("DEBUG render type %s +v", element.Type, element.Value))
		}
	}

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
			return lineNo, token, buf
		}
	}
}

func parseValue(lineNo int, buf []byte) (int, []byte, []byte, error) {
	var (
		quoteValue []byte
		i          int
		token      *tok.Token
		val        []byte
	)
	//FIXME: Need to handle case where a numeric value is not quoted
	lineNo, token, buf = skipSpaces(lineNo, buf)
	quoteValue = token.Value
	for i = 0; i < len(buf); i++ {
		//FIXME: Need to handle escaping the quoteValue...
		if bytes.Equal(buf[i:i+1], quoteValue) {
			break
		}
	}
	// Copy out the value from buffer and advance
	val, buf = buf[0:i], buf[i-1:]
	fmt.Printf("DEBUG %d quoteValue: %s, val: %s, buf: %s\n", lineNo, quoteValue, val, buf)
	return lineNo, val, buf, nil
}

func parseKeysAndAttributes(lineNo int, buf []byte) (int, string, map[string]string, []byte, error) {
	var (
		citationKey string
		attributes  map[string]string
		ky          string
		val         []byte
		token       *tok.Token
		err         error
	)
	attributes = make(map[string]string)
	for {
		if len(buf) == 0 {
			break
		}
		lineNo, token, buf = skipSpaces(lineNo, buf)
		if token.Type == "AlphaNumeric" {
			ky = fmt.Sprintf("%s", token.Value)
		}
		lineNo, token, buf = skipSpaces(lineNo, buf)
		fmt.Printf("DEBUG %d token: %s\n", lineNo, token)
		if token != nil {
			switch token.Type {
			case tok.EqualSign:
				// FIXME: Cut until a trailing comma is encountered.
				lineNo, val, buf, err = parseValue(lineNo, buf)
				if err != nil {
					return lineNo, citationKey, attributes, buf, err
				}
				fmt.Printf("DEBUG ky/val: %s -> %s\n", ky, val)
				if val != nil {
					attributes[ky] = fmt.Sprintf("%s", val)
				}
			case "Comma":
				citationKey = ky
			case tok.CloseCurlyBracket:
				break
			}
		} else {
			// FIXME: Should not get here???? should have seen a CloseCurlyBracket before then.
			break
		}
	}

	return lineNo, citationKey, attributes, buf, nil
}

func mkArticle(citationKey string, attributes map[string]string) *Article {
	article := new(Article)
	article.CitationKey = citationKey
	if val, ok := attributes["author"]; ok == true {
		article.Author = val
	}
	if val, ok := attributes["title"]; ok == true {
		article.Title = val
	}
	if val, ok := attributes["journal"]; ok == true {
		article.Journal = val
	}
	if val, ok := attributes["year"]; ok == true {
		article.Year = val
	}
	if val, ok := attributes["volume"]; ok == true {
		article.Volume = val
	}
	if val, ok := attributes["number"]; ok == true {
		article.Number = val
	}
	if val, ok := attributes["pages"]; ok == true {
		article.Pages = val
	}
	if val, ok := attributes["month"]; ok == true {
		article.Month = val
	}
	if val, ok := attributes["note"]; ok == true {
		article.Note = val
	}
	if val, ok := attributes["key"]; ok == true {
		article.Key = val
	}

	fmt.Printf("DEBUG mkArticle() -> %s\n", article)
	return article
}

func mkElement(elementType string, citationKey string, attributes map[string]string) (*Element, error) {
	element := new(Element)
	switch strings.ToLower(elementType) {
	//FIXME: Need to populate Element
	case "article":
		element.Type = "article"
		element.Value = mkArticle(citationKey, attributes)
	default:
		return nil, fmt.Errorf("unknown elemenet type: %s", elementType)
	}

	return element, nil
}

// Parse a BibTeX file into appropriate structures
func Parse(buf []byte) ([]*Element, error) {
	var (
		lineNo      int
		token       *tok.Token
		elements    []*Element
		citationKey string
		attributes  map[string]string
		err         error
	)
	lineNo = 1
	for {
		if len(buf) == 0 {
			break
		}
		lineNo, token, buf, err = advanceTo(tok.AtSign, lineNo, buf)
		if err != nil {
			return elements, err
		}
		if token.Type == tok.AtSign {
			element := new(Element)
			lineNo, token, buf, err = advanceTo("AlphaNumeric", lineNo, buf)
			if err != nil {
				return elements, err
			}
			fmt.Printf("DEBUG element type: %s -> %s\n", token, buf)
			elementType := fmt.Sprintf("%s", token.Value)
			lineNo, token, buf, err = advanceTo(tok.OpenCurlyBracket, lineNo, buf)
			if err != nil {
				return elements, err
			}
			/*
				lineNo, token, buf = skipSpaces(lineNo, buf)
				if token.Type != tok.OpenCurlyBracket {
					return elements, fmt.Errorf("line %d, expected an open curly bracket", lineNo)
				}
			*/

			lineNo, citationKey, attributes, buf, err = parseKeysAndAttributes(lineNo, buf)
			if err != nil {
				return elements, err
			}
			// Assemble our element
			element, err := mkElement(elementType, citationKey, attributes)
			if err != nil {
				return elements, fmt.Errorf("line %d, %s", lineNo, err)
			}

			// Add to the list
			elements = append(elements, element)
		}
	}
	return elements, nil
}
