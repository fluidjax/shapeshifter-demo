package main

import "fmt"
import "flag"
import "crypto/rand"
import "encoding/hex"
import "amcl"
import "amcl/BN254"

func printBinary(array []byte) {
	for i:=0;i<len(array);i++ {
		fmt.Printf("%02x", array[i])
	}
	fmt.Printf("\n")
} 

func mpin254(rng *amcl.RAND) {

	var sha = BN254.HASH_TYPE
	const MGS=BN254.MGS
	var S [MGS]byte

	fmt.Println(sha)
	/* Trusted Authority set-up */

	fmt.Printf("\nTesting MPIN\n")
	BN254.MPIN_RANDOM_GENERATE(rng,S[:])
	fmt.Printf("Master Secret s: 0x");  printBinary(S[:])
}

func main() {
	createMasterKeyPtr := flag.Bool("createMasterKey", false, "Create a new master key")

	flag.Parse()

	if *createMasterKeyPtr {
		rng := amcl.NewRAND()
		var raw [100]byte
		for i := 0; i < 100; i++ {
			raw[i] = byte(i + 1)
		}
		rng.Seed(100, raw[:])

		mpin254(rng)

	} else {
		fmt.Println("Gonna Make a Random Number")
		mk := make([]byte, 32)
		rand.Read(mk)
		mkStr := hex.EncodeToString(mk)

		fmt.Println(string(mkStr))
	}

}
