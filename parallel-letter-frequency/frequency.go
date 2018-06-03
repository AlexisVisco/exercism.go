// Package letter count the frequency of a letter in a text
package letter

import (
	"sync"
)

type FreqMap map[rune]int

var mutex sync.Mutex
var waitGroup sync.WaitGroup

// FrequencyMapLocked do the same job as Frequency but lock the
// map.
// When job is done, set waitGroup to done
func FrequencyMapLocked(m FreqMap, s string) {
	for _, r := range s {
		mutex.Lock()
		m[r]++
		mutex.Unlock()
	}
	waitGroup.Done()
}

// Frequency add 1 to m[b] where b is a letter in string s
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency return a map of letter frequency in
// texts passed as param.
// Create a wait group of len of the number of texts passed.
// Iterate thought those texts and execute a go routine.
// Then before returning the map, force the program to wait
// the end of each go routines.
func ConcurrentFrequency(texts []string) FreqMap {
	m := FreqMap{}
	waitGroup.Add(len(texts))
	for e := range texts {
		go FrequencyMapLocked(m, texts[e])
	}
	waitGroup.Wait()
	return m
}
