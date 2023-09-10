#!/bin/zsh

case $OSTYPE in
	darwin*)
		alias date="/opt/homebrew/bin/gdate"
		;;
esac
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
host_users="$host/users?itemsPerPage=2550"
updated_at_n=`date --iso-8601=seconds -d '1 days ago'`
now_at=`date --iso-8601=seconds`
raid_at_n=`date --iso-8601=seconds -d '1 days ago'`
data=`curl -sL "$host_users"|jq .`
nd=`date +"%Y%m%d"`

n=`echo $data|jq length`
n=$((n - 1))

count=1
for ((i=0;i<=$n;i++))
do
	username=`echo $data|jq ".[$i]"|jq -r .username`
	uid=`echo $data|jq ".[$i]"|jq -r .id`
	acard=`curl -sL "$host/users/$uid/card?itemsPerPage=3000"|jq ".[0]"`
	cid=`echo $acard|jq -r .id`
	curl -X PATCH -H "Content-Type: application/json" -d "{\"author\":\"$username\", \"count\":$count,\"token\":\"$token\"}" $host/cards/$cid
done
