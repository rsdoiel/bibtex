[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Go Report Card](http://goreportcard.com/badge/rsdoiel/bibtex)](http://goreportcard.com/report/rsdoiel/bibtex)


# bibtex

A quick and dirty BibTeX package. Includes *bibfilter* which can parse and filter tag entries of a BibTeX file.

## Usage

```
 bibfilter [OPTION] BIBFILE
```

This command will output *BIBFILE* to the console applying the options specified.

## Options

 + -exclude a comma separated list of tags to exclude
 + -include a comma separated list of tags to include
 + -h display help information
 + -l display license
 + -v display version information

## Examples

Output **my.bib** without any comment entries

```
    bibfilter -exlude=comment my.bib
```

Output **my.bib** only including articles

```
    bibfilter -include=article my.bib
```

Output **my.bib** only including articles and proceedings

```
    bibfilter -include=article,inproceedings my.bib
```


## Prior art

+ [makebib.perl](http://www.snowelm.com/~t/doc/tips/makebib.perl) - Converts plain text through a series regexp rules into BibTeX
  + Perl script includes self test that maybe helpful 
+ [pybtex](https://bitbucket.org/pybtex-devs/pybtex/src/1819074df33a?at=master) - Python BibTeX processor
+ [peer2](https://github.com/njwilson23/peer2) - Golang port of a Python tool called peer, it can do basic BibTeX format parsing
+ [r59-lex.go](https://talks.golang.org/2011/lex/r59-lex.go) - Golang, simple lexer example using goroutines
    + [meling/biblexer](https://github.com/meling/biblexer) - a Golang BibTeX lexer based on Rob Pike's r59-lex.go
+ [nickng/bibtex](https://github.com/nickng/bibtex) - Golang, BibTeX parser package
+ [tmc/bibtex](https://github.com/tmc/bibtex) - Golang, BibTeX parser package
+ [sotetsuk/gobibtext](https://github.com/sotetsuk/gobibtex) - Golang, a BibTeX parser implementing the Decode, Encode function approach

## About BibTeX

+ [Wikipedia page on BibTeX](https://en.m.wikipedia.org/wiki/BibTeX) 
    + includes good description of currently used fields
+ [Bibliographies with BibTeX](https://getpocket.com/a/read/98701243)
    + article explaining practical usage
+ [bibtex.org](http://www.bibtex.org/)
    + [format](http://www.bibtex.org/Format/)
    + [usage](http://www.bibtex.org/Using/)

## Open Source Citation Tools

+ [JabRef](http://www.jabref.org/) - an open source bibliography reference manager. 
    + The native file format used by JabRef is BibTeX, the standard LaTeX bibliography format. 
+ [Zotero](https://www.zotero.org/)
    + Provides a hosted solution
    + [License](https://www.zotero.org/support/licensing) 
    + [Source Code](https://www.zotero.org/support/dev/source_code)
        + [Developer docs](https://www.zotero.org/support/dev/client_coding)

