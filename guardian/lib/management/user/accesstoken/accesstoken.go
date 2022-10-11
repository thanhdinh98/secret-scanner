package accesstoken

import (
	"crypto"
	"crypto/hmac"
	"encoding/hex"
)

type AccessToken struct {
	Hash crypto.Hash
	Raw  string
	key  []byte
}

func (h *AccessToken) SetKey(key []byte) {
	h.key = key
}

func (h *AccessToken) GenerateFrom(message []byte) string {
	mac := hmac.New(h.Hash.New, h.key)
	mac.Write(message)
	h.Raw = hex.EncodeToString(mac.Sum(nil))
	return h.Raw
}

func (h *AccessToken) IsSameToken(token []byte) bool {
	mac := hmac.New(h.Hash.New, h.key)
	mac.Write(token)
	macEncoded := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(h.Raw), []byte(macEncoded))
}
