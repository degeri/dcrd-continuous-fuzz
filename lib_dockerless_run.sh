#!/usr/bin/env bash
cd "$(dirname "$0")"

make_bins() {
    if [ -d lib_fuzzbins/ ] 
    then
        rm -rf lib_fuzzbins/ 
    fi
    mkdir -p lib_output
    ./lib_build.sh
    cp -r $HOME/src/fuzzdcrd/ lib_fuzzbins/
    for folder in lib_fuzzbins/*/
        do
            mkdir -p lib_output/$(basename $folder)
            mkdir -p lib_output/crash_$(basename $folder)
        done
}

check_master() {
    sha_web=$(curl -s 'https://api.github.com/repos/decred/dcrd/commits/master' | jq -r '.sha')
    if [ "$sha_web" == "null" ]; then
        return 0
    fi
    echo "Checking master"
    if [ -f lib_shafile ]; then
        sha_stored=`cat lib_shafile`
        if [ "$sha_web" == "$sha_stored" ]; then
            return 0
        else
            echo "Found new master hash updating"
            make_bins
            echo $sha_web > lib_shafile
        fi
    else
        echo $sha_web > lib_shafile
    fi
}


check_crashes() {
    echo "Checking crashes"
    for folder in lib_output/crash_*/
        do
        print_string=$(($(ls "$folder/" | wc -l)))
        if [ 0 != "$print_string" ]
        then
            echo -e '\e[0;31m'$print_string" crash/es found in "$folder'\033[0m'
        fi
        done
}


if [ ! -d lib_fuzzbins/ ] 
then
    echo "Fuzzbins dont exist making them"
    make_bins
fi

if [ -d lib_fuzzbins/ ] 
then
    echo "fuzzer starting"
    while true
    echo "==================="
    check_master
    check_crashes
    echo "==================="
    do
    for folder in lib_fuzzbins/*/
        do   
            echo "==================="
            (echo "Running $folder")
            (echo "Doing for $1")
            timeout --foreground $1 $folder/libbin_$(basename $folder)  -max_len=16384 lib_output/$(basename $folder)
            sleep 5
            pkill -f go-fuzz
            echo "==================="
        done
    done
fi
