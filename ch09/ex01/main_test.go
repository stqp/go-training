// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		Deposit(200)
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	go func() {
		result := Withdraw(100)
		if !result {
			t.Errorf("not enough money..")
		}
		result = Withdraw(500)
		if result {
			t.Errorf("too much money..")
		}
		done <- struct{}{}
	}()

	<-done

	if got, want := Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
