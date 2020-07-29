package apollo

import (
	"github.com/douyu/jupiter/pkg/conf"
	"github.com/douyu/jupiter/pkg/datasource/manager"
	"github.com/douyu/jupiter/pkg/flag"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/philchia/agollo"
	"net/url"
)

// DataSourceApollo defines apollo scheme
const DataSourceApollo = "apollo"

func init() {
	manager.Register(DataSourceApollo, func() conf.DataSource {
		var (
			configAddr = flag.String("config")
		)
		if configAddr == "" {
			xlog.Panic("new apollo dataSource, configAddr is empty")
			return nil
		}
		// configAddr is a string in this format:
		// apollo://ip:port?appId=XXX&cluster=XXX&namespaceName=XXX&key=XXX
		urlObj, err := url.Parse(configAddr)
		if err != nil {
			xlog.Panic("parse configAddr error", xlog.FieldErr(err))
			return nil
		}
		apolloConf := agollo.Conf{
			AppID:          urlObj.Query().Get("appId"),
			Cluster:        urlObj.Query().Get("cluster"),
			NameSpaceNames: []string{urlObj.Query().Get("namespaceName")},
			IP:             urlObj.Host,
		}
		return NewDataSource(&apolloConf, urlObj.Query().Get("namespaceName"), urlObj.Query().Get("key"))
	})
}