package cors
import "net/http"

//Enablecors Enable cors
func Enablecors(res *http.ResponseWriter){
	(*res).Header().Set("Access-Control-Allow-Origin","myhost");
}