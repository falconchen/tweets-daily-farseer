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
		client: &http.Client{Timeout: 10 * time.Second},
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

func (c *Client) SendComment(id string, content string) error {

	data := map[string]interface{}{"blog": id, "content": content}

	formData, err := form.EncodeToString(data)

	req, err := http.NewRequest("POST", addCommentUrl, strings.NewReader(formData))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "_user_behavior_=f593857c-23a6-4e1c-b6de-d85330da3d23; Hm_lvt_a411c4d1664dd70048ee98afe7b28f0b=1536028574; oscid=V1RqlbaVAFpbxBsxcj3XptlomrvkuzuqTuy22x0pRNcaDvaWxvhwXL6IWpMgqhjF0dXpy%2F0QeK8f6JvThBgPaqp0QURJalk3vyzpbEkNKm1svLrF37W8RzyhDUqigfz%2FHpok3jQZ9bhFlYz58nZGBYp6diZh4QE7JEZEggF1Jbw%3D; Hm_lvt_cb47adfe0fabd7059a2a90a495077efe=1538147391,1538286249,1538286341,1538288930; Hm_lvt_d237257153dcc02ba4647b448cbafcb8=1535951354; Hm_lpvt_a411c4d1664dd70048ee98afe7b28f0b=1538288926; aliyungf_tc=AQAAAJ5jAmYvxwYAyljleJtPVLrqlF/X; Hm_lpvt_cb47adfe0fabd7059a2a90a495077efe=1538368874")
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
