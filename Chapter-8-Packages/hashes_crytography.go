package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {

	// cryptoHash()
	// crc32Hash()
	test_getHash()
}

func test_getHash() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	filepath1 := dir + "/Chapter-8-Packages/test1.txt"
	filepath2 := dir + "/Chapter-8-Packages/test2.txt"

	h1, err1 := getHash(filepath1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	h2, err2 := getHash(filepath2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}

func getHash(fileName string) (uint32, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("file close ok")
	}(f)

	h := crc32.NewIEEE()
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func crc32Hash() {
	hasher := crc32.NewIEEE()
	hasher.Write([]byte("test"))
	v := hasher.Sum32()
	fmt.Println(v)

	hasher2 := crc32.NewIEEE()
	hasher2.Write([]byte("test"))
	hash := hasher2.Sum32()
	fmt.Println(hash)

	fmt.Printf("%v == %v\n", v, hash)
}

func cryptoHash() {
	hash1 := sha1.New()
	hash1.Write([]byte("test"))
	bs1 := hash1.Sum([]byte{})
	fmt.Printf("%x\n", bs1)

	hash2 := sha1.New()
	hash2.Write([]byte("test"))
	bs2 := hash2.Sum([]byte{})
	fmt.Printf("%x\n", bs2)

	fmt.Printf("%x\n%x\n", bs1, bs2)
}

/*
both crc32 and sha1 implement the "hash.Hash" interface.
The main difference is that whereas crc32 computes a 32-bit hash, sha1 computes a 160-bit hash.
There is no native type to represent a 160-bit number, so we use a slice of 20 bytes instead.
*/
