// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package band

import (
	"time"

	"github.com/TheThingsNetwork/go-utils/rate"
)

// FrequencyUtilization interface used in duty cycle calculation
type FrequencyUtilization interface {
	GetFrequency(frequency uint64, past time.Duration) float32
}

// SubDutyCycle contains the dutycycle configuration for a sub-band
type SubDutyCycle struct {
	frequencies    []uint64
	timeSpan       time.Duration
	maxUtilization float32
}

// Available returns the availability of the sub-band given its utilization
func (d SubDutyCycle) Available(utilization FrequencyUtilization) bool {
	var total float32
	for _, frequency := range d.frequencies {
		total += utilization.GetFrequency(frequency, d.timeSpan)
	}
	return total <= d.maxUtilization
}

// DutyCycle contains the dutycycle configuration for a band
type DutyCycle []SubDutyCycle

func (d DutyCycle) frequencies() (frequencies []uint64) {
	fMap := make(map[uint64]struct{})
	for _, sub := range d {
		for _, freq := range sub.frequencies {
			fMap[freq] = struct{}{}
		}
	}
	for f := range fMap {
		frequencies = append(frequencies, f)
	}
	return
}

func (d DutyCycle) subs(frequency uint64) (subs DutyCycle) {
next:
	for _, sub := range d {
		if len(sub.frequencies) == 0 {
			subs = append(subs, sub)
			continue next
		}
		for _, freq := range sub.frequencies {
			if freq == frequency {
				subs = append(subs, sub)
				continue next
			}
		}
	}
	return
}

// Available returns the availability of the band given its utilization. If any
// sub-band that this frequency belongs to is unavailable, false will be returned
func (d DutyCycle) Available(frequency uint64, utilization FrequencyUtilization) bool {
	for _, sub := range d.subs(frequency) {
		if !sub.Available(utilization) {
			return false
		}
	}
	return true
}

const bucketResolution = 10

// Default values for the utilization counters
var (
	DefaultRetention  = time.Hour
	DefaultBucketSize = DefaultRetention / bucketResolution
)

// BuildUtilization builds utilization counters for the duty cycle configuration
func (d DutyCycle) BuildUtilization() Utilization {
	bucket := make(map[uint64]time.Duration)
	retention := make(map[uint64]time.Duration)

	// Initialize bucket and retention durations to defaults
	for _, freq := range d.frequencies() {
		bucket[freq] = DefaultBucketSize
		retention[freq] = DefaultRetention
	}

	// Make retention longer and buckets shorter if needed
	for _, sub := range d {
		for _, freq := range sub.frequencies {
			if retention[freq] < sub.timeSpan {
				retention[freq] = sub.timeSpan
			}
			if bucket[freq] > retention[freq]/bucketResolution {
				bucket[freq] = retention[freq] / bucketResolution
			}
		}
	}

	// Fill it up
	u := make(Utilization)
	for freq, bucket := range bucket {
		u[freq] = rate.NewCounter(bucket, retention[freq])
	}
	return u
}
