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
