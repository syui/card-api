#!/bin/zsh
case $OSTYPE in
	darwin*)
		alias date="/opt/homebrew/bin/gdate"
		;;
esac

card_status=second
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`

data=`curl -sL "https://api.syui.ai/users?itemsPerPage=2555"`
fav=`echo $data|jq ".[]|select(.fav != 0)|.fav"`
id=`echo $data|jq ".[]|select(.fav != 0)|.id"`

n=`echo $fav|wc -l`
echo $n

for ((i=1;i<=$n;i++))
do
	cid=`echo $fav|awk "NR==$i"`
	uid=`echo $id|awk "NR==$i"`
	#check
	u_data=`curl -sL "https://api.syui.ai/users/$uid/card?itemsPerPage=2555"|jq -r ".[]|select(.status == \"second\" or \"second\")"`
	if [ -z "$u_data" ];then
		d_data=`curl -sL $host/cards/$cid|jq -r "select(.status == \"first\")"`
		echo $d_data
		if [ -z "$d_data" ];then
			curl -X PATCH -H "Content-Type: application/json" -d "{\"status\":\"$card_status\",\"token\":\"$token\"}" $host/cards/$cid
		fi
	fi
done

