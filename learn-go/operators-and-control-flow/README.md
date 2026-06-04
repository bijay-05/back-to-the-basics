# Operators and Control Flow

Operators are symbols used to represent mathematical and logical computations. In the expression "A + B", the plus sign ( `+` ) is the operator, while A and B are the operands.

## Operator Categories in Go

- **Comparison Operators** : Used to compare two values, returning a boolean result. Equality operators ( `==`, `!=` ) require operands to be of comparable types

> [!Caution]
> Not all types are comparable using `==` and `!=`. Types such as slices, maps, and functions cannot be compared with equality operators — attempting to do so results in a compile-time error. Structs and arrays are comparable only if all their fields/elements are comparable.

- **Arithmetic Operators** : Perform mathematical calculations such as addition, division, etc.

| Operation      | Symbol | Remarks                                                                                       |
| -------------- | ------ | --------------------------------------------------------------------------------------------- |
| Addition       | +      | Can be used for string concatenations                                                         |
| Subtraction    | -      | Cannot be used with string data type                                                          |
| Multiplication | \*     |
| Division       | /      | with integers, the operator performs integer division, which may truncate any fractional part |
| Modulus        | %      |
| Increment      | ++     | Always use as post-fix operator                                                               |
| Decrement      | --     |

- **Assignment Operators** : Assign values to variables, often combining arithmetic operations with assignment.

| Operation                 | Symbol |
| ------------------------- | ------ |
| Addition Assignment       | +=     |
| Subtraction Assignment    | -=     |
| Multiplication Assignment | \*=    |
| Division Assignment       | /=     |
| Modulus Assignment        | %=     |

- **Logical Operators**: Evaluate boolean expressions and combine conditions

| Operation | Symbol  | Remarks                                                        |
| --------- | ------- | -------------------------------------------------------------- |
| AND       | `&&`    | Returns true only when both conditions are true                |
| OR        | `\|\| ` | Returns true when either one of the conditions is true         |
| NOT       | `!`     | Unary operator that inverts the boolean value of an expression |

- **Bitwise Operators** : Operate on binary representations of numbers, useful for low-level programming.

Bitwise operators work directly on the binary representation of integer values, manipulating one bit at a time. In Go, these operators are essential for low-level programming, bit masks, flags, and efficient arithmetic (shifts are equivalent to multiplying/dividing by powers of two for non-negative integers).

Left shift (<<) Left shifting moves bits to the left, inserting zeros on the right. For non-negative integers, shifting left by n is equivalent to multiplying by 2^n (subject to type limits and overflow).

Right shift (>>) Right shifting moves bits to the right, discarding bits shifted out on the right. For unsigned or non-negative signed integers, right shifting by n is equivalent to integer division by 2^n, rounding down.

| Operation                                                       | Symbol | Example           |
| --------------------------------------------------------------- | ------ | ----------------- |
| Bitwise AND - 1 if both bits are 1                              | `&`    | `12 & 25 // 8`    |
| Bitwise OR - 1 if at least one bit is 1                         | `\|`   | `12 & 25 // 29`   |
| Bitwise XOR - 1 if bits differ                                  | `^`    | `12 ^ 25 // 21`   |
| Left shift - shift bits left (multiply by 2^n for non-negative) | `<<`   | `212 << 1 // 424` |
| Right shift - shift bits right (divide by 2^n for non-negative) | `>>`   | `212 >> 2 // 53`  |

```markdown
bitwise AND (&)

12 = 00001100 (In binary)
25 = 00011001 (In binary)

0 0 0 0 1 1 0 0
& 0 0 0 1 1 0 0 1
0 0 0 0 1 0 0 0 = 8 (In decimal)
```

> [!Important]
> In Go, `^` is also the unary bitwise NOT (complement) when used with a single operand ( for example `^x` flips all bits of x ).

> [!Caution]
> Be careful with signed integers and right shifts: the fill behavior for the leftmost bits (sign extension vs. zero fill) depends on the type and implementation. For predictable behavior, use unsigned integers ( `uint`, `uint32`, `uint64` ) when performing bit-level shifts.

## Conditionals ( if, if-else and else-if )

> [!Important]
> Parantheses around condition is optional. And Avoid formatting the else block on a new line after the closing brace:

## Switch Statements

> [!Important]
> Switch statements provide a clean and efficient way to handle multiple conditions, especially when comparing a single variable to several possible values.

**fallthrough**
An interesting feature of Go’s switch statement is the use of the fallthrough keyword. This keyword forces the execution to continue into the next case block, regardless of whether that subsequent case matches the switch expression.

## Loops with `for`

```go
for initialization; condition; post {
  // statements
}
```

- **Initialization Statement (optional)**: Executed once before the loop starts. It typically declares and assigns a starting value.
- **Condition Statement**: A Boolean expression evaluated before each iteration. If it evaluates to true, the loop body executes; if false, the loop exits.
- **Post Statement (Optional)**: Executed at the end of each iteration. This is often used to update the counter variable.

### Using `break` and `continue`

The `break` statement immediately terminates the loop when a particular condition is met.

The `continue` statement allows your loop to skip over specific iterations without exiting the entire loop, providing greater control over loop execution.
