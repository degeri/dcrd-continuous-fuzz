dcrd-continuous fuzzing
===

## Install the following prerequisites



- **Go 1.13 or 1.14**

  Installation instructions can be found here: https://golang.org/doc/install.
  Ensure Go was installed properly and is a supported version:
  ```sh
  $ go version
  $ go env GOROOT GOPATH
  ```
  NOTE: `GOROOT` and `GOPATH` must not be on the same path. Since Go 1.8 (2016),
  `GOROOT` and `GOPATH` are set automatically, and you do not need to change
  them. However, you still need to add `$GOPATH/bin` to your `PATH` in order to
  run binaries installed by `go get` and `go install` (On Windows, this happens
  automatically).

  Unix example -- add these lines to .profile:

  ```
  PATH="$PATH:/usr/local/go/bin"  # main Go binaries ($GOROOT/bin)
  PATH="$PATH:$HOME/go/bin"       # installed Go projects ($GOPATH/bin)
  ```

 - **Gofuzz**

    ```go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build```

 - **Docker**

    Install: https://docs.docker.com/engine/install/. 
        
    Post Install setup docker group  https://docs.docker.com/engine/install/linux-postinstall/#manage-docker-as-a-non-root-user.

</details>

# To run 



```
chmod +x run.sh

./run.sh 10m
```

The above command runs each fuzzer (currently eleven) for 10 minutes. It is recommended to not make this value too low.

On first run the script will use docker to generate the go-fuzz bins and place them in the fuzzbins folder. 

Once one full loop of fuzzing is done the script will check dcrd master for any changes and will update the bins if needed. It also outputs alerts for any crashes found.

corpus and crashes can be found in the relative folder inside output/