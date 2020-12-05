package client

import (
	govclient "github.com/dbchaincloud/cosmos-sdk/x/gov/client"
	"github.com/dbchaincloud/cosmos-sdk/x/upgrade/client/cli"
	"github.com/dbchaincloud/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
