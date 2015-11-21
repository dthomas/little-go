package ipv4

// Header IPv4 Packet header
type Header struct {
	Version        int
	IHL            int
	DSCP           int
	ECN            int
	TotalLength    int
	Identification int
	Flags          int
	FragmentOffset int
	TTL            int
	Protocol       int
	HeaderChecksum int
	SourceIP       int
	DestinationIP  int
}

func parseHeader(packet []byte) (header Header) {
	header.Version = int(packet[0]) >> 4
	header.IHL = (int(packet[0]) & 0x0F)
	header.DSCP = (int(packet[1]) & 0xFC) >> 2
	header.ECN = (int(packet[1]) & 0x03)
	header.TotalLength = (int(packet[2]) << 8) + int(packet[3])
	header.Identification = (int(packet[4]) << 8) + int(packet[5])
	header.Flags = (int(packet[6]) & 0x70) >> 5
	header.FragmentOffset = (int(packet[6]) & 0x20) + int(packet[7])
	header.TTL = int(packet[8])
	header.Protocol = int(packet[9])
	header.HeaderChecksum = (int(packet[10]) << 8) + int(packet[11])
	header.SourceIP = (int(packet[12]) << 24) + (int(packet[13]) << 16) + (int(packet[14]) << 8) + (int(packet[15]))
	header.DestinationIP = (int(packet[16]) << 24) + (int(packet[17]) << 16) + (int(packet[18]) << 8) + (int(packet[19]))

	return header
}
