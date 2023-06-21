#!/bin/zsh

username=$1

case $OSTYPE in
	darwin*)
		alias date="/opt/homebrew/bin/gdate"
		;;
esac
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
host_users="$host/users?itemsPerPage=255"
data=`curl -sL "$host_users"|jq .`
nd=`date +"%Y%m%d"`
nd=20230101
id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .id`

echo username did
read
curl -X PATCH -H "Content-Type: application/json" -d "{\"did\":\"$2\",\"token\":\"$token\"}" -s $host/users/$id
