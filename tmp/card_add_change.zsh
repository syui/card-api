#!/bin/zsh
case $OSTYPE in
	darwin*)
		alias date="/opt/homebrew/bin/gdate"
		;;
esac

echo cid
read cid

card_status=book
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`

data=`curl -sL "https://api.syui.ai/users?itemsPerPage=2555"`
book=`echo $data|jq ".[]|select(.book == true)|.book"`
id=`echo $data|jq ".[]|select(.fav != 0)|.id"`
echo $book

curl -X PATCH -H "Content-Type: application/json" -d "{\"status\":\"$card_status\",\"token\":\"$token\"}" $host/cards/$cid

