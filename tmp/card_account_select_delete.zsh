#!/bin/zsh

pass=`cat ~/.config/atr/api_card.json|jq -r .password`
d=${0:a:h}
f=$d/t.txt

if [ -z "$1" ];then
	exit
fi

if [ ! -f $f ];then
	echo $f
	exit
fi

id=$1
n=`cat $f|wc -l`

for ((i=1;i<=$n;i++))
do
	card=`cat $d/t.txt|awk "NR==$i"`
	echo $card
	read
	curl -X DELETE -H "Content-Type: application/json" -d "{'owner':\"$id\"}" https://api.syui.ai/cards/$card
done

rm $d/t.txt
