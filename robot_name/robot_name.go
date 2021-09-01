package robotname

import (
	"math/rand"
)

// Robot type defines a robot
type Robot struct {
	name string
}

// Name method returns the Robot name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.Reset()
	}
	return r.name, nil
}

// Reset method resets the Robot name
func (r *Robot) Reset() {
	random := rand.New(rand.NewSource(rand.Int63()))
	rNum := func() byte {
		return byte(random.Intn(10) + 48)
	}
	rChar := func() byte {
		return byte(random.Intn(26) + 65)
	}
	bytes := []byte{rChar(), rChar(), rNum(), rNum(), rNum()}
	r.name = string(bytes)
}
