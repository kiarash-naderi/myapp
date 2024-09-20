package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/golang/protobuf/proto"
    "errors"
)

// MsgSendToken defines a message for sending tokens from one address to another
type MsgSendToken struct {
    Sender   string    `json:"sender" yaml:"sender"`
    Receiver string    `json:"receiver" yaml:"receiver"`
    Amount   sdk.Coins `json:"amount" yaml:"amount"`
}

type MsgSendTokenResponse struct{}

type Transaction struct {
    ID     string    `json:"id"`
    Amount sdk.Coins `json:"amount"`
    // سایر فیلدها می‌توانند اضافه شوند
}

// NewMsgSendToken creates a new MsgSendToken instance
func NewMsgSendToken(sender string, receiver string, amount sdk.Coins) MsgSendToken {
    return MsgSendToken{
        Sender:   sender,
        Receiver: receiver,
        Amount:   amount,
    }
}

// Route returns the name of the module
func (msg MsgSendToken) Route() string { return RouterKey }

// Type returns the action type of the message
func (msg MsgSendToken) Type() string { return "SendToken" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSendToken) ValidateBasic() error {
    if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
        return errors.New("invalid sender address")
    }
    if _, err := sdk.AccAddressFromBech32(msg.Receiver); err != nil {
        return errors.New("invalid receiver address")
    }
    if (!msg.Amount.IsAllPositive()) {
        return errors.New("amount must be positive")
    }
    return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSendToken) GetSignBytes() []byte {
    return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSendToken) GetSigners() []sdk.AccAddress {
    sender, err := sdk.AccAddressFromBech32(msg.Sender)
    if err != nil {
        panic(err)
    }
    return []sdk.AccAddress{sender}
}

// Implementing ProtoMessage for MsgSendToken
func (msg MsgSendToken) ProtoMessage() {}

func (msg MsgSendToken) Reset() {
    msg = MsgSendToken{}
}

func (msg MsgSendToken) String() string {
    return string(sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg)))
}

// Implementing ProtoMessage for Transaction
func (tx *Transaction) ProtoMessage() {}

func (tx *Transaction) Reset() {
    *tx = Transaction{}
}

func (tx *Transaction) String() string {
    return proto.CompactTextString(tx)
}

func (tx *Transaction) Marshal() ([]byte, error) {
    return proto.Marshal(tx)
}

func (tx *Transaction) Unmarshal(data []byte) error {
    return proto.Unmarshal(data, tx)
}

// RouterKey is used for routing the messages in the module
const RouterKey = ModuleName