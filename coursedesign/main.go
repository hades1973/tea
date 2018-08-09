package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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

const (
	SIG  = "本软件由福建工程学院白凤军开发。本软件生成指导教师成绩评定表。用法:\t \n"
	SIG2 = "程序将会生成\"x.tex\" 文件.\n"
)

type Info struct {
	StudyYear, Term           string
	StudyYearAndTerm          string
	Major, Speciality         string
	Class, Student, StudentId string
	Teacher                   string
	B, L, Qk                  float64
}

type Parm struct {
	B, L, Qk float64
}

func MakeParms() []Parm {
	parms := make([]Parm, 0)
	bs := []float64{6.3, 6.6, 6.9, 7.2, 7.5, 7.8}
	for _, b := range bs {
		ls := []float64{b - 0.3, b}
		for _, l := range ls {
			qs := []float64{2, 3, 4, 5, 6, 8}
			for _, q := range qs {
				parms = append(parms, Parm{B: b, L: l, Qk: q})
			}
		}
	}
	return parms
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(SIG+"%s <inputtable>.xlsx \n"+SIG2, filepath.Base(os.Args[0]))
		return
	}

	fileName := os.Args[1]
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		fmt.Printf("Can't open file: %s\n", fileName)
		return
	}

	var sheet *xlsx.Sheet = xlFile.Sheets[0]
	if sheet == nil {
		return
	}

	tmpl, err := template.ParseFiles("begin.tex.tmpl")
	if err != nil {
		return
	}
	file, err := os.Create("x.tex")
	if err != nil {
		return
	}
	defer file.Close()
	tmpl.Execute(file, nil)

	tmpl, err = template.ParseFiles("coursedesign.tex.tmpl")
	if err != nil {
		return
	}

	info := Info{}
	parms := MakeParms()
	j := 0
	for i := 5; i < len(sheet.Rows); i++ {
		//		fmt.Println(sheet.Cell(1, A).String())
		id := sheet.Cell(i, B).String()
		if id == "" {
			continue
		}
		info.StudentId = id
		info.Student = sheet.Cell(i, C).String()
		bookName := sheet.Cell(0, A).String()
		info.StudyYearAndTerm = strings.TrimSuffix(bookName, "点名册")
		teacherName := sheet.Cell(3, D).String()
		info.Teacher = strings.TrimPrefix(teacherName, "任课教师：")
		info.Major = sheet.Cell(i, D).String()
		info.Speciality = "建造与安全工程"
		info.Class = sheet.Cell(i, E).String()
		info.B = parms[j].B
		info.L = parms[j].L
		info.Qk = parms[j].Qk
		j++
		if j > 71 {
			log.Fatal("参数不足，请修改程序\n")

		}
		tmpl.Execute(file, &info)
	}

	tmpl, err = template.ParseFiles("end.tex.tmpl")
	if err != nil {
		return
	}
	tmpl.Execute(file, &info)
}
