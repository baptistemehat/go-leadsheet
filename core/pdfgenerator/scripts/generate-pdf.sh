#!/bin/bash
set -e

BASE_DIR=$(dirname "$0")

LATEX_DIR=$BASE_DIR/../../../tmp/latex

# copy latex static files
cp $BASE_DIR/../../dataprocessing/formatting/latexformatting/latex $BASE_DIR/../../../tmp/ -ar

LATEX_TMP_DIR=$LATEX_DIR/tmp
LATEX_OUT_DIR=$LATEX_TMP_DIR/out
LATEX_SONGS_DIR=$LATEX_TMP_DIR/songs

# generate file_list.tex from template and file list
ls $LATEX_SONGS_DIR | awk '{print "\\input{tmp/songs/"$0"}\\newpage"}' > $LATEX_TMP_DIR/file_list.tex

# run pdflatex
pushd $LATEX_DIR
pwd
pdflatex -output-directory="tmp/out" main.tex
popd
