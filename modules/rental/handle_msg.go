package rental

import (
	"encoding/base64"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"
	rentaltypes "github.com/terramirum/mirumd/x/rental/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *rentaltypes.MsgDeployNftRequest:
		return m.HandleMsgDeployNftRequest(index, tx, cosmosMsg)
	case *rentaltypes.MsgMintNftRequest:
		return m.HandleMsgMintNftRequest(index, tx, cosmosMsg)
	case *rentaltypes.MsgBurnNftRequest:
		return m.HandleMsgBurnNftRequest(index, tx, cosmosMsg)
	case *rentaltypes.MsgMintRentRequest:
		return m.HandleMsgMintRentRequest(index, tx, cosmosMsg)
	case *rentaltypes.MsgBurnRentRequest:
		return m.HandleMsgBurnRentRequest(index, tx, cosmosMsg)
	case *rentaltypes.MsgAccessNftRequest:
		return m.HandleMsgAccessNftRequest(index, tx, cosmosMsg)
	case *rentaltypes.MsgRentGiveAccessRequest:
		return m.HandleMsgRentGiveAccessRequest(index, tx, cosmosMsg)
	case *rentaltypes.MsgSendSessionRequest:
		return m.HandleMsgSendSessionRequest(index, tx, cosmosMsg)
	}

	return nil
}

func (m *Module) HandleMsgDeployNftRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgDeployNftRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeDeployNft)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyClassId)
	if err != nil {
		resultData = ""
	}
	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveDeployNft(msg.ContractOwner, msg.Name, msg.Symbol, msg.Description, msg.Uri, msg.Detail.JsonData,
		string(resultDataBz), tx.Height, timestamp)
}

func (m *Module) HandleMsgMintNftRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgMintNftRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeMintNft)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyNftId)
	if err != nil {
		resultData = ""
	}

	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveMintNft(msg.ContractOwner, msg.ClassId, msg.NftId, msg.Reciever, msg.Uri, msg.Detail.JsonData,
		string(resultDataBz), tx.Height, timestamp)
}

func (m *Module) HandleMsgBurnNftRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgBurnNftRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeBurnNft)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyNftId)
	if err != nil {
		resultData = ""
	}

	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveBurnNft(msg.ContractOwner, msg.ClassId, msg.NftId,
		string(resultDataBz), tx.Height, timestamp)
}

func (m *Module) HandleMsgMintRentRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgMintRentRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeRentNft)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyNftId)
	if err != nil {
		resultData = ""
	}

	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveRentMintNft(msg.ContractOwner, msg.ClassId, msg.NftId, msg.StartDate, msg.EndDate,
		string(resultDataBz), tx.Height, timestamp)
}

func (m *Module) HandleMsgBurnRentRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgBurnRentRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeBurnRentNft)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyNftId)
	if err != nil {
		resultData = ""
	}

	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveRentBurnNft(msg.ContractOwner, msg.ClassId, msg.NftId, msg.SessionId,
		string(resultDataBz), tx.Height, timestamp)
}

func (m *Module) HandleMsgAccessNftRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgAccessNftRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeAccessNft)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyNftId)
	if err != nil {
		resultData = ""
	}

	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveAccessNft(msg.ClassId, msg.NftId, msg.Renter,
		string(resultDataBz), tx.Height, timestamp)
}

func (m *Module) HandleMsgRentGiveAccessRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgRentGiveAccessRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeAccessGiveNft)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyNftId)
	if err != nil {
		resultData = ""
	}

	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveGiveAccessNft(msg.ClassId, msg.NftId, msg.SessionId, msg.Renter, msg.NewRenter,
		string(resultDataBz), tx.Height, timestamp)
}

func (m *Module) HandleMsgSendSessionRequest(index int, tx *juno.Tx, msg *rentaltypes.MsgSendSessionRequest) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, rentaltypes.EventTypeRentSend)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, rentaltypes.AttributeKeyNftId)
	if err != nil {
		resultData = ""
	}

	resultDataBz, err := base64.StdEncoding.DecodeString(resultData)
	if err != nil {
		return fmt.Errorf("error while decoding result data: %s", err)
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return fmt.Errorf("error while parsing time: %s", err)
	}

	return m.db.SaveSendRent(msg.ClassId, msg.NftId, msg.SessionId, msg.FromRenter, msg.ToRenter,
		string(resultDataBz), tx.Height, timestamp)
}
