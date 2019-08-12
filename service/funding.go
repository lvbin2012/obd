package service

import (
	"LightningOnOmni/config/chainhash"
)

//type: -34 (funding_created)
type Funding_created struct {
	Id                   int            `storm:"id,increment" `
	Temporary_channel_id chainhash.Hash `json:"temporary_channel_id"`
	Funder_pubKey        chainhash.Hash `json:"funder_pub_key"`
	Property_id          int            `json:"property_id"`
	Max_assets           float64        `json:"max_assets"`
	Amount_a             float64        `json:"amount_a"`
}

type Funding_Service struct {
}

var FundingService Funding_Service

func (service *Funding_Service) CreateFunding() error {
	db, e := DB_Manager.GetDB()
	if e != nil {
		return e
	}
	tempId, _ := Channel_Service.getTemporayChaneelId()
	hashes, _ := chainhash.NewHashFromStr("abc")
	node := &Funding_created{
		Temporary_channel_id: *tempId,
		Funder_pubKey:        *hashes,
		Property_id:          31,
		Max_assets:           1000,
		Amount_a:             20,
	}
	return db.Save(node)
}

//type: -35 (funding_signed)
type Funding_signed struct {
	Id int `storm:"id,increment" `
	//the same as the temporary_channel_id in the open_channel message
	Temporary_channel_id chainhash.ChainHash `json:"temporary_channel_id"`
	//the omni address of funder Alice
	Funder_pubKey chainhash.Hash `json:"funder_pub_key"`
	// the id of the Omni asset
	Property_id int `json:"property_id"`
	//amount of the asset on Alice side
	Amount_a float64 `json:"amount_a"`
	//the omni address of fundee Bob
	Fundee_pubKey chainhash.Hash `json:"fundee_pub_key"`
	//amount of the asset on Bob side
	Amount_b float64 `json:"amount_b"`
	//signature of fundee Bob
	Fundee_signature chainhash.Signauture `json:"fundee_signature"`
	//redeem script used to generate P2SH address
	RedeemScript string `json:"redeem_script"`
	//hash of redeemScript
	P2sh_address chainhash.Hash `json:"p_2_sh_address"`
	//final global channel id generated
	Channel_id chainhash.Hash `json:"channel_id"`
}