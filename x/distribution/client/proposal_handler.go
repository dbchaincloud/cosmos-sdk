package client

import (
	"github.com/dbchaincloud/cosmos-sdk/x/distribution/client/cli"
	"github.com/dbchaincloud/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/dbchaincloud/cosmos-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
