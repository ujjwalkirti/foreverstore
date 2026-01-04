package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: defaultPathTransformFunc,
	}

	store := NewStore(opts)

	data := bytes.NewReader([]byte("Some jpg bytes"))

	if err := store.writeStream("bin/myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
