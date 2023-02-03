package extractors

import (
	"net/http"
	"strings"
)

func ExtractUserIP(req *http.Request) string {
	IPArray := req.Header["X-Forwarded-For"]
	backupIPArray := req.Header["X-Appengine-Remote-Addr"]
	if len(IPArray) < 1 {
		if len(backupIPArray) < 1 {
			return ""
		}
		return backupIPArray[0]
	} else {
		IPWithPotentialProxies := IPArray[0]
		primaryIP := strings.Split(IPWithPotentialProxies, ",")
		return primaryIP[0]
	}
}
