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

// Field types
type Address string
type Annotation string
type Author string
type BookTitle string
type Chapter string
type CrossRef string
type Edition string
type Editor string
type HowPublished string
type Institution string
type Journal string
type Key string
type Month string
type Note string
type Number string
type Organization string
type Pages string
type Publisher string
type School string
type Series string
type Title string
type Type string
type Volume string
type Year string
type CitationKey string

// Generic Entries
type Entry struct {
	XMLName xml.Name          `json:"-"`
	Type    string            `xml:"type" json:"type"`
	Key     string            `xml:"citation-key,omitempty" json:"citation-key,omitempty"`
	Fields  map[string]string `xml:"fields" json:"fields"`
}

// Entry types
type Article struct {
	// Required fields: author, title, journal, year, volume
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"`
	Title       *Title       `xml:"title" json:"title"`
	Journal     *Journal     `xml:"journal" json:"journal"`
	Year        *Year        `xml:"year", json:"year"`
	Volume      *Volume      `xml:"volume" json:"volume"`

	// Optional fields: number, pages, month, note, key
	Number *Number `xml:"number,omitempty" json:"number,omitempty"`
	Pages  *Pages  `xml:"pages,omitempty" json:"pages,omitempty"`
	Month  *Month  `xml:"month,omitempty" json:"month,omitempty"`
	Note   *Note   `xml:"note,omitempty" json:"note,omitempty"`
	Key    *Key    `xml:"key,omitempty" json:"key,omitempty"`
}

