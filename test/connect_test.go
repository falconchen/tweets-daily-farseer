package test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestConnect(t *testing.T) {
	url := "https://my.oschina.net/xxiaobian"
	statusCode := 200

	t.Log("Given the need to test to tweets daily list")

	{
		t.Logf("\t Test connect to \"%s\" for \"%d\" ", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t\t connect status \"%s\" for %d ", ballotX, statusCode)

			}
			t.Log("\t\t connect status ", checkMark, statusCode)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\t connect \"%s\" status \"%v\" ", checkMark, statusCode)
			} else {
				t.Errorf("\t\t connect \"%s\" status \"%v\" ", ballotX, resp.StatusCode)
			}

		}
	}

}
