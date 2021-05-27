package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimizedOrderCost(t *testing.T) {
	t.Run("minimized cost of an amount and exchange information test", func(t *testing.T) {
		price, cost, exchange, err := MinimizedOrderCost(1.2)
		require.Nil(t, err)
		require.NotEqual(t, 0, cost)
		require.NotEqual(t, 0, price)
		require.NotEqual(t, "", exchange)
	})
}
