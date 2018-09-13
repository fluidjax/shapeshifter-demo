package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"amcl"
	"amcl/BN254"
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

//Config is used to store persistent keys in /shapeshifter/config.json
type Config struct {
	TimeStamp    time.Time `json:"timeStamp"`
	MasterSecret string    `json:"mastersecret"`
	ServerSecret string    `json:"serversecret"`
}

func printBinary(array []byte) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("%02x", array[i])
	}
	fmt.Printf("\n")
}

func genMasterKey() string {
	const MGS = BN254.MGS
	var S [MGS]byte
	rng := amcl.NewRAND()
	var raw [100]byte
	for i := 0; i < 100; i++ {
		raw[i] = byte(i + 1)
	}
	rng.Seed(100, raw[:])
	BN254.MPIN_RANDOM_GENERATE(rng, S[:])
	SHex := hex.EncodeToString(S[:])
	return SHex

}

func genServerKey(S string) string {

	var SSTHex string

	if S != ""{
		const MFS = BN254.MFS
		const G2S = 4 * MFS /* Group 2 Size */
		var SST [G2S]byte
		SByte, _ := hex.DecodeString(S)
		BN254.MPIN_GET_SERVER_SECRET(SByte[:], SST[:])
		SSTHex = hex.EncodeToString(SST[:])
	}else{
		fmt.Println("No Master Key Found")
	}
	return SSTHex
}

func writeConfigFile(serverVals Config) {
	serverVals.TimeStamp = time.Now()
	configJSON, _ := json.MarshalIndent(serverVals, "", "")
	err := ioutil.WriteFile("../config.json", configJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	//Read params from commandline. ToDo: Add some security to this
	createMasterKeyPtr := flag.Bool("createMasterKey", false, "Create a new master key")
	createServerKeyPtr := flag.Bool("createServerKey", false, "Create a new server key")
	flag.Parse()

	//Read serverConfig from Config file
	var serverConfig Config
	config, err := ioutil.ReadFile("../config.json")
	if err != nil {
		//ToDo: need to check for "no such file or directory"
		fmt.Println(err)
	} else {
		json.Unmarshal(config, &serverConfig)
	}

	if *createMasterKeyPtr {
		if serverConfig.MasterSecret != ""{
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("You already have a Master Secret, are you sure? Press Y to continue")
			decision, _ := reader.ReadString('\n')
			if decision == "Y\n"{
				serverConfig.MasterSecret = genMasterKey()
				writeConfigFile(serverConfig)
			}else{
				fmt.Println("A Wise Choice")
			}
		}else{
			serverConfig.MasterSecret = genMasterKey()
			writeConfigFile(serverConfig)
		}
	}

	if *createServerKeyPtr {
		if serverConfig.ServerSecret != ""{
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("You already have a Server Secret, are you sure? Press Y to continue")
			decision, _ := reader.ReadString('\n')
			if decision == "Y\n"{
				serverConfig.ServerSecret = genServerKey(serverConfig.MasterSecret)
				writeConfigFile(serverConfig)
			}else{
				fmt.Println("A Wise Choice")
			}
		}else{
			serverConfig.ServerSecret = genServerKey(serverConfig.MasterSecret)
			writeConfigFile(serverConfig)
		}
	}
}
