#!/bin/zsh

host=https://api.syui.ai
api=localhost:8080

host_users="$host/users?itemsPerPage=255"
d=${0:a:h}
dd=${0:a:h:h}
pass=`cat $dd/token.json|jq -r .password`

function l_users() {
	curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"$name\",\"password\":\"$pass\"}" $api/users
	sleep 1
}

function l_cards() {
	curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":$card,\"status\":\"$s\",\"cp\":$cp,\"password\":\"$pass\"}" $api/cards
	sleep 1
}

data=`curl -sL "$host_users"|jq .`
n=`echo $data|jq length`
n=$((n - 1))
for ((i=0;i<=$n;i++))
do

	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	echo "{\"username\":\"$name\"} localhost:8080/users"
	if [ "$1" = "-a" ];then
		l_users
	fi

	data_card=`curl -sL "$host/users/$id/card?itemsPerPage=255"`
	nn=`echo $data_card|jq length`
	nn=$((nn - 1))

	for ((ii=0;ii<=$nn;ii++))
	do
		card=`echo $data_card|jq -r ".[$ii].card"`
		s=`echo $data_card|jq -r ".[$ii].status"`
		cp=`echo $data_card|jq -r ".[$ii].cp"`
		echo "{\"owner\":$id,\"card\":\"$card\",\"status\":\"$s\",\"cp\":\"$cp\"} localhost:8080/cards"
		if [ "$1" = "-a" ];then
			l_cards
		fi
	done

done

