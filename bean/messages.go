package bean

import (
	"LightningOnOmni/bean/chainhash"
	"LightningOnOmni/bean/enum"
)

type RequestMessage struct {
	Type            enum.MsgType `json:"type"`
	SenderPeerId    string       `json:"sender_peer_id"`
	RecipientPeerId string       `json:"recipient_peer_id"`
	Data            string       `json:"data"`
	PubKey          string       `json:"pub_key"`
	Signature       string       `json:"signature"`
}
type ReplyMessage struct {
	Type   enum.MsgType `json:"type"`
	Status bool         `json:"status"`
	Sender string       `json:"sender"`
	Result interface{}  `json:"result"`
}

//type = 1
type User struct {
	PeerId   string         `json:"peer_id"`
	Password string         `json:"password"`
	State    enum.UserState `json:"state"`
}

//https://github.com/LightningOnOmnilayer/Omni-BOLT-spec/blob/master/OmniBOLT-03-RSMC-and-OmniLayer-Transactions.md
//type = -32
type OpenChannelInfo struct {
	ChainHash                chainhash.ChainHash `json:"chain_hash"`
	TemporaryChannelId       chainhash.Hash      `json:"temporary_channel_id"`
	FundingSatoshis          uint64              `json:"funding_satoshis"`
	PushMsat                 uint64              `json:"push_msat"`
	DustLimitSatoshis        uint64              `json:"dust_limit_satoshis"`
	MaxHtlcValueInFlightMsat uint64              `json:"max_htlc_value_in_flight_msat"`
	ChannelReserveSatoshis   uint64              `json:"channel_reserve_satoshis"`
	HtlcMinimumMsat          uint64              `json:"htlc_minimum_msat"`
	FeeRatePerKw             uint32              `json:"feerate_per_kw"`
	ToSelfDelay              uint16              `json:"to_self_delay"`
	MaxAcceptedHtlcs         uint16              `json:"max_accepted_htlcs"`
	FundingPubKey            string              `json:"funding_pubkey"`
	RevocationBasePoint      chainhash.Point     `json:"revocation_basepoint"`
	PaymentBasePoint         chainhash.Point     `json:"payment_basepoint"`
	DelayedPaymentBasePoint  chainhash.Point     `json:"delayed_payment_basepoint"`
	HtlcBasePoint            chainhash.Point     `json:"htlc_basepoint"`
}

//type = -33
type AcceptChannelInfo struct {
	OpenChannelInfo
	Attitude bool `json:"attitude"`
}

//type: -38 (close_channel)
type CloseChannel struct {
	ChannelId    ChannelID           `json:"channel_id"`
	Len          uint16              `json:"len"`
	ScriptPubKey []byte              `json:"script_pub_key"`
	Signature    chainhash.Signature `json:"signature"`
}

//type: -34 (funding_created)
type FundingCreated struct {
	TemporaryChannelId chainhash.Hash `json:"temporary_channel_id"`
	FunderPubKey       string         `json:"funder_pub_key"`
	FunderPubKey2      string         `json:"funder_pub_key2"`
	PropertyId         int64          `json:"property_id"`
	MaxAssets          float64        `json:"max_assets"`
	AmountA            float64        `json:"amount_a"`
	FundingTxid        string         `json:"funding_txid"`
	FundingOutputIndex uint32         `json:"funding_output_index"`
}

//type: -35 (funding_signed)
type FundingSigned struct {
	ChannelId ChannelID `json:"channel_id"`
	//the omni address of funder Alice
	FunderPubKey string `json:"funder_pub_key"`
	// the id of the Omni asset
	PropertyId int `json:"property_id"`
	//amount of the asset on Alice side
	AmountA float64 `json:"amount_a"`
	//the omni address of fundee Bob
	FundeePubKey string `json:"fundee_pub_key"`
	//amount of the asset on Bob side
	AmountB float64 `json:"amount_b"`
	//signature of fundee Bob
	FundeeSignature string `json:"fundee_signature"`
	//redeem script used to generate P2SH address
	RedeemScript string `json:"redeem_script"`
	//hash of redeemScript
	P2shAddress string `json:"p2sh_address"`
	Attitude    bool   `json:"attitude"`
}

//type: -351 (commitment_tx)
type CommitmentTx struct {
	ChannelId                ChannelID `json:"channel_id"`  //the global channel id.
	PropertyId               int       `json:"property_id"` //the id of the Omni asset
	Amount                   float64   `json:"amount"`      //amount of the payment
	TempPubKey               string    `json:"temp_pub_key"`
	EncrptedAlice2PrivateKey string    `json:"encrpted_alice2_private_key"` //private key of Alice2, encrypted by Bob's public key
}

//type: -352 (commitment_tx_signed)
type CommitmentTxSigned struct {
	//the global channel id.
	ChannelId ChannelID `json:"channel_id"`
	//the id of the Omni asset.
	PropertyId int `json:"property_id"`
	//amount of the payment.
	Amount float64 `json:"amount"`
	//signature of Bob.
	ReceiverSignature string `json:"receiver_signature"`
	Attitude          bool   `json:"attitude"`
}

//type: -353 (get_balance_request)
type GetBalanceRequest struct {
	//the global channel id.
	ChannelId ChannelID `json:"channel_id"`
	//the p2sh address generated in funding_signed message.
	P2shAddress string `json:"p2sh_address"`
	// the channel owner, Alice or Bob, can query the balance.
	Who chainhash.Hash `json:"who"`
	//the signature of Alice or Bob
	Signature chainhash.Signature `json:"signature"`
}

//type: -354 (get_balance_respond)
type GetBalanceRespond struct {
	//the global channel id.
	ChannelId ChannelID `json:"channel_id"`
	//the asset id generated by Omnilayer protocol.
	PropertyId int `json:"property_id"`
	//the name of the asset.
	Name string `json:"name"`
	//balance in this channel
	Balance float64 `json:"balance"`
	//currently not in use
	Reserved float64 `json:"reserved"`
	//currently not in use
	Frozen float64 `json:"frozen"`
}
