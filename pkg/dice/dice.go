package dice

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strconv"

	"github.com/rs/zerolog/log"
)

func rollHelper(d int8) (result int8) {
	bigInt := big.NewInt(int64(d))
	roll, err := rand.Int(rand.Reader, bigInt)
	if err != nil {
		log.Error().Err(err).Msgf("could not get d%d", d)
		return -1
	}
	result = int8(roll.Int64()) + 1
	log.Debug().Int("int", int(d)).Int("result", int(result)).Msg("rolling...")
	return
}

type DiceFunction func() int8

func D4() int8 {
	return rollHelper(4)
}

func D6() int8 {
	return rollHelper(6)
}

func D8() int8 {
	return rollHelper(8)
}

func D10() int8 {
	return rollHelper(10)
}

func D12() int8 {
	return rollHelper(12)
}

func D20() int8 {
	return rollHelper(20)
}

func D100() int8 {
	return rollHelper(100)
}

func ParseDiceCmd(s string) (result int64) {

	re := regexp.MustCompile(`^([0-9]*)d([0-9]*)([+-]?)([0-9]*)$`)

	if re.MatchString(s) {
		log.Info().Str("cmd", s).Msg("winner winner chicken dinner")

		matches := re.FindStringSubmatch(s)

		var mult int64
		var err error
		if len(matches[1]) > 0 {
			mult, err = strconv.ParseInt(matches[1], 10, 8)
			if err != nil {
				log.Error().Err(err).Str("string", matches[1]).Msg("problem parsing integer value")
			}
			log.Debug().Int("mult", int(mult)).Msg("getting number of dice")
		}

		helper := func(limit int64, df DiceFunction) {
			log.Debug().Int("limit", int(limit)).Msgf("result... %+v", df)
			for n := limit; n > 0; n-- {
				result = result + int64(df())
			}
		}

		switch matches[2] {
		case "4":
			helper(mult, D4)
		case "6":
			helper(mult, D6)
		case "8":
			helper(mult, D8)
		case "10":
			helper(mult, D10)
		case "12":
			helper(mult, D12)
		case "20":
			helper(mult, D20)
		case "100", "percent":
			helper(mult, D100)
		default:
			log.Error().Msgf("unknown dice type %s", matches[2])
		}

		var amt int64
		if len(matches[4]) > 0 {
			amt, err = strconv.ParseInt(matches[4], 10, 8)
			if err != nil {
				log.Error().Err(err).Str("string", matches[4]).Msgf("problem parsing integer value")
			}
		}

		switch matches[3] {
		case "+":
			Add(&result, amt)
		case "-":
			Subtract(&result, amt)
		}

	} else {
		log.Error().Msg("invalid dice command")
		result = -1
	}

	return result
}

func Add(result *int64, n int64) {
	*result += n
}

func Subtract(result *int64, n int64) {
	*result -= n
}
