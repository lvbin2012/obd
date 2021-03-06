package lightclient

import (
	"encoding/json"
	"github.com/omnilaboratory/obd/bean"
	"github.com/omnilaboratory/obd/bean/enum"
	"github.com/omnilaboratory/obd/service"
	"strconv"
)

func (client *Client) hdWalletModule(msg bean.RequestMessage) (enum.SendTargetType, []byte, bool) {
	status := false
	var sendType = enum.SendTargetType_SendToNone
	var data string
	switch msg.Type {
	case enum.MsgType_GetMnemonic_2004:
		mnemonic, err := service.HDWalletService.Bip39GenMnemonic(256)
		if err == nil { //get  successful.
			data = mnemonic
			status = true
		} else {
			data = err.Error()
		}
		client.sendToMyself(msg.Type, status, data)
		sendType = enum.SendTargetType_SendToSomeone
	case enum.MsgType_Mnemonic_CreateAddress_3000:
		wallet, err := service.HDWalletService.CreateNewAddress(client.User)
		if err == nil { //get  successful.
			bytes, err := json.Marshal(wallet)
			if err != nil {
				data = err.Error()
			} else {
				data = string(bytes)
				status = true
			}
			status = true
		} else {
			data = err.Error()
		}
		client.sendToMyself(msg.Type, status, data)
		sendType = enum.SendTargetType_SendToSomeone
	case enum.MsgType_Mnemonic_GetAddressByIndex_3001:
		index, err := strconv.Atoi(msg.Data)
		if err != nil || index < 0 {
			data = "error index"
		} else {
			wallet, err := service.HDWalletService.GetAddressByIndex(client.User, uint32(index))
			if err == nil { //get  successful.
				bytes, err := json.Marshal(wallet)
				if err != nil {
					data = err.Error()
				} else {
					data = string(bytes)
					status = true
				}
				status = true
			} else {
				data = err.Error()
			}
		}
		client.sendToMyself(msg.Type, status, data)
		sendType = enum.SendTargetType_SendToSomeone
	}
	return sendType, []byte(data), status
}
