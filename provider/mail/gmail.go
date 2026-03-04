package mail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"os"
	"sync"
	"time"
)

type smtpConfig struct {
	host string
	port string
	user string
	pass string
	from string
}

var (
	cfg  smtpConfig
	once sync.Once
)

func loadConfig() {
	cfg = smtpConfig{
		host: mustEnv("SMTP_HOST"),
		port: mustEnv("SMTP_PORT"),
		user: mustEnv("SMTP_USER"),
		pass: mustEnv("SMTP_PASS"),
		from: getEnv("SMTP_FROM", os.Getenv("SMTP_USER")),
	}
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("missing env: " + key)
	}
	return v
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Send(to, subject, html string) error {
	once.Do(loadConfig)

	addr := fmt.Sprintf("%s:%s", cfg.host, cfg.port)

	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, cfg.host)
	if err != nil {
		return err
	}
	defer client.Quit()

	tlsConfig := &tls.Config{
		ServerName: cfg.host,
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", cfg.user, cfg.pass, cfg.host)
	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(cfg.user); err != nil {
		return err
	}

	if err = client.Rcpt(to); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n"+
			"%s",
		cfg.from, to, subject, html,
	)

	if _, err = w.Write([]byte(msg)); err != nil {
		return err
	}

	return w.Close()
}
