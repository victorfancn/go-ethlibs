package eth_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/INFURA/ethereum-interaction/pkg/eth"
	"github.com/INFURA/ethereum-interaction/pkg/jsonrpc"
)

func TestQuantityFromUInt64(t *testing.T) {
	q := eth.QuantityFromUInt64(uint64(0x1234))

	b, err := json.Marshal(&q)
	require.NoError(t, err)
	require.Equal(t, []byte(`"0x1234"`), b)

	params := jsonrpc.MustParams(&q, true)
	b, err = json.Marshal(&params)
	require.NoError(t, err)

	require.Equal(t, []byte(`["0x1234",true]`), b)

	zero := eth.Quantity{}
	b, err = json.Marshal(&zero)
	require.NoError(t, err)
	require.Equal(t, []byte(`"0x0"`), b)

	fromBig := eth.QuantityFromBigInt(big.NewInt(0x4567))
	b, err = json.Marshal(&fromBig)
	require.NoError(t, err)

	require.Equal(t, []byte(`"0x4567"`), b)

	fromUI64 := eth.QuantityFromUInt64(0xABCD)
	b, err = json.Marshal(&fromUI64)
	require.NoError(t, err)

	require.Equal(t, []byte(`"0xabcd"`), b)

	require.Equal(t, fromUI64.Big().Int64(), eth.QuantityFromInt64(fromUI64.Int64()).Int64())

	{
		invalid, err := eth.NewQuantity("bad")
		require.Error(t, err)
		require.Nil(t, invalid)
	}

	{
		invalid, err := eth.NewQuantity("0xinvalid")
		require.Error(t, err)
		require.Nil(t, invalid)
	}

	{
		invalid, err := eth.NewQuantity("0x00")
		require.Error(t, err)
		require.Nil(t, invalid)
	}

	{
		invalid, err := eth.NewQuantity("0x")
		require.Error(t, err)
		require.Nil(t, invalid)
	}

	{
		invalid, err := eth.NewQuantity("0")
		require.Error(t, err)
		require.Nil(t, invalid)
	}
}
