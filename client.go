package tweetsDailyFarseer

import (
	"compress/gzip"
	"errors"
	"github.com/ajg/form"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
		client: &http.Client{Timeout: 30 * time.Second},
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
	if err != nil {
		return "", err
	}

	setCommonHeaders(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)

	return string(body), nil

}

func setCommonHeaders(req *http.Request) {
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Referer", "https://www.oschina.net/tweets")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
}

func (c *Client) SendComment(id string, content string, cookieStr string) error {

	data := map[string]interface{}{"blog": id, "content": content}

	formData, err := form.EncodeToString(data)

	req, err := http.NewRequest("POST", addCommentUrl, strings.NewReader(formData))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", cookieStr)
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	setCommonHeaders(req)
	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("comment error code: " + resp.Status)
	}
	return nil

}
