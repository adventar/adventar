# How to install protobuf
* I tried to install these packages on Ubuntu 18.04 LTS version.
## Download and install protoc(protobuf compiler)
* You can download relaease version of protobuf [Here](https://github.com/protocolbuffers/protobuf/releases) `https://github.com/protocolbuffers/protobuf/releases`
* If you use Linux, you can use wget for download file.
* You can get information about how to install this release files from [Here](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)
* But, to save your time, I summarized the install commands. 
```
./configure
make
make check
sudo make install
sudo ldconfig
```

## Download and install golang-protobuf(for protobuf to go-lang complier)
* You can download Go support protobuf compiler from [Here](https://github.com/golang/protobuf)
* Read instructions and install golang-protobuf.
* We need to install normal protobuf firstly. If you didn't install normal protoc, please install by following above instrunction.
* You will need to install `golang compiler` to install `golang/protobuf`.
* You can install golang complier with this command
```sudo apt-get install golang-go```
* If this command doesn't work well, you can get more information [Here](https://github.com/golang/go/wiki/Ubuntu)

* And please add these code under the `~/.bashrc`
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```

## To execute proto.sh
* You need install `protoc-gen-grpc-web` to execute proto.sh on protobuf directory.
* You can download binary file from [Here](https://github.com/grpc/grpc-web/releases)
* You can use these commands to enroll this binary file to your ubuntu command.
```
$ sudo mv ~/Downloads/protoc-gen-grpc-web-1.0.7-darwin-x86_64 \  # This file name can be changed.
  /usr/local/bin/protoc-gen-grpc-web
$ chmod +x /usr/local/bin/protoc-gen-grpc-web
```

## Download gsed package for hot-fix on updated `*.proto` file
* You can download GNU sed from [Here](http://ftp.gnu.org/gnu/sed/)
* Download the latest version of GNU sed with `wget`
* Build with these commands
```
./configure
make
make check
sudo make install
sudo ldconfig
```

* But, gsed isn't work..


