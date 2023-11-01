package statequery

import (
	"encoding/json"
	"math/big"

	"github.com/SundaeSwap-finance/ogmigo/v6/ouroboros/shared"
)

type EraStart struct {
	Time  EraSeconds `json:"time,omitempty"`
	Slot  big.Int    `json:"slot,omitempty"`
	Epoch big.Int    `json:"epoch,omitempty"`
}

type EraSeconds struct {
	Seconds big.Int `json:"seconds"`
}

type EraMilliseconds struct {
	Milliseconds big.Int `json:"milliseconds"`
}

type Utxo struct {
	Transaction UtxoTxID        `json:"transaction"`
	Index       uint32          `json:"index"`
	Address     string          `json:"address"`
	Value       shared.Value    `json:"value"`
	DatumHash   string          `json:"datumHash,omitempty"`
	Datum       string          `json:"datum,omitempty"`
	Script      json.RawMessage `json:"script,omitempty"`
}

type UtxoTxID struct {
	ID string `json:"id"`
}
