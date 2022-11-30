package ff_sdk

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	ENABLED  = "ENABLED"
	DISABLED = "DISABLED"
)

// flags of flagd configuration
type flags struct {
	Flags map[string]flag `json:"flags"`
}

// flag - one entry of flag
type flag struct {
	State   string                 `json:"state"`
	DefVar  string                 `json:"defaultVariant"`
	Variant map[string]interface{} `json:"variants"`
}

// FeatureFlag for SDK users
type FeatureFlag struct {
	Id      string
	State   string
	DefVar  string
	Variant map[string]interface{}
}

func Create() flags {
	holder := make(map[string]flag)

	return flags{
		Flags: holder,
	}
}

func (f *flags) Define(id string, state string, variants map[string]interface{}, defVar string) {
	f.Flags[id] = flag{
		State:   state,
		DefVar:  defVar,
		Variant: variants,
	}
}

func (f *flags) Add(featureFlag FeatureFlag) {
	f.Flags[featureFlag.Id] = flag{
		State:   featureFlag.State,
		DefVar:  featureFlag.DefVar,
		Variant: featureFlag.Variant,
	}
}

func (f *flags) Generate(fName string) {
	data, err := json.Marshal(f)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}

	err = os.WriteFile(fName, data, 777)
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
}
