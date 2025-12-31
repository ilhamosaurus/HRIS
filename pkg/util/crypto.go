package util

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ilhamosaurus/HRIS/pkg/setting"
)

type Hasher struct {
	Secret string
}

func NewHasher(secret string) *Hasher {
	return &Hasher{Secret: secret}
}

func (h *Hasher) GenerateSHAHash(password string) string {
	secret := []byte(setting.Server.Secret)
	passwordByte := []byte(password)

	hasher := sha256.New()
	passwordByte = append(passwordByte, secret...)
	hasher.Write(passwordByte)

	hashedPassword := hasher.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}

func (h *Hasher) VerifySHAHash(password, hashedPassword string) bool {
	return h.GenerateSHAHash(password) == hashedPassword
}
