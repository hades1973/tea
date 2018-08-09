package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type PlanItem struct {
	Date    string
	Content string
}

type DataStruct struct {
	X    map[string]string //refer widths
	Plan []PlanItem
}

var cells = map[string]string{
	"AA": `@{\vline}m{1.3cm}<{\centering}`,
	"AB": `@{\vline}m{2.0cm}<{\centering}`,
	"AC": `@{\vline}m{1.0cm}<{\centering}`,
	"AD": `@{\vline}m{2.0cm}<{\centering}`,
	"AE": `@{\vline}m{1.3cm}<{\centering}`,
	"AF": `@{\vline}m{2.0cm}<{\centering}`,
	"AG": `@{\vline}m{1.5cm}<{\centering}`,
	"AH": `@{\vline}m{2.0cm}<{\centering}`,
	"AI": `@{\vline}m{1.3cm}<{\centering}`,
	"AJ": `@{\vline}m{1.7cm}<{\centering}`,
	"BA": `\multicolumn{2}{|c|}{设计(论文)题目}`,
	"BB": `\multicolumn{8}{c|}{仓上小学结构设计}`,
	"C1": `\begin{enumerate}
		\item 进一步掌握建筑与结构设计全过程基本方法和步骤，认真考虑影响设计的各项因素。
		\item 认真处理好结构与建筑的总体与细部关系。
		\item 理解和掌握与本设计有关的设计规范和规定，并在设计中正确运用它们。
		\item 正确选择合理的结构形式、结构体系和结构布置，掌握结构设计的计算方法和基本构造要求。
		\item 熟练使用适当的结构辅助设计软件完成结构建模、计算及施工图绘制。
		\item 认真编写设计说明书、计算书，并绘制相应的设计图纸。
		\end{enumerate}
		`,
	"C2": `\begin{enumerate}
		\item 设计总说明
		\item 基础平面布置图(1:100或1:150)及基础详图(1:20)。
		\item 首层、标准层及屋面层结构平面布置图及配筋图(1:100)。
		\item 框架梁平法配筋图(1:100)。
		\item 框架柱平法配筋图(1:100)。
		\item 楼梯配筋图（1:50）
		\item 结构计算书
		\end{enumerate}`,
	"CA":  `\multicolumn{2}{|c|}{起\quad~讫\quad~日\quad~期}`,
	"CB":  `\multicolumn{7}{|c|}{工\quad~作\quad~内\quad~容}`,
	"CC":  `\multicolumn{1}{|c|}{备\quad~注}`,
	"CAH": `\multicolumn{2}{|c|}`,
	"CBH": `\multicolumn{7}{|c|}`,
	"CAC": `\multicolumn{1}{|c|}`,
	"CankaoWenxian": `\begin{enumerate}[label={[}\arabic*{]}]
		\item 《建筑结构荷载规范》GB50009-2012.
		\item 《混凝土结构设计规范》GB50010-2010.
		\item 《建筑抗震设计规范》GB50011-2010.
		\item 《建筑地基基础设计规范》GB50007-2011.
		\item 《建筑结构制图标准》GB/T50105-2010.
		\item 中国有色工程有限公司.~《混凝土结构构造手册（第5版）》[M]. 北京：中国建筑工业出版社，2016年1月.
		\item 梁兴文.~《土木工程毕业设计指导-房屋建筑工程卷》[M]. 北京：中国建筑工业出版社，2014年1月.
		\item 陈萌,于秋波.~《土木工程毕业设计指南---手算与电算实例详解》[M]. 武汉：武汉理工大学出版社，2015年11月.
		\item 林同炎,斯托特斯伯里.~《结构概念和体系》[M]. 北京：中国建筑工业出版社，1999年1月.
		\item 深圳市斯维尔科技有限公司.~《结构设计软件高级实例教程》[M]. 北京：中国建筑工业出版社，2013年10月.
		\end{enumerate}`,
	"Name":       "",
	"Class":      "",
	"Projection": "",
}

var contents = []string{
	`结构布置与选型、手算荷载`,
	`使用计算机辅助结构设计软件完成结构建模`,
	`使用计算机辅助结构设计软件完成结构内力计算、配筋设计并手算一块双向板`,
	`绘制各楼层结构平面布置施工图`,
	`绘制梁平法、柱平法施工图`,
	`利用基础辅助设计软件设计基础并绘制基础施工图，手算一个基础`,
	`手算楼梯并绘制楼梯施工图`,
	`撰写设计说明、计算书`,
	`毕业答辩`,
	`修改、完善设计成果`,
}
var Plan = make([]PlanItem, 0)

func main() {

	// 自动生成日期，填充Plan
	layout := "2006.01.02"
	oneDay, _ := time.ParseDuration("24h")
	sevenDay := oneDay * 7
	sixDay := oneDay * 6
	beginDate, _ := time.Parse(layout, "2018.04.16")
	endDate := beginDate.Add(sixDay)
	for _, content := range contents {
		Plan = append(Plan,
			PlanItem{
				Date:    fmt.Sprintf("\\multicolumn{2}{|l|}{%s-%s}", beginDate.Format(layout), endDate.Format(layout)),
				Content: fmt.Sprintf("\\multicolumn{7}{|l|}{%s}", content),
			})
		beginDate = beginDate.Add(sevenDay)
		endDate = endDate.Add(sevenDay)
	}

	// 根据参数读取excel文件，读取学生姓名、班级、工程名称
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s\n%s\n", "gradtask 毕业设计登记表.xlsx", "自行在R列添加班级名称")
	}
	names, class, prjs := readNamesClassProjs(os.Args[1])

	// 根据模板生成tex文件
	templ, err := template.ParseFiles("templ.tex")
	if err != nil {
		log.Panic(err)
	}
	data := DataStruct{
		X:    cells,
		Plan: Plan,
	}
	for i := range names {
		f, err := os.Create("out/" + names[i] + ".tex")
		if err != nil {
			log.Fatalln(err)
		}
		data.X["Name"] = names[i]
		data.X["Class"] = class[i]
		data.X["Projection"] = prjs[i]
		templ.Execute(f, &data)
		f.Close()
	}
}
