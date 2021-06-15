#!/bin/bash
#Install dcrd as per instructions. 
cd "$(dirname "$0")"
if [ -d $HOME/src/fuzzdcrd/ ] 
    then
        rm -rf $HOME/src/fuzzdcrd/ 
    fi

git clone https://github.com/decred/dcrd $HOME/src/fuzzdcrd

go get -u github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest

echo "Dcrd cloned. Copying over fuzzers and compiling"

cp -r fuzzers/fuzz_* $HOME/src/fuzzdcrd/

(cd $HOME/src/fuzzdcrd/ && go get github.com/dvyukov/go-fuzz/go-fuzz-dep && go get github.com/decred/dcrd/dcrec/edwards/v2@v2.0.1)

for folder in $HOME/src/fuzzdcrd/fuzz_*/
do  
    (cd "$folder" && go-fuzz-build)
done

cd $HOME/src/fuzzdcrd/
shopt -s extglob
rm -rf !(fuzz_*)
shopt -u extglob