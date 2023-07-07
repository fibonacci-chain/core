package tests

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/fibonacci-chain/core/testutil/processblock"
	"github.com/fibonacci-chain/core/testutil/processblock/verify"
)

func TestEpoch(t *testing.T) {
	app := processblock.NewTestApp()
	_ = processblock.CommonPreset(app)
	app.FastEpoch()
	for i, testCase := range []TestCase{
		{
			description: "first epoch",
			input:       []signing.Tx{},
			verifier: []verify.Verifier{
				verify.Epoch,
			},
			expectedCodes: []uint32{},
		},
		{
			description: "second epoch",
			input:       []signing.Tx{},
			verifier: []verify.Verifier{
				verify.Epoch,
			},
			expectedCodes: []uint32{},
		},
	} {
		if i > 0 {
			time.Sleep(6 * time.Second)
		}
		testCase.run(t, app)
	}
}
