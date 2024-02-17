# Gin Web Dice Rolling Server

Playing with gin-gonic

## Dice

- d4
- d6
- d8
- d10
- d12
- d20
- d100

## Dice Command Examples

- 3d8+5: d8 + d8 + d8 + 5
- d20-1: d20 - 1

## Regular Expression

```go
    re := regexp.MustCompile(`^([0-9]*)d([0-9]*)([+-]?)([0-9]*)$`)
```
