package util
import(
	"github.com/dgrijalva/jwt-go"
	"icourses/config"
	"icourses/defination"
	"time"
)
func SignJwtToken(user defination.User)(string,error){

	tokenStr := jwt.New(jwt.SigningMethodHS256)
	claims := defination.JwtToken{
		Uid: user.Uid,
		User: user.User,
		Type:user.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Config.JwtConfig.TokenDuration)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	tokenStr.Claims = claims
	token, err := tokenStr.SignedString([]byte(config.Config.JwtConfig.Secret))
	if err!=nil{
		return "",err
	}
	return token,nil
}