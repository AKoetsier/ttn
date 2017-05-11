// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package band

import (
	"time"

	"github.com/TheThingsNetwork/go-utils/rate"
)

// Utilization for channels, as well as aggregated utilization
type Utilization map[uint64]rate.Counter

func (u Utilization) Add(frequency uint64, t time.Duration) {
	now := time.Now()
	d := uint64(t.Nanoseconds())
	if channel, ok := u[frequency]; ok {
		channel.Add(now, d)
	}
}

// GetFrequency returns the utilization of a frequency
func (u Utilization) GetFrequency(frequency uint64, past time.Duration) float32 {
	if channel, ok := u[frequency]; ok {
		ns, _ := channel.Get(time.Now(), past)
		return float32(ns) / float32(past)
	}
	return 0
}
