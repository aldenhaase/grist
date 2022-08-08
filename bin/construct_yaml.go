package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

func main() {
	path := "/workspace/src/server/go-app/"
	f, _ := ioutil.ReadFile(path + os.Args[1])
	secret := os.Getenv("SERVER_SIG")
	output := bytes.Replace(f, []byte("ServerSignature"), []byte(secret), -1)
	ioutil.WriteFile(path+"app.yaml", output, 0666)
}
