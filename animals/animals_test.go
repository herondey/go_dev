package animals

import(
	"testing"
)

func TestTigerFeed(t *testing.T){
	expect := "meat"
	actual := TigerFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestMonkeyFeed(t *testing.T){
	expect := "banana"
	actual := MonkeyFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}


func TestRabbitFeed(t *testing.T){
	expect := "carrot"
	actual := RabbitFeed()

	if expect != actual{
		t.Errorf("%s != %s", expect, actual)
	}
}

