package packetization

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacketize(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/Lenna.png")
	assert.NoError(t, err)

	packets := Packetize(data, 32)
	origin := Depacketize(packets, 32)

	assert.Equal(t, len(data), len(origin))
	assert.Equal(t, data, origin)
}
