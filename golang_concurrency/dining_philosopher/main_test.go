package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thingTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished := []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length ;expected 5 got %d", len(orderFinished))
		}
	}

}
