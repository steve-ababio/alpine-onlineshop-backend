package cookieconfig


import(
	"config"
	"os"
	"github.com/gorilla/sessions"
)

//Getcookiestore creates cookie store and returns it 
func Getcookiestore() *sessions.CookieStore{
	config.Loadenvfile();
	cookiesecret := os.Getenv("COOKIE_SECRET");
	key := []byte(cookiesecret);
	store := sessions.NewCookieStore(key); 
	return store;
}