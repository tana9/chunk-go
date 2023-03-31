package chunk

import "testing"

func TestChunk(t *testing.T) {
	type data struct {
		name string
	}
	datas := []data{
		{name: "1"},
		{name: "2"},
		{name: "3"},
	}
	results := Chunk(datas, 2)
	if len(results) != 2 {
		t.Fail()
	}
	if len(results[0]) != 2 {
		t.Fail()
	}
	if len(results[1]) != 1 {
		t.Fail()
	}
}
