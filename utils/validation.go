package utils

import (
	"fmt"
	"github.com/oliverziegert/dccmd-go/config"
)

// IsValidTarget ToDo: Target Validation
func IsValidTarget(target string) bool {
	cp := fmt.Sprintf("%s.%s", config.ALIASES, target)
	return config.IsSet(cp)
}

func AreValidTargets(targets []string) bool {
	if len(targets) == 0 {
		return false
	}
	for _, target := range targets {
		if !IsValidTarget(target) {
			return false
		}
	}
	return true
}
