package grapher

import "net/http"

type Explorer interface {
	http.Handler
}
