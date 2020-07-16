if [ -d temp/ ] 
    then
        rm -rf temp/ 
    fi

#Install dcrd as per instructions. 
git clone https://github.com/decred/dcrd temp/dcrd
(cd temp/dcrd && go install . ./...)

go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

cp -r fuzzers/ temp/fuzzers

for folder in temp/fuzzers/*/
do 
    (cd "$folder" && go-fuzz-build)
done

cp -r temp/fuzzers/ fuzzbins/

if [ -d temp/ ] 
    then
        rm -rf temp/ 
    fi