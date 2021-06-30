package main

import (
	"context"

	ory "github.com/ory/kratos-client-go"
	"github.com/ory/kratos/examples/go/pkg"
	"github.com/ory/kratos/x"
)

// If you use Open Source this would be:
//
// var client = pkg.NewSDKForSelfHosted("http://127.0.0.1:4433")
var client = pkg.NewSDK("playground")

func createIdentity() *ory.Identity {
	ctx := context.Background()

	identity, _, err := client.V0alpha0Api.AdminCreateIdentity(ctx).AdminCreateIdentityBody(ory.AdminCreateIdentityBody{
		SchemaId: "default",
		Traits: map[string]interface{}{
			"email": "dev+" + x.NewUUID().String() + "@ory.sh",
		},
	}).Execute()
	pkg.FatalOnError(err)

	return identity
}

func main() {
	pkg.PrintJSONPretty(createIdentity())
}
