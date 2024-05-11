#!/bin/bash
sudo perf record ./parseint < randomgenerator/input.txt
sudo perf report