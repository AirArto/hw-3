package stringcount

import (
	"testing"
)

func TestCount(t *testing.T) {
	packedArr := "oNe \\seven\\tWo, tHree, fOur,oNe \\seven\\tWo,   seven  tHree, fOur,//six nine, two, three-foUr/seven/, nine.../seven///eight//'.four'" +
		" nine.nine.nine.nine.nine//nine eight//. eight  seven six five oNe \\\\tWo, tHree, fOur/seven/six nine.,;',nine three'/seven///eight//>/seven/'.four123nine" +
		"//eight////eight//six five five;five//eight//eight//;<>0]1354six//.six['six///eight//five';five]]['six*five//eight,cat,cat,... dog some more random words lorem ipsum"
	UnpackedArr := []string{"nine", "eight", "seven", "six", "five", "four", "three", "two", "one", "cat"}
	data, err := Do(&packedArr)
	for i, str := range UnpackedArr {
		if data[i] != str {
			t.Errorf("Unexpected result while string unpacking \nExpected:%s \nReceived:%s", UnpackedArr, data)
		}
		if err != nil {
			t.Errorf("Unexpected result while string unpacking \n%s", err)
		}
	}
}
