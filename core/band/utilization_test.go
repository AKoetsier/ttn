// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package band

import (
	"testing"
	"time"

	"github.com/TheThingsNetwork/go-utils/rate"
	. "github.com/smartystreets/assertions"
)

func TestUtilization(t *testing.T) {
	a := New(t)
	u := Utilization{
		868100000: rate.NewCounter(5*time.Minute, time.Hour),
		868300000: rate.NewCounter(5*time.Minute, time.Hour),
	}
	a.So(u.GetFrequency(868100000, time.Hour), ShouldEqual, 0)

	u.Add(868100000, 2*time.Second)
	a.So(u.GetFrequency(868100000, time.Hour), ShouldEqual, 2.0/3600.0)
	a.So(u.GetFrequency(868300000, time.Hour), ShouldEqual, 0)
}
