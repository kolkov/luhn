## An implementation of Luhn check-digit mod10 algorithm

[![Go Report Card](https://goreportcard.com/badge/kolkov/luhn)](https://goreportcard.com/report/kolkov/luhn)

Algorithm will detect any single-digit error, as well as almost all transpositions of adjacent digits. It will not, however, detect transposition of the two-digit sequence 09 to 90 (or vice versa).

It is not intended to be a cryptographically secure hash function. It is mostly used for pre-flight credit card number validation as specified in [ISO/IEC 7812-1:2015](http://www.iso.org/iso/catalogue_detail?csnumber=66011)

### Usage ###

```
import "github.com/kolkov/luhn"

result, _ := luhn.Validate("00123014764700968325")

signed := luhn.Generate("1")
```

test on your own by running `make benchmark`

### Resources ###

* [Credit Card Validation - Check Digits](https://web.eecs.umich.edu/~bartlett/credit_card_number.html)
* [Wikipedia - Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm)