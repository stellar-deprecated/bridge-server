// Copyright 2016 Stellar Development Foundation and contributors. Licensed
// under the Apache License, Version 2.0. See the COPYING file at the root
// of this distribution or at http://www.apache.org/licenses/LICENSE-2.0

package archivist

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheckpointPrefix(t *testing.T) {
	assert.Equal(t, DirPrefix{0xaa, 0xbb, 0xcc}, CheckpointPrefix(0xaabbcc12))
}

func TestDirPrefixPath(t *testing.T) {
	assert.Equal(t, "aa/bb/cc", (&DirPrefix{0xaa, 0xbb, 0xcc}).Path())
}

func TestRangePaths(t *testing.T) {
	r := Range{ Low:0x0010001f, High:0x0014001b, }
	assert.Equal(t, RangePaths(r), []string{
		"00/10",
		"00/11",
		"00/12",
		"00/13",
		"00/14",
	})
	r = Range{ Low:0x00100000, High:0x0010ff00, }
	rps := RangePaths(r)
	assert.Equal(t, rps[0], "00/10/00")
	assert.Equal(t, rps[255], "00/10/ff")
}
