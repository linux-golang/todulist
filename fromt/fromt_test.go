package fromt

import "testing"

func TestWriterTable(t *testing.T) {

	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}
	WriterTable(data)

}
