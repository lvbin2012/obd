package enum

type UserState int

const (
	UserState_ErrorState UserState = -1
	UserState_Offline    UserState = 0
	UserState_OnLine     UserState = 1
)

type SendTargetType int

const (
	SendTargetType_SendToNone     SendTargetType = -1
	SendTargetType_SendToAll      SendTargetType = 0
	SendTargetType_SendToSomeone  SendTargetType = 1
	SendTargetType_SendToExceptMe SendTargetType = 2
)

type MsgType int

const (
	MsgType_UserLogin  MsgType = 1
	MsgType_UserLogout MsgType = 2

	MsgType_Core_GetNewAddress  MsgType = 1001
	MsgType_Core_GetMiningInfo  MsgType = 1002
	MsgType_Core_GetNetworkInfo MsgType = 1003

	MsgType_ChannelOpen                 MsgType = -32
	MsgType_ChannelOpen_ItemByTempId    MsgType = -3201
	MsgType_ChannelOpen_AllItem         MsgType = -3202
	MsgType_ChannelOpen_Count           MsgType = -3203
	MsgType_ChannelOpen_DelItemByTempId MsgType = -3204

	MsgType_ChannelAccept MsgType = -33

	MsgType_FundingCreate_Edit     MsgType = -34
	MsgType_FundingCreate_ItemById MsgType = -3401
	MsgType_FundingCreate_Count    MsgType = -3402
	MsgType_FundingCreate_DelById  MsgType = -3403
	MsgType_FundingCreate_DelAll   MsgType = -3404

	MsgType_FundingSign_Edit     MsgType = -35
	MsgType_FundingSign_ItemById MsgType = -3501
	MsgType_FundingSign_Count    MsgType = -3502
	MsgType_FundingSign_DelById  MsgType = -3503
	MsgType_FundingSign_DelAll   MsgType = -3504

	MsgType_CommitmentTx_Edit         MsgType = -351
	MsgType_CommitmentTx_ItemByChanId MsgType = -35101
	MsgType_CommitmentTx_ItemById     MsgType = -35102
	MsgType_CommitmentTx_Count        MsgType = -35103

	MsgType_CommitmentTxSigned_Edit         MsgType = -352
	MsgType_CommitmentTxSigned_ItemByChanId MsgType = -35201
	MsgType_CommitmentTxSigned_ItemById     MsgType = -35202
	MsgType_CommitmentTxSigned_Count        MsgType = -35203

	MsgType_GetBalanceRequest MsgType = -353
	MsgType_GetBalanceRespond MsgType = -354
)
