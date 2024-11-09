package mail

import (
	"github.com/emersion/go-smtp"
	"io"
	"log"
)

type Session struct {
	host string
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
