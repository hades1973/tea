\documentclass[12pt]{article}
\usepackage{ctex}
\usepackage{lastpage}
\usepackage{layout}
\usepackage{color}
\usepackage{placeins}
\usepackage{ulem}
\usepackage{titlesec}
\usepackage{multirow}
\usepackage{graphicx}
\usepackage{float}
\usepackage{colortbl}
\usepackage{listings}
\usepackage{indentfirst}
\usepackage{fancyhdr}
\usepackage{amsmath,amssymb}
\usepackage{tikz}
\usepackage{tabularx}
\usepackage{makecell}
\usepackage{booktabs}
\usepackage{multirow}
\usepackage[a4paper,hmargin={2.18cm,2.18cm},vmargin={2.54cm, 2.54cm}]{geometry}
\usepackage{wallpaper}

%------------------------------设置字体大小------------------------%  
\newcommand{\chuhao}{\fontsize{42pt}{\baselineskip}\selectfont}     %初号  
\newcommand{\xiaochuhao}{\fontsize{36pt}{\baselineskip}\selectfont} %小初号  
\newcommand{\yihao}{\fontsize{28pt}{\baselineskip}\selectfont}      %一号  
\newcommand{\erhao}{\fontsize{21pt}{\baselineskip}\selectfont}      %二号  
\newcommand{\xiaoerhao}{\fontsize{18pt}{\baselineskip}\selectfont}  %小二号  
\newcommand{\sanhao}{\fontsize{15.75pt}{\baselineskip}\selectfont}  %三号  
\newcommand{\sihao}{\fontsize{14pt}{\baselineskip}\selectfont}%     四号  
\newcommand{\xiaosihao}{\fontsize{12pt}{\baselineskip}\selectfont}  %小四号  
\newcommand{\wuhao}{\fontsize{10.5pt}{\baselineskip}\selectfont}    %五号  
\newcommand{\xiaowuhao}{\fontsize{9pt}{\baselineskip}\selectfont}   %小五号  
\newcommand{\liuhao}{\fontsize{7.875pt}{\baselineskip}\selectfont}  %六号  
\newcommand{\qihao}{\fontsize{5.25pt}{\baselineskip}\selectfont}    %七号  
%-------------------------------------------------------------------------
\newcommand{\blank}[1]{\makebox[2.7cm][l]{\uline{\hfill #1 \hfill}}}
\newcommand{\blankk}[1]{\makebox[3cm][l]{\uline{\hfill #1 \hfill}}}

%------------------------------表格列简写---------------------------%  
\newcommand{\firstC}[1]{\parbox[c]{.5cm}{#1}}
\newcommand{\secondC}[1]{\parbox[c]{2.0cm}{#1}}
\newcommand{\thirdC}[1]{\parbox[c]{10cm}{#1}}
%-------------------------------------------------------------------%

\pagestyle{empty}

\begin{document}
	\xiaosihao \songti
	\centering{\yihao \songti 福{\quad}建{\quad}工{\quad}程{\quad}学{\quad}院}\vspace{0.3cm}

	\centering{\xiaoerhao \songti 本科毕业（论文）答辩委员会决议书} \vspace{0.5cm}

	\begin{tabular}{@{}p{16.5cm}@{}}
		\CJKunderline{土木工程}学院~\CJKunderline{土木工程}专业~\CJKunderline{~{{.Year}}~}届~\CJKunderline{~{{.Class}}~}班学生\blank{~{{.StudentName}}~}学号\blank{~{{.StudentNumber}}~}\vspace{0.2cm}
	\end{tabular}
\renewcommand\arraystretch{1.3}
\begin{tabular}{@{}|p{5cm}<{\centering}|@{}p{11.5cm}<{\centering}|@{}}
	\hline
	毕业设计（论文）题目 & {{.DesignName}} \\
	\hline
	答辩成绩评定 &
	\begin{tabular}{@{}p{3.5cm}|@{}p{4cm}|@{}p{3.8cm}@{}}
		\makecell{报告、讲解\\(满分40分)} &  & \makecell{答辩成绩\\(百分制)} \\ 
		\hline
		\makecell{答辩情况\\(满分50分)} &  & \multirowcell{2}{~{{.DaBianChengJi}}~} \\
		\cline{1-2}
		\makecell{创新\\(满分10分)} &  \\
	\end{tabular}\\
	\hline
	指导教师评定成绩 &  {{.ZhiDaoChengJi}}\\
	\hline
	评阅教师评定成绩 &  {{.PingYueChengJi}}\\
	\hline
	\parbox{6em}{~~答辩委员会\\综合评定成绩} & {{.ZongHeChengJi}}\\
\end{tabular}

\begin{tabular}{@{}|p{2cm}@{}|p{14.5cm}<{\centering}|@{}}
	\hline
	\multirowcell{4}{答辩\\委员会\\评语} 
	& \\
	& 	\parbox[l][5cm][t]{12.5cm} {~{{.DaBianWeiYuanHuiPingYu}}~} \\
	&	\hfill 答辩委员会主任签名：\makebox[2cm][l]{} \\
	&	\hfill 年 \quad 月 \quad 日 \makebox[1cm][l]{} \\
	\hline
\end{tabular}
\renewcommand\arraystretch{1}
\begin{tabular}{@{}p{16.5cm}@{}}
	{\wuhao 备注：毕业设计（论文）综合评定成绩=指导教师成绩*40\%+评阅教师成绩*20\%+答辩小组成绩*40\%，最终以百分制记分。}
\end{tabular}

\end{document}                  %
