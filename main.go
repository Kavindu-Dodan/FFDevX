package main

// Generate FFs of this application. To invoke, run `go generate ./...`. Output - ff.json
//go:generate go run ./non_build/generator.go

import (
	"context"
	cts "dev/common"
	"fmt"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/pkg/openfeature"
)

func main() {
	openfeature.SetProvider(flagd.NewProvider())
	client := openfeature.NewClient("dev-test")

	a(client)
	b(client)
}

func a(client *openfeature.Client) {
	resolve, err := client.StringValue(context.TODO(), cts.FlagA.Id, cts.FlagB.DefVar, openfeature.EvaluationContext{})
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	if resolve == cts.FlagA.DefVar {
		fmt.Println("Default FF resolved: " + cts.FlagA.DefVar)
	} else {
		fmt.Println("Non default FF resolved: " + resolve)
	}
}

func b(client *openfeature.Client) {
	// todo - clarify default value usage
	resolve, err := client.BooleanValue(context.TODO(), cts.FlagB.Id, false, openfeature.EvaluationContext{})
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	if !resolve {
		fmt.Printf("Default FF resolved : %t", resolve)
	} else {
		fmt.Printf("Non default FF resolved : %t", resolve)
	}
}
