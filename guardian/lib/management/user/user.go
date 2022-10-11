package user

import (
	"context"
	"fmt"

	"guardian/guardian/lib/management/user/session"
)

func GenerateAccessToken(ctx context.Context, sess *session.Session) (hmac string, err error) {
	exsits, err := sess.IsExists(ctx)
	if err != nil || exsits {
		err = fmt.Errorf("session status: %v | err=%s", exsits, err.Error())
		err = fmt.Errorf("session is exsits or failed | err=%s", err.Error())
		return
	}
	encodedKey := sess.Encode(sess.Password.GenerateHash())
	hmac = sess.AccessToken.GenerateFrom([]byte(encodedKey))
	if err = sess.Memorize(ctx); err != nil {
		err = fmt.Errorf("save session failed | err=%s", err.Error())
		return
	}
	return
}

func CheckSessionExists(ctx context.Context, sess *session.Session) (exsits bool, err error) {
	if err = sess.GetFromCache(ctx); err != nil {
		err = fmt.Errorf("get session from cache failed | err=%s", err.Error())
		return
	}
	exsits = true
	return
}
