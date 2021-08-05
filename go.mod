module github.com/i9code/xcho

go 1.16

require (
	github.com/casbin/casbin/v2 v2.34.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-fed/httpsig v1.1.0
	github.com/go-playground/validator/v10 v10.8.0
	github.com/i9code/xutils v0.0.0-20210803105152-8c01de13cf70
	github.com/json-iterator/go v1.1.11
	github.com/labstack/echo/v4 v4.5.0
	github.com/mcuadros/go-defaults v1.2.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/rs/xid v1.3.0
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
)

replace (
	github.com/i9code/xutils  => ../xutils
)
