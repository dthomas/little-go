package ipv4

import "testing"

func TestParseHeader(t *testing.T) {
	packet := []byte{0x45, 0x10, 0x00, 0x3c, 0x83, 0x1b, 0x40, 0x00, 0x40, 0x06, 0x15, 0x0a, 0xc0, 0xa8, 0x14, 0x46, 0x4a, 0x7d, 0x83, 0x1b}
	result := parseHeader(packet)
	expected := Header{
		Version:        4,
		IHL:            5,
		DSCP:           4,
		ECN:            0,
		TotalLength:    60,
		Identification: 33563,
		Flags:          2,
		FragmentOffset: 0,
		TTL:            64,
		Protocol:       6,
		HeaderChecksum: 5386,
		SourceIP:       3232240710,
		DestinationIP:  1249739547,
	}

	if result != expected {
		t.Error("Expected ", expected, " got ", result)
	}
}

func BenchmarkParseHeader(b *testing.B) {
	packet := []byte{0x45, 0x10, 0x00, 0x3c, 0x83, 0x1b, 0x40, 0x00, 0x40, 0x06, 0x15, 0x0a, 0xc0, 0xa8, 0x14, 0x46, 0x4a, 0x7d, 0x83, 0x1b}
	var r, result Header
	for n := 0; n < b.N; n++ {
		r = parseHeader(packet)
	}
	result = r
	r = result
}
