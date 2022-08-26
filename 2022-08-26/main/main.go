package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = `MIIEpAIBAAKCAQEAtqRsXsVbl5LM2kVp3Pgfgsyukuom2LdksfjgCHLI3GN02TDq
HfT7mx4OGolUakB3ZT2LcRyP0OdFZXzD1HW9TlEISZLi8hzwljip0hupcWGep/fG
XT+uB4YrRINcX5cu3W7z2zXc5DUSwnBaJNuG6v8EpGzh2wnau5FP5F3Bjkl31NUk
NEmjNdBQArLowjWVEnX3kGzwKmBlZxTleFuqqLcYNJdzNGEsCujvU07zVspqmffQ
NqDjy7cw/NVW61UXNZDK4TMz5VxYanmXOvU78ysn2Tj4PD/Zd4AcS0nlc9Y0hKzR
MltTeOx/FmtcfKTLDRWCjTrgc9CF/3HYs+N3EwIDAQABAoIBAQCPOZdmJjwyO8N7
M03WI5mKV/paaVZ3QjGrwu/kkCbldPTP0ST2wwN/2+zRcHoLLGy0rrOnyu8YshPg
hxuGg6IlRTRLhe2EUZ91HhBGHL0elZ/2Nj6PJ33Qlq1dd+m7aHIu4XBHqhCDwP7q
DXMTTjeaFsvMKTpgK3Uk+8n72m4LpktSvS27CH2dXliyFwwGItMFjJHijpCLkSdF
DXQ6/4DRy9uwQX93S4JR713EAu1Oh2KAD6XPdNBqoo2PQga1no8hshgrJWXrwNKE
Wkl8SI1yiOfkLfE/LPz5vOTPpi3VDuMvvS4dVdYX+9OgSbEbPGtZk72DRIoAabCL
PmXCVaXxAoGBAOiKdUAC4JsHbmVniMzfwHQZYH8mmPNoCwUeMbLhoN1e7h21d9JR
9BUnzdX+mjCmUEKmMkaym8tEPAIaAThgHfkKzMEsaGGA6TP6a+rXk8alS6AI/7WA
fpK85RjaieimIHe/byI6tG9x7hE5aV9OPCdoiRGZXlX338KmusMhOUSLAoGBAMkR
TBazmTJ485tfYtgwWYRgXJH3aDpgTfW6+AXPcTcXx7nOC4zvirRmskaMbXZxyj54
WeN7JnSo4WRTV285+GAm7JNmpiUPXJ7x7giRYIsDDa3vKvjvnkmUXYXBqPY/QNjG
2mDncXmVnyKkK6FLA2clVStLjPl7IUziKEmbrICZAoGAe9OO0AE4PRVd8d2J+R6E
ys3glpSlCagzhgwoBssi7/5m6acCIRrG0KUbdIJY9OL7BiKdzwu47iptkejrEWwN
Sdo4Yf0VsCYHCEinQqx84mCOvq9MCwhbXiP8Epn8qcgcredgdGeyQU33qXBa7gco
/QsrQhXbKAgQPQSvmsyMRDUCgYAEFJfItihbv9yhwJPO81w9tX7rb2vsE1xBqmOy
Kn2PsqnY+Xd+irXz50mi6OKnzrNeBS890Jf9Mhgw4wgZN8H0oZWXgDPK+L7Wcu5z
ug+NgqhaaUoj9yjtMVeciUuWg74bKB3ybX/+Ca1LFK3V/iG5jCZoVIYt4fPRDZ3n
sIMVsQKBgQCHGB7CKeGP7YEN8oxuuZN/pJbtN0f2tBG6EeAKspaMCuHiiJiP1hDx
0P0sd63atc2A36SgNGCjls6vc6WgXVUl+McVfKJHvmBqwawfmKHLsCb3c2UgZEdY
egwhRkBp06mLgc8JyOlOZ9pV6aIoOWCYW+7yv7im0tfTw0gjBaQIyg==`

func main() {
	token := encryption()
	fmt.Println(token)
	decryption(token)
}

func encryption() string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)
	return tokenString
}

func decryption(tokenString string) {
	// sample token string taken from the New example
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

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
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
