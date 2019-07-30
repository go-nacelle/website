#!/usr/bin/env bash

declare -A repos
repos["nacelle"]="./content/docs/core/overview.md"
repos["config"]="./content/docs/core/config.md"
repos["log"]="./content/docs/core/log.md"
repos["process"]="./content/docs/core/process.md"
repos["service"]="./content/docs/core/service.md"
repos["grpcbase"]="./content/docs/base processes/grpcbase.md"
repos["httpbase"]="./content/docs/base processes/httpbase.md"
repos["lambdabase"]="./content/docs/base processes/lambdabase.md"
repos["workerbase"]="./content/docs/base processes/workerbase.md"
repos["awsutil"]="./content/docs/libraries/awsutil.md"
repos["pgutil"]="./content/docs/libraries/pgutil.md"
repos["chevron"]="./content/docs/frameworks/chevron.md"
repos["scarf"]="./content/docs/frameworks/scarf.md"

function add_content() {
    echo "Updating $2"
    header=`cat "${2}" | sed '/<!-- Fold -->/q'`
    content=`curl -s "https://raw.githubusercontent.com/go-nacelle/${1}/master/README.md" | sed '1,/---/d'`
    echo -e "${header}\n${content}" > "${2}"
}

repo=${1:-"all"}
target=${repos[$repo]}

if [ $repo == "all" ]; then
    for repo in ${!repos[@]}; do
        add_content $repo "${repos[$repo]}"
    done
else
    if [ -z "$target" ]; then
        echo "Unknown repo $repo"
        exit 1
    fi

    add_content $repo "$target"
fi
