package common

import "dev/ff_sdk"

var variantA = make(map[string]interface{})
var FlagA = ff_sdk.FeatureFlag{
	Id:      "FlagA",
	State:   ff_sdk.ENABLED,
	DefVar:  "this",
	Variant: variantA,
}

var variantB = make(map[string]interface{})
var FlagB = ff_sdk.FeatureFlag{
	Id:      "FlagB",
	State:   ff_sdk.ENABLED,
	DefVar:  "true",
	Variant: variantB,
}

func init() {
	variantA["this"] = "this"
	variantA["that"] = "that"

	variantB["true"] = true
	variantB["false"] = false
}
