package kernel

import (
	"errors"
	"strconv"
	"strings"
)

type Profile struct {
	env      string // 当前环境，小写，缺省：pro
	debug    bool   // 测试？缺省：false
	name     string // 项目名，小写
	port     int    // 运行端口，缺省：8080
	timezone string // 时区，缺省：Asia/Shanghai
	i18n     string // 语言包，缺省：zh-CN
}

func NewProfile(data map[string]string) (*Profile, error) {
	if data == nil {
		return nil, errors.New("data can't be nil")
	}

	env := Env
	debug := Debug
	name := ""
	port := Port
	timezone := Timezone
	i18n := I18n

	if data[ProfileEnv] != "" {
		if s := strings.TrimSpace(data[ProfileEnv]); s != "" {
			env = strings.ToLower(s)
		}
	}

	if data[ProfileDebug] != "" {
		if b, err := strconv.ParseBool(data[ProfileDebug]); err == nil {
			debug = b
		} else {
			return nil, err
		}
	}

	if data[ProfileName] != "" {
		if s := strings.TrimSpace(data[ProfileName]); s != "" {
			name = strings.ToLower(s)
		}
	}

	if data[ProfilePort] != "" {
		if n, err := strconv.ParseInt(data[ProfilePort], 10, 32); err == nil {
			port = int(n)
		} else {
			return nil, err
		}
	}

	if data[ProfileTimezone] != "" {
		if s := strings.TrimSpace(data[ProfileTimezone]); s != "" {
			timezone = s
		}
	}

	if data[ProfileI18n] != "" {
		if s := strings.TrimSpace(data[ProfileI18n]); s != "" {
			i18n = s
		}
	}

	if name == "" {
		return nil, errors.New("name can't be empty")
	}

	return &Profile{
		env:      env,
		debug:    debug,
		name:     name,
		port:     port,
		timezone: timezone,
		i18n:     i18n,
	}, nil
}

func (x *Profile) GetEnv() string {
	return x.env
}

func (x *Profile) GetDebug() bool {
	return x.debug
}

func (x *Profile) GetName() string {
	return x.name
}

func (x *Profile) GetPort() int {
	return x.port
}

func (x *Profile) GetTimezone() string {
	return x.timezone
}

func (x *Profile) GetI18n() string {
	return x.i18n
}
