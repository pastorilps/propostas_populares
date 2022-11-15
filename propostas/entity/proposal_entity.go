package entity

type Send_Proposal_Data struct {
	ProposalTitle       string `json:"proposaltitle" example:"Proposal Title"`
	ProposalPictures    int64  `json:"proposalpictures" example:"2"`
	ProposalAttachments int64  `json:"proposalattachments" example:"1"`
	ProposalDescription string `json:"proposaldescription" example:"Proposal Description"`
	ProposalStatus      bool   `json:"proposalstatus" example:"true"`
	ProposalUserID      int16  `json:"proposaluserid" example:"2"`
}

type Receive_Proposal_Data struct {
	ProposalTitle       string `json:"proposaltitle" example:"Proposal Title"`
	ProposalId          int16
	ProposalPictures    int64  `json:"proposalpictures" example:"2"`
	ProposalAttachments int64  `json:"proposalattachments" example:"1"`
	ProposalDescription string `json:"proposaldescription" example:"Proposal Description"`
	ProposalStatus      bool   `json:"proposalstatus" example:"true"`
	ProposalUserID      int16  `json:"proposaluserid" example:"2"`
	Token               string `json:"token"`
	UserID              int16  `json:"userid"`
}
