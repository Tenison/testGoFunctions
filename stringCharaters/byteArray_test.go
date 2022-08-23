package stringCharaters

import (
	"reflect"
	"testing"
)

func TestReplaceString(t *testing.T) {
	replacement := ReplaceCharater("Hii There!!", 5, "t")

	if replacement != "Hii there!!" {
		t.Error(map[string]string{"Failed": "Expected 'Hii there!!'"})
	}
}

func TestFirstCharacter(t *testing.T) {
	first := FirstStringCharater("Hello")

	if first != "H" {
		t.Error(map[string]string{"Failed": "Expected 'H'"})
	}

}

func TestConvertStringToByteArray(t *testing.T) {
	Barray := ConvertStringToByteArray("hello")

	if reflect.TypeOf(Barray) != reflect.TypeOf([]byte{}) {
		t.Error(map[string]string{"Failed": "Expected []byte"})
	}
}
