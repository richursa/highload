#!/bin/bash
go build parseint.go
sudo perf record ./parseint < randomgenerator/input.txt
sudo perf report