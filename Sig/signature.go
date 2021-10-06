package Sig

import (
	"log"

	"github.com/xlcetc/cryptogm/sm/sm9"
)

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
