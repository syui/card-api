#!/bin/zsh
host=https://api.syui.ai
uid=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .id`
did=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .did`
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
card=36
cp=`echo $(($RANDOM % 1000 + 400))`
s=$(($RANDOM % 7))
if [ $s -eq 1 ];then
	s=super
	skill=post
	plus=$(($RANDOM % 1000 + 300))
	cp=$((cp + plus))
else
	s=normal
	skill=ten
fi
tmp=`curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$uid,\"card\":$card,\"status\":\"$s\",\"cp\":$cp,\"password\":\"$pass\",\"skill\":\"$skill\"}" -sL $host/cards`

card=`echo $tmp|jq -r .card`
cp=`echo $tmp|jq -r .cp`
body="[card]\nid : $card\ncp : $cp"

atr @ $did -p "`echo $body`"
