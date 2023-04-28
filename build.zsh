#!/bin/zsh

d=${0:a:h}
cd $d
go generate ./...
cp -rf $d/tmp/ogent ent/
cp -rf $d/tmp/openapi.json ent/
