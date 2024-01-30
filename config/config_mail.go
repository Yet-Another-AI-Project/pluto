package config

type MailConfig struct {
	Key string `kiper_value:"name:key;help:mailer send api key"`
}

type SMTP struct {
	s string
}

func (smtp *SMTP) Set(s string) error {
	smtp.s = s
	return nil
}

func (smtp *SMTP) String() string {
	return smtp.s
}

func newMailConfig() *MailConfig {
	return &MailConfig{}
}
