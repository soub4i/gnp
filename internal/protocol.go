package internal

// WIP

type frame struct {
	Length   int
	Version  int
	Data     []byte
	Checksum int
}

func NewFrame(number int, length int, version int, data []byte) *frame {
	return &frame{
		Length:  length,
		Version: version,
		Data:    data,
	}
}

func (f *frame) checksum() int {
	sum := 0
	for _, b := range f.Data {
		sum += int(b)
	}
	return sum
}

func (f *frame) VerifyChecksum() bool {
	return f.Checksum == f.checksum()
}
