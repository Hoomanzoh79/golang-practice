package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String()string{
	parts := make([]string,len(ip))
	for i,v := range ip {
		parts[i] = fmt.Sprintf("%d",v)
	}
	return strings.Join(parts,".")
}
