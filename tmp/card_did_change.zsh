#!/bin/zsh
echo id
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
echo $token
pass=`cat ~/.config/atr/api_card.json|jq -r .pass`
if [ -z "$1" ];then
	echo 1 : id
	exit
fi
if [ -z "$2" ];then
	echo 2 : did
	exit
fi

id=$1
did=$2

id=$1
curl -X PATCH -H "Content-Type: application/json" -d "{\"did\":\"$did\",\"token\":\"$token\"}" -s $host/users/$id
