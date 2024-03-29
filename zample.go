package zample

import (
	"math/rand"
)

// DefaultRate controls the normal probability preservation rate of each line.
const DefaultRate = float64(0.1)

// Zample selects strings from a random source.
//
// rate specifies the probability of preserving each string.
//
// skip specifies deterministic skipping each nth file entry.
// This option disables probabilistic rate behavior.
//
// Sampling is tuneable via the Seed function from math/rand.
//
// Returns an input channel for submitting population strings;
// an output channel for receiving select strings.
func Zample(rate *float64, skip *int64) (chan<- string, <-chan string, chan<- struct{}) {
	chIn := make(chan string)
	chOut := make(chan string)
	chDone := make(chan struct{})

	go func() {
		i := int64(1)
		var line string

		for {
			select {
			case <-chDone:
				return
			case line = <-chIn:
				if skip == nil {
					if rand.Float64() < *rate {
						chOut <- line
					}
				} else {
					if i == *skip {
						i = 1
					} else {
						chOut <- line
						i++
					}
				}
			}
		}
	}()

	return chIn, chOut, chDone
}
