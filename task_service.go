package main

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

func NewCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}

func GenerateRandomArray(size int) []int {
	rand.Seed(time.Now().Unix())
	resultArray := make([]int, size)
	for i := 0; i < size; i++ {
		resultArray[i] = int(NewCryptoRand())
	}
	return resultArray
}
