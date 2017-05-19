##Generate the protocol buffer
###Step 1:
cd $GOPATH\src

###Step 2: 
protoc --proto_path=$GOPATH\src --go_out=. $GOPATH\src\github.com\crazybits\x\blockchain\blockchain.proto