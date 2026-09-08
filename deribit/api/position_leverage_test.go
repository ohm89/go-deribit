package api

import (
	"encoding/json"
	"testing"
)

// 🔴 Deribit returns `leverage` as a JSON NUMBER, and for this account it is
// 50.0 — a float literal. Declaring the field as `int` makes encoding/json
// refuse the whole response:
//
//	json: cannot unmarshal number 50.0 into Go struct field
//	Position.result.leverage of type int
//
// Because GetPosition returns that error, every caller of it fails. In
// go-deribit-rebalance that is RestartThresholdSwingOrderImbalance ->
// GetAssetPositionSizeAndMarketValue -> GetPosition, which runs every 5 minutes
// and is what decides whether to place rebalancing orders. Observed live on
// gdrb-85a59587 (prod, BTC-PERPETUAL, ~1,100-contract short): the bot ran
// healthy for hours and placed NOTHING.
//
// The payload below is the real one from that account, trimmed.
func TestPositionUnmarshal_leverageMayBeFractional(t *testing.T) {
	const body = `{"jsonrpc":"2.0","result":{"size":-1.1e3,"kind":"future",` +
		`"direction":"sell","instrument_name":"BTC-PERPETUAL","leverage":50.0,` +
		`"average_price":68122.67,"mark_price":79920.94,"size_currency":-0.013763602}}`

	var resp GetPositionDetailsResponse
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		t.Fatalf("Deribit's own payload did not decode: %v", err)
	}
	if resp.Result.Leverage != 50 {
		t.Errorf("Leverage = %v, want 50", resp.Result.Leverage)
	}
	if resp.Result.InstrumentName != "BTC-PERPETUAL" {
		t.Errorf("InstrumentName = %q", resp.Result.InstrumentName)
	}
}
