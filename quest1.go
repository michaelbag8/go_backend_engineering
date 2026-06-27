// Understanding Go toolchain
package main

import (
	"time"
)

func myName(name, message string, date string) (string, string, time.Time) {
	return name, message, time.Now()
}

