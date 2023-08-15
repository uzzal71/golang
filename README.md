# golang

Go Language Practices

## Golang Types

### Numbers

#### Integers

###### Unsigned Integers (no negatives)

- uint8 / byte (0 - 255)
- uint16 (0 - 65535)
- uint32 (0 - 4294967295)
- uint64 (0 to 18446744073709551615)

###### Signed Integers

- int8 (-128 to 127)
- int16 (-32768 to 32767)
- int32 / rune (-2147483648 to 2147483647)
- int64 (-9223372036854775808 to 9223372036854775807)

###### Machine Dependent Types

- uint (32 or 64 bits)
- int (same size as uint)
- uintptr (an unsigned integer to store the uninterpreted bits of a pointer value)

# Floating Point Numbers

##### Floats

- float32 (IEEE-754 32-bit floating-point numbers)
- float64 (IEEE-754 64-bit floating-point numbers)

##### Complex (Imaginary Parts)

- complex64 (Complex numbers with float32 real and imaginary parts)
- complex128 (Complex numbers with float64 real and imaginary parts)

# Strings

- "Hello World"

# Boleans

- true
- false

# FMT Cheat Sheet

## General

- %v (value in default format)
- %T (type)
- %% (literal %)

## Boolean

- %t (true or false)

## Integer

- %b (base 2)
- %o (base 8)
- %d (base 10)
- %x (base 16)

## Floating Points

- %e (scientific notation)
- %f / %F (decimal no exponent)
- %g (for large exponents)

## Strings

- %s (default)
- %q (double quoted string)

## Width & Percision

- %f (default width, default precision)
- %9f (width 9, default precision)
- % 2f (default width precision 2f)
- %9.2f (width 9, precision 2)
- %9.f (width 9, precision 0)

## Padding

- %08d (pads digit to length 9 with preceeding 0's)
- %-4d (Pads with spaces (width 4, left justified))

## Methods

- Sprintf() format without printing
- Printf() format with printing
