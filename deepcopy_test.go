package deepcopy

import "testing"

type TestStruct struct {
	Data []string
	Ptr  *string
}

func TestDeepCopy(t *testing.T) {
	testStr := "test"
	orig := TestStruct{
		Data: []string{"a", "b", "c"},
		Ptr:  &testStr,
	}
	copyBuffer := NewCopyBuffer[TestStruct]()
	copy, copy_err := copyBuffer.DeepCopy(&orig)
	if copy_err != nil {
		t.Errorf("error copying object: %v", copy_err)
	}
	if &orig == copy {
		t.Errorf("original and copy are the same object")
	}
	if orig.Ptr == copy.Ptr {
		t.Errorf("orig.Ptr and copy.Ptr are the same object")
	}
	if *orig.Ptr != *copy.Ptr {
		t.Errorf("orig.Ptr and copy.Ptr are different")
	}
	if &orig.Data == &copy.Data {
		t.Errorf("orig.Data and copy.Data are the same object")
	}
	if len(orig.Data) != len(copy.Data) {
		t.Errorf("length of orig.Data and copy.Data are different")
	}
	for i := range orig.Data {
		if orig.Data[i] != copy.Data[i] {
			t.Errorf("orig.Data[%d] and copy.Data[%d] are different", i, i)
		}
		if &orig.Data[i] == &copy.Data[i] {
			t.Errorf("orig.Data[%d] and copy.Data[%d] are the same object", i, i)
		}
	}
}

func BenchmarkDeepCopy(b *testing.B) {
	testStr := "test"
	orig := TestStruct{
		Data: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
		Ptr:  &testStr,
	}
	copyBuffer := NewCopyBuffer[TestStruct]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyBuffer.DeepCopy(&orig)
	}
}
