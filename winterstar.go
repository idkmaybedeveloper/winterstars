// lain (c) 2026
//
// this Go code is one of the openDS components
// decided to pull it out into a separate module since i think
// that someday itll be useful to someone other than just me
package winterstars

import (
	"crypto/rand"
	"sync"
	"time"
)

// epoch is 2015-01-01 00:00:00 UTC in milliseconds - same /epoch/ discord uses.
const epoch int64 = 1420070400000

var (
	mu     sync.Mutex
	lastMs int64
	seq    int64
)

// next returns a discord-compatible snowflake id string.
// layout (63 usable bits):
// 41 bits - ms since Discord epoch
// 10 bits - node (fixed 1)
// 12 bits - per-ms sequence counter
func Next() string {
	mu.Lock()
	now := time.Now().UnixMilli() - epoch
	if now == lastMs {
		seq = (seq + 1) & 0xFFF
		if seq == 0 {
			/* seq rolled over - wait for next ms... */
			for now <= lastMs {
				now = time.Now().UnixMilli() - epoch
			}
		}
	} else {
		seq = 0
	}
	lastMs = now
	id := (now << 22) | (1 << 12) | seq
	mu.Unlock()

	return int64tostring(id)
}

func int64tostring(v int64) string {
	if v == 0 {
		return "0"
	}
	// v is ALWAYS pos (discord epoch is in the past)
	u := uint64(v)
	buf := [20]byte{}
	i := len(buf)
	for u > 0 {
		i--
		buf[i] = byte('0' + u%10)
		u /= 10
	}
	return string(buf[i:])
}

// WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALER
// WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALER
// SessionID returns an opaque session identifier safe for clients that may
// attempt to BigInt/big.js-parse it. We deliberately avoid the letter 'e'
// (hex session IDs routinely contain 'e', which trips parsers that treat the
// string as a number with scientific-notation exponent: e.g. "1ecafe..." is
// read as 1e<cafe> -> "NaN is not a valid exponent").
// Format: 32 chars from the set [0-9a-df-z] (base32-ish without 'e').
func SessionID() string {
	const alphabet = "0123456789abcdfghijklmnopqrstuvwxyz" // no 'e'
	var raw [32]byte
	_, _ = rand.Read(raw[:])
	out := make([]byte, 32)
	for i, b := range raw {
		out[i] = alphabet[int(b)%len(alphabet)]
	}
	return string(out)
}

// WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALER
// WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALERT // WHEELCHAIR_ALER
