package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func initDefaults(v *viper.Viper) {
	initDefaultsRecursive(v, reflect.TypeOf(cfg), "")
}

func initDefaultsRecursive(v *viper.Viper, t reflect.Type, prefix string) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type

		if f.Anonymous {
			initDefaultsRecursive(v, ft, prefix)
		} else {
			name := strings.ToLower(f.Name)
			if value, ok := f.Tag.Lookup("mapstructure"); ok {
				name = value
			}
			if ft.Kind() == reflect.Struct {
				initDefaultsRecursive(v, ft, fmt.Sprintf("%s%s.", prefix, name))
			} else {
				if value, ok := f.Tag.Lookup("default"); ok {
					key := prefix + name
					v.SetDefault(key, value)
				}
			}
		}
	}
}

func init() {
	v := viper.New()

	initDefaults(v)

	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

}
