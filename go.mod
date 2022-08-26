module github.com/thoohv5/test

go 1.17

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.6
)

require (
	github.com/bytedance/sonic v1.3.4 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20220526154910-8bf9453eb81a // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.10 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.0.14 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/arch v0.0.0-20220412001346-fc48f9fe4c15 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//github.com/golang-jwt/jwt v3.2.2+incompatible => ../../thoohv5/jwt
replace github.com/gin-gonic/gin v1.8.1 => ../../thoohv5/gin
