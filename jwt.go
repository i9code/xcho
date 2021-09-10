package xcho

import (
	`encoding/json`
	`reflect`
	`time`

	`github.com/dgrijalva/jwt-go`
	`github.com/labstack/echo/v4`
	`github.com/labstack/echo/v4/middleware`
	`github.com/rs/xid`
)

// Jwt Jwt配置
type Jwt struct {
	// 确定是不是要走中间件
	skipper middleware.Skipper
	// 执行前的操作
	beforeHandler middleware.BeforeFunc
	// 成功后操作
	successHandler jwtSuccessHandler
	// 签名密钥
	// 必须字段
	key interface{}
	// 签名方法
	// 非必须，默认是HS256
	method string
	// 存储用户信息的键
	// 非必须，默认值是"user"
	context string
	// 存储数据的类型
	// 非必须，默认值是jwt.StandardClaims
	claims jwt.Claims
	// 定义从哪获得Token
	// 非必须，默认值是"header:Authorization和query:token"
	// 可能的值：
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	lookups []string
	// Token分隔字符串
	// 非必须，默认值是"Bearer"
	scheme string

	keyFunc   jwt.Keyfunc
	extractor []jwtExtractor
}

// NewJwt 创建Jwt配置，快捷方式
func NewJwt(signingKey string) *Jwt {
	return NewJwtWithConfig(
		middleware.DefaultSkipper,
		signingKey, AlgorithmHS256,
		"user", &jwt.StandardClaims{},
		"Bearer",
		nil, nil,
		"header:"+echo.HeaderAuthorization, "query:token",
	)
}

// NewJwtWithConfig 创建Jwt配置
func NewJwtWithConfig(
	skipper middleware.Skipper,
	key interface{}, method string,
	context string,
	claims jwt.Claims,
	scheme string,
	beforeHandler middleware.BeforeFunc, successHandler jwtSuccessHandler,
	lookups ...string,
) *Jwt {
	return &Jwt{
		skipper:        skipper,
		beforeHandler:  beforeHandler,
		successHandler: successHandler,
		key:            key,
		method:         method,
		context:        context,
		claims:         claims,
		lookups:        lookups,
		scheme:         scheme,
	}
}

func (j *Jwt) Subject(ctx *Context, subject interface{}) (err error) {
	var (
		token  string
		claims jwt.Claims
	)

	if token, err = j.runExtractor(ctx); nil != err {
		return
	}
	if claims, _, err = j.Parse(token); nil != err {
		return
	}
	// 从Token中反序列化主题数据
	err = json.Unmarshal([]byte(claims.(*jwt.StandardClaims).Subject), &subject)

	return
}

func (j *Jwt) runExtractor(ctx echo.Context) (token string, err error) {
	for _, extractor := range j.extractor {
		if token, err = extractor(ctx); nil == err || "" != token {
			break
		}
	}

	return
}

func (j *Jwt) Parse(token string) (claims jwt.Claims, header map[string]interface{}, err error) {
	jwtToken := new(jwt.Token)
	if _, ok := j.claims.(jwt.MapClaims); ok {
		jwtToken, err = jwt.Parse(token, j.keyFunc)
	} else {
		elem := reflect.ValueOf(j.claims).Type().Elem()
		claims := reflect.New(elem).Interface().(jwt.Claims)
		jwtToken, err = jwt.ParseWithClaims(token, claims, j.keyFunc)
	}
	if nil == err && jwtToken.Valid {
		claims = jwtToken.Claims
		header = jwtToken.Header
	}

	return
}

func (j *Jwt) MakeToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(j.method), claims)

	return token.SignedString([]byte(j.key.(string)))
}

func (j *Jwt) Token(
	domain string,
	subject interface{},
	expire time.Duration,
) (token string, id string, err error) {
	// 序列化User对象为JSON
	var subjectBytes []byte
	if subjectBytes, err = json.Marshal(subject); nil != err {
		return
	}

	id = xid.New().String()
	token, err = j.MakeToken(jwt.StandardClaims{
		// 代表这个JWT的签发主体
		Issuer: domain,
		// 代表这个JWT的主体，即它的所有人
		Subject: string(subjectBytes),
		// 代表这个JWT的接收对象
		Audience: domain,
		// 是一个时间戳，代表这个JWT的签发时间
		IssuedAt: time.Now().Unix(),
		// 是一个时间戳，代表这个JWT生效的开始时间，意味着在这个时间之前验证JWT是会失败的
		NotBefore: time.Now().Unix(),
		// 是一个时间戳，代表这个JWT的过期时间
		ExpiresAt: time.Now().Add(expire).Unix(),
		// 是JWT的唯一标识
		Id: id,
	})

	return
}
