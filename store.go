package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type PathKey struct {
	PathName string
	Original string
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha256.Sum256([]byte(key))
	hashString := hex.EncodeToString(hash[:])

	blockSize := 6
	sliceLen := len(hashString)/blockSize + 1
	paths := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		if to > len(hashString) {
			to = len(hashString)
		}
		paths[i] = hashString[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		Original: hashString,
	}
}

type PathTransformFunc func(string) PathKey

func (p PathKey) Filename() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Original)
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

// var defaultPathTransformFunc = func(key string) string {
// 	return key
// }

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathKey.PathName, os.ModePerm); err != nil {
		return err
	}

	pathAndFilename := pathKey.Filename()

	f, err := os.Create(pathAndFilename)

	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)

	if err != nil {
		return err
	}

	log.Printf("Written (%d) bytes to disk: %s", n, pathAndFilename)

	return nil

}
