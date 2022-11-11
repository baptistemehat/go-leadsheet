#!/bin/bash
set -e

source_file=$1

BASE_DIR=$(dirname "$0")
LATEX_DIR=$BASE_DIR/../latex

LATEX_TMP_DIR=$LATEX_DIR/tmp
LATEX_SONGS_DIR=$LATEX_TMP_DIR/songs
LATEX_OUT_DIR=$LATEX_TMP_DIR/out

mkdir -p $LATEX_TMP_DIR $LATEX_SONGS_DIR $LATEX_OUT_DIR

# convert custom format to latex for all files

base_file_name=$(echo "$(basename ${source_file})" | cut -f 1 -d '.')
latex_file="${LATEX_SONGS_DIR}/${base_file_name}.tex"
sed -E \
's/^([^:]*)$/\1 \\\\/g;
s/\[\]/^/g;
s/\[([^]]+)\]/^{\1}/g;
s/:begin properties/\{/g;
s/:end properties/}/g;
s/:begin ([^ ]+) *$/\\begin{\1}/g;
s/:begin ([^ ]+) ?(.*)?$/\\begin{\1} [\2]/g;
s/:end ([^ ]+).*$/\\end{\1}/g;
s/:chord ([^ ]+)/~\\writechord{\1}/g;
s/\|\|:/\\leftrepeat/g;
s/:\|\|/\\rightrepeat/g;
s/\|\|/\\doublebar/g;
s/\|/\\normalbar/g;
s/\(\(([^)]*)\)\)/\\hspace{\\fill}(\1)/g;
s/\(([^)]*)\)/~~\\color{gray}\\textit{(\1)}\\color{black}~~/g;
s/:rememberchords ([^ ]*)/\\setleadsheets{remember-chords = \1}/g;
s/:([a-z]+) (.*)$/\1={\2},/g;
s/\^\s*\\+/^~ \\\\/g;
s/^\s*\\*$//g;' \
$source_file > $latex_file

# generate main.tx from template and file list
ls $LATEX_SONGS_DIR | awk '{print "\\input{tmp/songs/"$0"}\\newpage"}' > $LATEX_TMP_DIR/file_list.tex