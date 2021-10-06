package Sig

import (
	"log"

	"github.com/xlcetc/cryptogm/sm/sm9"
)

//Verify
func Verify(id string, message []byte, hid byte, mk *sm9.MasterKey, sign *sm9.Sm9Sig) bool {
	if !sm9.Verify(sign, message, []byte(id), hid, &mk.MasterPubKey) {
		log.Fatal("Verify invaild")
		return false
	}
	return true
}
