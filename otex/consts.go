// Package main provides ...
package main

const head = `
\documentclass[12pt]{article}
\usepackage{ctex}
\usepackage{lastpage}
\usepackage{color}
\usepackage{placeins}
\usepackage{ulem}
\usepackage{titlesec}
\usepackage{graphicx}
\usepackage{colortbl}
\usepackage{listings}
\usepackage{amsmath}%
%\usepackage{MnSymbol}%
%\usepackage{wasysym}%
\usepackage{amssymb}
\usepackage{indentfirst}
\usepackage{fancyhdr}
\usepackage{tikz}
%%\usepackage[a4paper,hmargin={5cm,2cm},vmargin={2cm, 2cm},landscape]{geometry}
\usepackage{wallpaper}


%%\columnsep=2cm
%%\marginparsep=3.5cm

\newcommand{\blank}{\uline{\rule{3cm}{0pt}}}
\newcommand{\tcir}[1]{\text{\textcircled{#1}}}

\renewcommand{\headrulewidth}{0pt}
\pagestyle{fancy}

\fancyhf{}
\fancyfoot[CO,CE]{\vspace*{1mm}第\,\thepage\,页 , 共 ~\pageref{LastPage} 页}

\providecommand{\CJKnumber}[1]{
    \ifcase#1\or{一}\or{二}\or{三}\or{四}\or{五}\or{六}\or{七}\or{八}\or{九}\or{十}
    \fi
}

\newcounter{Problem}
\renewcommand{\theProblem} {
    \CJKnumber{\arabic{Problem}}
}

\newcommand{\Problem}[1] {
     \par{\vspace{1em}}
     \noindent \refstepcounter{Problem} \textbf{\large \theProblem \hspace{0.35em} #1} 
     \par{\vspace{1em}}
}

\renewcommand{\labelenumii}{\arabic{enumi}.\arabic{enumii}}

    \newboolean{showanswer}
    
    \newcounter{xzno}
    \newcounter{xx}
    \newcounter{ctemp}
    
    \newlength{\WidthQuestion} % 选项行最大文字宽度，默认为页面宽度
    \newlength{\lengtha} % 依次为A B C D选项长度
    \newlength{\lengthb}
    \newlength{\lengthc}
    \newlength{\lengthd}
    \newlength{\maxlength} % maxlength{A,B,C,D}
    \newlength{\lengthhalf}
    \newlength{\lengthquart}
    \newlength{\xxjg} % 两个选项间最小间隔
    \newlength{\LengthIndent} % 选项行的悬挂长度
    \newlength{\lengthtemp}
    
    \newcommand{\showanswer}{\setboolean{showanswer}{true}}
    \newcommand{\notshowanswer}{\setboolean{showanswer}{false}}
    \newcommand{\setxxsep}[1]{\setlength{\xxjg}{#1}}
    \newcommand{\setxxindent}[1]{\setlength{\LengthIndent}{#1}}
    
    \newcommand{\setWidthQuestion}[1]{\setlength{\WidthQuestion}{#1}}
    
    \newcommand{\zqxxyc}{\color{blue}} % 选项正确答颜色
    \newcommand{\xxyc}{\color{black}}
    
    \newcommand{\xx}[1]{\addtocounter{xx} {1} \xxyc(\Alph{xx})~#1}
    \newcommand{\zqxx}[1]{\addtocounter{xx} {1} \ifthenelse{\boolean{showanswer}} {\zqxxyc}{\xxyc}(\Alph{xx})~#1}
    
    %
    \newcommand{\maxlen}[4]{% 
    	\setcounter{ctemp}{\arabic{xx}}%
    	\settowidth{\lengtha}{#1}%
    	\settowidth{\lengthb}{#2}%
    	\settowidth{\lengthc}{#3}%
    	\settowidth{\lengthd}{#4}%
    	\setlength{\maxlength}{\lengtha}%
    	\ifthenelse{\lengthtest{\lengthb>\maxlength}}{\setlength{\maxlength}{\lengthb}}{}%
    	\ifthenelse{\lengthtest{\lengthc>\maxlength}}{\setlength{\maxlength}{\lengthc}}{}%
    	\ifthenelse{\lengthtest{\lengthd>\maxlength}}{\setlength{\maxlength}{\lengthd}}{}%
    	\addtolength{\maxlength}{\xxjg} \setcounter{xx}{\arabic{ctemp}}
    }
    
    %
    \newcommand{\xxind}{\rule{\LengthIndent}{0pt}}
    %\newcommand{\xxind}{}
    %
    
    \newcommand{\xzA}[4]{%
    	\maxlen{#1}{#2}{#3}{#4}%
    	\setlength{\lengthtemp}{\WidthQuestion}%
    	\addtolength{\lengthtemp}{-\LengthIndent}%
    	\setlength{\lengthhalf}{0.5\lengthtemp}%
    	\setlength{\lengthquart}{0.25\lengthtemp}%
    	\ifthenelse {\lengthtest{\maxlength>\lengthhalf}} %if
            {\xxind#1\\ \xxind#2\\ \xxind#3\\ \xxind#4} % then
    		{\ifthenelse{\lengthtest{\maxlength>\lengthquart}} %else-if
                 {\xxind \makebox[\lengthhalf][l]{#1} \makebox[\lengthhalf][l]{#2} \\%\\% then
                  \xxind \makebox[\lengthhalf][l]{#3}  \makebox[\lengthhalf][l]{#4}%
                 }% else-then
    	         {\xxind\makebox[\lengthquart][l]{#1}\makebox[\lengthquart][l]{#2}\makebox[\lengthquart][l]{#3}\makebox[\lengthquart][l]{#4}%
                 }%
            }% else-else
    }
    
    \newcommand{\xzB}[4]{% 试题号不变
    	%%\xxyc\par\arabic{xzno}.~#1\\%
		\par
    	\setcounter{xx}{0}%
        \xzA{#1}{#2}{#3}{#4}
    }
    
    \newcommand{\xzC}[5]{% 试题号自动增加
    	\addtocounter{xzno}{1}%
        \xzB{#1}{#2}{#3}{#4}{#5}
    	\vspace{2mm}
    }
    
    \newcommand{\kh} {\nolinebreak(\nolinebreak\qquad\nolinebreak)}
    
    \setWidthQuestion{0.8\columnwidth}
    
    \parindent=0mm
    \setxxindent{6mm}
    \setxxsep{1cm}
\begin{document}
\begin{enumerate}

`

const tail = `
\end{enumerate}
\end{document}
`
