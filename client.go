package tweetsDailyFarseer

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client       *http.Client
	proxyType    int
	proxyUrl     string
	socks5Proto  string
	socks5IpPort string
}

/// Init client structure
func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c *Client) GetTest(url string) (string, error) {
	if c.client == nil {
		return "", errors.New("Client is nil, should init with proxy")
	}
	resp, err := c.client.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *Client) Get(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	req.Header.Add("Referer", "https://twitter.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil

}
