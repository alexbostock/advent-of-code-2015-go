package p4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

type Puzzle4 struct {
	key string
}

func (p *Puzzle4) LoadInput(input io.ReadSeeker) error {
	buffer, err := io.ReadAll(input)
	if err != nil {
		return nil
	}
	p.key = string(buffer)
	return nil
}

func (p *Puzzle4) SolvePart1() interface{} {
	return searchHashWithZeros(p.key, 5)
}

func (p *Puzzle4) SolvePart2() interface{} {
	return searchHashWithZeros(p.key, 6)
}

func searchHashWithZeros(prefix string, numZeros int) string {
	hashHex := "111111"
	isValid := func(hash string) bool { return hash[0:5] == "00000" }
	if numZeros == 6 {
		isValid = func(hash string) bool { return hash[0:6] == "000000" }
	}
	var i int
	for i = 0; !isValid(hashHex); i++ {
		buffer := []byte(fmt.Sprintf("%v%d", prefix, i))
		hash := md5.Sum(buffer)
		hashHex = hex.EncodeToString(hash[:])
	}
	return fmt.Sprintf("%d", i-1)
}