type Book struct {
	// Required fields: author/editor, title, publisher, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"` // You need at least one Author or Editor, can also have both
	Editor      *Editor      `xml:"editor" json:"editor"`
	Title       *Title       `xml:"title" json:"title"`
	Publisher   *Publisher   `xml:"publisher" json:"publisher"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: volume/number, series, address, edition, month, note, key
	Volume  *Volume  `xml:"volume,omitempty" json:"volume,omitempty"`
	Number  *Number  `xml:"number,omitempty" json:"number,omitempty"`
	Series  *Series  `xml:"series,omitempty" json:"series,omitempty"`
	Address *Address `xml:"address,omitempty" json:"address,omitempty"`
	Edition *Edition `xml:"edition,omitempty" json:"edition,omitempty"`
	Month   *Month   `xml:"month,omitempty" json:"month,omitempty"`
	Note    *Note    `xml:"note,omitempty" json:"note,omitempty"`
	Key     *Key     `xml:"key,omitempty" json:"key,omitempty"`
}

type Booklet struct {
	// Required fields: title
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Title       *Title       `xml:"title" json:"title"`

	// Optional fields: author, howpublished, address, month, year, note, key
	Author       *Author       `xml:"author,omitempty" json:"author,omitempty"`
	HowPublished *HowPublished `xml:"howpublished,omitempty" json:"howpublished,omitempty"`
	Address      *Address      `xml:"address,omitempty" json:"address,omitempty"`
	Month        *Month        `xml:"month,omitempty" json:"month,omitempty"`
	Year         *Year         `xml:"year,omitempty" json:"year,omitempty"`
	Note         *Note         `xml:"note,omitempty" json:"note,omitempty"`
	Key          *Key          `xml:"key,omitempty" json:"key,omitempty"`
}

type InBook struct {
	// Reuqired fields: author/editor, title, chapter/pages, publisher, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"` // You need at least one Author or Editor, can also have both
	Editor      *Editor      `xml:"editor" json:"editor"`
	Title       *Title       `xml:"title" json:"title"`
	Chapter     *Chapter     `xml:"chapter" json:"chapter"` // You need at least Chapter or Pages, can also have both
	Pages       *Pages       `xml:"pages" json:"pages"`
	Publisher   *Publisher   `xml:"publisher" json:"publisher"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: volume/number, series, type, address, edition, month, note, key
	Volume  *Volume  `xml:"volume,omitempty" json:"volume,omitempty"` // You may have Volune, Number or both
	Number  *Number  `xml:"number,omitempty" json:"number,omitempty"`
	Series  *Series  `xml:"series,omitempty" json:"series,omitempty"`
	Type    *Type    `xml:"type,omitempty" json:"type,omitempty"`
	Address *Address `xml:"address,omitempty" json:"address,omitempty"`
	Edition *Edition `xml:"edition,omitempty" json:"edition,omitempty"`
	Month   *Month   `xml:"month,omitempty" json:"month,omitempty"`
	Note    *Note    `xml:"note,omitempty" json:"note,omitempty"`
	Key     *Key     `xml:"key,omitempty" json:"key,omitempty"`
}

type InCollection struct {
	// Reuqired fields: author, title, booktitle, publisher, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"`
	Title       *Title       `xml:"title" json:"title"`
	BookTitle   *BookTitle   `xml:"booktitle" json:"booktitle"`
	Publisher   *Publisher   `xml:"publisher" json:"publisher"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: editor, volume/number, series, type, chapter, pages, address, edition, month, note, key
	Editor  *Editor  `xml:"editor,omitempty" json:"editor,omitempty"`
	Volume  *Volume  `xml:"volume,omitempty" json:"volume,omitempty"`
	Number  *Number  `xml:"number,omitempty" json:"number,omitempty"`
	Series  *Series  `xml:"series,omitempty" json:"series,omitempty"`
	Type    *Type    `xml:"type,omitempty" json:"type,omitempty"`
	Chapter *Chapter `xml:"chapter,omitempty" json:"chapter,omitempty"`
	Pages   *Pages   `xml:"pages,omitempty" json:"pages,omitempty"`
	Address *Address `xml:"address,omitempty" json:"address,omitempty"`
	Edition *Edition `xml:"edition,omitempty" json:"edition,omitempty"`
	Month   *Month   `xml:"month,omitempty" json:"month,omitempty"`
	Note    *Note    `xml:"note,omitempty" json:"note,omitempty"`
	Key     *Key     `xml:"key,omitempty" json:"key,omitempty"`
}

type InProceedings struct {
	// Required fields: author, title, booktitle, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"`
	Title       *Title       `xml:"title" json:"title"`
	BookTitle   *BookTitle   `xml:"booktitle" json:"booktitle"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: editor, volume/number, series, pages, address, month, organization, publisher, note, key
	Editor       *Editor       `xml:"editor,omitempty" json:"editor,omitempty"`
	Volume       *Volume       `xml:"volume,omitempty" json:"volume,omitempty"`
	Number       *Number       `xml:"number,omitempty" json:"number,omitempty"`
	Series       *Series       `xml:"series,ommitempty" json:"series,omitempty"`
	Pages        *Pages        `xml:"pages,omitempty" json:"pages,omitempty"`
	Address      *Address      `xml:"address,omitempty" json:"address,omitempty"`
	Month        *Month        `xml:"month,omitempty" json:"month,omitempty"`
	Organization *Organization `xml:"organization,omitempty" json:"organization,omitempty"`
	Publisher    *Publisher    `xml:"publisher,omitempty" json:"publisher,omitempty"`
	Note         *Note         `xml:"note,omitempty" json:"note,omitempty"`
	Key          *Key          `xml:"key,omitempty" json:"key,omitempty"`
}

type Conference InProceedings

type Manual struct {
	// Required fields: title
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Title       *Title       `xml:"title" json:"title"`

	// Optional fields: author, organization, address, edition, month, year, note, key
	Author       *Author       `xml:"author,omitempty" json:"author,omitempty"`
	Organization *Organization `xml:"organization,omitempty" json:"organization,omitempty"`
	Address      *Address      `xml:"address,omitempty" json:"address,omitempty"`
	Edition      *Edition      `xml:"edition,omitempty" json:"edition,omitempty"`
	Month        *Month        `xml:"month,omitempty" json:"month,omitempty"`
	Year         *Year         `xml:"year,omitempty" json:"year,omitempty"`
	Note         *Note         `xml:"note,omitempty" json:"note,omitempty"`
	Key          *Key          `xml:"key,omitempty" json:"key,omitempty"`
}

type MastersThesis struct {
	// Required fields: author, title, school, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"`
	Title       *Title       `xml:"title" json:"title"`
	School      *School      `xml:"school" json:"school"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: type, address, month, note, key
	Type    *Type    `xml:"type,omitempty" json:"type,omitempty"`
	Address *Address `xml:"address,omitempty" json:"address,omitempty"`
	Month   *Month   `xml:"month,omitempty" json:"month,omitempty"`
	Note    *Note    `xml:"note,omitempty" json:"note,omitempty"`
	Key     *Key     `xml:"key,omitempty" json:"key,omitempty"`
}

type Misc struct {
	// Required fields: none
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`

	// Optional fields: author, title, howpublished, month, year, note, key
	Author       *Author       `xml:"author,omitempty" json:"author,omitempty"`
	Title        *Title        `xml:"title,omitempty" json:"title,omitempty"`
	HowPublished *HowPublished `xml:"how_published,omitempty" json:"how_published,omitempty"`
	Month        *Month        `xml:"month,omitempty" json:"month,omitempty"`
	Year         *Year         `xml:"year" json:"year"`
	Note         *Note         `xml:"note,omitempty" json:"note,omitempty"`
	Key          *Key          `xml:"key,omitempty" json:"key,omitempty"`
}

