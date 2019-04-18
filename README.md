# Golang Rules Engine  [![Build Status][ci-img]][ci]
Rules engine written in golang with the help of antlr.

This package will be very helpful in situations where you have a generic rule and want to verify if your values (specified using `map[string]interface{}`) satisfy the rule. 


Here are some examples:

```
  type obj map[string]interface{}
  
  parser.Evaluate("x eq 1", obj{"x": 1})
  parser.Evaluate("x lt 1", obj{"x": 1})
  parser.Evaluate("x gt 1", obj{"x": 1})
  
  parser.Evaluate("x.a eq 1 and x.b.c le 2", obj{
    "x": obj{
       "a": 1,
       "b": obj{
          "c": 2,
       },
    },
  })
  

  parser.Evaluate("y eq 4 and (x gt 1)", obj{"x": 1})

  parser.Evaluate("y eq 4 and (x IN [1 2 3])", obj{"x": 1})

  parser.Evaluate("y eq 4 and (x eq 1.2.3)", obj{"x": "1.2.3"})
  
```

## Operations
All the operations can be written capitalized or lowercase (ex: `eq` or `EQ` can be used)

Logical Operations supported are `AND OR`

Compare Expression and their definitions:
```
eq: equals to 
ne: not equals to
lt: less than 
gt: greater than
le: less than equal to
ge: greater than equal to 
co: contains 
sw: starts with 
ew: ends with
in: in a list
pr: present
not: not of a logical expression
```

[ci-img]: https://api.travis-ci.org/nikunjy/rules.svg?branch=master
[ci]: https://travis-ci.org/nikunjy/rules
