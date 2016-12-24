#
# Simple Makefile
#

PROJECT = bibtex

PROG_FILES = bibfilter bibmerge

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

build:
	go build -o bin/bibfilter cmds/bibfilter/bibfilter.go
	go build -o bin/bibmerge cmds/bibmerge/bibmerge.go

install:
	env GOBIN=$(HOME)/bin go install cmds/bibfilter/bibfilter.go
	env GOBIN=$(HOME)/bin go install cmds/bibmerge/bibmerge.go

test:
	go test

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f index.html ]; then /bin/rm -f *.html; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then /bin/rm $(PROJECT)-$(VERSION)-release.zip; fi

release:
	./mk-release.bash

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash
