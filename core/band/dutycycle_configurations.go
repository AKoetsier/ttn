// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package band

import "time"

// DutyCycle_EU_863_870 is the duty cycle configuration for the EU_863_870 band
var DutyCycle_EU_863_870 = DutyCycle{
	SubDutyCycle{ // g band
		frequencies:    []uint64{867100000, 867300000, 867500000, 867700000, 867900000},
		timeSpan:       time.Hour,
		maxUtilization: 0.01,
	},
	SubDutyCycle{ // g1 band
		frequencies:    []uint64{868100000, 868300000, 868500000, 868800000},
		timeSpan:       time.Hour,
		maxUtilization: 0.01,
	},
	SubDutyCycle{ // g3 band
		frequencies:    []uint64{869525000},
		timeSpan:       time.Hour,
		maxUtilization: 0.1,
	},
}

// DutyCycle_CN_779_787 is the duty cycle configuration for the CN_779_787 band
var DutyCycle_CN_779_787 = DutyCycle{
	SubDutyCycle{
		timeSpan:       time.Hour,
		maxUtilization: 0.01,
	},
}

// DutyCycle_EU_433 is the duty cycle configuration for the EU_433 band
var DutyCycle_EU_433 = DutyCycle{
	SubDutyCycle{
		timeSpan:       time.Hour,
		maxUtilization: 0.01,
	},
}

// DutyCycle_AS_923 is the duty cycle configuration for the AS_923 band
var DutyCycle_AS_923 = DutyCycle{
	SubDutyCycle{
		frequencies:    []uint64{9232000000, 9234000000},
		timeSpan:       time.Hour,
		maxUtilization: 0.01,
	},
}
