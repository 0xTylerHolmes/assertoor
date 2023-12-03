package buildinfo

import "fmt"

var BuildVersion string
var BuildRelease string

func GetVersion() string {
	if BuildRelease == "" {
		return fmt.Sprintf("git-%v", BuildVersion)
	} else {
		return fmt.Sprintf("%v (git-%v)", BuildRelease, BuildVersion)
	}
}
