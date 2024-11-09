package smtp

import (
	"fmt"
	"github.com/changhoi/slake/internal/adapter/config"
	"github.com/emersion/go-smtp"
	"time"
)

type Adapter struct {
	addr              string
	domain            string
	readTimeout       time.Duration
	writeTimeout      time.Duration
	maxMessageBytes   int64
	maxRecipients     int
	allowInsecureAuth bool
}

func (a *Adapter) NewServer(backend smtp.Backend) *smtp.Server {
	s := smtp.NewServer(backend)
	s.Addr = a.addr
	s.Domain = a.domain
	s.AllowInsecureAuth = a.allowInsecureAuth
	return s
}

func NewAdapter(cfg *config.Config) *Adapter {
	a := &Adapter{}
	a.addr = fmt.Sprintf(":%d", cfg.SMTP.Port) // 서버가 수신 대기할 포트 번호
	a.domain = cfg.SMTP.Domain                 // 서버의 도메인
	a.allowInsecureAuth = true                 // TLS 없이도 인증을 허용
	return a
}
