package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	buffer := bytes.NewBufferString("word1 word2 word3 word4\n")

	expected := 4
	result := count(buffer, false, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	buffer := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	expected := 3
	result := count(buffer, true, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	buffer := bytes.NewBufferString("123456789\n")

	expected := 10
	result := count(buffer, false, true)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}
