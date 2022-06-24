package utils

import (
	"encoding/hex"
	"encoding/json"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	hash := "e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746"
	s := struct{ Test string }{Test: "test"}
	t.Run("Hash is always same", func(t *testing.T) {
		x := Hash(s)
		t.Log(x)
		if x != hash {
			t.Errorf("Expected %s, got %s", hash, x)
		}
	})
	t.Run("Hash is hex encoded", func(t *testing.T) {
		x := Hash(s)
		_, err := hex.DecodeString(x)
		if err != nil {
			t.Error("Hash should be hex encoded")
		}
	})
}

func TestToBytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("ToBytes should return a slice of bytes got %s", reflect.TypeOf(b).Kind())
	}
}

func TestFromBytes(t *testing.T) {
	type testStruct struct {
		Test string
	}
	var restored testStruct
	ts := testStruct{"test"}
	b := ToBytes(ts)
	FromBytes(&restored, b)
	if !reflect.DeepEqual(ts, restored) {
		t.Error("FromBytes() should restore struct.")
	}
}

func TestToJSON(t *testing.T) {
	type testStruct struct{ Test string }
	s := testStruct{"test"}
	b := ToJSON(s)

	if reflect.TypeOf(b).Kind() != reflect.Slice {
		t.Errorf("expected %v and got %v", reflect.Slice, b)
	}

	var restored testStruct
	json.Unmarshal(b, &restored)

	if !reflect.DeepEqual(s, restored) {
		t.Error("ToJson() should restore struct.")
	}
}
