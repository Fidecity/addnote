package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	"github.com/Assetsadapter/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCommitteeMemberUpdateGlobalParameters] = func() types.Operation {
		op := &CommitteeMemberUpdateGlobalParametersOperation{}
		return op
	}
}

type ChainParameters struct {
	AllowNonMemberWhitelists         bool              `json:"allow_non_member_whitelists"`
	CountNonMemberVotes              bool              `json:"count_non_member_votes"`
	Extensions                       types.Extensions  `json:"extensions"`
	CurrentFees                      types.FeeSchedule `json:"current_fees"`
	AccountFeeScaleBitshifts         types.UInt8       `json:"account_fee_scale_bitshifts"`
	BlockInterval                    types.UInt8       `json:"block_interval"`
	MaintenanceSkipSlots             types.UInt8       `json:"maintenance_skip_slots"`
	MaxAuthorityDepth                types.UInt8       `json:"max_authority_depth"`
	MaximumAssetFeedPublishers       types.UInt8       `json:"maximum_asset_feed_publishers"`
	MaximumAssetWhitelistAuthorities types.UInt8       `json:"maximum_asset_whitelist_authorities"`
	AccountsPerFeeScale              types.UInt16      `json:"accounts_per_fee_scale"`
	LifetimeReferrerPercentOfFee     types.UInt16      `json:"lifetime_referrer_percent_of_fee"`
	MaxPredicateOpcode               types.UInt16      `json:"max_predicate_opcode"`
	MaximumAuthorityMembership       types.UInt16      `json:"maximum_authority_membership"`
	MaximumCommitteeCount            types.UInt16      `json:"maximum_committee_count"`
	MaximumWitnessC