package main

import "sync"
import "strconv"

type Station struct {
	id int
	delayTime int
	mutex *sync.RWMutex

}

func newStation(id int, delay string) *Station{
	ptr := new(Station)
	ptr.id = id
	d ,_:= strconv.Atoi(delay)
	ptr.delayTime = d

	return ptr
}
