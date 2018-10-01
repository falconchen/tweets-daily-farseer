package test

import (
	tf "github.com/falconchen/tweets-daily-farseer"
	"strconv"
	"testing"
)

func TestGetRemote(t *testing.T) {
	t.Log("Given  the need for getRemote \n")
	{

		//client := tf.NewClient()

		td := tf.New(tf.NewClient())
		item, err := td.GetRemote()
		if err != nil {
			t.Error(err)
		}

		t.Log(item)

	}

}

func TestGetLoal(t *testing.T) {
	t.Log("Given the test for GetLocal")
	{
		td := tf.New(tf.NewClient())
		item, err := td.GetLocal()
		if err != nil {
			t.Error(err)
		}

		t.Log(item)
	}
}

func TestCompare(t *testing.T) {
	t.Log("Given the test for Compare")
	{
		td := tf.New(tf.NewClient())
		local, err := td.GetLocal()
		if err != nil {
			t.Error(err)
		}

		t.Log("local item", local)

		remote, err := td.GetRemote()
		if err != nil {
			t.Error(err)
		}
		t.Log("remote item", remote)
	}
}

func TestUpdateLocal(t *testing.T) {
	t.Log("Given the test for Update local")
	{
		td := tf.New(tf.NewClient())
		num := 6234123
		id := []byte(strconv.Itoa(num))
		t.Log("Update local item")

		err := td.UpdateLocal(id)
		if err != nil {
			t.Error(err)
		}
		t.Log("update item")
	}
}

func TestUpdateLocalViaRemote(t *testing.T) {
	t.Log("Given the test for Update local")
	{
		td := tf.New(tf.NewClient())

		item, err := td.GetRemote()
		if err != nil {
			t.Error(err)
		}
		t.Logf("Remote item %v", item)
		id := []byte(item[1])
		t.Log("Update local id with ", item[1])

		err = td.UpdateLocal(id)
		if err != nil {
			t.Error(err)
		}
		t.Log("update item")

		item, err = td.GetLocal()
		if err != nil {
			t.Error(err)
		}

		t.Log("Get local data", item)

	}
}

/*
func TestMutiRemote(t *testing.T) {
	t.Log("Test for muti remote ")
	{
		var i int
		for {
			td := tf.New(tf.NewClient())
			remote, err := td.GetRemote()
			text := fmt.Sprintf("the %d Get %s \n", i+1, remote[0])
			fmt.Printf(text)

			if err != nil {
				t.Error(err)
			}

			local, err := td.GetLocal()
			if err != nil {
				t.Error(err)
			}

			if remote[1] != local[1] {
				td.UpdateLocal([]byte(local[1]))
				fmt.Printf("update: %s change to %s \n", local[1], remote[1])
				break
			} else {
				fmt.Printf("local %s==%s remote \n", local[1], remote[1])
			}

			time.Sleep(3 * time.Second)

		}

	}

}
*/

func TestSendComment(t *testing.T) {
	t.Log("Start testing send comment")
	{
		td := tf.New(tf.NewClient())
		id, content := "2222208", "国庆节快乐!"
		err := td.SendComment(id, content)
		if err != nil {
			t.Error(err)
		}

	}
}
