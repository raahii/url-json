package cmd

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "url-json <url>",
	Short: "url-json print decomposed parameters of a url in json format",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := printUrlComponents(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type urlComponents struct {
	Scheme   string                 `json:"scheme"`
	User     *userInfo              `json:"user"`
	Host     string                 `json:"host"`
	Port     string                 `json:"port"`
	Path     string                 `json:"path"`
	Fragment string                 `json:"fragment"`
	Queries  map[string]interface{} `json:"queries"`
}

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func printUrlComponents(rawInputURL string) error {
	inputURL, err := parseURL(rawInputURL)
	if err != nil {
		return fmt.Errorf("failed to parse url: %w", err)
	}

	components := newURLComponents(inputURL)
	jsonBytes, err := json.Marshal(components)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	fmt.Println(string(jsonBytes))

	return nil
}

func parseURL(raw string) (*url.URL, error) {
	urlString := strings.TrimSpace(raw)

	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func newURLComponents(inputUrl *url.URL) *urlComponents {
	var err error
	c := new(urlComponents)

	c.Scheme = inputUrl.Scheme
	c.Path = inputUrl.Path
	c.Fragment = inputUrl.Fragment

	c.Host, c.Port, err = net.SplitHostPort(inputUrl.Host)
	if err != nil {
		c.Host = inputUrl.Host
	}

	c.User = new(userInfo)
	c.User.Username = inputUrl.User.Username()
	c.User.Password, _ = inputUrl.User.Password()

	c.Queries = make(map[string]interface{})
	for k, v := range inputUrl.Query() {
		if len(v) == 1 {
			c.Queries[k] = v[0]
		} else {
			c.Queries[k] = v
		}
	}

	return c
}
