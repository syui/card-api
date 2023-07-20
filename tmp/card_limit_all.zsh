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

if [ -n "$1" ];then
	id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .id`
	if [ "ai" = "$1" ] || [ "yui" = "$1" ];then
		curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\", \"updated_at\":\"$updated_at_n\", \"raid_at\":\"$raid_at_n\", \"token\":\"$token\", \"luck_at\": \"$now_at\", \"luck\": 7, \"like\":0,\"aiten\":1000}" -s $host/users/$id
	else
		curl -X PATCH -H "Content-Type: application/json" -d "{\"raid_at\": \"$updated_at_n\",\"token\": \"$token\"}" -s $host/users/$id
		#curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\", \"updated_at\":\"$updated_at_n\", \"raid_at\":\"$raid_at_n\", \"luck_at\": \"$updated_at_n\", \"ten_at\": \"$updated_at_n\",\"token\": \"$token\"}" -s $host/users/$id
	fi
	exit
fi

for ((i=0;i<=$n;i++))
do
	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	echo "{\"updated_at\":\"$updated_at_n\"} -s $host/users/$id"
	if [ "ai" = "$1" ];then
		curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\", \"updated_at\":\"$updated_at_n\", \"raid_at\":\"$raid_at_n\", \"token\":\"$token\", \"luck_at\": \"$now_at\", \"luck\": 7}" -s $host/users/$id
	else
		curl -X PATCH -H "Content-Type: application/json" -d "{\"raid_at\":\"$raid_at_n\",\"token\": \"$token\", \"ten_at\": \"$updated_at_n\"}" -s $host/users/$id
		#curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\", \"updated_at\":\"$updated_at_n\", \"raid_at\":\"$raid_at_n\", \"token\":\"$token\", \"luck_at\": \"$now_at\", \"ten_at\": \"$updated_at_n\"}" -s $host/users/$id
	fi
done
