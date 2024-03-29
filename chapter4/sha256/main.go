package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("davide"))
	c2 := sha256.Sum256([]byte("Davide"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	zero(&c2)
	fmt.Printf("%x\n", c2)
}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
