package middlewares

import (
	"log"
	"net/http"
)

// SetContentType is an unecessarily complex middlware function written
// like this because i really dont want to declare multiple functions
// and also go has closures soo everybody is just having fun.
// This simple function took me like an hour to write because...
// Errm type http.HandlerFunc is `func(http.ResponseWriter, *http.Request)`
// mux.Router.Use(...MiddlewareFunc) -> it takes a MiddleWareFunc, so what the fuck is
// a MiddleWareFunc => it is func(http.Handler) http.Handler .. alright.
// Okay wtf is http.Handler -> it is an interface{ ServeHTTP(http.ResponseWriter, *http.Request) }.
// are we done?.. i cant even tell. Okay so http.HandlerFunc implicitly satisfies the
// http.Handler interface.
// I guess i'm done now but i'm more confused than ever. :(
// chaley the thing didnt even work LOL!
func SetContentType(key, value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			next.ServeHTTP(w, r)
		})
	}
}

func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}
