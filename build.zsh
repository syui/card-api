#!/bin/zsh

d=${0:a:h}
cd $d
su=4000

go generate ./...
sed -i '' "s/255/$su/g" $d/ent/ogent/oas_parameters_gen.go
sed -i '' "s/255/$su/g" $d/ent/openapi.json
cp -rf $d/ent/ogent/oas_parameters_gen.go $d/tmp/ogent/oas_parameters_gen.go
cp -rf $d/ent/ogent/openapi.json $d/tmp/openapi.json

cp -rf $d/tmp/ogent ent/
cp -rf $d/tmp/openapi.json ent/
