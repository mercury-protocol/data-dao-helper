package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func Encrypt(key []byte, data []byte) (encoded []byte, err error) {
	//Create a new AES cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(data))

	//iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	//Encrypt the data:
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], data)

	return cipherText, nil
}

func Decrypt(key []byte, encoded []byte) (decoded []byte, err error) {
	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return nil, err
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(encoded) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return nil, err
	}

	iv := encoded[:aes.BlockSize]
	encoded = encoded[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encoded, encoded)
	return encoded, nil
}
