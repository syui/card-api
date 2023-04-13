#!/bin/zsh
d=${0:a:h}
PASS=`cat token.json|jq -r .password` go run -mod=mod main.go
