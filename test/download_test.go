package test

import (
	tf "github.com/falconchen/tweets-daily-farseer"
	"regexp"
	"testing"
)

type testClients struct {
	client *tf.Client
	ip     string
}

func TestDownload(t *testing.T) {

	t.Log("Given the need to test downloaders.")
	{

		client := tf.NewClient()

		proxyClient := tf.NewClient()
		//fmt.Printf("%v", client)
		//Setup SOCK5
		_, err := proxyClient.ClientWithSOCKS5("tcp", "127.0.0.1:1080")
		if err != nil {
			t.Fatalf("SOCK5 Init failed: %s", err.Error())
		}

		//匿名结构体
		//clients := []struct {
		//	client *tf.Client
		//	ip     string
		//}{
		//	{client, "1.2.3.4"},
		//	{proxyClient, "1.3.5.8"},
		//}

		clientsGroup := []*tf.Client{
			client,
			proxyClient,
		}

		t.Log("Check downloaders' ip")
		{
			for i, c := range clientsGroup {
				resp, err := c.Get("https://ip.cn")
				if err != nil {
					t.Errorf("fetch error: %s", err.Error())
				}
				reg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)

				//ip := reg.FindAllString(resp, -1)

				ip := reg.FindString(resp)
				t.Logf("%T,%v", ip, ip)
				t.Logf("===finish client %d test\n", i+1)

			}
		}

	}

}
