#!/bin/zsh

host=https://api.syui.ai
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
if [ -z "$1" ];then
	exit
fi

echo username
read
id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .id`
card=43
cp=0
s=book
curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":$card,\"status\":\"$s\",\"cp\":$cp,\"password\":\"$pass\"}" -sL $host/cards
