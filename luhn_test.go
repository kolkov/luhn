// Test for the Luhn package
package luhn

import "testing"

// Test some valid numbers
func TestValidNums(t *testing.T) {
	validNums := []string{"4507990000000010", "5323601111111112", "376411234531007", "45847712", "458477122"}
	for _, item := range validNums {
		if err := Validate(item); err == nil {
			if err == ErrValueNotLuhn {
				t.Error("Valid number validated as invalid", item)
			}
		} else if err == ErrValueNotLuhn {
			t.Error("Not a digits")
		}
	}
}

// Test some invalid numbers
func TestInvalidNums(t *testing.T) {
	invalidNums := []string{"4507990000000011", "5323601111111113", "376411234531004", "45847172", "458477121"}
	for _, item := range invalidNums {
		if err := Validate(item); err == nil {
			if err == ErrValueNotLuhn {
				t.Error("Invalid number validated as valid", item)
			}
		} else if err != ErrValueNotLuhn {
			t.Error("Not a digits")
		}
	}
}

// Test generating numbers
func TestLuhn(t *testing.T) {
	expectSignature := func(a string, b string) {
		c, err := Generate(a)
		if err != nil {
			t.Errorf("%v", err)
		} else if b != c {
			t.Errorf("for input " + a + " expected signature " + b + " but got " + c + " instead")
		}

		if err := Validate(c); err == nil {
			if err == ErrValueNotLuhn {
				t.Errorf("Unable to validate signature that was generated")
			}
		}
	}

	expectSignature("00123014764700968321002", "001230147647009683210024")
}

func BenchmarkLuhnSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate("123")
	}
}

func BenchmarkLuhnLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate("00123014764700968325")
	}
}
