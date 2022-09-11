package signature

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Generate(rawInput string) common.Hash {
	input := strings.ReplaceAll(rawInput, " ", "")

	return crypto.Keccak256Hash([]byte(input))
}
