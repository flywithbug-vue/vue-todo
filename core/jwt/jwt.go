package jwt

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/flywithbug/log4go"
	"os"
	"time"
)

var jwtAuthBackend *JWTAuthenticationBackend

func JWTAuthKey() *JWTAuthenticationBackend {
	if jwtAuthBackend == nil{
		jwtAuthBackend = new(JWTAuthenticationBackend)
	}
	return jwtAuthBackend
}

//读取密钥文件
func ReadSigningKey(private_path, public_path string)  {
	JWTAuthKey().privateKey = getPrivateKey(private_path)
	JWTAuthKey().PublicKey = getPublicKey(public_path)
}

const (
	notBeforeDuration  = 1000
	expiresOffset  = 3600*24*7
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

// 一些常量
var (
	TokenExpired       = errors.New("Token is expired")
	TokenNotValidYet   = errors.New("Token not active yet")
	TokenMalformed     = errors.New("That's not even a token")
	TokenInvalid       = errors.New("Couldn't handle this token:")
)

//自定义载荷
type CustomClaims struct {
	jwt.StandardClaims
	ID   		string  	`json:"user_id"`
	Account 	string		`json:"account"`
}

func NewCustomClaims(user_id,account string)CustomClaims  {
	now := time.Now().Unix()
	claims := CustomClaims{
		ID:user_id,
		Account:account,
		StandardClaims:jwt.StandardClaims{
			NotBefore:int64(now - notBeforeDuration), 	// 	签名生效时间
			ExpiresAt:int64(now + expiresOffset),	// 	过期时间 一小时
			IssuedAt:now,
			Issuer:"flyWithBug",					   	//	签名的发行者
		},
	}
	return claims
}

func GenerateToken(claims CustomClaims)(string, error)  {
	if jwtAuthBackend == nil{
		log4go.Fatal("signingKey not readed")
	}
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = claims
	return token.SignedString(jwtAuthBackend.privateKey)
}

func ParseToken(tokenString string)(*CustomClaims, error)  {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtAuthBackend.PublicKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func RefreshToken(tokenString string)(string, error)  {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0,0)
	}
	token , err := jwt.ParseWithClaims(tokenString,&CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtAuthBackend.privateKey,nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(2*24 *time.Hour).Unix()
		return GenerateToken(*claims)
	}
	return "",TokenInvalid
}


func getPrivateKey(path string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	pemfileinfo, _ := privateKeyFile.Stat()
	var size  = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		panic(err)
	}
	return privateKeyImported
}

func getPublicKey(path string) *rsa.PublicKey {
	publicKeyFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	publicKeyFile.Close()
	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		panic(err)
	}
	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		panic(err)
	}
	return rsaPub
}
