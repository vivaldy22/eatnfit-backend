// +heroku install ./cmd/...
// +heroku goVersion go1.14.4
module github.com/vivaldy22/eatnfit-auth-service-grpc

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	google.golang.org/grpc v1.31.1
	google.golang.org/protobuf v1.24.0
)
