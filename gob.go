package main

import (
	"bytes"
	_ "embed"
	"encoding/gob"
	"fmt"
)

var (
	// File value.gob contains some complicated data which we have precomputed and saved.
	//go:embed value.gob
	b []byte
	s = func() (s struct{
		Number float64
		Weather string
		Alphabet []string
	}) {
		dec := gob.NewDecoder(bytes.NewReader(b))
		if err := dec.Decode(&s); err != nil {
			panic(err)
		}
		return
	}()
)

func main()  {
	fmt.Printf("s: %+v\n", s)
}