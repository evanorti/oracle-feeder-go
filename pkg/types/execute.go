package types

import (
	"github.com/terra-money/oracle-feeder-go/internal/types"
)

type MsgUpdateChainsInfo struct {
	UpdateChainsInfo UpdateChainsInfo `json:"update_chains_info"`
}

type UpdateChainsInfo struct {
	ChainsInfo types.AllianceProtocolRes `json:"chains_info"`
}

func NewMsgUpdateChainsInfo(data types.AllianceProtocolRes) MsgUpdateChainsInfo {
	return MsgUpdateChainsInfo{
		UpdateChainsInfo: UpdateChainsInfo{
			ChainsInfo: data,
		},
	}
}

type MsgAllianceRedelegate struct {
	Redelegations types.AllianceRedelegateReq `json:"redelegations"`
}

func NewMsgAllianceRedelegate(redelegations types.AllianceRedelegateReq) MsgAllianceRedelegate {
	return MsgAllianceRedelegate{
		Redelegations: redelegations,
	}
}
