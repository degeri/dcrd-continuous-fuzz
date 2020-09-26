#!/usr/bin/env bash
cd "$(dirname "$0")"

make_bins() {
    if [ -d fuzzbins/ ] 
    then
        rm -rf fuzzbins/ 
    fi
    mkdir -p output
    ./build.sh
    cp -r $HOME/src/fuzzdcrd/ fuzzbins/
    if [ -d $HOME/src/fuzzdcrd/ ] 
    then
        rm -rf $HOME/src/fuzzdcrd/ 
    fi
}

check_master() {
    sha_web=$(curl -s 'https://api.github.com/repos/decred/dcrd/commits/master' | jq -r '.sha')
    if [ "$sha_web" == "null" ]; then
        return 0
    fi
    echo "Checking master"
    if [ -f shafile ]; then
        sha_stored=`cat shafile`
        if [ "$sha_web" == "$sha_stored" ]; then
            return 0
        else
            echo "Found new master hash updating"
            make_bins
            echo $sha_web > shafile
        fi
    else
        echo $sha_web > shafile
    fi
}

check_crashes() {
    echo "Checking crashes"
    for folder in output/*/
        do
        print_string=$(($(ls "$folder/crashers" | wc -l) / 3 ))
        if [ 0 != "$print_string" ]
        then
            echo -e '\e[0;31m'$print_string" crash/es found in "$folder'\033[0m'
        fi
        done
}


if [ ! -d fuzzbins/ ] 
then
    echo "Fuzzbins dont exist making them"
    make_bins
fi

if [ -d fuzzbins/ ] 
then
    echo "fuzzer starting"
    while true
    echo "==================="
    check_master
    check_crashes
    echo "==================="
    do
    for folder in fuzzbins/*/
        do   
            echo "==================="
            (echo "Running $folder")
            (echo "Doing for $1")
            timeout --foreground $1 go-fuzz -bin=$folder/$(basename $folder)-fuzz.zip -workdir=output/$(basename $folder)
            sleep 5
            pkill -f go-fuzz
            echo "==================="
        done
    done
fi
