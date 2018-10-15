package vault

import (
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

// A new set of documentation-only paths have been created for those endpoints
// that bypass the normal logical.Framework structure and have no Paths defined.
var documentationPaths = []*framework.Path{
	&framework.Path{
		Pattern: "health",

		HelpSynopsis: "Returns the health status of Vault.",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: &framework.PathOperation{
				Callback: framework.NullCallback,
				Responses: map[int]framework.Response{
					200: framework.Response{Description: "Initialized, unsealed, and active"},
					429: framework.Response{Description: "Unsealed and standby"},
					472: framework.Response{Description: "Data recovery mode replication secondary and active"},
					501: framework.Response{Description: "Not initialized"},
					503: framework.Response{Description: "Sealed"},
				},
			},
		},
	},
}
