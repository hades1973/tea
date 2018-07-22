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



	A:= (0*30u, 0*6.22u);
	B:= (0*30u+28u, 0*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 魏宁 etex, .5[A,B]);

	A:= (0*30u, 1*6.22u);
	B:= (0*30u+28u, 1*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 郭芳忠 etex, .5[A,B]);

	A:= (0*30u, 2*6.22u);
	B:= (0*30u+28u, 2*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 刘锦 etex, .5[A,B]);

	A:= (0*30u, 3*6.22u);
	B:= (0*30u+28u, 3*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 李昊洋 etex, .5[A,B]);

	A:= (0*30u, 4*6.22u);
	B:= (0*30u+28u, 4*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 董建辉 etex, .5[A,B]);

	A:= (0*30u, 5*6.22u);
	B:= (0*30u+28u, 5*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 李嘉荣 etex, .5[A,B]);

	A:= (0*30u, 6*6.22u);
	B:= (0*30u+28u, 6*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 许智军 etex, .5[A,B]);

	A:= (0*30u, 7*6.22u);
	B:= (0*30u+28u, 7*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 陈淅 etex, .5[A,B]);

	A:= (0*30u, 8*6.22u);
	B:= (0*30u+28u, 8*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 骆韬 etex, .5[A,B]);

	A:= (0*30u, 9*6.22u);
	B:= (0*30u+28u, 9*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 谭毅 etex, .5[A,B]);

	A:= (0*30u, 10*6.22u);
	B:= (0*30u+28u, 10*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 林新剑 etex, .5[A,B]);

	A:= (0*30u, 11*6.22u);
	B:= (0*30u+28u, 11*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 翁丽丽 etex, .5[A,B]);

	A:= (0*30u, 12*6.22u);
	B:= (0*30u+28u, 12*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 占德炜 etex, .5[A,B]);

	A:= (0*30u, 13*6.22u);
	B:= (0*30u+28u, 13*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 吕培榕 etex, .5[A,B]);

	A:= (0*30u, 14*6.22u);
	B:= (0*30u+28u, 14*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 郑君松 etex, .5[A,B]);

	A:= (0*30u, 15*6.22u);
	B:= (0*30u+28u, 15*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 沈泽林 etex, .5[A,B]);

	A:= (0*30u, 16*6.22u);
	B:= (0*30u+28u, 16*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 张晓捷 etex, .5[A,B]);

	A:= (0*30u, 17*6.22u);
	B:= (0*30u+28u, 17*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 陈凌凡 etex, .5[A,B]);

	A:= (0*30u, 18*6.22u);
	B:= (0*30u+28u, 18*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 林永銮 etex, .5[A,B]);

	A:= (0*30u, 19*6.22u);
	B:= (0*30u+28u, 19*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 赵楚楚 etex, .5[A,B]);

	A:= (0*30u, 20*6.22u);
	B:= (0*30u+28u, 20*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 闫朕 etex, .5[A,B]);

	A:= (0*30u, 21*6.22u);
	B:= (0*30u+28u, 21*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 邓万枝 etex, .5[A,B]);

	A:= (0*30u, 22*6.22u);
	B:= (0*30u+28u, 22*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 林张杰 etex, .5[A,B]);

	A:= (0*30u, 23*6.22u);
	B:= (0*30u+28u, 23*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 许泳森 etex, .5[A,B]);

	A:= (0*30u, 24*6.22u);
	B:= (0*30u+28u, 24*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 章鹏山 etex, .5[A,B]);

	A:= (0*30u, 25*6.22u);
	B:= (0*30u+28u, 25*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 陈泽川 etex, .5[A,B]);

	A:= (0*30u, 26*6.22u);
	B:= (0*30u+28u, 26*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 游朝裕 etex, .5[A,B]);

	A:= (0*30u, 27*6.22u);
	B:= (0*30u+28u, 27*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 文豪 etex, .5[A,B]);

	A:= (0*30u, 28*6.22u);
	B:= (0*30u+28u, 28*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 陈凡楠 etex, .5[A,B]);

	A:= (0*30u, 29*6.22u);
	B:= (0*30u+28u, 29*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 李广智 etex, .5[A,B]);

	A:= (0*30u, 30*6.22u);
	B:= (0*30u+28u, 30*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 庄志威 etex, .5[A,B]);

	A:= (0*30u, 31*6.22u);
	B:= (0*30u+28u, 31*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 黄俊亮 etex, .5[A,B]);

	A:= (0*30u, 32*6.22u);
	B:= (0*30u+28u, 32*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 刘蔚文 etex, .5[A,B]);

	A:= (0*30u, 33*6.22u);
	B:= (0*30u+28u, 33*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 林楠 etex, .5[A,B]);

	A:= (0*30u, 34*6.22u);
	B:= (0*30u+28u, 34*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 冯凌俊 etex, .5[A,B]);



	A:= (1*30u, 0*6.22u);
	B:= (1*30u+28u, 0*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex  etex, .5[A,B]);

	A:= (1*30u, 1*6.22u);
	B:= (1*30u+28u, 1*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex  etex, .5[A,B]);

	A:= (1*30u, 2*6.22u);
	B:= (1*30u+28u, 2*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex  etex, .5[A,B]);

	A:= (1*30u, 3*6.22u);
	B:= (1*30u+28u, 3*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex  etex, .5[A,B]);

	A:= (1*30u, 4*6.22u);
	B:= (1*30u+28u, 4*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex  etex, .5[A,B]);

	A:= (1*30u, 5*6.22u);
	B:= (1*30u+28u, 5*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 郑攀攀 etex, .5[A,B]);

	A:= (1*30u, 6*6.22u);
	B:= (1*30u+28u, 6*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 林杰 etex, .5[A,B]);

	A:= (1*30u, 7*6.22u);
	B:= (1*30u+28u, 7*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 黄翔 etex, .5[A,B]);

	A:= (1*30u, 8*6.22u);
	B:= (1*30u+28u, 8*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 张祥亿 etex, .5[A,B]);

	A:= (1*30u, 9*6.22u);
	B:= (1*30u+28u, 9*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 张锦城 etex, .5[A,B]);

	A:= (1*30u, 10*6.22u);
	B:= (1*30u+28u, 10*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 牛昊 etex, .5[A,B]);

	A:= (1*30u, 11*6.22u);
	B:= (1*30u+28u, 11*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 朱年志 etex, .5[A,B]);

	A:= (1*30u, 12*6.22u);
	B:= (1*30u+28u, 12*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 周培松 etex, .5[A,B]);

	A:= (1*30u, 13*6.22u);
	B:= (1*30u+28u, 13*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 王松明 etex, .5[A,B]);

	A:= (1*30u, 14*6.22u);
	B:= (1*30u+28u, 14*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 庄凌展 etex, .5[A,B]);

	A:= (1*30u, 15*6.22u);
	B:= (1*30u+28u, 15*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 王赟 etex, .5[A,B]);

	A:= (1*30u, 16*6.22u);
	B:= (1*30u+28u, 16*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 张万松 etex, .5[A,B]);

	A:= (1*30u, 17*6.22u);
	B:= (1*30u+28u, 17*6.22u);
	draw A--B;
	A:= A+(0, 2u);
	B:= B+(0, 2u);
	label(btex 邱伟鹏 etex, .5[A,B]);



endfig;

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
verbatimtex
\end{CJK}
\end{document}
etex end
