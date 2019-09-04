package bash

import "fmt"

func BoatswainRelease(release string, options string) string {
	return fmt.Sprintf("boatswain release %s %s --assume-yes", release, options)
}

