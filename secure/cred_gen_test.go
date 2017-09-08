package secure

import "testing"

func TestGenCa(t *testing.T) {
	g := &Generator{
		Country:      "US",
		Organization: "Engineering",
	}
	ca, err := g.CA()
	if err != nil {
		t.Fatal(err)
	}
	if len(ca.Raw) == 0 {
		t.Error(len(ca.Raw))
	}

	c, err := g.Cert(ca)
	if err != nil {
		t.Fatal(err)
	}
	if len(c.Raw) == 0 {
		t.Error(len(c.Raw))
	}
}
