[![Go Report Card](https://goreportcard.com/badge/github.com/nikunjy/rules?style=flat-square)](https://goreportcard.com/report/github.com/nikunjy/rules)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/github.com/nikunjy/rules)](https://pkg.go.dev/github.com/nikunjy/rules)
[![Release](https://img.shields.io/github/v/release/nikunjy/rules?sort=semver&style=flat-square)](https://github.com/nikunjy/rules/releases/latest)
[![Release](https://img.shields.io/github/go-mod/go-version/nikunjy/rules?style=flat-square)](https://github.com/nikunjy/rules/releases/latest)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg?style=flat-square)](https://github.com/nikunjy/rules/commits)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://img.shields.io/github/stars/nikunjy/rules?style=flat-square&label=Star&maxAge=2592000)](https://github.com/nikunjy/rules/stargazers/)

# Golang Rules Engine

Rules engine written in golang with the help of antlr.

This package will be very helpful in situations where you have a generic rule and want to verify if your values (specified using `map[string]interface{}`) satisfy the rule.

Here are some examples:

```
  parser.Evaluate("x eq 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x == 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x lt 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x < 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x gt 1", map[string]interface{}{"x": 1})

  parser.Evaluate("x.a == 1 and x.b.c <= 2", map[string]interface{}{
    "x": map[string]interface{}{
       "a": 1,
       "b": map[string]interface{}{
          "c": 2,
       },
    },
  })


  parser.Evaluate("y == 4 and (x > 1)", map[string]interface{}{"x": 1})

  parser.Evaluate("y == 4 and (x IN [1,2,3])", map[string]interface{}{"x": 1})

  parser.Evaluate("y == 4 and (x eq 1.2.3)", map[string]interface{}{"x": "1.2.3"})

```

## Operations

All the operations can be written capitalized or lowercase (ex: `eq` or `EQ` can be used)

Logical Operations supported are `and` `or`

Compare Expression and their definitions

| expression | meaning   | 
-------------|------------
| eq         | equals to |
| ==         | equals to |
| ne         | not equal to |
| !=         | not equal to |
| lt         | less than |
| <          | less than |
| gt         | greater than |
| >          | greter than |
| le         | less than or equal to |
| <=          | less than or equal to  |
| ge          | greater than or equal to|
| >=          | greater than or equal to|
| co         | contains  |
| sw         | starts with  |
| ew         | ends with  |
| in         | in a list  |
| pr         | present, will be true if you have a key as true  |
| not         | not of a logical expression  |

## How to use it

Use your dependency manager to import `github.com/nikunjy/rules/parser`. This will let you parse a rule and keep the parsed representation around.
Alternatively, you can also use `github.com/nikunjy/rules` directly to call the root `Evaluate(string, map[string]interface{})` method.

I would recommend importing `github.com/nikunjy/rules/parser`

## How to extend the grammar

1. Please look at this [antlr tutorial](https://tomassetti.me/antlr-mega-tutorial/#setup-antlr), the link will show you how to setup antlr.
   The article has a whole lot of detail about antlr I encourage you to read it, you might also like [my blog post](https://medium.com/@nikunjyadav/generic-rules-engine-in-golang-using-antlr-d30a0d0bb565) about this repo.
2. After taking a look at the antlr tutorial, you can extend the [JsonQuery.g4 file](https://github.com/nikunjy/rules/blob/master/parser/JsonQuery.g4).
3. Compile the parser `antlr4 -Dlanguage=Go -visitor -no-listener JsonQuery.g4 -o ./` (Note: `-o` is the output directory, make sure all the stuff it generates is in the `parser` directory of the root repo folder)

[ci-img]: https://api.travis-ci.org/nikunjy/rules.svg?branch=master
[ci]: https://travis-ci.org/nikunjy/rules
