package dice

import (
	"crypto/rand"
	"math/big"

	"github.com/rs/zerolog/log"
)

func rollHelper(d int8) int8 {
	bigInt := big.NewInt(d)
	roll, err := rand.Int(rand.Reader, bigInt)
	if err != nil {
		log.Error().Err(err).Msgf("could not get d%d", d)
		return -1
	}
	return int8(roll.Int64())
}
func d4() int8 {
	return rollHelper(4)
}

func d6() int8 {
	return rollHelper(6)
}

func d8() int8 {
	return rollHelper(8)
}

func d10() int8 {
	return rollHelper(10)
}

func d12() int8 {
	return rollHelper(12)
}

func d20() int8 {
	return rollHelper(20)
}

func d100() int8 {
	return rollHelper(100)
}
