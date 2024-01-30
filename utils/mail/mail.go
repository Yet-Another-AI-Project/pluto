package mail

import (
	"bytes"
	"errors"
	"log"

	"pluto/utils/view"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"pluto/config"
	"pluto/utils/jwt"

	"github.com/resend/resend-go/v2"

	perror "pluto/datatype/pluto_error"
)

type Mail struct {
	config *config.Config
	bundle *i18n.Bundle
	resend *resend.Client
}

func (m *Mail) Send(recv, subj, contentType, body string) error {

	params := &resend.SendEmailRequest{
		From:    "noreply@mail.kiwiworlds.com",
		To:      []string{recv},
		Subject: subj,
	}

	if contentType == "text/plain" {
		params.Text = body
	} else if contentType == "text/html" {
		params.Html = body
	}

	sent, err := m.resend.Emails.Send(params)

	if err != nil {
		return err
	}

	log.Printf("Mail sent with message id: %s\n", sent.Id)

	return nil
}

func (m *Mail) SendPlainText(address, subject, text string) *perror.PlutoError {
	if err := m.Send(address, subject, "text/plain", text); err != nil {
		return perror.ServerError.Wrapper(errors.New("Mail sending failed: " + err.Error()))
	}

	return nil
}

func (m *Mail) SendRegisterVerify(userID uint, address string, baseURL string, language string, appName string) *perror.PlutoError {
	rvp := jwt.NewRegisterVerifyPayload(userID, m.config.Token.RegisterVerifyTokenExpire)
	token, perr := jwt.GenerateRSA256JWT(rvp)
	if perr != nil {
		return perr.Wrapper(errors.New("JWT token generate failed"))
	}

	vw, err := view.GetView()
	if err != nil {
		return perror.ServerError.Wrapper(err)
	}

	t, err := vw.Parse(language, "register_verify_mail.html")
	if err != nil {
		return perror.ServerError.Wrapper(err)
	}
	var buffer bytes.Buffer
	type Data struct {
		AppName string
		BaseURL string
		Token   string
	}
	t.Execute(&buffer, Data{AppName: appName, Token: token.B64String(), BaseURL: baseURL})
	localizer := i18n.NewLocalizer(m.bundle, language)
	subject, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "VerifyMailSubject",
		TemplateData: map[string]interface{}{
			"AppName": appName,
		},
	})
	if err != nil {
		subject = "[Kiwi] Mail Confirmation"
	}
	if err := m.Send(address, subject, "text/html", buffer.String()); err != nil {
		return perror.ServerError.Wrapper(errors.New("Mail sending failed: " + err.Error()))
	}

	return nil
}

// appName 是字符串名称，用于邮件模板
func (m *Mail) SendResetPassword(appID, address string, baseURL string, userLanguage string, appName string) *perror.PlutoError {
	prp := jwt.NewPasswordResetPayload(appID, address, m.config.Token.ResetPasswordTokenExpire)
	token, perr := jwt.GenerateRSA256JWT(prp)
	if perr != nil {
		return perr.Wrapper(errors.New("JWT token generate failed"))
	}

	vw, err := view.GetView()
	if err != nil {
		return perror.ServerError.Wrapper(err)
	}

	t, err := vw.Parse(userLanguage, "password_reset_mail.html")
	if err != nil {
		return perror.ServerError.Wrapper(err)
	}

	var buffer bytes.Buffer
	type Data struct {
		AppName string
		BaseURL string
		Token   string
	}
	t.Execute(&buffer, Data{AppName: appName, Token: token.B64String(), BaseURL: baseURL})
	localizer := i18n.NewLocalizer(m.bundle, userLanguage)
	subject, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "ResetPasswordMailSubject",
		TemplateData: map[string]interface{}{
			"AppName": appName,
		},
	})
	if err != nil {
		subject = "[Kiwi] Password Reset"
	}
	if err := m.Send(address, subject, "text/html", buffer.String()); err != nil {
		return perror.ServerError.Wrapper(errors.New("Mail sending failed: " + err.Error()))
	}
	return nil
}

func NewMail(config *config.Config, bundle *i18n.Bundle) (*Mail, *perror.PlutoError) {
	c := config.Mail
	if c.Key == "" {
		return nil, perror.ServerError.Wrapper(errors.New("mail key is not set"))
	}

	resend := resend.NewClient(c.Key)

	mail := &Mail{
		config: config,
		bundle: bundle,
		resend: resend,
	}
	return mail, nil
}
