package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func ErrQuit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func markdownToLatex(s string) string {
	const imageFormat = "\\begin{center} \\resizebox{%s pt}{!}{\\includegraphics{%s}} \\end{center}\n"
	result := ""
	s = strings.Replace(s, "_blank", "{\\blank}", -1)
	root := Parse(s)
	for node := root.FirstChild; node != nil; node = node.NextSibling {
		if node.Type == NodePara {
			result += fmt.Sprintf("%s\n", node.Data)
		} else if node.Type == NodeImag {
			result += fmt.Sprintf(imageFormat, node.Attr, node.Data[strings.Index(node.Data, "/")+1:])
		} else if node.Type == NodeOl {
			result += "\\begin{enumerate}"
			for nodeli := node.FirstChild; nodeli != nil; nodeli = nodeli.NextSibling {
				result += fmt.Sprintf("\\item")
				for p := nodeli.FirstChild; p != nil; p = p.NextSibling {
					if p.Type == NodePara {
						result += fmt.Sprintf("%s\n", p.Data)
					} else if p.Type == NodeImag {
						result += fmt.Sprintf(imageFormat,
							p.Attr,
							p.Data[strings.Index(p.Data, "/")+1:])
					}

				}
			}
			result += "\\end{enumerate}"
		}
	}
	return result
}

func ProduceTestPaper(info *TestInfoForm) (qfileName, afileName string) {
	tQuestion, err := template.ParseFiles("./question.tex.template")
	ErrQuit(err)
	tAnswer, err := template.ParseFiles("./answer.tex.template")
	ErrQuit(err)

	// 使用模板生成试卷的非试题部分
	qfileName = path.Join("./out/", info.PaperFileName+"Q.tex")
	afileName = path.Join("./out/", info.PaperFileName+"A.tex")
	qfile, err := os.Create(qfileName)
	ErrQuit(err)
	afile, err := os.Create(afileName)
	ErrQuit(err)
	tQuestion.Execute(qfile, &info)
	tAnswer.Execute(afile, &info)

	// 逐个生成题目
	//base := DBConfigByName(info.DatabaseName)
	//fmt.Println(base.Path)
	//db, err := sql.Open("sqlite3", base.Path)
	db, err := OpenDB(info.DatabaseName)
	ErrQuit(err)
	qtypes := []string{"blank", "select", "judge", "caculate", "report"}
	fmt.Errorf("%v\n", info.Questions)
	for _, qtype := range qtypes {
		item, ok := info.Questions[qtype]
		if !ok {
			continue
		}
		fmt.Fprintf(qfile, "\\Problem{%s}\n", item.QuestNote)
		fmt.Fprintf(afile, "\\Answer{%s}\n", item.QuestNote)
		if qtype == "选择题" {
			//fmt.Println("选择题")
			fmt.Fprintf(qfile, "\\renewcommand{\\labelenumii}{\\Alph{enumii}}\n")
		} else {
			fmt.Fprintf(qfile, "\\renewcommand{\\labelenumii}{\\arabic{enumi}.\\arabic{enumii}}\n")
		}

		fmt.Fprintf(qfile, "\\begin{enumerate}")
		fmt.Fprintf(afile, "\\begin{enumerate}")

		query := fmt.Sprintf("SELECT question, answer FROM '%s' where id=?", qtype)
		question, answer := "", ""
		for _, n := range item.Questions {
			err := db.QueryRow(query, n).Scan(&question, &answer)
			if err != nil {
				fmt.Errorf("%v\n", err)
				continue
			}
			fmt.Fprintf(qfile, "\\item{%s}", markdownToLatex(question))
			fmt.Fprintf(afile, "\\item{%s}", answer)
		}
		fmt.Fprintf(qfile, "%s\n", "\\end{enumerate}")
		fmt.Fprintf(afile, "%s\n", "\\end{enumerate}")

	}

	fmt.Fprintf(qfile, "%s\n", "\\end{document}")
	fmt.Fprintf(afile, "%s\n", "\\end{document}")
	db.Close()

	qfile.Close()
	afile.Close()

	return qfileName, afileName
}