type PhDThesis struct {
	// Required fields: author, title, school, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"`
	Title       *Title       `xml:"title" json:"title"`
	School      *School      `xml:"school" json:"school"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: type, address, month, note, key
	Type    *Type    `xml:"type,omitempty" json:"type,omitempty"`
	Address *Address `xml:"address,omitempty" json:"address,omitempty"`
	Month   *Month   `xml:"month,omitempty" json:"month,omitempty"`
	Note    *Note    `xml:"note,omitempty" json:"note,omitempty"`
	Key     *Key     `xml:"key,omitempty" json:"key,omitempty"`
}

type Proceedings struct {
	// Required fields: title, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Title       *Title       `xml:"title" json:"title"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: editor, volume/number, series, address, month, publisher, organization, note, key
	Editor       *Editor       `xml:"editor" json:"editor"`
	Volume       *Volume       `xml:"volume,omitempty" json:"volume,omitempty"`
	Number       *Number       `xml:"number,omitempty" json:"number,omitempty"`
	Series       *Series       `xml:"series,omitempty" json:"series,omitempty"`
	Address      *Address      `xml:"address,omitempty" json:"address,omitempty"`
	Month        *Month        `xml:"month,omitempty" json:"month,omitempty"`
	Publisher    *Publisher    `xml:"publisher,omitempty" json:"publisher,omitempty"`
	Organization *Organization `xml:"organization,omitempty" json:"organization,omitempty"`
	Note         *Note         `xml:"note,omitempty" json:"note,omitempty"`
	Key          *Key          `xml:"key,omitempty" json:"key,omitempty"`
}

type TechReport struct {
	// Required fields: author, title, institution, year
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"`
	Title       *Title       `xml:"title" json:"title"`
	Institution *Institution `xml:"institution" json:"institution"`
	Year        *Year        `xml:"year" json:"year"`

	// Optional fields: type, number, address, month, note, key
	Type    *Type    `xml:"type,omitempty" json:"type,omitempty"`
	Number  *Number  `xml:"number,omitempty" json:"number,omitempty"`
	Address *Address `xml:"address,omitempty" json:"address,omitempty"`
	Month   *Month   `xml:"month,omitempty" json:"month,omitempty"`
	Note    *Note    `xml:"note,omitempty" json:"note,omitempty"`
	Key     *Key     `xml:"key,omitempty" json:"key,omitempty"`
}

type Unpublished struct {
	// Required fields: author, title, note
	XMLName     xml.Name     `json:"-"`
	CitationKey *CitationKey `xml:"key", json:"key"`
	Author      *Author      `xml:"author" json:"author"`
	Title       *Title       `xml:"title" json:"title"`
	Note        *Note        `xml:"note" json:"note"`

	// Optional fields: month, year, key
	// month, year, key
	Month *Month `xml:"month,omitempty" json:"month,omitempty"`
	Year  *Year  `xml:"year,omitempty" json:"year,omitempty"`
	Key   *Key   `xml:"key,omitempty" json:"key,omitempty"`
}

