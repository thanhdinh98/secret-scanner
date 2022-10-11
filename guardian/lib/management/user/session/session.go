package session

import (
	"context"
	"encoding/hex"
	"time"

	"guardian/common"
	"guardian/common/cache"
	"guardian/common/log"
	"guardian/guardian/lib/management/user/accesstoken"
	"guardian/guardian/lib/management/user/password"
)

type Session struct {
	Email       string
	Password    *password.Password
	AccessToken *accesstoken.AccessToken
	Out         *log.Logger

	cacheDuration time.Duration
	mem           cache.Handler
}

func Initialize(email, plaintext string) *Session {
	session := &Session{
		Email:         email,
		mem:           cache.Get(),
		cacheDuration: 24 * time.Hour,
	}
	session.InitLogger()
	session.InitAccessToken()
	session.InitPassword(plaintext)
	return session
}

func (session *Session) InitPassword(plaintext string) {
	session.Password = &password.Password{
		Hash: password.DefaultHash,
	}
	session.Password.SetPlaintext(plaintext)
}

func (session *Session) InitLogger() {
	session.Out = &log.Logger{}
}

func (session *Session) InitAccessToken() {
	session.AccessToken = &accesstoken.AccessToken{
		Hash: accesstoken.DefaultHash,
	}
	session.AccessToken.SetKey(common.RandomBytesF(8))
}

func (session Session) Encode(message string) string {
	return hex.EncodeToString([]byte(session.Email + ":" + message))
}

func (session Session) Memorize(ctx context.Context) error {
	clonedSession := Session{
		Email:       session.Email,
		Password:    session.Password,
		AccessToken: session.AccessToken,
	}
	return session.mem.Set(ctx, session.Email, clonedSession, session.cacheDuration)
}

func (session Session) IsExists(ctx context.Context) (bool, error) {
	return session.mem.Exists(ctx, session.Email)
}

func (session *Session) GetFromCache(ctx context.Context) (err error) {
	var clonedSession *Session
	if err = session.mem.Get(ctx, session.Email, &clonedSession); err != nil {
		return
	}
	session.Password = clonedSession.Password
	session.AccessToken = clonedSession.AccessToken
	return
}
