package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

var entropy = ulid.Monotonic(
	rand.New(rand.NewSource(time.Now().UnixNano())),
	0,
)

func NewID(prefix string) string {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	return prefix + "_" + strings.ToLower(id)
}
