%%input macros;
outputtemplate:="%j-%c.mps";
verbatimtex
%&latex
\documentclass{article}
%%\usepackage{ctex}
\usepackage{CJKutf8}
\begin{CJK}{UTF8}{gbsn}
\begin{document}
etex

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%% fig1. 3x9 网格
beginfig(1);
	numeric u, v;
	u=1mm;
	pair A, B;
{{$space := .Space}}
{{range $j, $CS := .Groups}}
{{range $i, $name := $CS}}
	A:= ({{$j}}*30u, {{$i}}*{{$space}}u);
	B:= ({{$j}}*30u+28u, {{$i}}*{{$space}}u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex {{$name}} etex, .5[A,B]);
{{end}}
{{end}}

endfig;

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
verbatimtex
\end{CJK}
\end{document}
etex end
