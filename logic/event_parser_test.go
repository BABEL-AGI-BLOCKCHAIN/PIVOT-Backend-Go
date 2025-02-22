package logic

import (
	"math/big"
	"testing"
)

func TestAmountConversion(t *testing.T) {
	amount := new(big.Int)
	amount.SetString("10000000000000000000", 10)

	amountInt64 := amount.Int64()
	amountInt := int(amountInt64)

	t.Logf("Original amount: %s\n", amount.String())
	t.Logf("Converted to int64: %d\n", amountInt64)
	t.Logf("Converted to int: %d\n", amountInt)

	expectedInt64 := int64(922337203685477587)
	if amountInt64 != expectedInt64 {
		t.Errorf("Expected int64: %d, but got: %d", expectedInt64, amountInt64)
	}

	expectedInt := int(expectedInt64)
	if amountInt != expectedInt {
		t.Errorf("Expected int: %d, but got: %d", expectedInt, amountInt)
	}
}
