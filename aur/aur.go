package aur

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type response struct {
	Resultcount int       `json:"resultcount"`
	Results     []Package `json:"results"`
	Type        string    `json:"type"`
	Version     int       `json:"version"`
	Error       string    `json:"error"`
}

const (
	baseURL    = "https://aur.archlinux.org/rpc/"
	RPCVersion = "5"
)

// https://aur.archlinux.org/rpc/?v=5&type=info&arg[]=pkg1&arg[]=pkg2
func Info(query ...string) ([]Package, error) {
	url, err := makeRpcInfoUrl(RPCVersion, query)
	if err != nil {
		return []Package{}, fmt.Errorf("error creating url: %w", err)
	}

	result, err := rpcRequest(url)

	return result.Results, nil
}

func search(query, keyword string) ([]Package, error) {
	url, err := makeRpcSerarchUrl(RPCVersion, "search", keyword, query)
	if err != nil {
		return []Package{}, fmt.Errorf("error creating url: %w", err)
	}

	result, err := rpcRequest(url)

	return result.Results, nil
}

func rpcRequest(url string) (response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response{}, fmt.Errorf("could not create request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response{}, fmt.Errorf("error making http request: %w", err)
	}
	defer res.Body.Close()

	var result response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return response{}, fmt.Errorf("could not unmarshal JSON: %w", err)
	}
	log.Println(result)

	if result.Type == "error" {
		log.Println(result.Error)
		return response{}, errors.New("Error type")
	}

	return result, nil
}

func SearchByName(query string) ([]Package, error) {
	return search(query, "name")
}

func SearchByNameDesc(query string) ([]Package, error) {
	return search(query, "name-desc")
}

func SearchByMaintainer(query string) ([]Package, error) {
	return search(query, "maintainer")
}

func Install(pkg ...Package) error {
	return nil
}

func Update() error {
	return nil
}

func Delete(pkg ...Package) error {
	return nil
}

// https://aur.archlinux.org/rpc/?v=5&type=info&arg[]=pkg1&arg[]=pkg2
func makeRpcInfoUrl(version string, args []string) (string, error) {
	queryParams := url.Values{}
	queryParams.Set("type", "info")
	queryParams.Set("v", version) //RPC version
	for _, arg := range args {
		queryParams.Set("arg", arg)
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("Error parsing base URL:%w", err)
	}
	u.RawQuery = queryParams.Encode()

	log.Println("AUR URL:", u.String())
	return u.String(), nil
}

// makeRPCURL makes URL for GET request to AUR API.
// It returns URL like:
// https://aur.archlinux.org/rpc?arg=<arg>&by=<field>&type=<queryType>&v=<version>
func makeRpcSerarchUrl(version, queryType, field, arg string) (string, error) {
	queryParams := url.Values{}
	queryParams.Set("arg", arg)
	queryParams.Set("by", field)
	queryParams.Set("type", queryType)
	queryParams.Set("v", version) //RPC version

	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("Error parsing base URL:%w", err)
	}
	u.RawQuery = queryParams.Encode()

	log.Println("AUR URL:", u.String())
	return u.String(), nil
}
