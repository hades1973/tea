package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
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

const (
	SIG  = "本软件由福建工程学院白凤军开发。本软件生成课程报告任务书。用法:\t \n"
	SIG2 = "程序将会生成\"x.tex\" 文件.\n"
)

type Info struct {
	StudyYear, Term           string
	StudyYearAndTerm          string
	Major, Speciality         string
	Class, Student, StudentId string
	Teacher                   string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(SIG+"%s <inputtable>.xlsx \n"+SIG2, filepath.Base(os.Args[0]))
		return
	}

	// 获取学生信息
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

	students := make([]Info, 0)
	for i := 5; i < len(sheet.Rows); i++ {
		info := Info{}
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
		students = append(students, info)
	}

	// 将学生按班级分组
	groups := make(map[string][]string, 0)
	for _, stu := range students {
		grp := groups[stu.Class]
		grp = append(grp, stu.Student)
		groups[stu.Class] = grp
	}
	//fmt.Println("%v", groups)

	// 生成文档
	tmpl, err := template.ParseFiles("begin.tex.tmpl", "body.tex.tmpl", "end.tex.tmpl")
	if err != nil {
		log.Panic(err)
		return
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "begin.tex.tmpl", nil)
	if err != nil {
		log.Panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "body.tex.tmpl", groups)
	if err != nil {
		return
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "end.tex.tmpl", nil)
}
