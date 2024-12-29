/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package uuid

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	const n = 100000
	previous := make(map[UUID]struct{}, n)

	for i := 0; i < n; i++ {
		id, err := New()
		if err != nil {
			t.Fatal(err) //.............................................................................................
		}

		if !Validate(id) {
			t.Fatalf("UUID not valid: %s", id) //.......................................................................
		}

		if _, ok := previous[id]; ok {
			t.Fatal("a duplicate UUID has been generated") //...........................................................
		}

		previous[id] = struct{}{}
	}
}

func TestErrorCase(t *testing.T) {
	_, err := generate(strings.NewReader(""))
	if err == nil {
		t.Fatal("error expected") //....................................................................................
	}
}

/*
####### END ############################################################################################################
*/
