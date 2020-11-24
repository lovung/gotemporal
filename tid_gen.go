package gotemporal

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

func GenTID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	tid := ulid.MustNew(ulid.Timestamp(t), entropy)
	return tid.String()
}
