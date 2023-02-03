package auth

import (
	"net/http"
	"os"
	"server/lystrTypes"
	"server/serializers"
	"time"
)

func GenerateSessionCookie(userInfo lystrTypes.UserQuery) *http.Cookie {

	expiration, formatedExpiration := getExpiration()
	signature, _ := generateSignature(userInfo.Username, formatedExpiration)

	val := serializers.PackSCookie(lystrTypes.SessionCookie{
		Username:   userInfo.Username,
		Expiration: formatedExpiration,
		Signature:  signature,
	})
	return &http.Cookie{
		Name:     lystrTypes.SCookie_t,
		Value:    val,
		Expires:  expiration,
		Secure:   true,
		HttpOnly: true,
		Domain:   os.Getenv("DOMAIN"),
		SameSite: http.SameSiteStrictMode,
	}
}

func GenerateRegCookie(userIP string) *http.Cookie {
	expiration, formatedExpiration := getExpiration()
	signature, _ := generateSignature(userIP, formatedExpiration)

	regCookie := serializers.PackRCookie(lystrTypes.RegistrationCookie{
		UserIP:     userIP,
		Expiration: formatedExpiration,
		Signature:  signature,
	})
	return &http.Cookie{
		Name:     lystrTypes.RCookie_t,
		Value:    regCookie,
		Expires:  expiration,
		Secure:   true,
		HttpOnly: true,
		Domain:   os.Getenv("DOMAIN"),
		SameSite: http.SameSiteStrictMode,
	}
}

func getExpiration() (time.Time, string) {
	expiration := time.Now().AddDate(0, 0, 15)
	return expiration, expiration.Format(time.RFC3339)
}
