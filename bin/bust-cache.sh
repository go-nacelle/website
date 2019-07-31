#!/bin/sh

function refresh() {
    local file=$(echo "$1" | sed 's#public/##' | sed 's#/index.html$##')
    echo "bypassing $file"
    curl --silent -H 'X-No-Cache: true' "https://nacelle.dev/$file" > /dev/null
}

find public -type f | while read file; do
    refresh "$file"
done
