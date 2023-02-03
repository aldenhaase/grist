package serializers

import "server/lystrTypes"

func PackSCookie(source lystrTypes.SessionCookie) string {
	return source.Username + "|" + source.Expiration + "|" + source.Signature
}
func PackRCookie(source lystrTypes.RegistrationCookie) string {
	return source.UserIP + "|" + source.Expiration + "|" + source.Signature
}
