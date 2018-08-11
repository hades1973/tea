package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	Year                   string
	Class                  string
	StudentName            string
	StudentNumber          string
	DesignName             string
	DaBianChengJi          string
	ZhiDaoChengJi          string
	PingYueChengJi         string
	ZongHeChengJi          string
	DaBianWeiYuanHuiPingYu string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(SIG+"%s <inputtable>.xlsx \n"+SIG2, filepath.Base(os.Args[0]))
		return
	}

	// Read student's info
	fileName := os.Args[1]
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Fatalf("Can't open file: %s\n", fileName)
	}

	var sheet *xlsx.Sheet = xlFile.Sheets[0]
	if sheet == nil {
		log.Fatalln(fmt.Errorf("Sheets[0] is nil"))
	}
	infos := make([]*Info, 0)
	for i := 1; i < len(sheet.Rows); i++ {
		info := Info{}
		info.StudentName = sheet.Cell(i, A).String()
		info.StudentNumber = sheet.Cell(i, B).String()
		info.Class = sheet.Cell(i, C).String()
		info.DesignName = sheet.Cell(i, D).String()
		info.DaBianChengJi = sheet.Cell(i, E).String()
		info.ZhiDaoChengJi = sheet.Cell(i, F).String()
		info.PingYueChengJi = sheet.Cell(i, G).String()
		info.ZongHeChengJi = sheet.Cell(i, H).String()
		info.DaBianWeiYuanHuiPingYu = sheet.Cell(i, I).String()
		if len(info.DaBianWeiYuanHuiPingYu) < 5 {
			info.DaBianWeiYuanHuiPingYu = `
		该生毕业设计选题具有实际工程意义，设计质量较好，设计成果完整、具有一定的工程价值。
		答辩准备充分、陈述问题清晰、回答问题较好。
			`
		}
		infos = append(infos, &info)
	}
	if len(infos) == 0 {
		log.Fatalln("No student's info in excel file!!")
	}

	// Produce tex file from template
	tmpl, err := template.ParseFiles("./exm.tpl")
	if err != nil {
		log.Fatalln(err)
	}
	for _, info := range infos {
		file, err := os.Create(info.StudentName + ".tex")
		if err != nil {
			log.Fatalln(err)
		}
		tmpl.Execute(file, &info)
		file.Close()
	}
}
