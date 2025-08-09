package policy

import (
	"fmt"
	"strings"
)

type Policy int

const (
	ROUND_ROBIN Policy = iota
	IP_HASH
	LEAST_CONN
	RANDOM
)

func ParsePolicy(policy string) (Policy, error) {
	switch strings.ToLower(policy) {
	case "round_robin":
		return ROUND_ROBIN, nil
	case "ip_hash":
		return IP_HASH, nil
	case "least_conn":
		return LEAST_CONN, nil
	case "random":
		return RANDOM, nil
	}
	var p Policy
	return p, fmt.Errorf("not a valid policy: %q", policy)
}

func (policy Policy) MarshalText() (string, error) {
	switch policy {
	case ROUND_ROBIN:
		return "round_robin", nil
	case IP_HASH:
		return "ip_hash", nil
	case LEAST_CONN:
		return "least_conn", nil
	case RANDOM:
		return "random", nil
	}
	return "", fmt.Errorf("not a valid policy: %q", policy)
}
