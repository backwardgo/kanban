package passwords

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/backwardgo/kanban/env"

	"golang.org/x/crypto/pbkdf2"
)

const iter = 4096

var salt []byte

func init() {
	salt = []byte(env.PASSWORD_SALT)

	if len(salt) < 32 {
		panic("invalid PASSWORD_SALT")
	}
}

func Digest(pass string) []byte {
	return pbkdf2.Key([]byte(pass), salt, iter, sha256.Size, sha256.New)
}

func Equal(pass1, pass2 []byte) bool {
	return len(pass1) > 0 && len(pass2) > 0 && subtle.ConstantTimeCompare(pass1, pass2) == 1
}
