package tmquery_test

import (
	"testing"

	query "github.com/fibonacci-chain/core/types/tmquery"
	"github.com/tendermint/tendermint/abci/types"
)

const testQuery = `tm.events.type='NewBlock' AND abci.account.name='Igor'`

var testEvents = []types.Event{
	{
		Type: "tm.events",
		Attributes: []types.EventAttribute{{
			Key:   []byte("index"),
			Value: []byte("25"),
		}, {
			Key:   []byte("type"),
			Value: []byte("NewBlock"),
		}},
	},
	{
		Type: "abci.account",
		Attributes: []types.EventAttribute{{
			Key:   []byte("name"),
			Value: []byte("Anya"),
		}, {
			Key:   []byte("name"),
			Value: []byte("Igor"),
		}},
	},
}

func BenchmarkParseCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := query.New(testQuery)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMatchCustom(b *testing.B) {
	q, err := query.New(testQuery)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !q.Matches(testEvents) {
			b.Error("no match")
		}
	}
}
