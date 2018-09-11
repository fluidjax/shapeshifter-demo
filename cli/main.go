package main

import "fmt"
import "flag"
import "crypto/rand"
import "encoding/hex"

func main() {
	createMasterKeyPtr := flag.Bool("createMasterKey", false, "Create a new master key")

	flag.Parse()

	if *createMasterKeyPtr {
		mk := make([]byte, 32)
		rand.Read(mk)
		mkStr := hex.EncodeToString(mk)

		fmt.Println(string(mkStr))
	} else {
		fmt.Println("Not Gonna Make a Master Key")
	}

}
