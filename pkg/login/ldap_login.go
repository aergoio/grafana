package login

import (
	m "github.com/aergoio/grafana/pkg/models"
	"github.com/aergoio/grafana/pkg/setting"
)

var loginUsingLdap = func(query *m.LoginUserQuery) (bool, error) {
	if !setting.LdapEnabled {
		return false, nil
	}

	for _, server := range LdapCfg.Servers {
		author := NewLdapAuthenticator(server)
		err := author.Login(query)
		if err == nil || err != ErrInvalidCredentials {
			return true, err
		}
	}

	return true, ErrInvalidCredentials
}
