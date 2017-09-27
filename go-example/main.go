package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	priv, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Generate Keys Error: %s\n", err)
		return
	}

	fmt.Printf("D: %v \n", priv.D)
	fmt.Printf("E: %v \n", priv.E)
	fmt.Printf("N: %v \n", priv.N)
	fmt.Printf("Primes: %v \n", priv.Primes)

	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&priv.PublicKey,
		[]byte("Cryptography & Security Training"),
		[]byte("iTexico"),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Encryption Error: %s\n", err)
		return
	}

	// Since encryption is a randomized function, ciphertext will be different each time.
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	message, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, []byte("iTexico"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Decryption Error: %s\n", err)
		return
	}

	fmt.Printf("Message: %v\n", string(message))

}
