package main

import (
	"encoding/json"
	"fmt"

	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Tx is the transactional information between two parties.
type Tx struct {
	FromID string `json:"from"`
	ToID   string `json:"to"`
	Value  uint64 `json:"value"`
}

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	privateKey, err := crypto.LoadECDSA("zblock/accounts/kennedy.ecdsa")
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	tx := Tx{
		FromID: "Bill",
		ToID:   "Bob",
		Value:  100,
	}

	data, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("unable to marshal: %w", err)
	}

	v := crypto.Keccak256(data)

	sig, err := crypto.Sign(v, privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign: %w", err)
	}

	fmt.Println("SIG:", hexutil.Encode(sig))

	// ========================================================
	// OVER THE WIRE

	publicKey, err := crypto.SigToPub(v, sig)
	if err != nil {
		return fmt.Errorf("unable to get public key: %w", err)
	}

	fmt.Println("PUB:", crypto.PubkeyToAddress(*publicKey).String())

	// ========================================================

	return nil
}
