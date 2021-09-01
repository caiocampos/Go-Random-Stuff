package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var generated map[string]bool

func init() {
	generated = make(map[string]bool, 676000)
}

func randomChar() string {
	return string(rand.Intn(26) + 'A')
}

func randomNum() int {
	return rand.Intn(1000)
}

func randomName() (string, error) {
	if len(generated) == 676000 {
		return "", errors.New("Namespace is exhausted")
	}
	for {
		name := fmt.Sprintf("%s%s%03d", randomChar(), randomChar(), randomNum())
		if !generated[name] {
			generated[name] = true
			return name, nil
		}
	}
}

// Robot type defines a robot
type Robot struct {
	name string
}

// Name method returns the Robot name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := randomName()
		if err != nil {
			return "", err
		}
		r.name = name
	}
	return r.name, nil
}

// Reset method resets the Robot name
func (r *Robot) Reset() {
	r.name = ""
}
