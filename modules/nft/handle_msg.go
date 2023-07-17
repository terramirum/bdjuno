package nft

import (
	"encoding/base64"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	juno "github.com/forbole/juno/v5/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *nfttypes.MsgSend:
		return m.HandleMsgSend(index, tx, cosmosMsg)
	}

	return nil
}

func (m *Module) HandleMsgSend(index int, tx *juno.Tx, msg *nfttypes.MsgSend) error {
	// Get Execute Contract event
	event, err := tx.FindEventByType(index, nft.TypeMsgSend)
	if err != nil {
		return fmt.Errorf("error while searching for EventTypeExecute: %s", err)
	}

	// Get result data
	resultData, err := tx.FindAttributeByKey(event, nft.TypeMsgSend)
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

	return m.db.SaveNftSend(msg.Sender, msg.Receiver, msg.ClassId, msg.Id,
		string(resultDataBz), tx.Height, timestamp)
}
