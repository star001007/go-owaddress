package tv

import (
	"github.com/blocktree/go-owcrypt"

	"github.com/star001007/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "tv"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		Prefix = []byte{0x1D, 0x3B}
	)

	if address == "" {
		return false
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 26 {
		return false
	}

	if decodeBytes[0] != Prefix[0] || decodeBytes[1] != Prefix[1] {
		return false
	}

	check := owcrypt.Hash(owcrypt.Hash(decodeBytes[:22], 32, owcrypt.HASH_ALG_BLAKE2B), 32, owcrypt.HASH_ALG_KECCAK256)[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[22+i] {
			return false
		}
	}

	return true
}
