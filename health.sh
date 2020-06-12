#!/bin/sh


landscape-sysinfo | jq -sR '[sub("\n$";"") | splits("\n") | sub("^:[ \t]+";"") | [splits(":[ \t]+")]]' | jq '[.[] | {"type": .[0], "status": .[1]}]' | sed 's/" */"/g'
 
