package sqlstore

import (
	"github.com/aergoio/grafana/pkg/bus"
	m "github.com/aergoio/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
