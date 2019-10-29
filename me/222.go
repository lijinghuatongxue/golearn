#! /bin/bash
for ((i=0; i<200; ))
do
	i=$(expr $i + 1)
    go run ./test.go
	sleep 1
done