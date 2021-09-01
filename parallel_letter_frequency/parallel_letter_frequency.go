package letter

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

// ConcurrentFrequency function counts the frequency of each rune in a given text list and returns a FreqMap
func ConcurrentFrequency(list []string) FreqMap {
	chList := make(chan FreqMap, len(list))
	for _, s := range list {
		go func(s string) {
			chList <- Frequency(s)
		}(s)
	}
	res := FreqMap{}
	for range list {
		for r, freq := range <-chList {
			res[r] += freq
		}
	}
	return res
}
