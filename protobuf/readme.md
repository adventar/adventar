# How to install protobuf
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
