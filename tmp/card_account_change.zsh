#!/bin/zsh
echo old_id old_did new_name
host=https://api.syui.ai
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
if [ -z "$1" ];then
	exit
fi
if [ -z "$2" ];then
	exit
fi
if [ -z "$3" ];then
	exit
fi
if [ -z "$4" ];then
	exit
fi

id=$1
data=`curl -sL "$host/users/$id/card?itemsPerPage=2550"`
echo $data
id_n=$4
echo $id_n
read

n=`echo $data|jq length`
n=$((n - 1))

for ((i=0;i<=$n;i++))
do
	card=`echo $data|jq -r ".[$i].card"`
	s=`echo $data|jq -r ".[$i].status"`
	cp=`echo $data|jq -r ".[$i].cp"`
	curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id_n,\"card\":$card,\"status\":\"$s\",\"cp\":$cp,\"password\":\"$pass\"}" -sL $host/cards
done

