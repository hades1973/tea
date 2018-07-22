package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//var questSeps = []string{".c.", ".s.", ".b."}
var questHandle = map[string]func(string){
	".c.": caculateHandle,
	".s.": selectHandle,
	".b.": blankHandle,
}

func caculateHandle(par string) {
	fmt.Println("\\item{")
	idx := strings.Index(par, " ")
	fmt.Print(par[idx:])
	fmt.Println("}")
}

func selectHandle(par string) {
	lines := strings.Split(par, "\n")
	idx := strings.Index(lines[0], " ")
	fmt.Printf("\\item{%s\n", lines[0][idx:])
	//fmt.Printf("\\xzC{%s", lines[0][idx:])
	firstItem := true
	for _, text := range lines[1:] {
		cs := text
		if !isItem(cs) { // 题干
			fmt.Printf("%s\n", cs)
		} else if firstItem { // 第一个题支
			firstItem = false
			fmt.Printf("\\xzB{\\xx~%s\n}", cs[2:])
		} else { // 非第一个题支
			fmt.Printf("{\\xx~%s\n}", cs[2:])
		}
	}
	fmt.Println("}")
}

func blankHandle(par string) {
}

// 用来根据文本文件生成tex格式
// 方便写题目。题目放在一个文件
// 按照：
// 1. 题目
// 2. 题目
// 来写。
// 选择题格式如下,num表示题号
//s.1. 题干
//A. 题支
//B. 题支
//C. 题支
//D.题支
// 然后使用命令mathtest inputfile > outputfile 生成tex文件

// NewParScanner 将输入reader转换为bufio.Scanner，用来对输入进行分段
func NewParScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(
		// 分段函数.
		func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}

			if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
				if len(data) > 0 && data[len(data)-1] == '\r' {
					data = data[0 : len(data)-1]
				}
				// We have a full newline-terminated line.
				return i + 2, (data[0:i]), nil
			}
			// If we're at EOF, we have a final, non-terminated line. Return it.
			if atEOF {
				return len(data), data, nil
			}
			// Request more data.
			return 0, nil, nil
		})
	return scanner
}

func main() {
	// 打开文件
	if len(os.Args) != 2 {
		fmt.Println("usage: doquest 文件名")
	}
	//	fname := os.Args[1]

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//  打印文件头
	fmt.Println(head)

	// 逐行扫描，识别并打印题目
	scanner := NewParScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		parseQuestion(text)
	}

	// 打印文件尾
	fmt.Println(tail)
}

// 用来根据文本文件生成tex格式
// 方便写题目。题目放在一个文件
// 按照：
// 1. 题目
// 2. 题目
// 来写。
// 选择题格式如下,num表示题号
//s.1. 题干，可以有分行符
//A. 题支，一个题支一行
//B. 题支
//C. 题支
//D.题支
// 然后使用命令mathtest inputfile > outputfile 生成tex文件
func isItem(cs string) bool { // 判断是否是题支，题支形如：A. xxxx, 有效字符至少四个
	if len(cs) >= 4 { // 题支需要进一步处理
		if cs[0] == 'A' || cs[0] == 'B' || cs[0] == 'C' || cs[0] == 'D' {
			if cs[1] == '.' {
				return true
			}
		}
	}
	return false
}

// 将图替换为tex格式，然后交由questionHandle处理
func parseQuestion(par string) {
	// 将图替换为tex格式
	buf := new(bytes.Buffer)
	lines := strings.Split(par, "\n")
	for _, text := range lines {
		if len(text) >= 5 && text[:5] == ".fig." { //图，单独占一行，格式：.fig. 图文件名
			figname := strings.Fields(text)[1]
			fmt.Fprintf(buf, "\\begin{center}\n\\scalebox{1.0}{\\includegraphics{./fig/%s}}\n\\end{center}\n", figname)
		} else {
			fmt.Fprintln(buf, text)
		}
	}
	par = buf.String()

	// 处理试题
	if len(par) > 3 {
		fn, ok := questHandle[par[:3]]
		if ok {
			fn(par)
		}
	}
}
