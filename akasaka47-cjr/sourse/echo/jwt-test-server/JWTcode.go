package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"math/big"
	"time"
)

// 加密
func NewTokens(name string, id int) string {
	key := "akasaka"
	fmt.Println("time1:", time.Now().Unix())
	fmt.Println("time2:", time.Now().Add(time.Minute*15).Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 自定义键值对
		"una": name,
		"uid": id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(key)) // 传入秘钥 []byte() 拿到string
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(tokenString)
	return tokenString
}

// 解密
func decode(tokenString string) {
	// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJ4eHgiLCJuYmYiOjE0NDQ0Nzg0MDB9.YXnq53IqsbOVMRJ8vfYFPqbRH2HLM9ezwwQo6TqHBW0"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("akasaka"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 获取分区内容
		fmt.Println("okk", claims["uid"], claims["iat"])
	} else {
		fmt.Println(err)
	}
}

//func random() int{
//	rand.Seed(time.Now().UnixNano())
//	return rand.Intn(1000)
//}

func CreateRandomString(len int) string {
	var container string
	//var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var str = "1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func Redis() (redis.Conn, error) {
	re, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return re, err
	}
	defer re.Close()
	return re, nil
}

func SetRedis(re redis.Conn, mail string, check string, time string) error {
	_, err := re.Do("SET", "mail", mail, "EX", time)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	_, err = re.Do("SET", "check", check, "EX", time)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	fmt.Println("set-redis success")
	return nil
}

func FindRedis(re redis.Conn, mail string, check string) (bool, error) {
	m, err := redis.String(re.Do("GET", "mail"))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return false, err
	}
	c, err := redis.String(re.Do("GET", "check"))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return false, err
	}
	fmt.Println(m)
	fmt.Println(c)
	if mail == m || c == check {
		return true, nil
	}
	return false, nil
}

