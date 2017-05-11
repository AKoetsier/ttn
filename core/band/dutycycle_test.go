// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package band

import (
	"testing"
	"time"

	. "github.com/smartystreets/assertions"
)

func TestDutyCycle(t *testing.T) {
	a := New(t)

	d := DutyCycle{
		SubDutyCycle{
			timeSpan:       time.Hour,
			maxUtilization: 0.1,
		},
		SubDutyCycle{
			frequencies:    []uint64{868100000, 868300000, 868500000},
			timeSpan:       time.Minute,
			maxUtilization: 0.01,
		},
	}
	u := d.BuildUtilization()

	a.So(d.Available(868100000, u), ShouldBeTrue)
	u.Add(868100000, 10*time.Minute)
	a.So(d.Available(868100000, u), ShouldBeFalse)
	a.So(d.Available(868800000, u), ShouldBeTrue)

}
