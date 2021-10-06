sm9-based sign && verf

import "github.com/xlcetc/cryptogm/sm/sm9"
then use func MasterKeyGen(rand io.Reader) (mk *MasterKey, err error) to get mk
such as

    mk, err := sm9.MasterKeyGen(rand.Reader)
	    if err != nil {
		    log.Fatal(err)
		    return
	    }


