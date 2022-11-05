#!/bin/bash
set -e


BASE_DIR=$(dirname "$0")
LATEX_DIR=$BASE_DIR/latex
LATEX_TMP_DIR=$LATEX_DIR/tmp
LATEX_OUT_DIR=$LATEX_TMP_DIR/out

# run pdflatex
pushd $LATEX_DIR
pwd
pdflatex -output-directory="tmp/out" main.tex
popd

mv $LATEX_OUT_DIR/main.pdf $BASE_DIR/output.pdf

rm -rf $LATEX_TMP_DIR