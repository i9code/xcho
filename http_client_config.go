package xcho

import (
	`fmt`
	`net/http`
	`net/url`
	`time`
)

const (
	// URISchemeSocksV4 Socks协议
	URISchemeSocksV4 uriScheme = "socks4"
	// URISchemeSocksV5 Socks协议
	URISchemeSocksV5 uriScheme = "socks4"
	// URISchemeHttp Socks协议
	URISchemeHttp uriScheme = "http"
	// URISchemeHttps Socks协议
	URISchemeHttps uriScheme = "https"
)

const (
	// AuthTypeBasic 基本Http授权验证
	AuthTypeBasic authType = "basic"
	// AuthTypeToken 基本传Token的授权验证
	AuthTypeToken authType = "token"
)

type (
	HttpClientConfig struct {
		// 超时
		Timeout time.Duration `json:"timeout" yaml:"timeout"`
		// 代理
		Proxy proxyConfig `json:"proxy" yaml:"proxy" validate:"structonly"`
		// 授权配置
		Auth authConfig `json:"auth" yaml:"auth" validate:"structonly"`
		// Body数据传输控制
		Payload payload `json:"payload" yaml:"payload" validate:"structonly"`
		// 秘钥配置
		Certificate certificateConfig `json:"certificate" yaml:"certificate" validate:"structonly"`
		// 通用的查询参数
		Queries map[string]string `json:"queries" yaml:"queries"`
		// 表单参数，只对POST和PUT方法有效
		Forms map[string]string `json:"forms" yaml:"forms"`
		// 通用头信息
		Headers map[string]string `json:"headers" yaml:"headers"`
		// 通用Cookie
		Cookies []*http.Cookie `json:"cookies" yaml:"cookies"`
	}

	// uriScheme 协议
	uriScheme string
	// authType 授权类型
	authType string

	// authConfig 授权信息
	authConfig struct {
		// Type 授权类型
		Type authType `default:"type" json:"type" yaml:"type" validate:"oneof=basic token"`
		// Username 用户名
		Username string `json:"username" yaml:"username"`
		// Password 密码
		Password string `json:"password" yaml:"password"`
		// Token 授权码
		Token string `json:"token" yaml:"token"`
		// Scheme 身份验证方案类型
		Scheme string `json:"scheme" yaml:"scheme"`
	}

	// proxyConfig 代理配置
	proxyConfig struct {
		// Host 主机（可以是Ip或者域名）
		Host string `json:"ip" yaml:"ip" validate:"required"`
		// Port 端口
		Port int `default:"80" json:"port" yaml:"port" validate:"required"`
		// Scheme 代理类型
		Scheme uriScheme `default:"scheme" json:"scheme" yaml:"type" validate:"required,oneof=socks4 socks5 http https"`
		// Username 代理认证用户名
		Username string `json:"username" yaml:"username"`
		// Password 代理认证密码
		Password string `json:"password" yaml:"password"`
	}

	payload struct {
		// 是否允许Get方法使用Bogy传输数据
		Get bool `default:"true" json:"get" yaml:"get"`
	}

	// clientCertificate 客户端秘钥
	clientCertificate struct {
		// Public 公钥文件路径
		Public string `json:"public" yaml:"public" validate:"required,file"`
		// Private 私钥文件路径
		Private string `json:"private" yaml:"private" validate:"required,file"`
	}

	// certificateConfig 秘钥
	certificateConfig struct {
		// Skip 是否跳过TLS检查
		Skip bool `default:"true" json:"skip" yaml:"skip"`
		// Root 根秘钥文件路径
		Root string `json:"root" yaml:"root" validate:"required,file"`
		// Clients 客户端
		Clients []clientCertificate `json:"clients" yaml:"clients" validate:"structonly"`
	}
)

func (p *proxyConfig) Addr() (addr string) {
	if "" != p.Username && "" != p.Password {
		addr = fmt.Sprintf(
			"%s://%s:%s@%s:%d",
			p.Scheme,
			url.QueryEscape(p.Username), url.QueryEscape(p.Password),
			p.Host, p.Port,
		)
	} else {
		addr = fmt.Sprintf("%s://%s:%d", p.Scheme, p.Host, p.Port)
	}

	return
}
