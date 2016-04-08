package memo

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/stellar/go-stellar-base/keypair"
)

type BlockType string

const (
	BlockTypeSimple    BlockType = "simple"
	BlockTypeEncrypted BlockType = "encrypted"
	BlockTypeEnvelope  BlockType = "envelope"
)

type Memo struct {
	SenderInfo     Block  `json:"sender_info"`
	OperationIndex int    `json:"op_index"`
	Note           Block  `json:"note"`
	Route          Block  `json:"route"`
	Blocks         []Memo `json:"blocks"`
}

type Block struct {
	Type      BlockType `json:"type"`
	Value     *string
	Encrypted *PackedBlockEncrypted
	Envelope  *PackedBlockEnvelope
}

type PackedBlockEncrypted struct {
	Method string `json:"method"`
	For    string `json:"for"`
	Value  string `json:"value"`
}

type PackedBlockEnvelope struct {
	From      string `json:"from"`
	Signature string `json:"sig"`
	Value     string `json:"value"`
}

// To implement json.Unmarshaler interface.
func (b *Block) UnmarshalJSON(bytes []byte) error {
	if bytes[0] == '"' {
		// String
		var value string
		err := json.Unmarshal(bytes, &value)
		if err != nil {
			return err
		}

		*b = Block{
			Type:  BlockTypeSimple,
			Value: &value,
		}
	} else if bytes[0] == '{' {
		// Packed block
		type UnmarshalBlockType struct {
			Type BlockType `json:"type"`
		}

		var blockType UnmarshalBlockType
		err := json.Unmarshal(bytes, &blockType)
		if err != nil {
			return err
		}

		switch blockType.Type {
		case BlockTypeEncrypted:
			// Encrypted block
			*b = Block{
				Type: BlockTypeEncrypted,
			}
		case BlockTypeEnvelope:
			// Signed block
			envelope := PackedBlockEnvelope{}
			err := json.Unmarshal(bytes, &envelope)
			if err != nil {
				return err
			}

			*b = Block{
				Type: BlockTypeEnvelope,
				Envelope: &envelope,
			}
		default:
			return errors.New("Invalid block type")
		}
	}
	return nil
}

// Returns the value of the block.
// If the block is encrypted block then `key` argument must be a private key to decrypt it.
// If the block is envelope block then `key` argument must be a public key to verify it.
func (b *Block) GetValue(key ...string) (value string, err error) {
	switch {
	case b.Type == BlockTypeSimple && b.Value != nil:
		return *b.Value, nil
	case b.Type == BlockTypeEncrypted && b.Encrypted != nil:
		return b.Encrypted.Value, nil
	case b.Type == BlockTypeEnvelope && b.Envelope != nil:
		kp, err := keypair.Parse(b.Envelope.From)
		if err != nil {
			return "", err
		}

		signatureBytes, err := base64.StdEncoding.DecodeString(b.Envelope.Signature)
		if err != nil {
			return "", err
		}

		err = kp.Verify([]byte(b.Envelope.Value), signatureBytes)
		if err != nil {
			return "", err
		}

		return b.Envelope.Value, nil
	default:
		return "", errors.New("Invalid block")
	}
}
