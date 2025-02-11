package modelexts

import (
	"strconv"

	"pluto/models"
)

type User struct {
	User        *models.User
	Bindings    []*models.Binding
	Role        string `json:"role"`
	AppID       string `json:"app_id"`
	PasswordSet bool
}

type UserInfo struct {
	Sub           uint       `json:"sub"`
	UserID        string     `json:"user_id"`
	UserUpdated   bool       `json:"user_updated"`
	Name          string     `json:"name"`
	Role          string     `json:"role"`
	AppID         string     `json:"app_id"`
	Avatar        string     `json:"avatar"`
	Verified      bool       `json:"verified"`
	CreatedAt     int64      `json:"created_at"`
	UpdatedAt     int64      `json:"update_at"`
	IsPasswordSet bool       `json:"is_password_set"`
	Bindings      []*Binding `json:"bindings"`
}

type Binding struct {
	LoginType string `json:"login_type"`
	Mail      string `json:"mail"`
}

func (u User) UserInfo() UserInfo {
	userInfo := UserInfo{
		Sub:           u.User.ID,
		UserID:        u.User.UserID,
		UserUpdated:   u.User.UserIDUpdated,
		Name:          u.User.Name,
		Avatar:        u.User.Avatar.String,
		Role:          u.Role,
		AppID:         u.AppID,
		Verified:      u.User.Verified.Bool,
		CreatedAt:     u.User.CreatedAt.Time.Unix(),
		UpdatedAt:     u.User.UpdatedAt.Time.Unix(),
		IsPasswordSet: u.PasswordSet,
	}

	bindings := make([]*Binding, 0)

	for _, binding := range u.Bindings {
		b := &Binding{}
		b.LoginType = binding.LoginType
		b.Mail = binding.Mail
		bindings = append(bindings, b)
	}

	userInfo.Bindings = bindings

	return userInfo
}

func (u User) Format() map[string]interface{} {
	res := make(map[string]interface{})
	res["sub"] = u.User.ID
	res["name"] = u.User.Name
	res["app_id"] = u.AppID
	res["avatar"] = u.User.Avatar
	res["role"] = u.Role
	res["verified"] = u.User.Verified
	res["created_at"] = u.User.CreatedAt.Time.Unix()
	res["updated_at"] = u.User.UpdatedAt.Time.Unix()
	res["is_password_set"] = strconv.FormatBool(u.PasswordSet)
	res["user_id"] = u.User.UserID
	res["user_id_updated"] = u.User.UserIDUpdated
	bindings := make([]interface{}, 0)
	for _, binding := range u.Bindings {
		b := make(map[string]interface{})
		b["login_type"] = binding.LoginType
		b["mail"] = binding.Mail
		bindings = append(bindings, b)
	}
	res["bindings"] = bindings
	return res
}

func (u User) PublicInfo() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = u.User.ID
	res["name"] = u.User.Name
	res["avatar"] = u.User.Avatar
	return res
}

type Scopes struct {
	Application *models.Application `json:"application"`
	Scopes      []*models.RbacScope `json:"scopes"`
}

func (scopes Scopes) Format() map[string]interface{} {
	s := make([]map[string]interface{}, 0)

	app := make(map[string]interface{})
	app["id"] = scopes.Application.ID
	app["name"] = scopes.Application.Name
	app["default_role"] = scopes.Application.DefaultRole

	for _, scope := range scopes.Scopes {
		m := make(map[string]interface{})
		m["id"] = scope.ID
		m["name"] = scope.Name
		s = append(s, m)
	}

	res := make(map[string]interface{})
	res["application"] = app
	res["scopes"] = s
	return res
}

type Roles struct {
	Application *models.Application `json:"application"`
	Roles       []Role              `json:"roles"`
}

type Role struct {
	*models.RbacRole
	Scopes []*models.RbacScope `json:"scopes"`
}

func (roles Roles) Format() map[string]interface{} {
	res := make(map[string]interface{})
	app := make(map[string]interface{})
	app["name"] = roles.Application.Name
	app["id"] = roles.Application.ID
	app["default_role"] = roles.Application.DefaultRole

	res["application"] = app

	rs := make([]interface{}, 0)
	for _, role := range roles.Roles {
		r := make(map[string]interface{})
		r["id"] = role.ID
		r["name"] = role.Name
		r["default_scope"] = role.DefaultScope

		scopes := make([]interface{}, 0)
		for _, scope := range role.Scopes {
			s := make(map[string]interface{})
			s["id"] = scope.ID
			s["name"] = scope.Name
			scopes = append(scopes, s)
		}
		r["scopes"] = scopes
		rs = append(rs, r)
	}

	res["roles"] = rs

	return res
}

type UserApplicationRole struct {
	*models.Application
	Roles []*models.RbacRole `json:"roles"`
}

type FindUser struct {
	User         *models.User
	Bindings     []*models.Binding
	Applications []*UserApplicationRole `json:"applications"`
}

func (f FindUser) Format() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = f.User.ID
	res["name"] = f.User.Name
	res["avatar"] = f.User.Avatar.String

	applications := make([]interface{}, 0)
	for _, application := range f.Applications {
		a := make(map[string]interface{})
		a["id"] = application.ID
		a["name"] = application.Name
		rs := make([]interface{}, 0)
		for _, role := range application.Roles {
			r := make(map[string]interface{})
			r["id"] = role.ID
			r["name"] = role.Name
			rs = append(rs, r)
		}
		a["roles"] = rs
		applications = append(applications, a)
	}

	res["applications"] = applications

	bindings := make([]interface{}, 0)
	for _, binding := range f.Bindings {
		b := make(map[string]interface{})
		b["login_type"] = binding.LoginType
		b["mail"] = binding.Mail
		bindings = append(bindings, b)
	}
	res["bindings"] = bindings

	return res
}

type OauthClient struct {
	Client       *models.OauthClient
	OriginSecret string
}

func (oc *OauthClient) Format() map[string]interface{} {
	res := make(map[string]interface{})
	res["key"] = oc.Client.Key
	res["status"] = oc.Client.Status
	res["redirect_uri"] = oc.Client.RedirectURI
	res["origin_secret"] = oc.OriginSecret
	return res
}

type ApplicationI18nName struct {
	Language string `json:"tag"`
	Name     string `json:"i18n_name"`
}

type ApplicationI18nNameInfo struct {
	AppName   string
	AppId     uint
	I18nNames *[]ApplicationI18nName
}

func (aii *ApplicationI18nNameInfo) Format() map[string]interface{} {
	res := make(map[string]interface{})
	res["app_id"] = aii.AppId
	res["app_name"] = aii.AppName
	res["i18n_names"] = aii.I18nNames
	return res
}

type GoogleLogin struct {
	Aud string `json:"aud"`
}

type WechatLogin struct {
	// 移动端应用
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`

	// 网站应用
	WebID     string `json:"web_id"`
	WebSecret string `json:"web_secret"`
}

type AppleLogin struct {
	TeamID        string `json:"team_id"`
	BundleID      string `json:"bundle_id"`
	ClientID      string `json:"client_id"`
	KeyID         string `json:"key_id"`
	P8CertContent string `json:"p8_cert_content"`
	RedirectURL   string `json:"redirect_url"`
}
