package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// GenerateRoomId creates a unique room ID
func GenerateRoomId() (roomID string) {
	id := uuid.New()
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(999999)
	roomID = fmt.Sprintf("%s-%06d", id, rn)
	return
}
