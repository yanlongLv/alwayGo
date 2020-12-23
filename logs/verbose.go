package logs

import "strconv"

// VerboseKey ..
const VerboseKey = "verbose"

// Verbose ..
type Verbose int

// String ...
func (v Verbose) String() string {
	return strconv.Itoa(int(v))
}

// Enabled ...
func (v Verbose) Enabled(n Verbose) bool {
	return v > n
}
