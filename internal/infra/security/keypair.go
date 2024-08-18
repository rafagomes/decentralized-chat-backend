package security

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/curve25519"
	"io"
)

var Basepoint = [32]byte{9}

type Keygen struct{}

func NewKeygen() *Keygen {
	return &Keygen{}
}

func (k *Keygen) Generate() ([32]byte, [32]byte, error) {
	var privateKey, publicKey [32]byte

	if _, err := io.ReadFull(rand.Reader, privateKey[:]); err != nil {
		return privateKey, publicKey, err
	}

	pub, err := curve25519.X25519(privateKey[:], Basepoint[:])
	if err != nil {
		return privateKey, publicKey, err
	}
	copy(publicKey[:], pub)

	return privateKey, publicKey, nil
}

func (k *Keygen) GetPublicAsString(publicKey [32]byte) string {
	return hex.EncodeToString(publicKey[:])
}

func (k *Keygen) SharedSecret(privateKey, publicKey [32]byte) ([32]byte, error) {
	var sharedSecret [32]byte

	secret, err := curve25519.X25519(privateKey[:], publicKey[:])
	if err != nil {
		return sharedSecret, err
	}
	copy(sharedSecret[:], secret)

	return sharedSecret, nil
}
