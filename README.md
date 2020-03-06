# Lorikeet

Lorikeet is a simple scripting language written in GoLang. \
The language was initially created as a learning experience and is now
used as base to experiment with new programing concepts.

# Syntax

## Basics

### Case Sensitivity

Lorikeet is case sensitive. \
Examples:

```
let apple = 1;
let Apple = 2;
let aPpLe = 3;
say(apple); // 1
say(Apple); // 2
say(aPpLe); // 3
```

```
let applePie = 10;
say(applepie); // compiler error: undefined variable applepie; line=2
```

### Whitespace and Semicolons
Spaces and newlines used outside of strings in Lorikeet are ignored.
This has an impact on semantics because Lorikeet uses automatic
semicolon insertion (ASI). \
Example:

```
let a = 10 let b = 3 say(a+b)
// Treated as:
// let a = 10;
// let b = 3;
// say(a + b);
```

Although ASI can be useful, you should keep in consideration that
sometimes it can lead to unintended effects:

```
let a = 5
let b = 1 + c
((2 + 3) / 2) |> say()
// Treated as:
// let a = 5;
// let b = 1 + b((c + 2) / 2) |> say(); // vm error: calling non-closure and non-builtin
```

This will lead to b to being treated as a function call although its value is not
callable.

### Comments

Lorikeet only supports single-line comments.\
Example:

```
// this is a comment
```

## Variables

Currently variables in Lorikeet have no type attached, any value can be
stored in any variable although this is subject to change. All variables
are scoped to block level. \
A variable in Lorikeet can be declared with `let` for an immutable variable
or `let mut` for a mutable variable.
```
let apple = 6;
apple = 10; // compiler error: can't mutate constant symbol apple; line=1

let mut pie = 6;
pie = 10;
say(pie); // 10
```

## Types

Lorikeet has the following types:

| Type     | Syntax                                | GoLang Type |
|----------|---------------------------------------|-------------|
| INTEGER  | `0 96 1234 -10`                       | int64       |
| FLOAT    | `1.0 10.03 -22.2`                     | float64     |
| STRING   | `"" "\\" quotes \\" \n new line"`     | string      |
| BOOLEAN  | `true false`                          | bool        |
| ARRAY    | `[] [1, 2] ["test", 10, true]`        | array       |
| HASH     | `{} { "key": "val" } { 96: "apple" }` | map         |
| FUNCTION | `fn() {}`                             | N/A         |
| NULL     | `return; [undefined index]`           | N/A         |

## Operators

Operators can only used with the left and right side of the same type if left side is applicable. \
The following is valid combinations: `INTEGER : INTEGER`, `FLOAT : FLOAT`, `STRING : STRING` \
The following is not valid: `INTEGER : FLOAT`

### Arithmetic

| Operator | Arithmetic     |
|----------|----------------|
| +        | addition       |
| -        | subtraction    |
| *        | multiplication |
| /        | division       |

Example:
```
    1 + 2;     // 3
    2 - 1;     // 1
    2 * 5;     // 10
    5 / 2;     // 2
    5.0 / 2.0; // 2.5
    "Hello, " + "World!"; // Hello, World! 
```


