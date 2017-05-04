// Copyright (c) 2017 Hank Donnay. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type keyifyRes struct {
	// {"start":2641,"end":2656,"replacement":"dataWriter{\n\tWin: win,\n}"}
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Replacement string `json:"replacement"`
}

// keyify turns unkeyed struct literals into keyed ones.
func keyify(s selection, args []string) {
	js := runWithStdin(archive(s), "keyify", "-modified", "-json", s.pos())
	r := keyifyRes{}
	if err := json.NewDecoder(bytes.NewBufferString(js)).Decode(&r); err != nil {
		log.Fatal(err)
	}
	start, end := bytesToRunes(archive(s), r.Start), bytesToRunes(archive(s), r.End)
	if err := s.win.Addr("#%d,#%d", start, end); err != nil {
		log.Fatal(err)
	}
	if _, err := s.win.Write("data", []byte(r.Replacement)); err != nil {
		log.Fatal(err)
	}

	reloadShowAddr(s.win, s.start)
}

// bytesToRunes counts the number of runes n bytes into rd.
func bytesToRunes(rd io.Reader, n int) int {
	r := bufio.NewReader(rd)
	nr := 0
	for ; n > 0; nr++ {
		_, w, err := r.ReadRune()
		if err != nil {
			panic(err)
		}
		n -= w
	}
	return nr
}
