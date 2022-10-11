package common

import "crypto/rand"

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func RandomBytes(length int) ([]byte, error) {
	var (
		data   = make([]byte, length)
		_, err = rand.Read(data)
	)
	return data, err
}

func RandomBytesF(length int) []byte {
	bytes, err := RandomBytes(length)
	PanicOnError(err)
	return bytes
}
