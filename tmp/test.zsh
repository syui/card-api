#!/bin/zsh

case $OSTYPE in
	darwin*)
		alias date="/opt/homebrew/bin/gdate"
		;;
esac
u=ai
s=normal
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
host_users="$host/users?itemsPerPage=2550"
updated_at_n=`date --iso-8601=seconds -d '1 days ago'`
now_at=`date --iso-8601=seconds`
raid_at_n=`date --iso-8601=seconds -d '1 days ago'`
data=`curl -sL "$host_users"|jq .`
nd=`date +"%Y%m%d"`
card=0
cp=0
body="next[y]"
id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$u\")"|jq -r .id`
echo $id

#echo no token
#curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\", \"updated_at\":\"$updated_at_n\", \"raid_at\":\"$raid_at_n\", \"ten_su\": 1, \"luck_at\": \"$now_at\", \"ten_at\": \"$updated_at_n\"}" -s $host/users/$id
#echo $body
#read
#
#echo yes token
#echo $token
#curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\", \"updated_at\":\"$updated_at_n\", \"raid_at\":\"$raid_at_n\", \"ten_su\": 1, \"luck_at\": \"$now_at\", \"ten_at\": \"$updated_at_n\",\"token\": \"$token\"}" -s $host/users/$id
#echo $body
#read
#
#echo account delete id no token
#curl -X PATCH -H "Content-Type: application/json" -d "{\"delete\":true}" -s $host/users/381
#echo $body
#read
#
#echo account delete id token
#curl -X PATCH -H "Content-Type: application/json" -d "{\"delete\":false,\"token\":\"$token\"}" -s $host/users/381
#echo $body
#read
#
#echo add card no pass
#curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":$card,\"status\":\"$s\",\"cp\":$cp}" -sL $host/cards
#echo $body
#read
#
#echo add card no pass
#echo $id
#curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":$card,\"status\":\"$s\",\"cp\":$cp,\"password\":\"$pass\"}" -sL $host/cards
#echo $body
#read
#
#echo add user
#curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"test\",\"did\":\"t\"}" -s "$host/users"
#echo $body
#read
#
#echo add user pass
#curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"test\",\"did\":\"t\",\"password\":\"$pass\"}" -s "$host/users"
#echo $body
#read

#echo card draw no pass
#curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id}" -s $host/cards
#echo $body
#read
#
#echo card draw ok pass
#curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"password\":\"$pass\"}" -s $host/cards
#echo $body
#read
