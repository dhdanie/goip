package goip

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	TWO32 = 4294967296
)

type CIDR struct {
	addr *V4Address

	bitmask uint
}

//NewCIDR returns a new instance of CIDR based on the passed parameter.
//The single import parameter must represent a CIDR in the form of N.N.N.N/M where each N is an integer value between
//0 and 255 (inclusive) and M is the CIDR mask whose value must fall between 0 and 32 (inclusive)
func NewCIDR(cidr string) (*CIDR, error) {
	var err error

	toks := strings.Split(cidr, "/")
	if len(toks) != 2 {
		return nil, fmt.Errorf("invalid CIDR block (token count)")
	}

	cidrBlock := CIDR{}
	if cidrBlock.addr, err = NewV4Address(toks[0]); err != nil {
		return nil, err
	}
	if bitmask, err := strconv.Atoi(toks[1]); err != nil {
		return nil, err
	} else {
		cidrBlock.bitmask = uint(bitmask)
	}
	if cidrBlock.bitmask < 0 || cidrBlock.bitmask > 32 {
		return nil, fmt.Errorf("invalid bitmask value - %d", cidrBlock.bitmask)
	}
	cidrBlock.calcOctetMasks()

	return &cidrBlock, nil
}

//NewCIDRFromRange returns a new instance of CIDR based on the passed parameter.
//The import parameters are the lowest and highest IP addresses of the range of IPs of the resulting CIDR.
func NewCIDRFromRange(low *V4Address, high *V4Address) (*CIDR, error) {
	highIps := high.getIpCount()
	lowIps := low.getIpCount()

	diff := highIps - lowIps + 1
	mask := 32 - uint(math.Log2(float64(diff)))

	cidrBlock := CIDR{
		addr:    low,
		bitmask: mask,
	}
	cidrBlock.addr, _ = cidrBlock.ToRange()

	return &cidrBlock, nil
}

//ToRange returns the lowest and highest (in that order) IPs represented by this CIDR
func (b *CIDR) ToRange() (*V4Address, *V4Address) {
	o0m, o1m, o2m, o3m := b.calcOctetMasks()
	return b.calcLow(o0m, o1m, o2m, o3m), b.calcHigh(o0m, o1m, o2m, o3m)
}

//String returns the string representation of this CIDR in the format N.N.N.N/M where each N is an integer value between
//0 and 255 (inclusive) and M is the CIDR mask whose value will fall between 0 and 32 (inclusive)
func (b *CIDR) String() string {
	return fmt.Sprintf("%s/%d", b.addr.String(), b.bitmask)
}

func (b *CIDR) calcOctetMasks() (uint, uint, uint, uint) {
	diff := 32 - b.bitmask

	little := math.Exp2(float64(diff))

	result := uint(TWO32 - little)

	octet3Mask := result >> 24
	octet2Mask := result >> 16 & 255
	octet1Mask := result >> 8 & 255
	octet0Mask := result & 255

	return octet0Mask, octet1Mask, octet2Mask, octet3Mask
}

func (b *CIDR) calcLow(o0m uint, o1m uint, o2m uint, o3m uint) *V4Address {
	octet3 := b.addr.octet3 & o3m
	octet2 := b.addr.octet2 & o2m
	octet1 := b.addr.octet1 & o1m
	octet0 := b.addr.octet0 & o0m

	low, err := NewV4AddressFromOctets(octet3, octet2, octet1, octet0)
	if err != nil {
		panic(fmt.Errorf("unexpected error"))
	}
	return low
}

func (b *CIDR) calcHigh(o0m uint, o1m uint, o2m uint, o3m uint) *V4Address {
	octet3 := b.addr.octet3 | (o3m ^ 255)
	octet2 := b.addr.octet2 | (o2m ^ 255)
	octet1 := b.addr.octet1 | (o1m ^ 255)
	octet0 := b.addr.octet0 | (o0m ^ 255)

	high, err := NewV4AddressFromOctets(octet3, octet2, octet1, octet0)
	if err != nil {
		panic(fmt.Errorf("unexpected error"))
	}
	return high
}
