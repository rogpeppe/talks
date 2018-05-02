#!/bin/sh

echo '##' Some talks given by Roger Peppe
echo
for i in */*.slide; do
	name=$(echo $i | sed 's:/.*::')
	echo "- [$name](https://talks.godoc.org/github.com/rogpeppe/talks/$i)"
done
