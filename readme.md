# godc
A dc implementation written in golang

---

dc is a very old **reverse-polish notation** calculator, written even before the C language was made.

In reverse-polish notation, instead of writing `2 + 2`, you write `2 2 +`. The operator comes later, while awkward at first glance, it can help avoid confusion at times

All commands **must be space seperated**, the only exception is conditional commands, like `=m`

---

### Features
- Decimal numbers supported
- Basic Arithmetic + modulo, square root, modular exponentiation
- A Stack to store numbers
- Registers, strings, and macros for extra command
- Conditional operators
- Loops can be implemented using registers, macros and conditionals
- For a full list of supported commands, do `?` inside godc
--- 

### Limitations
- Unlike dc where the register can store multiple values, in godc a register can only store one value at a time
- In nested macros, you cannot just come out of 1 macro, you exit out of all of them (`Q`)
- Precision control not implemented
- Script (`-e`) and file mode (`-f`) not implemented, only interactive mode supported for now
- certain dc commands like `P, S, L, a, Z, X, :, ;` not implemented
