// TODO: Move this into the api/pkg/models
package models

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Tlp int

const (
	RED Tlp = iota + 1
	AMBER
	GREEN
	WHITE
)

func (t Tlp) String() string {
	switch t {
	case RED:
		return "RED"
	case AMBER:
		return "AMBER"
	case GREEN:
		return "GREEN"
	case WHITE:
		return "WHITE"
	default:
		panic(fmt.Sprintf("%d is not a valid Tlp value", t))
	}
}

func ParseTlp(t string) (Tlp, error) {
	switch strings.ToLower(t) {
	case "red":
		return RED, nil
	case "amber":
		return AMBER, nil
	case "green":
		return GREEN, nil
	case "white":
		return WHITE, nil
	default:
		return 0, errors.Errorf("%s cannot be parsed to a valid Tlp value", t)
	}
}
