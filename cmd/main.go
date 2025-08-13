package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"

	"github.com/jantytgat/go-netscaleradc-graphviz/pkg/adcviz"
)

var (
	NITRO_ENV_NAME     string
	NITRO_ENV_ADDRESS  string
	NITRO_ENV_USERNAME string
	NITRO_ENV_PASSWORD string
)

func main() {
	var err error
	var file *os.File

	if file, err = os.Open("test.env"); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "=")
		switch values[0] {
		case "NITRO_ENV_NAME":
			NITRO_ENV_NAME = values[1]
			// fmt.Println("NITRO_ENV_NAME:", NITRO_ENV_NAME)
		case "NITRO_ENV_ADDRESS":
			NITRO_ENV_ADDRESS = values[1]
			// fmt.Println("NITRO_ENV_ADDRESS:", NITRO_ENV_ADDRESS)
		case "NITRO_ENV_USERNAME":
			NITRO_ENV_USERNAME = values[1]
			// fmt.Println("NITRO_ENV_USERNAME:", NITRO_ENV_USERNAME)
		case "NITRO_ENV_PASSWORD":
			NITRO_ENV_PASSWORD = values[1]
			if NITRO_ENV_PASSWORD == "" {
				panic(errors.New("NITRO_ENV_PASSWORD is empty"))
			}
		default:
			panic(values)
		}
		line++
	}
	var client *nitro.Client

	if client, err = nitro.NewClient(
		NITRO_ENV_NAME,
		NITRO_ENV_ADDRESS,
		nitro.Credentials{
			Username: NITRO_ENV_USERNAME,
			Password: NITRO_ENV_PASSWORD,
		},
		nitro.ConnectionSettings{
			UseSsl:                    true,
			Timeout:                   0,
			UserAgent:                 "dev-nitro",
			ValidateServerCertificate: false,
			LogTlsSecrets:             false,
			LogTlsSecretsDestination:  "",
			AutoLogin:                 true,
		},
		nitro.StrictSerializationMode); err != nil {
		panic(errors.New("error creating client: " + err.Error()))
	}

	defer func() {
		if client.IsLoggedIn() {
			if err = client.Logout(); err != nil {
				panic(err)
			}
		}
	}()

	// fmt.Println(client.BaseUrl())
	ctx := context.Background()

	var conf = adcviz.Config{
		LbVserverFields: []string{
			config.LbVserverFieldNames.Name,
			config.LbVserverFieldNames.Ipv46,
			config.LbVserverFieldNames.Port,
			config.LbVserverFieldNames.ServiceType,
			config.LbVserverFieldNames.ListenPolicy,
			config.LbVserverFieldNames.ListenPriority,
			config.LbVserverFieldNames.LbMethod,
		},
		LbVserverServiceGroupBindingFields: []string{
			config.LbVserverServiceGroupBindingFieldNames.ServiceGroupName,
			config.LbVserverServiceGroupBindingFieldNames.Order,
		},
	}

	var graph *adcviz.Graph
	if graph, err = adcviz.Generate(ctx, conf, client); err != nil {
		panic(err)
	}

	for _, lb := range graph.LbVservers {
		if err = lb.SaveAsPng(lb.Name + ".png"); err != nil {
			panic(err)
		}
		fmt.Println(lb.String())
		break
	}
	// time.Sleep(1 * time.Second)
	// var json string
	// if json, err = graph.ToJson(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(json)
	fmt.Println("Done.")
	os.Exit(0)
}
