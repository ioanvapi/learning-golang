package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
	"strconv"
)

func main() {
	// Public and Private Key Encryption
	// create a public/private keypair
	// NOTE: Use crypto/rand not math/rand
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	}

	// log.Printf("private key: %v", privKey)

	plainText := []byte("The bourgeois human is a virus on the hard drive of the working robot!")

	// use the public key to encrypt the message
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &privKey.PublicKey, plainText)
	if err != nil {
		log.Fatalf("could not encrypt data: %v", err)
	}
	log.Printf("%s\n", strconv.Quote(string(cipherText)))

	// decrypt with the private key
	decryptedText, err := rsa.DecryptPKCS1v15(nil, privKey, cipherText)
	if err != nil {
		log.Fatalf("error decrypting cipher text: %v", err)
	}
	log.Printf("decrypt: %s\n", decryptedText)

	// Digital Signatures
	// compute the hash of our original plain text
	hash := sha256.Sum256(plainText)
	log.Printf("The hash of my message is: %#x\n", hash)

	// generate a signature using the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash[:])
	if err != nil {
		log.Fatalf("error creating signature: %v", err)
	}

	// let's not print the signature, it's big and ugly
	log.Printf("length of the signature: %v", len(signature))

	// use a public key to verify the signature for a message was created by the private key
	verify := func(pub *rsa.PublicKey, msg, signature []byte) error {
		hash := sha256.Sum256(msg)
		return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hash[:], signature)
	}

	log.Println(verify(&privKey.PublicKey, plainText, []byte("a bad signature")))

	log.Println(verify(&privKey.PublicKey, []byte("a different plain text"), signature))

	log.Println(verify(&privKey.PublicKey, plainText, signature))
}
