#
# Simple Makefile
#
build:
	go build -o bin/bibfilter cmds/bibfilter/bibfilter.go
	go build -o bin/bibmerge cmds/bibmerge/bibmerge.go
	./mk-webapp.sh
	./mk-website.sh

install:
	env GOBIN=$(HOME)/bin go install cmds/bibfilter/bibfilter.go
	env GOBIN=$(HOME)/bin go install cmds/bibmerge/bibmerge.go

test:
	go test

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f index.html ]; then rm -f *.html; fi

release:
	./mk-release.sh

