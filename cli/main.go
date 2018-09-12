package main

import "fmt"
import "flag"
import "crypto/rand"
import "encoding/hex"
import "amcl/BN254"

func mpin254() {

	var sha = BN254.HASH_TYPE

	fmt.Print(sha)
}

func main() {
	createMasterKeyPtr := flag.Bool("createMasterKey", false, "Create a new master key")

	flag.Parse()

	if *createMasterKeyPtr {
		mpin254()
	} else {
		fmt.Println("Gonna Make a Random Number")
		mk := make([]byte, 32)
		rand.Read(mk)
		mkStr := hex.EncodeToString(mk)

		fmt.Println(string(mkStr))
	}

}
