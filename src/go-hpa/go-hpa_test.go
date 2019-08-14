package main

import "testing"

func TestGreeting(t *testing.T) {
     loop := LoopHole("Code!!!")
     var expected string = "<b>Code!!!</b>";
     if greeting != expected {
            t.Errorf("Result inesperado, foi: %s, tinha que ser: %s.", greeting, expected)
     }
 }