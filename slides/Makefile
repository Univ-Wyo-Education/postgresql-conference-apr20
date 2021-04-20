
DIR=~/go/src/github.com/Univ-Wyo-Education/F20-1015
PP=`pwd`
PY=
IMG=

MD= protocol.md
HTML= protocol.html


all: template.html pre $(MD) $(HTML) 

template.html: raw_template.html slides.html
	m4 -P raw_template.html slides.html >template.html

run:
	89-http-server --dir . --port 12125 &


# protocal.md

pre:
	cp /dev/null open.1
.PHONEY: pre

%.md: %.raw.md $(PY) $(IMG)
	m4 -P $< >$@

%.html: %.md
	blackfriday-tool ./$< $@
	echo cat ${DIR}/md.css $@ >/tmp/$@
	cat ${DIR}/css/pre ${DIR}/css/markdown.css ${DIR}/css/post ${DIR}/md.css ${DIR}/css/hpre $@ ${DIR}/css/hpost >/tmp/$@
	mv /tmp/$@ ./$@
	echo "file://${PP}/$@" >>open.1

