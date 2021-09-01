package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// FrequencySync counts the frequency of each rune in a given text and updates a sync map
func FrequencySync(s string, sMap *sync.Map) {
	for _, r := range s {
		if freq, ok := sMap.LoadOrStore(r, 1); ok {
			sMap.Store(r, 1+freq.(int))
		}
	}
}

// ConcurrentFrequency function counts the frequency of each rune in a given text list and returns a FreqMap
func ConcurrentFrequency(list []string) FreqMap {
	var sMap sync.Map
	var wg sync.WaitGroup
	wg.Add(len(list))
	for _, s := range list {
		in := s
		go func(s string) {
			defer wg.Done()
			FrequencySync(s, &sMap)
		}(in)
	}
	wg.Wait()
	res := FreqMap{}
	sMap.Range(func(key, value interface{}) bool {
		res[key.(rune)] = value.(int)
		return true
	})
	return res
}
