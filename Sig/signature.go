package Sig

// import "github.com/xlcetc/cryptogm/sm/sm9"
// then use func MasterKeyGen(rand io.Reader) (mk *MasterKey, err error) to get mk
// such as
// mk, err := sm9.MasterKeyGen(rand.Reader)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

import (
	"log"

	"github.com/xlcetc/cryptogm/sm/sm9"
)

//@ import "github.com/xlcetc/cryptogm/sm/sm9"
//@ then use func MasterKeyGen(rand io.Reader) (mk *MasterKey, err error) to get mk
//@ such as main.go

//Signature with sm9 to get a sign
func Signature(id string, message []byte, hid byte, mk *sm9.MasterKey) (*sm9.Sm9Sig, error) {
	var uid = []byte(id)
	uk, err := sm9.UserKeyGen(mk, uid, hid)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	sig, err := sm9.Sign(uk, &mk.MasterPubKey, message)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return sig, nil
}
