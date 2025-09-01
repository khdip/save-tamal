Protocol Buffer Compiler Installation: https://protobuf.dev/installation/
```
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v32.0/protoc-32.0-osx-x86_64.zip
```
```
unzip protoc-32.0-osx-x86_64.zip -d $HOME/.local
```
Update environment’s path variable to include the path to the protoc executable. PATH="$PATH:$HOME/.local/bin"
```
open ~/.zshrc
source ~/.zshrc
```
To install protocol buffers add-in:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
To generate the pb.go file:
```
protoc -I=./proto/users --go_out=. ./proto/users/users.proto 
```
To generate both pb.go and grpc.pb.go file:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./proto/users/users.proto 
```
To run the migration:
```
cd tamal
go run migrations/migrate.go up
```
To run the server:
```
go run tamal/main.go
```
To run the cms:
```
go run cms/main.go
```
