#Install dcrd as per instructions. 
git clone https://github.com/decred/dcrd $HOME/src/dcrd
(cd $HOME/src/dcrd && go install . ./...)

go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

for folder in /root/fuzzers/*/
do 
    (cd "$folder" && go-fuzz-build)
done