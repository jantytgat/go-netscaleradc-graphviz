package adcviz

import (
	"context"
	"encoding/json"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

func Generate(ctx context.Context, c Config, client *nitro.Client) (*Graph, error) {
	var err error
	var g = &Graph{}
	if g.LbVservers, err = generateLbVserverGraph(ctx, c, client); err != nil {
		return nil, err
	}
	return g, nil
}

type Graph struct {
	LbVservers map[string]LbVserver
}

func (g Graph) ToJson() (string, error) {
	var err error
	var jsonBytes []byte
	if jsonBytes, err = json.MarshalIndent(g, "", "  "); err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func generateLbVserverGraph(ctx context.Context, c Config, client *nitro.Client) (map[string]LbVserver, error) {
	var err error
	var vservers []config.LbVserver

	if vservers, err = client.LbVserver.List(ctx, c.LbVserverFields, nil); err != nil {
		return nil, err
	}

	var lbVservers = make(map[string]LbVserver, len(vservers))
	for _, vserver := range vservers {
		lbVservers[vserver.Name] = LbVserver{
			Name:           vserver.Name,
			IpAddress:      vserver.Ipv46,
			Port:           vserver.Port,
			Type:           vserver.ServiceType,
			ListenPolicy:   vserver.ListenPolicy,
			ListenPriority: vserver.ListenPriority,
			LbMethod:       vserver.LbMethod,
		}
	}
	return lbVservers, nil
}
