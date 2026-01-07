package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "Jai Shree Krishna"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "a7442b7f32617feec0b2dfee87d7d03d2fdee163ba7db0de63e6a63b30210217"
	expectedPathName := "a7442b/7f3261/7feec0/b2dfee/87d7d0/3d2fde/e163ba/7db0de/63e6a6/3b3021/0217"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Original != expectedOriginalKey {
		t.Errorf("have %s want %s", pathKey.Original, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)

	data := bytes.NewReader([]byte("Some jpg bytes"))

	if err := store.writeStream("bin/myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
