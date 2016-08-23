#!/bin/bash

function softwareCheck() {
    for CMD in $@; do
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

    echo "Rendering $html"
    mkpage \
        "pageTitle=markdown:$pageTitle" \
        "pageContent=$content" \
        "copyright=text:copyright &copy; $YEAR, R. S. Doiel, all rights reserved" \
        "nav=$nav" \
        page.tmpl > $html
}

echo "Checking software"
softwareCheck mkpage
echo "Generating index.html"
mkPage nav.md "Experimental prototype BibTeX tools" README.md index.html
echo "Generating license.html"
mkPage nav.md "Experimental prototype BibTeX tools" "markdown:$(cat LICENSE)" license.html
echo "Generating install.html"
mkPage nav.md "Experimental prototype BibTeX tools" INSTALL.md install.html

