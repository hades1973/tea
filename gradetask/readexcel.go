package main

import (
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
)

func readNamesClassProjs(fname string) (names, class, prjs []string) {
	sheets, err := xlsx.FileToSlice(fname)
	if err != nil {
		log.Fatal(err)
	}
	beginRow := 8
	for beginRow < len(sheets[0]) {
		teachers := sheets[0][beginRow][J]
		if strings.Contains(teachers, "白凤军") == true {
			names = append(names, sheets[0][beginRow][L])
			class = append(class, sheets[0][beginRow][R])
			prjs = append(prjs, sheets[0][beginRow][B])
		}
		beginRow++
	}
	return
}
