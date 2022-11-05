sudo apt intall texlive-full

sudo mkdir -p /usr/share/texlive/texmf-dist/tex/plain/leadsheets/
sudo cp *.tex /usr/share/texlive/texmf-dist/tex/plain/leadsheets/

sudo cp *.cls /usr/share/texlive/texmf-dist/tex/latex/base/

sudo mkdir -p /usr/share/texlive/texmf-dist/tex/latex/leadsheets/
sudo cp *.sty /usr/share/texlive/texmf-dist/tex/latex/leadsheets/

sudo texhash