package hash

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
	"strings"
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength = 6
)

func GenerateShortCode(originalURL string) string {
	hash := md5.Sum([]byte(originalURL))
	hashString := hex.EncodeToString(hash[:])

	hashBig := new(big.Int)
	hashBig.SetString(hashString, 16)

	return toBase62(hashBig, codeLength)
}

func toBase62(num *big.Int, length int) string {
	if num.Sign() == 0 {
		return strings.Repeat("a", length)
	}

	base := big.NewInt(int64(len(charset)))
	result := make([]byte, 0, length)

	for num.Sign() > 0 {
		remainder := new(big.Int)
		num.DivMod(num, base, remainder)
		result = append([]byte{charset[remainder.Int64()]}, result...)
	}

	for len(result) < length {
		result = append([]byte{'a'}, result...)
	}

	if len(result) > length {
		result = result[:length]
	}

	return string(result)
}
