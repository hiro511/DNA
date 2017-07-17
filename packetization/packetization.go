package packetization

type (
	// Packet represents size and payload
	Packet struct {
		Size    int
		Payload []byte
	}
)

// Packetize data with the specified size
func Packetize(data []byte, packetSize int) []Packet {

	packetNum := (len(data) / packetSize) + 1
	packets := make([]Packet, packetNum)

	for i := range packets {
		packets[i].Payload = make([]byte, 0, packetSize)
		packets[i].Size = packetSize

		begin := packetSize * i
		end := begin + packetSize
		if end > len(data) {
			end = len(data)
			packets[i].Size = end - begin
		}
		packets[i].Payload = append(packets[i].Payload, data[begin:end]...)
	}

	return packets
}

// Depacketize the specified packets
func Depacketize(packets []Packet, packetSize int) []byte {
	data := make([]byte, 0, len(packets)*packetSize)

	for i := range packets {
		data = append(data, packets[i].Payload...)
	}

	return data
}
