#!/bin/bash

function softwareCheck() {
    for CMD in shorthand; do
        APP=$(which $CMD)
        if [ "$APP" = "" ]; then
            echo "Skipping, missing $CMD"
            exit 1;
        fi 
    done
}

function mkPage () {
    nav="$1"
    pageTitle="$2"
    content="$3"
    html="$4"

    echo "Rendering $html from $content and $nav"
    shorthand \
        -e "{{pageTitle}} :markdown: $pageTitle" \
        -e "{{pageContent}} :import-markdown: $content" \
        -e "{{copyright}} :markdown: copyright &copy; $YEAR, R. S. Doiel, all rights reserved" \
        -e "{{nav}} :import-markdown: $nav" \
        page.shorthand > $html
}

softwareCheck
echo "Generating website with shorthand"
mkPage nav.md "Experimental prototype BibTeX tools" README.md index.html
mkPage nav.md "Experimental prototype BibTeX tools" LICENSE license.html

