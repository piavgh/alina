package signature

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	input := "Swapped (address,address,address,address,uint256, uint256)"
	want := common.HexToHash("0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8")
	got := Generate(input)

	assert.Equal(t, got, want)
}
