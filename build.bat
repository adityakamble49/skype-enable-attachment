go get -d ./...
go get github.com/akavel/rsrc
cd src
%GOPATH%\bin\rsrc -manifest main.exe.manifest -o main.syso
go build main.go