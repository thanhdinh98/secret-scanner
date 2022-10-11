package password

import (
	"crypto"
	"encoding/hex"
)

type Password struct {
	Hash      crypto.Hash
	Raw       string
	plaintext string
}

func (p *Password) SetPlaintext(plaintext string) {
	p.plaintext = plaintext
}

func (p *Password) GenerateHash() string {
	hash := p.Hash.New()
	hash.Write([]byte(p.plaintext))
	p.Raw = hex.EncodeToString(hash.Sum(nil))
	return p.Raw
}
