package main

import (
	"crypto/rand"
	"fmt"
	"go/Sig"
	"log"

	"github.com/xlcetc/cryptogm/sm/sm9"
)

func main() {
	//Gen master key
	mk, err := sm9.MasterKeyGen(rand.Reader)
	if err != nil {
		log.Fatal(err)
		return
	}
	var hid byte = 1
	message := []byte("hello")
	sign, _ := Sig.Signature("alice", message, hid, mk)
	if !Sig.Verify("alice", message, hid, mk, sign) {
		fmt.Println("failed")
	}
	fmt.Println("success")
}
