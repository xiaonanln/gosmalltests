package compress

import (
	"bytes"
	"compress/flate"
	"math/rand"
	"testing"
)

func randomData() []byte {
	rndlen := 1024 + rand.Intn(2048)
	data := make([]byte, rndlen)
	for i := 0; i < rndlen; i++ {
		data[i] = byte('a' + rand.Intn(26))
	}
	return data
}

func TestCompress(t *testing.T) {
	for i := 0; i < 1000; i++ {
		data := randomData()
		w := bytes.NewBuffer(nil)
		cw, err := flate.NewWriter(w, flate.BestSpeed)
		if err != nil {
			t.Error(err)
		}
		cw.Write(data)
		cw.Flush()
		cw.Close()
		cdata := w.Bytes()
		//t.Log("before compress", len(data), "after compress", len(cdata), "compress rate", len(cdata)*100/len(data), "%")

		cr := flate.NewReader(bytes.NewReader(cdata))
		rdata := make([]byte, len(data))
		n, err := cr.Read(rdata)
		if err != nil {
			t.Error(err)
		}
		if n != len(data) {
			t.Errorf("size is wrong")
		}

		if string(data) != string(rdata) {
			t.Errorf("restore data is wrong")
		}
	}
}

func TestUncompress(t *testing.T) {
	data := []byte{11, 0, 87, 87, 82, 50, 99, 54, 100, 95, 89, 121, 116, 73, 65, 65, 120, 79, 5, 0, 0, 0, 76, 111, 103, 105, 110, 2, 0, 8, 0, 0, 0, 167, 116, 101, 115, 116, 49, 54, 53, 7, 0, 0, 0, 166, 49, 50, 51, 52, 53, 54}

	t.Log("uncompress data len", len(data))

	cr := flate.NewReader(bytes.NewReader(data))
	rdata := make([]byte, len(data)*2)
	n, err := cr.Read(rdata)
	t.Logf("n %d err %v", n, err)
}
