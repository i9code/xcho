package xhttp

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
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

type (
	HttpClientConfig struct {
		// 超时
		Timeout time.Duration `json:"timeout" yaml:"timeout"`
		// 代理
		Proxy ProxyConfig `json:"proxy" yaml:"proxy" xvalidate:"structonly"`
		// 授权配置
		Auth AuthConfig `json:"auth" yaml:"auth" xvalidate:"structonly"`
		// Body数据传输控制
		Payload Payload `json:"Payload" yaml:"Payload" xvalidate:"structonly"`
		// 秘钥配置
		Certificate CertificateConfig `json:"certificate" yaml:"certificate" xvalidate:"structonly"`
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

	// AuthConfig 授权信息
	AuthConfig struct {
		// Type 授权类型
		Type AuthType `default:"type" json:"type" yaml:"type" xvalidate:"oneof=basic token"`
		// Username 用户名
		Username string `json:"username" yaml:"username"`
		// Password 密码
		Password string `json:"password" yaml:"password"`
		// Token 授权码
		Token string `json:"token" yaml:"token"`
		// Scheme 身份验证方案类型
		Scheme string `json:"scheme" yaml:"scheme"`
	}

	// ProxyConfig 代理配置
	ProxyConfig struct {
		// Host 主机（可以是Ip或者域名）
		Host string `json:"ip" yaml:"ip" xvalidate:"required"`
		// Port 端口
		Port int `default:"80" json:"port" yaml:"port" xvalidate:"required"`
		// Scheme 代理类型
		Scheme uriScheme `default:"scheme" json:"scheme" yaml:"type" xvalidate:"required,oneof=socks4 socks5 http https"`
		// Username 代理认证用户名
		Username string `json:"username" yaml:"username"`
		// Password 代理认证密码
		Password string `json:"password" yaml:"password"`
	}

	Payload struct {
		// 是否允许Get方法使用Bogy传输数据
		Get bool `default:"true" json:"get" yaml:"get"`
	}

	// clientCertificate 客户端秘钥
	clientCertificate struct {
		// Public 公钥文件路径
		Public string `json:"public" yaml:"public" xvalidate:"required,file"`
		// Private 私钥文件路径
		Private string `json:"private" yaml:"private" xvalidate:"required,file"`
	}

	// CertificateConfig 秘钥
	CertificateConfig struct {
		// Skip 是否跳过TLS检查
		Skip bool `default:"true" json:"skip" yaml:"skip"`
		// Root 根秘钥文件路径
		Root string `json:"root" yaml:"root" xvalidate:"required,file"`
		// Clients 客户端
		Clients []clientCertificate `json:"clients" yaml:"clients" xvalidate:"structonly"`
	}
)

func (p *ProxyConfig) Addr() (addr string) {
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
