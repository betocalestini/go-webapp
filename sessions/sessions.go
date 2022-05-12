package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("S3CR3TK3Y"))

func Flash(r *http.Request, w http.ResponseWriter) (string, string, string) {
	var message, alert, active string = "", "", ""

	session, _ := Store.Get(r, "session")
	untypedMessage := session.Values["MESSAGE"]
	message, ok := untypedMessage.(string)
	if !ok {
		return "", "", ""
	}
	untypedAlert := session.Values["ALERT"]
	alert, ok = untypedAlert.(string)
	if !ok {
		return "", "", ""
	}
	untypedActive := session.Values["ACTIVE"]
	active, ok = untypedActive.(string)
	if !ok {
		return "", "", ""
	}
	delete(session.Values, "MESSAGE")
	delete(session.Values, "ALERT")
	delete(session.Values, "ACTIVE")
	session.Save(r, w)
	return message, alert, active
}
