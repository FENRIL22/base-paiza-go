package main

import (
	"testing"
	"os"
	"bytes"
	"bufio"
	"io/ioutil"
	"fmt"
	//"strings"
)

type Tester struct {
	filename string
}

func NewTester(filename string) *Tester {
	s := new(Tester)

	s.filename = filename

	return s
}

func (s *Tester) Process() (string, error) {
	filename := s.filename
	fp, _ := os.Open("./"+ filename)
	defer fp.Close()

	buf := &bytes.Buffer{}

	err := Question(fp, buf)
	out := buf.String()

	if err != nil {
		return "", err
	}

	return out, err
}

func (s *Tester) Judge() (bool, error) {
	ret, err := s.Process()
	if err != nil {
		return false, err
	}

	fp, _ := os.Open("./"+ s.filename + "_out")
	defer fp.Close()
	r := bufio.NewReader(fp)
	rb, err := ioutil.ReadAll(r)
	if err != nil {
		return false, err
	}

	if ret == string(rb) {
		return true, nil
	} else {
		return false, nil
	}
}

func (s *Tester) Debug() (string, string, error) {
	ret, err := s.Process()
	if err != nil {
		return "", "", err
	}

	fp, _ := os.Open("./"+ s.filename+"_out")
	defer fp.Close()
	r := bufio.NewReader(fp)
	rb, err := ioutil.ReadAll(r)
	if err != nil {
		return "", "", err
	}

	return ret, string(rb), nil
}


func TestQuestion(t *testing.T){
	testCount := 2
	for i:=0; i<testCount; i++ {
		s := NewTester(fmt.Sprintf("%v.dat", i))

		test, label, err := s.Debug()
		t.Log(fmt.Sprintf("-- Run Test: %v ---------------", i))
		t.Log("TEST:\n" + test)
		t.Log("LABL:\n" + label)
		if err != nil {
			t.Fatal(err)
		}

		jg, err := s.Judge()
		t.Log("JUDGE:" , jg)
		if err != nil {
			t.Fatal(err)
		}
		if !jg {
			t.Error("Invalid Ans")
		}
	}
}
