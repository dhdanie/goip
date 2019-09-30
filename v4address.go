package goip

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type V4Address struct {
	octet0 uint
	octet1 uint
	octet2 uint
	octet3 uint
}

//NewV4AddressFromOctets returns a new instance of V4Address whose octets correspond to the passed parameters
//The octets to-be-used should be passed in order of most-significance to the least-significant.
func NewV4AddressFromOctets(octet3 uint, octet2 uint, octet1 uint, octet0 uint) (*V4Address, error) {
	if octet3 < 0 || octet3 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", octet3)
	}
	if octet2 < 0 || octet2 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", octet2)
	}
	if octet1 < 0 || octet1 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", octet1)
	}
	if octet0 < 0 || octet0 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", octet0)
	}

	return &V4Address{
		octet0: octet0,
		octet1: octet1,
		octet2: octet2,
		octet3: octet3,
	}, nil
}

//NewV4Address returns a new instance of V4Address that corresponds to the IP address specified in the parameter.
//The address parameter must take the form N.N.N.N where the N's correspond to some number value between 0 and 255
func NewV4Address(address string) (*V4Address, error) {
	toks := strings.Split(address, ".")
	if len(toks) != 4 {
		return nil, fmt.Errorf("invalid IP address (token count)")
	}
	addr := V4Address{}
	if octet0, err := strconv.Atoi(toks[3]); err != nil {
		return nil, err
	} else {
		addr.octet0 = uint(octet0)
	}
	if addr.octet0 < 0 || addr.octet0 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", addr.octet0)
	}

	if octet1, err := strconv.Atoi(toks[2]); err != nil {
		return nil, err
	} else {
		addr.octet1 = uint(octet1)
	}
	if addr.octet1 < 0 || addr.octet1 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", addr.octet1)
	}

	if octet2, err := strconv.Atoi(toks[1]); err != nil {
		return nil, err
	} else {
		addr.octet2 = uint(octet2)
	}
	if addr.octet2 < 0 || addr.octet2 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", addr.octet2)
	}

	if octet3, err := strconv.Atoi(toks[0]); err != nil {
		return nil, err
	} else {
		addr.octet3 = uint(octet3)
	}
	if addr.octet3 < 0 || addr.octet3 > 255 {
		return nil, fmt.Errorf("invalid octet value - %d", addr.octet3)
	}

	return &addr, nil
}

func (a *V4Address) getIpCount() int {
	o3Count := int(math.Pow(256, 3)) * int(a.octet3)
	o2Count := int(math.Pow(256, 2)) * int(a.octet2)
	o1Count := 256 * int(a.octet1)
	o0Count := int(a.octet0)

	return o3Count + o2Count + o1Count + o0Count + 1
}

//String returns the string representation of this V4Address in the N.N.N.N format
func (a *V4Address) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a.octet3, a.octet2, a.octet1, a.octet0)
}
