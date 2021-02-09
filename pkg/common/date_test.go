package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestDate_MarshalYAML(t *testing.T) {
	// TODO: Figure out how to marshal a date without the enclosing quotes
	tt, err := time.Parse("2006-01-02", "1988-09-27")
	if err != nil {
		panic(err)
	}
	d := struct {
		TestDate Date `yaml:"date"`
	}{
		TestDate: Date{tt},
	}

	s, err := yaml.Marshal(&d)

	assert.NoError(t, err)
	assert.Equal(t, "date: \"1988-09-27\"\n", string(s))
}

func TestDate_UnmarshalYAML_NotNull(t *testing.T) {
	s := "date: \"1988-09-27\""

	d := struct {
		TestDate Date `yaml:"date"`
	}{}

	err := yaml.Unmarshal([]byte(s), &d)

	expected, err := time.Parse("2006-01-02", "1988-09-27")
	if err != nil {
		panic(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, Date{expected}, d.TestDate)
}

func TestDate_UnmarshalYAML_Null(t *testing.T) {
	s := "date: null"

	d := struct {
		TestDate *Date `yaml:"date"`
	}{}

	err := yaml.Unmarshal([]byte(s), &d)

	assert.NoError(t, err)
	assert.Nil(t, d.TestDate)
}
