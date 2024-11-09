package main

import (
	"io"
	"log"

	"github.com/emersion/go-smtp"
)

type Backend struct{}

func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	log.Println("새로운 세션을 생성합니다.")
	return &Session{}, nil
}

type Session struct {
	from string
	to   string
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Printf("발신자: %s\n", from)
	log.Printf("메일 옵션: %+v\n", opts)
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Printf("수신자: %s\n", to)
	log.Printf("수신자 옵션: %+v\n", opts)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	log.Printf("메일 본문: %s\n", data)
	return nil
}

func (s *Session) Reset() {
	log.Println("세션을 초기화합니다.")
}

func (s *Session) Logout() error {
	log.Println("사용자가 로그아웃했습니다.")
	return nil
}

func main() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":2025"
	s.Domain = "mail.changhoi.kim"
	s.AllowInsecureAuth = true

	log.Println("SMTP 서버가 시작되었습니다. 포트 2025에서 대기 중...")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
