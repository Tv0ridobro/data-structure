package sieve_of_eratosthenes

import (
	"github.com/Tv0ridobro/data-structure/util"
	"testing"
)

func TestSieveOfEratosthenes_GetDelimiters(t *testing.T) {
	s := New(101)
	if v := s.GetDelimiters(100); !util.Equal(v, []int{2, 2, 5, 5}) {
		t.Errorf("wrong answer %v %v", v, []int{2, 2, 5, 5})
	}
	if v := s.GetDelimiters(97); !util.Equal(v, []int{97}) {
		t.Errorf("wrong answer %v %v", v, []int{97})
	}
	if v := s.IsPrime(97); v != true {
		t.Errorf("wrong answer %v %v", v, true)
	}
	if v := s.IsPrime(98); v != false {
		t.Errorf("wrong answer %v %v", v, false)
	}
	if v := s.IsPrime(0); v != false {
		t.Errorf("wrong answer %v %v", v, false)
	}
	if v := s.IsPrime(1); v != false {
		t.Errorf("wrong answer %v %v", v, false)
	}
	if v := s.IsPrime(2); v != true {
		t.Errorf("wrong answer %v %v", v, true)
	}
	if v := s.Primes(); !util.Equal(v, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}) {
		t.Errorf("wrong answer %v %v", v, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97})
	}
}
