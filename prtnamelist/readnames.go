package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/tealeg/xlsx"
)

const (
	A = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

// read students names from ginve fnames
func readNames(fnames []string) []string {
	var (
		xlfile *xlsx.File
		sheet  *xlsx.Sheet
		names  []string
		err    error
	)
	for j := range fnames {
		fname := strings.TrimSpace(fnames[j])
		if xlfile, err = xlsx.OpenFile(fname); err != nil {
			fmt.Println("Can't openfile: ", err)
			log.Fatal(err)
		}
		sheet = xlfile.Sheets[0]
		for n := 5; n < len(sheet.Rows); n++ {
			name := sheet.Cell(n, C).String()
			names = append(names, strings.TrimSpace(name))
		}
	}
	return names
}
