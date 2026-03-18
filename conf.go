package goutil

import (
	"github.com/spf13/viper"
	"log"
)

var cfg = viper.New()

func ConfigInit(filepath string) {
	if err := ConfigInitE(filepath); err != nil {
		log.Fatal(err)
	}
}

func ConfigInit2(filepath string, in string) {
	if err := ConfigInit2E(filepath, in); err != nil {
		log.Fatal(err)
	}
}

func ConfigInitE(filepath string) error {
	return ConfigInit2E(filepath, "ini")
}

func ConfigInit2E(filepath string, in string) error {
	cfg.SetConfigFile(filepath)
	cfg.SetConfigType(in)
	return cfg.ReadInConfig()
}

func GetSystemValue(key string) string {
	return cfg.GetString("system." + key)
}

func GetSettingValue(key string) string {
	return cfg.GetString("setting." + key)
}

func GetSectionValue(section string, key string) string {
	return cfg.GetString(section + "." + key)
}

func GetSectionValueInt(section string, key string) int {
	return cfg.GetInt(section + "." + key)
}

func GetSection(section string) map[string]string {
	return cfg.GetStringMapString(section)
}

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return cfg.UnmarshalKey(key, rawVal, opts...)
}

func GetCfg() *viper.Viper {
	return cfg
}
