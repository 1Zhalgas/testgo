package main

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func TestPrompt(t *testing.T) {
	expected := "-> "

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	w.Close()
	os.Stdout = old

	in, _ := io.ReadAll(r)
	out := string(in)

	if out != expected {
		t.Errorf("expected %s but got %s", expected, out)
	}
}

func TestIntro(t *testing.T) {
	expected := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	w.Close()
	os.Stdout = old

	in, _ := io.ReadAll(r)
	out := string(in)

	if out != expected {
		t.Errorf("expected %s but got %s", expected, out)
	}
}

func TestCheckNumber(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  string
		expected bool
		msg      string
	}{
		{"q letter", "q", true, ""},
		{"another letter", "a", false, "Please enter a whole number!"},
		{"prime", "7", false, "7 is a prime number!"},
		{"not prime", "8", false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", "0", false, "0 is not prime, by definition!"},
		{"one", "1", false, "1 is not prime, by definition!"},
		{"negative number", "-11", false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		scanner := bufio.NewScanner(strings.NewReader(e.testNum))
		msg, result := checkNumbers(scanner)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
func TestReadUserInput(t *testing.T) {
	testNumbers := "7\nq\n"
	scanner := bufio.NewReader(strings.NewReader(testNumbers))
	doneChan := make(chan bool)

	go readUserInput(scanner, doneChan)

	var output []string

	for res := range doneChan {
		if res {
			return
		}
		line, _, err := scanner.ReadLine()
		if err != io.EOF {
			return
		}
		output = append(output, string(line))
	}

	expected := []string{"7 is a prime number", "Goodbye."}
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}
}
