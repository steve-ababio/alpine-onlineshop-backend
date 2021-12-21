package token

import (
	"config/env"
	"fmt"
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken creates a token and returns it
func CreateToken(userinfo interface{}) (string, error) {
	env.Loadenvfile()
	secret := os.Getenv("TOKEN_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims);
	claims["authorized"] = true;
	claims["user"] = userinfo;
	claims["exp"] = time.Now().Add(time.Hour * 24 * 28);

	tokenstring, err := token.SignedString([]byte(secret))
	return tokenstring, err
}

func VerifyToken(token_str string)(*jwt.Token, error){
	token,err := jwt.Parse(token_str,func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")),nil;
	})
	if err != nil {
		return nil, err
	 }
	 return token,nil;
	 
}
func ValidateToken(token *jwt.Token)bool{
	if _,ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid{
		return false;
	} 
	return true;
} 
func Getcustompayloadclaims(tokenstr string,ch chan interface{}){
	 token,err := VerifyToken(tokenstr);
	 if err != nil{
		ch<-nil;
	 }
	 isvalid := ValidateToken(token);

	 if !isvalid{
		fmt.Println("malformed token");
		ch<-nil;
	 }
	 claims,ok := token.Claims.(jwt.MapClaims);

	 if ok {
		userclaims,ok := claims["user"].(map[string]interface{});
		if !ok {
			fmt.Println("malformed token");
			ch<-nil;
		}
		ch <- userclaims["ID"];
	}
	ch <- nil;
}

