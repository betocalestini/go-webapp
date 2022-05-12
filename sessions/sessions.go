package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("S3CR3TK3Y"))

func Flash(r *http.Request, w http.ResponseWriter) (string, string) {
	var message string = ""
	var danger string = ""
	session, _ := Store.Get(r, "session")
	untypedMessage := session.Values["MESSAGE"]
	untypedMessage2 := session.Values["DANGER"]
	// danger = session.Values["DANGER"].(string)
	message, ok := untypedMessage.(string)
	if !ok {
		return "", ""
	}
	danger, ok = untypedMessage2.(string)
	if !ok {
		return "", ""
	}
	delete(session.Values, "MESSAGE")
	delete(session.Values, "DANGER")
	session.Save(r, w)
	return message, danger
}
