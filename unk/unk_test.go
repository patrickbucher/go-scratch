package unk

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReplaceUnk(t *testing.T) {
	expected, err := ioutil.ReadFile("expected")
	if err != nil {
		t.Error(err)
	}
	input, err := os.Open("input")
	if err != nil {
		t.Error(err)
	}
	output, err := os.Create("output")
	if err != nil {
		t.Error(err)
	}
	ReplaceUnk(input, output)
	defer input.Close()
	defer output.Close()
	actual, err := ioutil.ReadFile("output")
	if err != nil {
		t.Error(err)
	}
	if string(expected) != string(actual) {
		t.Errorf("expected:\n%s\nactual:\n%s", expected, actual)
	}
}

func BenchmarkReplaceUnk(b *testing.B) {
	input, err := os.Open("input")
	if err != nil {
		b.Error(err)
	}
	output, err := os.Create("output")
	if err != nil {
		b.Error(err)
	}
	err = ReplaceUnk(input, output)
	if err != nil {
		b.Error(err)
	}
}
