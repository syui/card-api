#!/bin/zsh
echo id
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
pass=`cat ~/.config/atr/api_card.json|jq -r .pass`
if [ -z "$1" ];then
	exit
fi

id=$1
curl -X PATCH -H "Content-Type: application/json" -d "{\"delete\":true,\"token\":\"$token\"}" -s $host/users/$id

