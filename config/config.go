package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/oliverziegert/dccmd-go/constant"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

const (
	DEBUG              = "debug"
	VERSION            = "version"
	ALIASES            = "aliases"
	DOMAIN             = "Domain"
	CLIENTID           = "ClientId"
	CLIENTSECRET       = "clientSecret"
	RETURNFLOW         = "ReturnFlow"
	BINDADDRESS        = "BindAddress"
	BINDPORT           = "BindPort"
	ACCESSTOKEN        = "accessToken"
	TOKENTYPE          = "tokenType"
	REFRESHTOKEN       = "refreshToken"
	EXPIRYACCESSTOKEN  = "expiryAccesstoken"
	EXPIRYREFRESHTOKEN = "expiryRefreshtoken"
)

type Config struct {
	debug   bool             `mapstructure:"debug"`
	version int              `mapstructure:"version"`
	aliases map[string]Alias `mapstructure:"aliases"`
}

type Alias struct {
	Domain       string     `mapstructure:"domain,omitempty"`
	ClientId     string     `mapstructure:"clientId,omitempty"`
	clientSecret string     `mapstructure:"clientSecret,omitempty"`
	ReturnFlow   ReturnFlow `mapstructure:"returnFlow,omitempty"`
	BindAddress  string     `mapstructure:"bindAddress,omitempty"`
	BindPort     uint16     `mapstructure:"bindPort,omitempty"`
	accessToken  string     `mapstructure:"accessToken,omitempty"`
	tokenType    string     `mapstructure:"tokenType,omitempty"`
	refreshToken string     `mapstructure:"refreshToken,omitempty"`
	expiry       time.Time  `mapstructure:"expiry,omitempty"`
}

type ReturnFlow string

const (
	ReturnFlowBrowser ReturnFlow = "browser"
	ReturnFlowCli     ReturnFlow = "cli"
)

func LoadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read config file!")
	}
	return nil
}

func Set(key string, value interface{}) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func IsSet(key string) bool {
	return viper.IsSet(key)
}

func AllSettings() (*Config, error) {
	var cfg Config
	err := viper.Unmarshal(&cfg)
	return &cfg, err
}

func GetTargets() []string {
	aliases := GetStringMap(ALIASES)
	var targets []string
	for alias := range aliases {
		targets = append(targets, alias)
	}
	return targets
}

func GetDebug() bool {
	return GetBool(DEBUG)
}

func SetDebug(debug bool) error {
	return Set(DEBUG, debug)
}

func GetVersion() int {
	return GetInt(VERSION)
}

func SetVersion(version int) error {
	return Set(VERSION, version)
}

func CreateDefaultConfigFile() error {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// Create config directory
	path := fmt.Sprintf("%s/%s", home, constant.ConfigSubPath)
	err = os.MkdirAll(path, 0700)
	if err != nil {
		return err
	}
	cfgFile := fmt.Sprintf("%s/%s.%s", path, constant.ConfigName, constant.ConfigType)
	return viper.SafeWriteConfigAs(cfgFile)
}

func NewAlias(
	domain string,
	clientId string,
	clientSecret string,
	returnFlow ReturnFlow,
	bindAddress string,
	bindPort uint16) *Alias {

	return &Alias{
		Domain:       domain,
		ClientId:     clientId,
		clientSecret: clientSecret,
		ReturnFlow:   returnFlow,
		BindAddress:  bindAddress,
		BindPort:     bindPort,
	}
}

func AddAlias(target string, alias *Alias) error {
	cp := fmt.Sprintf("%s.%s.", ALIASES, target)
	viper.Set(cp+DOMAIN, alias.Domain)
	viper.Set(cp+CLIENTID, alias.ClientId)
	viper.Set(cp+CLIENTSECRET, alias.clientSecret)
	viper.Set(cp+RETURNFLOW, alias.ReturnFlow)
	viper.Set(cp+BINDADDRESS, alias.BindAddress)
	viper.Set(cp+BINDPORT, alias.BindPort)
	return viper.WriteConfig()
}

func RemoveAlias(target string) error {
	cp := fmt.Sprintf("%s.%s", ALIASES, target)
	return unset(cp)
}

func GetAlias(target string) (*Alias, error) {
	var alias *Alias
	cp := fmt.Sprintf("%s.%s.", ALIASES, target)
	if !IsSet(fmt.Sprintf("%s.%s", ALIASES, target)) {
		return nil, fmt.Errorf("target %s is not valid", target)
	}
	alias = &Alias{
		Domain:       GetString(cp + DOMAIN),
		ClientId:     GetString(cp + CLIENTID),
		clientSecret: GetString(cp + CLIENTSECRET),
		ReturnFlow:   ReturnFlow(GetString(cp + RETURNFLOW)),
		BindAddress:  GetString(cp + BINDADDRESS),
		BindPort:     uint16(GetInt(cp + BINDPORT)),
	}
	return alias, nil
}

func GetAliases() (*map[string]*Alias, error) {
	aliases := make(map[string]*Alias)
	for _, target := range GetTargets() {
		alias, err := GetAlias(target)
		if err != nil {
			return nil, err
		}
		aliases[target] = alias
	}
	return &aliases, nil
}

// https://github.com/spf13/viper/issues/632
func unset(vars ...string) error {
	cfg := viper.AllSettings()
	vals := cfg

	for _, v := range vars {
		parts := strings.Split(v, ".")
		for i, k := range parts {
			v, ok := vals[k]
			if !ok {
				// Doesn't exist no action needed
				break
			}

			switch len(parts) {
			case i + 1:
				// Last part so delete.
				delete(vals, k)
			default:
				m, ok := v.(map[string]interface{})
				if !ok {
					return fmt.Errorf("unsupported type: %T for %q", v, strings.Join(parts[0:i], "."))
				}
				vals = m
			}
		}
	}

	b, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	if err = viper.ReadConfig(bytes.NewReader(b)); err != nil {
		return err
	}

	return viper.WriteConfig()
}
