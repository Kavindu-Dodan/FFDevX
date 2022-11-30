//go:build ignore

package main

import (
	. "dev/common"
	"dev/ff_sdk"
)

// Register Feature Flag definitions of this application.
func main() {
	F := ff_sdk.Create()
	defer F.Generate("ff.json")

	F.Add(FlagA)
	F.Add(FlagB)
}
