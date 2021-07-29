package main

import (
	"github.com/ltqthu/gotools/pkg/vscode"
	"log"
)

//go run cmd/vscode/main.go
//go build cmd/vscode/main.go
func main() {
	code := vscode.NewVscodeService()

	binPath, err := code.MingwSrv.GetMingwBin()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(binPath)

	includePath, err := code.MingwSrv.GetMingwInclude()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(includePath)
}
