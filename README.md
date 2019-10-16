## An implementation of Luhn check-digit mod10 algorithm

[![Go Report Card](https://goreportcard.com/badge/github.com/kolkov/luhn)](https://goreportcard.com/report/github.com/kolkov/luhn)
[![Coverage Status](https://coveralls.io/repos/github/kolkov/luhn/badge.svg?branch=master)](https://coveralls.io/github/kolkov/luhn?branch=master)
[![Build Status](https://travis-ci.com/kolkov/luhn.svg?branch=master)](https://travis-ci.com/kolkov/luhn)

Algorithm will detect any single-digit error, as well as almost all transpositions of adjacent digits. It will not, however, detect transposition of the two-digit sequence 09 to 90 (or vice versa).

It is not intended to be a cryptographically secure hash function. It is mostly used for pre-flight credit card number validation as specified in [ISO/IEC 7812-1:2015](http://www.iso.org/iso/catalogue_detail?csnumber=66011)

Compatible with the [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) package.

### Usage ###

```
import "github.com/kolkov/luhn"

err := luhn.Validate("00123014764700968325")

signed := luhn.Generate("1")
```

test on your own by running `make benchmark`

### Resources ###

* [Credit Card Validation - Check Digits](https://web.eecs.umich.edu/~bartlett/credit_card_number.html)
* [Wikipedia - Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm)
