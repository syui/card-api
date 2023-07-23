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

if [ -z "$1" ];then
	echo card
	read card
else
	echo card=$1
	read
	card=$1
fi

for ((i=0;i<=$n;i++))
do
	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	data_uu=`curl -sL "$host/users/$id/card?itemsPerPage=4000"`
	card_check=`echo $data_uu|jq -r ".[]|select(.card == $card)"`
	if [ -n "$card_check" ];then
		echo $name
	fi
done
