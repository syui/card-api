#!/bin/zsh

username=yui
id=381

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

## users
curl -sL "$host/users?itemsPerPage=2550"|jq .
read
curl -sL "$host/users/$id?itemsPerPage=2550"|jq .
read
curl -sL "$host/users/$id/card?itemsPerPage=2550"|jq .
read

## cards
curl -sL "$host/cards?itemsPerPage=2550"|jq .
read

## token
echo "\ntest token (no token)"
curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\"}" -s $host/users/$id
echo "\ntest token (yes token)"
curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\",\"token\":\"$token\"}" -s $host/users/$id

read
## delete
echo "\ntest delete"
data=`curl -sL https://api.syui.ai/users/$id/card`
n=`echo $data|jq length`
n=$((n - 1))

for ((i=0;i<=$n;i++))
do
	card=`echo $data|jq -r ".[$i].id"`
	echo "\ncard id : $card"
	curl -X DELETE -H "Content-Type: application/json" -d "{'owner':\"$id\"}" $host/cards/$card
done
curl -X DELETE -H "Content-Type: application/json" https://api.syui.ai/users/$id

## card select
echo "\ntest select card (no password)"
curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":0,\"status\":\"normal\",\"cp\":1}" $host/cards
echo "\ntest get card (no password)"
curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id}" $host/cards
#echo "\ntest get card (yes password)"
#curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"password\":\"$pass\"}" $host/cards

## create user
echo "\ntest create user (no password)"
curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"test\"}" $host/users
