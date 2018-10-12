package tweetsDailyFarseer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const listUrl = "https://my.oschina.net/xxiaobian"

//const listUrl = "https://my.oschina.net/falcon10086"
const addCommentUrl = "https://my.oschina.net/space/blog/add_comment"
const dataFile = "../data/id.txt"

type TwittersDaily struct {
	c         *Client
	cookieStr string
}

func New(c *Client, cookieStr string) *TwittersDaily {
	return &TwittersDaily{c, cookieStr}
}

//
func (t *TwittersDaily) GetRemote() ([]string, error) {
	content, err := t.c.Get(listUrl)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(listUrl + `/blog/(\d+)`)
	matches := re.FindStringSubmatch(content)

	if matches == nil {
		return nil, errors.New("blog article not found")
	}
	return matches, nil
}

func (t *TwittersDaily) GetLocal() ([]string, error) {

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return nil, err
		//fmt.Print(err)
	}

	id := string(data)
	if id == "" {
		return []string{"", ""}, err
	}
	return []string{t.makeArtilceUrl(id), id}, nil

}

func (t TwittersDaily) makeArtilceUrl(id string) string {
	return fmt.Sprintf("%s/blog/%s", listUrl, id)
}

func (t *TwittersDaily) UpdateLocal(id []byte) error {

	if err := ioutil.WriteFile(dataFile, id, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (t *TwittersDaily) SetCookieStr(cookieStr string) {
	t.cookieStr = cookieStr
}

func (t *TwittersDaily) CheckShafaExists(id string) (bool, error) {

	content, err := t.c.Get(t.makeArtilceUrl(id))
	if err != nil {
		return false, err
	}
	return strings.Contains(content, "data-comment-id="), nil

}

func (t *TwittersDaily) SendComment(id string, content string) error {
	return t.c.SendComment(id, content, t.cookieStr)
}
