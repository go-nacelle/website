#!/bin/bash

function add_content() {
    header=`cat "${2}" | sed '/<!-- Fold -->/q'`
    content=`curl -s "https://raw.githubusercontent.com/go-nacelle/${1}/master/README.md" | sed '1,/---/d'`
    echo -e "${header}\n${content}" > "${2}"
}

add_content "nacelle" "./content/docs/core/overview.md"

for repo in config log process service; do
    add_content "${repo}" "./content/docs/core/${repo}.md"
done

for repo in grpcbase httpbase lambdabase workerbase; do
    add_content "${repo}" "./content/docs/base processes/${repo}.md"
done

for repo in awsutil pgutil; do
    add_content "${repo}" "./content/docs/libraries/${repo}.md"
done

for repo in chevron scarf; do
    add_content "${repo}" "./content/docs/frameworks/${repo}.md"
done
