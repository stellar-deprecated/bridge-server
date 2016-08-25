// Copyright 2016 Stellar Development Foundation and contributors. Licensed
// under the Apache License, Version 2.0. See the COPYING file at the root
// of this distribution or at http://www.apache.org/licenses/LICENSE-2.0

package archivist

import (
	"testing"
	"encoding/json"
)

func TestUnmarshalState(t *testing.T) {

	var jsonBlob = []byte(`{
		"version": 1,
		"server": "v0.4.0-34-g2f015f6",
		"currentLedger": 2113919,
		"currentBuckets": [
			{
				"curr": "0000000000000000000000000000000000000000000000000000000000000000",
				"next": {
					"state": 0
				},
				"snap": "0000000000000000000000000000000000000000000000000000000000000000"
			},
			{
				"curr": "0000000000000000000000000000000000000000000000000000000000000000",
				"next": {
					"state": 1,
					"output": "0000000000000000000000000000000000000000000000000000000000000000"
				},
				"snap": "0000000000000000000000000000000000000000000000000000000000000000"
			}
		 ]
	}`)

	var state HistoryArchiveState
	err := json.Unmarshal(jsonBlob, &state)
	if err != nil {
		t.Error(err)
	} else if state.CurrentLedger != 2113919 {
		t.Error(state)
	}
}
