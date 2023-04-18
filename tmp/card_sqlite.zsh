#!/bin/zsh

host=https://api.syui.ai
api=localhost:8080

host_users="$host/users?itemsPerPage=2550"
d=${0:a:h}
dd=${0:a:h:h}
pass=`cat $dd/token.json|jq -r .password`
host_at=bsky.social

if [ -f $d/user.json ] && [ "$1" = "-s" ];then
	rm $d/user.json
fi
function did() {
	unset did
	url="https://$host_at/xrpc/com.atproto.repo.listRecords?repo=${name}.${host_at}&collection=app.bsky.actor.profile"
	if [ "`curl -sL $url| jq -r .error`" = "null" ];then
		t=`curl -sL $url | jq -r ".records|.[]|.uri"|cut -d / -f 3`
		did=$t
	else
		#did=`curl -sL "search.bsky.social/search/posts?q=$name" | jq -r ".[0].user.did"`
		did="null"
	fi
	echo "{\"name\":\"$name\",\"did\":\"$did\"}," >> $d/user.json
}

function l_users() {
	curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"$name\",\"password\":\"$pass\",\"did\":\"$did\"}" $api/users
	#sleep 1
}

function l_cards() {
	curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":$card,\"status\":\"$s\",\"cp\":$cp,\"password\":\"$pass\"}" $api/cards
	#sleep 1
}

data=`curl -sL "$host_users"|jq .`
n=`echo $data|jq length`
n=$((n - 1))
for ((i=0;i<=$n;i++))
do

	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	did=`echo $data|jq ".[$i]"|jq -r .did`

	echo "{\"username\":\"$name\", \"password\":\"$pass\",\"did\":\"$did\"} localhost:8080/users"
	if [ "$1" = "-a" ];then
		l_users
	fi

	data_card=`curl -sL "$host/users/$id/card?itemsPerPage=2550"`
	nn=`echo $data_card|jq length`
	nn=$((nn - 1))

	if [ "$1" != "-s" ];then
		for ((ii=0;ii<=$nn;ii++))
		do
			card=`echo $data_card|jq -r ".[$ii].card"`
			s=`echo $data_card|jq -r ".[$ii].status"`
			cp=`echo $data_card|jq -r ".[$ii].cp"`
			echo "{\"owner\":$id,\"card\":\"$card\",\"status\":\"$s\",\"cp\":\"$cp\", \"password\":\"$pass\"} localhost:8080/cards"
			if [ "$1" = "-a" ];then
				l_cards
			fi
		done
	fi
done
