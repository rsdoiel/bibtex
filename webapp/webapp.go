package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/rsdoiel/bibtex/webapp/bibfilter"
)

func main() {
	js.Global.Set("bibfilter", map[string]interface{}{
		"New": bibfilter.New,
	})
}
