#
# Simple Makefile
#
build:
	go build -o bin/mkbib cmds/mkbib/mkbib.go

install:
	env GOBIN=$HOME/bin go install cmds/mkbib/mkbib.go

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

release:
	./mk-release.sh

