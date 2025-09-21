package uri

import (
	"fmt"
	"net/url"
	"strings"
)

type uriBuilder struct {
	base   string
	params map[string]string
	build  string
}

func NewBuilder() *uriBuilder {
	return &uriBuilder{
		params: make(map[string]string),
	}
}

func (u *uriBuilder) BaseUrl(base string) {
	u.base = base
}

func (u *uriBuilder) Param(key string, values ...any) {
	var paramValues []string
	for _, value := range values {
		str := fmt.Sprintf("%v", value)
		if str != "" {
			paramValues = append(paramValues, str)
		}
	}

	if len(paramValues) > 0 {
		u.params[key] = strings.Join(paramValues, ",")
	}
}

func (u *uriBuilder) Build() string {
	if u.build != "" {
		return u.build
	}

	if len(u.params) == 0 {
		u.build = u.base
		return u.build
	}

	values := url.Values{}
	for key, value := range u.params {
		values.Add(key, value)
	}

	u.build = u.base + "?" + values.Encode()
	return u.build
}

func (u *uriBuilder) Reset() {
	u.base, u.build = "", ""
	u.params = make(map[string]string)
}
