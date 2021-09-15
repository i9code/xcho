module github.com/i9code/xcho

go 1.16

require (
	github.com/casbin/casbin/v2 v2.30.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-fed/httpsig v1.1.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-resty/resty/v2 v2.6.0
	github.com/i9code/xlog v0.0.0-20210909033202-bd0588e1680c
	github.com/i9code/xutil v0.0.0-20210909033641-ed0c48eb44d9
	github.com/labstack/echo/v4 v4.3.0
	github.com/mcuadros/go-defaults v1.2.0
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/rs/xid v1.3.0
	github.com/storezhang/validatorx v1.0.8
	github.com/valyala/fasttemplate v1.2.1
	github.com/vmihailenco/msgpack/v5 v5.3.4
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/protobuf v1.27.1
)

replace github.com/i9code/xutil => ../xutil
