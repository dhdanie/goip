##Go IP/CIDR Converter

A simple library for calculating ranges from CIDR and vice-versa.

###Example
####CIDR to IP Range
```go
cidr, err := goip.NewCIDR("10.42.0.0/21")
low, high := cidr.ToRange()
fmt.Printf("CIDR - %s\n", cidr)
fmt.Printf("Low  - %s\n", low)
fmt.Printf("High - %s\n", high)
```
Oubput would be
```
CIDR - 10.42.0.0/21
Low  - 10.42.0.0
High - 10.42.7.255
```
####IP Range to CIDR
```go
lowAddr, err := goip.NewV4Address("10.42.0.0")
if err != nil {
    fmt.Printf("Error - %s\n", err.Error())
    os.Exit(1)
}
highAddr, err := goip.NewV4Address("10.42.7.255")
if err != nil {
    fmt.Printf("Error - %s\n", err.Error())
    os.Exit(1)
}
cidr, err := goip.NewCIDRFromRange(lowAddr, highAddr)
if err != nil {
    fmt.Printf("Error - %s\n", err.Error())
    os.Exit(1)
}
fmt.Printf("CIDR - %s\n", cidr)
```
Output would be
```
CIDR - 10.42.0.0/21
```
