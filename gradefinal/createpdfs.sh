#!/bin/sh
suffix=".tex"
if [ ! -d "pdfs" ]; then
	mkdir pdfs
fi

./gradefinal studentsinfo.xlsx

mv *.tex pdfs/
cd pdfs

dir=$(eval pwd)
echo $dir

for file in $(ls $dir | grep .${suffix})
do
	xelatex $file
	xelatex $file
done
echo "Finished produce pdfs"
