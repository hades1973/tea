package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	SIG = `本软件由福建工程学院白凤军开发，用于打印名册。
	使用：prtnamelist -i 学生名册 -s 行间距(mm单位)
		mpost -tex=latex fig.mp
		pdflatex namelist.tex
	`
)

type Data struct {
	Groups [][]string
	Space  float64
}

var ifiles = flag.String("i", "stu.xlsx", "学生名册")
var space = flag.Float64("s", 6.22, "行间距")

func main() {
	if len(os.Args) < 3 {
		fmt.Println(SIG)
		return
	}
	flag.Parse()

	// 读取学生名单
	fnames := strings.Fields(*ifiles)
	names := readNames(fnames)

	data := Data{Space: *space}
	// 35人一组，计算因数与余数
	ng, mg := (len(names)-1)/35, (len(names)-1)%35
	for i := 0; i < ng; i++ {
		group := make([]string, 0)
		for j := (i+1)*35 - 1; j >= i*35; j-- {
			group = append(group, names[j])
		}
		data.Groups = append(data.Groups, group)
	}
	group := make([]string, 0)
	for j := len(names) - 1; j >= len(names)-mg; j-- {
		group = append(group, names[j])
	}
	data.Groups = append(data.Groups, group)

	// 利用模板生成fig.mp文件
	fout1, err := os.Create("fig.mp")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles("fig.tmpl.mp")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(fout1, &data)
	fout1.Close()
}
