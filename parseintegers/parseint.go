package main

import (
	"fmt"
	"runtime"
	"sync"
	"syscall"
)

const goroutines = 8
const chansize = 0
const buffersize = 1 << 20

var totSum uint64
var totLock sync.Mutex

func main() {
	runtime.GOMAXPROCS(1)
	var err error
	var gochan = make(chan []byte, chansize)
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for range goroutines {
		go goThreads(gochan, &wg)
	}
	var bytes []byte
	length := 1
	prevNumber := []byte{}
	for length != 0 && err == nil {
		bytes = make([]byte, buffersize)
		length, err = syscall.Read(0, bytes)
		for i, b := range bytes {
			if b == '\n' {
				gochan <- append(prevNumber, bytes[:i+1]...)
				gochan <- bytes[i+1 : length]
				if bytes[len(bytes)-1] != '\n' {
					for k := len(bytes) - 1; k >= 0; k-- {
						if bytes[k] == '\n' {
							prevNumber = bytes[k+1 : length]
							goto outer
						}
					}
				} else {
					prevNumber = []byte{}
				}
				goto outer
			}
		}
	outer:
	}
	close(gochan)
	wg.Wait()
	fmt.Printf("%d\n", totSum+stringToint(prevNumber))
}

func goThreads(gochan chan []byte, wg *sync.WaitGroup) {
	mysum := uint64(0)
	for bytes := range gochan {
		prevIndex := 0
		for i, char := range bytes {
			if char == '\n' {
				mysum += stringToint(bytes[prevIndex:i])
				prevIndex = i + 1
			}
		}
	}
	totLock.Lock()
	totSum += mysum
	totLock.Unlock()
	wg.Done()
}
func stringToint(bytes []byte) (num uint64) {
	length := len(bytes) - 1
	switch length {
	case 10:
		num = ((uint64(bytes[10]-'0') + uint64(bytes[9]-'0')*10) + (uint64(bytes[8]-'0')*100 + uint64(bytes[7]-'0')*1000)) + ((uint64(bytes[6]-'0')*10000 + uint64(bytes[5]-'0')*100000) + (uint64(bytes[4]-'0')*1000000 + uint64(bytes[3]-'0')*10000000)) + ((uint64(bytes[2]-'0')*100000000 + uint64(bytes[1]-'0')*1000000000) + uint64(bytes[0]-'0')*10000000000)
	case 9:
		num = ((uint64(bytes[9]-'0') + uint64(bytes[8]-'0')*10) + (uint64(bytes[7]-'0')*100 + uint64(bytes[6]-'0')*1000)) + ((uint64(bytes[5]-'0')*10000 + uint64(bytes[4]-'0')*100000) + (uint64(bytes[3]-'0')*1000000 + uint64(bytes[2]-'0')*10000000)) + (uint64(bytes[1]-'0')*100000000 + uint64(bytes[0]-'0')*1000000000)
	case 8:
		num = ((uint64(bytes[8]-'0') + uint64(bytes[7]-'0')*10) + (uint64(bytes[6]-'0')*100 + uint64(bytes[5]-'0')*1000)) + ((uint64(bytes[4]-'0')*10000 + uint64(bytes[3]-'0')*100000) + (uint64(bytes[2]-'0')*1000000 + uint64(bytes[1]-'0')*10000000)) + uint64(bytes[0]-'0')*100000000
	case 7:
		num = ((uint64(bytes[7]-'0') + uint64(bytes[6]-'0')*10) + (uint64(bytes[5]-'0')*100 + uint64(bytes[4]-'0')*1000)) + ((uint64(bytes[3]-'0')*10000 + uint64(bytes[2]-'0')*100000) + (uint64(bytes[1]-'0')*1000000 + uint64(bytes[0]-'0')*10000000))
	case 6:
		num = (uint64(bytes[6]-'0') + uint64(bytes[5]-'0')*10) + ((uint64(bytes[4]-'0')*100 + uint64(bytes[3]-'0')*1000) + (uint64(bytes[2]-'0')*10000 + uint64(bytes[1]-'0')*100000)) + uint64(bytes[0]-'0')*1000000
	case 5:
		num = ((uint64(bytes[5]-'0') + uint64(bytes[4]-'0')*10) + (uint64(bytes[3]-'0')*100 + uint64(bytes[2]-'0')*1000)) + (uint64(bytes[1]-'0')*10000 + uint64(bytes[0]-'0')*100000)
	case 4:
		num = ((uint64(bytes[4]-'0') + uint64(bytes[3]-'0')*10) + (uint64(bytes[2]-'0')*100 + uint64(bytes[1]-'0')*1000)) + uint64(bytes[0]-'0')*10000
	case 3:
		num = ((uint64(bytes[3]-'0') + uint64(bytes[2]-'0')*10) + (uint64(bytes[1]-'0')*100 + uint64(bytes[0]-'0')*1000))
	case 2:
		num = (uint64(bytes[2]-'0') + uint64(bytes[1]-'0')*10) + uint64(bytes[0]-'0')*100
	case 1:
		num = (uint64(bytes[1]-'0') + uint64(bytes[0]-'0')*10)
	case 0:
		num = uint64(bytes[0] - '0')
	}
	return num
}
