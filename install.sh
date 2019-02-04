#!/usr/bin/env bash
go get -u github.com/PuerkitoBio/goquery
go build subjs.go
mv subjs $GOPATH/bin

