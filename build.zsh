#!/bin/zsh

d=${0:a:h}
cd $d
su=5000

go generate ./...
sed -i '' "s/255/$su/g" $d/ent/ogent/oas_parameters_gen.go
sed -i '' "s/255/$su/g" $d/ent/openapi.json

cp -rf $d/tmp/ogent ent/
