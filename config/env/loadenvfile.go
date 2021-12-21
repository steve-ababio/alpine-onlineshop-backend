package env
import(
	"github.com/joho/godotenv"
)

//Loadenvfile gets env file
func Loadenvfile(){
	godotenv.Load(".env");
}