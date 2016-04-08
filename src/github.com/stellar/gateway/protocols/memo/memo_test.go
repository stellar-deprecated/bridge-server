package memo

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestMemo(t *testing.T) {
	Convey("Memo", t, func() {
		Convey("Unmarshal simple", func() {
			preimage := `{
    "note": "Hey thanks for all the Kelp.",
    "route": "6387622"
}`

			var memo Memo
			err := json.Unmarshal([]byte(preimage), &memo)
			if err != nil {
				panic(err)
			}

			noteValue, err := memo.Note.GetValue()
			if err != nil {
				panic(err)
			}

			routeValue, err := memo.Route.GetValue()
			if err != nil {
				panic(err)
			}

			assert.Equal(t, "Hey thanks for all the Kelp.", noteValue)
			assert.Equal(t, "6387622", routeValue)
		})

		Convey("Unmarshal envelope block", func() {
			// GCC7ZATN5OXH6XGEFM5WPTMA5NEJ4LT335FCKHW6GYCMKM3ZVZ5BJGGA
			// SCLJ7EDOWQGU7LQZC6KED25ZLTOLU36KRJLBSHRUT7Q7SOU26O3JVZYV
			
			preimage := `{
    "note" : "Note.",
    "blocks": [{
        "sender_info": {
            "type": "envelope",
            "from": "GCC7ZATN5OXH6XGEFM5WPTMA5NEJ4LT335FCKHW6GYCMKM3ZVZ5BJGGA",
            "sig": "68+umrzx9fIzhg/QIl73/0rP7FFYlDryDgEUSZuKtmd8KYX70QJ4guZx5yizH1V73epWcohMYPAUwBmFk8fDCA==",
            "value": "Sender Info JSON"
        }
    }]
}`

			var memo Memo
			err := json.Unmarshal([]byte(preimage), &memo)
			if err != nil {
				panic(err)
			}

			noteValue, err := memo.Note.GetValue()
			if err != nil {
				panic(err)
			}

			verifiedValue, err := memo.Blocks[0].SenderInfo.GetValue()
			if err != nil {
				panic(err)
			}

			assert.Equal(t, "Note.", noteValue)
			assert.Equal(t, BlockTypeEnvelope, memo.Blocks[0].SenderInfo.Type)
			assert.Equal(t, "GCC7ZATN5OXH6XGEFM5WPTMA5NEJ4LT335FCKHW6GYCMKM3ZVZ5BJGGA", memo.Blocks[0].SenderInfo.Envelope.From)
			assert.Equal(t, "68+umrzx9fIzhg/QIl73/0rP7FFYlDryDgEUSZuKtmd8KYX70QJ4guZx5yizH1V73epWcohMYPAUwBmFk8fDCA==", memo.Blocks[0].SenderInfo.Envelope.Signature)
			assert.Equal(t, "Sender Info JSON", memo.Blocks[0].SenderInfo.Envelope.Value)
			assert.Equal(t, "Sender Info JSON", verifiedValue)
		})

		Convey("Unmarshal envelope block (invalid signature)", func() {
			preimage := `{
    "note" : "Note.",
    "blocks": [{
        "sender_info": {
            "type": "envelope",
            "from": "GCC7ZATN5OXH6XGEFM5WPTMA5NEJ4LT335FCKHW6GYCMKM3ZVZ5BJGGA",
            "sig": "68+umrzx8fIzhg/QIl73/0rP7FFYlDryDgEUSZuKtmd8KYX70QJ4guZx5yizH1V73epWcohMYPAUwBmFk8fDCA==",
            "value": "Sender Info JSON"
        }
    }]
}`

			var memo Memo
			err := json.Unmarshal([]byte(preimage), &memo)
			if err != nil {
				panic(err)
			}

			noteValue, err := memo.Note.GetValue()
			if err != nil {
				panic(err)
			}

			assert.Equal(t, "Note.", noteValue)
			assert.Equal(t, BlockTypeEnvelope, memo.Blocks[0].SenderInfo.Type)
			assert.Equal(t, "GCC7ZATN5OXH6XGEFM5WPTMA5NEJ4LT335FCKHW6GYCMKM3ZVZ5BJGGA", memo.Blocks[0].SenderInfo.Envelope.From)
			assert.Equal(t, "Sender Info JSON", memo.Blocks[0].SenderInfo.Envelope.Value)

			_, err = memo.Blocks[0].SenderInfo.GetValue()
			assert.NotNil(t, err)
			assert.Equal(t, "signature verification failed", err.Error())
		})
	})
}
