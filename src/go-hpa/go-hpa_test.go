package main

import "testing"

func TestGreeting(t *testing.T) {
     loop := LoopHole("Code!!!")
     var expected string = "<b>Code!!!</b>";
     if loop != expected {
            t.Errorf("Result inesperado, foi: %s, tinha que ser: %s.", loop, expected)
     }
 }