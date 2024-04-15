package main
import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if resulted != expected {
      t.Errorf("Add(2, 3) = "%d; want %d", result, expected)
    }
  }
