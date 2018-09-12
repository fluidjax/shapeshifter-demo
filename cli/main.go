package main

import "fmt"
import "flag"
import "crypto/rand"
import "encoding/hex"
import "amcl"
import "amcl/BN254"
import "sync"
import "encoding/json"
import "io/ioutil"

var lock sync.Mutex

//Config is used to store persistent keys in /shapeshifter/config.json
type Config struct {
	MasterSecret [32]byte  `json:masterSecret`
	ServerSecret [128]byte `json:serverSecret`
}

func printBinary(array []byte) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("%02x", array[i])
	}
	fmt.Printf("\n")
}

func mpin254(rng *amcl.RAND) {

	const MGS = BN254.MGS
	var S [MGS]byte

	/* SST is needed to generate the server secret */
	const MFS = BN254.MFS
	const G2S = 4 * MFS /* Group 2 Size */
	var SST [G2S]byte

	/* Trusted Authority (Master Secret) set-up */
	fmt.Printf("\nTesting MPIN\n")
	BN254.MPIN_RANDOM_GENERATE(rng, S[:])
	// fmt.Printf("Master Secret s: 0x")
	// printBinary(S[:])

	var serverVals Config
	serverVals.MasterSecret = S

	/* TODO: This will have its own cli flag */

	BN254.MPIN_GET_SERVER_SECRET(S[:], SST[:])
	// fmt.Printf("Server Secret SS: 0x")
	// printBinary(SST[:])

	serverVals.ServerSecret = SST

	configJSON, _ := json.MarshalIndent(serverVals, "", "")
	// fmt.Println(string(configJSON))

	// writing json to file
	// TODO: Need to write this somewhere more central
	err := ioutil.WriteFile("../config.json", configJSON, 0644)

	if err != nil{
		fmt.Println(err)}
}

func main() {

	//ToDo: Add some security to this
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
