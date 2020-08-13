package timer

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	"math"
	rand "math/rand"
	"time"
)

type Timer struct{}

func (t *Timer) CurrentTimestamp() int64 {
	rand.Seed(time.Now().UnixNano())
	x := int64(GenerateNumberCode(3))
	return time.Now().UnixNano() + x
}

func GenerateNumberCode(length int) (code int) {
	var src cryptoSource
	rnd := rand.New(src)
	max := int(math.Pow10(length))
	x := rnd.Intn(max)
	if x >= max {
		x = max - 1
	}
	code = x
	return
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
