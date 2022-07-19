package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ GovHooks = MultiGovHooks{}

// MultiGovHooks combine multiple governance hooks, all hook functions are run in array sequence
type MultiGovHooks []GovHooks

func NewMultiGovHooks(hooks ...GovHooks) MultiGovHooks {
	return hooks
}

func (h MultiGovHooks) AfterProposalSubmission(ctx sdk.Context, proposalID uint64) {
	for i := range h {
		h[i].AfterProposalSubmission(ctx, proposalID)
	}
}

func (h MultiGovHooks) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	for i := range h {
		h[i].AfterProposalVote(ctx, proposalID, voterAddr)
	}
}

func (h MultiGovHooks) AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64) {
	for i := range h {
		h[i].AfterProposalVotingPeriodEnded(ctx, proposalID)
	}
}