// String conversions render in BibText format
func (a *Article) String() string {
	var (
		kv []string
	)
	kv = append(kv, fmt.Sprintf("%s", a.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "author", a.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", a.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "journal", a.Journal))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", a.Year))
	kv = append(kv, fmt.Sprintf("%s = %q", "volume", a.Volume))

	// number, pages, month, note, key
	if a.Number != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", a.Number))
	}
	if a.Pages != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", a.Pages))
	}
	if a.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", a.Month))
	}
	if a.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", a.Note))
	}
	if a.Key != nil {
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
	if b.Author != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", b.Author))
	}
	if b.Editor != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", b.Editor))
	}
	kv = append(kv, fmt.Sprintf("%s", b.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", b.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "publisher", b.Publisher))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", b.Year))

	//volume/number, series, address, edition, month, note, key
	if b.Volume != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", b.Volume))
	}
	if b.Number != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", b.Number))
	}
	if b.Series != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", b.Series))
	}
	if b.Edition != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", b.Edition))
	}
	if b.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", b.Month))
	}
	if b.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", b.Note))
	}
	if b.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", bl.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", bl.Title))
	// Optional fields: author, howpublished, address, month, year, note, key
	if bl.Author != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", bl.Author))
	}
	if bl.HowPublished != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "howpublished", bl.HowPublished))
	}
	if bl.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", bl.Address))
	}
	if bl.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", bl.Month))
	}
	if bl.Year != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "year", bl.Year))
	}
	if bl.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", bl.Note))
	}
	if bl.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", ib.CitationKey))
	if ib.Author != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", ib.Author))
	}
	if ib.Editor != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", ib.Editor))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "title", ib.Title))
	if ib.Chapter != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "chapter", ib.Chapter))
	}
	if ib.Pages != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", ib.Chapter))
	}
	kv = append(kv, fmt.Sprintf("%s = %q", "publisher", ib.Publisher))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", ib.Year))

	// Optional fields: volume/number, series, type, address, edition, month, note, key
	if ib.Volume != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", ib.Volume))
	}
	if ib.Number != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", ib.Number))
	}
	if ib.Series != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", ib.Series))
	}
	if ib.Type != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", ib.Type))
	}
	if ib.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", ib.Address))
	}
	if ib.Edition != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", ib.Edition))
	}
	if ib.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", ib.Month))
	}
	if ib.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", ib.Note))
	}
	if ib.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", ic.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "author", ic.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", ic.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "booktitle", ic.BookTitle))
	kv = append(kv, fmt.Sprintf("%s = %q", "publisher", ic.Publisher))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", ic.Year))

	// Optional fields: editor, volume/number, series, type, chapter, pages, address, edition, month, note, key
	if ic.Editor != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", ic.Editor))
	}
	if ic.Volume != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", ic.Volume))
	}
	if ic.Number != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", ic.Number))
	}
	if ic.Series != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", ic.Series))
	}
	if ic.Type != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", ic.Type))
	}
	if ic.Chapter != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "chapter", ic.Chapter))
	}
	if ic.Pages != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", ic.Pages))
	}
	if ic.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", ic.Address))
	}
	if ic.Edition != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", ic.Edition))
	}
	if ic.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", ic.Month))
	}
	if ic.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", ic.Note))
	}
	if ic.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", ip.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "author", ip.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", ip.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "booktitle", ip.BookTitle))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", ip.Year))

	// Optional fields: editor, volume/number, series, pages, address, month, organization, publisher, note, key
	if ip.Editor != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", ip.Editor))
	}
	if ip.Volume != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", ip.Volume))
	}
	if ip.Number != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", ip.Number))
	}
	if ip.Series != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", ip.Series))
	}
	if ip.Pages != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "pages", ip.Pages))
	}
	if ip.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", ip.Address))
	}
	if ip.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", ip.Month))
	}
	if ip.Organization != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "organization", ip.Organization))
	}
	if ip.Publisher != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "publisher", ip.Publisher))
	}
	if ip.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", ip.Note))
	}
	if ip.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", m.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", m.Title))
	// Optional fields: author, organization, address, edition, month, year, note, key
	if m.Author != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", m.Author))
	}
	if m.Organization != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "organization", m.Organization))
	}
	if m.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", m.Address))
	}
	if m.Edition != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "edition", m.Edition))
	}
	if m.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", m.Month))
	}
	if m.Year != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "year", m.Year))
	}
	if m.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", m.Note))
	}
	if m.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", mt.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "author", mt.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", mt.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "school", mt.School))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", mt.Year))

	// Optional fields: type, address, month, note, key
	if mt.Type != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", mt.Type))
	}
	if mt.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", mt.Address))
	}
	if mt.School != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "school", mt.School))
	}
	if mt.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", mt.Month))
	}
	if mt.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", mt.Note))
	}
	if mt.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", misc.CitationKey))
	// Optional fields: author, title, howpublished, month, year, note, key
	if misc.Author != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "author", misc.Author))
	}
	if misc.Title != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "title", misc.Title))
	}
	if misc.HowPublished != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "howpublished", misc.HowPublished))
	}
	if misc.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", misc.Month))
	}
	if misc.Year != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "year", misc.Year))
	}
	if misc.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", misc.Note))
	}
	if misc.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", phd.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "author", phd.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", phd.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "school", phd.School))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", phd.Year))

	// Optional fields: type, address, month, note, key
	if phd.Type != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "type", phd.Type))
	}
	if phd.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", phd.Address))
	}
	if phd.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", phd.Month))
	}
	if phd.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", phd.Note))
	}
	if phd.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", p.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", p.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", p.Year))

	// Optional fields: editor, volume/number, series, address, month, publisher, organization, note, key
	if p.Editor != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "editor", p.Editor))
	}
	if p.Volume != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "volume", p.Volume))
	}
	if p.Number != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "number", p.Number))
	}
	if p.Series != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "series", p.Series))
	}
	if p.Address != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "address", p.Address))
	}
	if p.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", p.Month))
	}
	if p.Publisher != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "publisher", p.Publisher))
	}
	if p.Organization != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "organization", p.Organization))
	}
	if p.Note != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "note", p.Note))
	}
	if p.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", t.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "author", t.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", t.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "institution", t.Institution))
	kv = append(kv, fmt.Sprintf("%s = %q", "year", t.Year))

	// Optional fields: type, number, address, month, note, key
	if t.Type != nil {
		kv = append(kv, fmt.Sprintf("%s =%q", "type", t.Type))
	}
	if t.Number != nil {
		kv = append(kv, fmt.Sprintf("%s =%q", "number", t.Number))
	}
	if t.Address != nil {
		kv = append(kv, fmt.Sprintf("%s =%q", "address", t.Address))
	}
	if t.Month != nil {
		kv = append(kv, fmt.Sprintf("%s =%q", "month", t.Month))
	}
	if t.Note != nil {
		kv = append(kv, fmt.Sprintf("%s =%q", "note", t.Note))
	}
	if t.Key != nil {
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
	kv = append(kv, fmt.Sprintf("%s", u.CitationKey))
	kv = append(kv, fmt.Sprintf("%s = %q", "author", u.Author))
	kv = append(kv, fmt.Sprintf("%s = %q", "title", u.Title))
	kv = append(kv, fmt.Sprintf("%s = %q", "note", u.Note))

	// Optional fields: month, year, key
	if u.Month != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "month", u.Month))
	}
	if u.Year != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "year", u.Year))
	}
	if u.Key != nil {
		kv = append(kv, fmt.Sprintf("%s = %q", "key", u.Key))
	}

	data := strings.Join(kv, ", ")
	return fmt.Sprintf(`@unpublished{ %s }`, data)
}

//
// Parser related structures
//

// Parse a BibTeX file into appropriate structures
func Parse(buf []byte) ([]Entry, error) {
	var (
		i     int
		token *tok.Token
	)
	for {
		if len(buf) == 0 {
			break
		}
		i++
		token, buf = tok.Tok2(buf, tok.Bib)
		fmt.Printf("DEBUG i: %d, token: %s\n", i, token)
	}
	return nil, nil
}
