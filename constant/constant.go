package constant

import (
	"fmt"
	"time"
)

const (
	Version   = "0.0.1"
	ShortName = "dccmd"
	LongName  = "DRACOON Commander"

	ConfigSubPath = "." + ShortName
	ConfigType    = "yaml"
	ConfigName    = "config"

	EnvPrefix = "DCCMD"

	OauthExpiryRefreshtokenOffset = 300 * time.Second
	OauthExpiryAccesstokenOffset  = 12 * time.Hour
	OauthExpiryAccesstoken        = 30 * 24 * time.Hour

	oauthAuthUrlFormat  = "https://%s/oauth/authorize"
	oauthTokenUrlFormat = "https://%s/oauth/token"
)

func OauthAuthUrl(domain string) string {
	return fmt.Sprintf(oauthAuthUrlFormat, domain)
}

func OauthTokenUrl(domain string) string {
	return fmt.Sprintf(oauthTokenUrlFormat, domain)
}

func OauthScops() []string {
	return []string{"all"}
}
