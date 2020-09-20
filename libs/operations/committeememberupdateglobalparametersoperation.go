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
	MaximumWitnessCount              types.UInt16      `json:"maximum_witness_count"`
	NetworkPercentOfFee              types.UInt16      `json:"network_percent_of_fee"`
	ReservePercentOfFee              types.UInt16      `json:"reserve_percent_of_fee"`
	CashbackVestingPeriodSeconds     types.UInt32      `json:"cashback_vesting_period_seconds"`
	CommitteeProposalReviewPeriod    types.UInt32      `json:"committee_proposal_review_period"`
	WitnessPayVestingSeconds         types.UInt32      `json:"witness_pay_vesting_seconds"`
	MaximumProposalLifetime          types.UInt32      `json:"maximum_proposal_lifetime"`
	MaximumTimeUntilExpiration       types.UInt32      `json:"maximum_time_until_expiration"`
	MaximumTransactionSize           types.UInt32      `json:"maximum_transaction_size"`
	MaintenanceInterval              types.UInt32      `json:"maintenance_interval"`
	MaximumBlockSize                 types.UInt32      `json:"maximum_block_size"`
	CashbackVestingThreshold         types.Int64       `json:"cashback_vesting_threshold"`
	WitnessPayPerBlock               types.Int64       `json:"witness_pay_per_block"`
	WorkerBudgetPerDay               types.Int64       `json:"worker_budget_per_day"`
	FeeLiquidationThreshold          types.Int64       `json:"fee_liquidation_threshold"`
}

func (p ChainParameters) Marshal(enc *util.TypeEncoder) error 