package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type Hasher struct {
	secret []byte
}

func NewHasher(secret []byte) *Hasher {
	return &Hasher{secret: secret}
}

func (hr *Hasher) GetHash(msg string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(msg))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(hr.secret)), nil
}
