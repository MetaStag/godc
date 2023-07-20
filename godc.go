package main

import (
	"bufio"   // to get input
	"fmt"     // to print
	"math"    // exponentiation
	"os"      // stdin manipulation
	"os/exec" // stdout manipulation
	"strconv" // check if input is a number
	"strings" // string manipulation
)

// Compute modulo of float
func modulo(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	for x > y {
		x -= y
	}
	return x
}

// Search element in array (return index)
func search(array []string, element string) int {
	for x, y := range array {
		if y == element {
			return x
		}
	}
	return -1
}

// Check if element exists in array (true/false)
func stringInSlice(slice []string, str string) bool {
	for _, b := range slice {
		if b == str {
			return true
		}
	}
	return false
}

func main() {
	stack := make([]any, 0)                                      // main stack
	registers := make(map[any]any)                               // registers
	conditionals := [...]string{"=", ">", "<", "!>", "!<", "!="} // conditional operators
	reader := bufio.NewReader(os.Stdin)                          // input

	fmt.Println("godc")
	fmt.Println("**************************")
	fmt.Println("Press ? for commands")

	// MAIN LOOP
	for 1 == 1 {
		fmt.Print("> ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSuffix(choice, "\n")
		if choice == "" {
			continue
		}

		input := strings.Split(choice, " ")

		for i := 0; i < len(input); i++ {
			// Commands
			if input[i] == "?" {
				fmt.Println("BASICS")
				fmt.Println("<num> - Add number to stack")
				fmt.Println("+-*/ - Basic Arithmetic")
				fmt.Println(`% - Modulo`)
				fmt.Println("~ - Quotient+remainder")
				fmt.Println("^ - Exponentiation")
				fmt.Println("| - Modular Exponentiation")
				fmt.Println("v - Square Root")
				fmt.Println("[ commands ] - Add string (macro) to stack")
				fmt.Println("STACK CONTROL")
				fmt.Println("p - Print top of stack")
				fmt.Println("n - Print+Pop top of stack")
				fmt.Println("f - Print full stack (topmost element at bottom)")
				fmt.Println("c - clear stack")
				fmt.Println("d - duplicate top element")
				fmt.Println("r - Swap top 2 elements")
				fmt.Println("R - Cyclically rotate no. of items equal to top of stack")
				fmt.Println("z - Push length of stack")
				fmt.Println("REGISTERS/MACROS")
				fmt.Println("s(x) - Push top of stack into register x (can be any letter)")
				fmt.Println("l(x) - Copy register x (any letter) onto stack")
				fmt.Println("x - Execute macro at top of stack")
				fmt.Println("Q - Break out of macro loop, does not quit program")
				fmt.Println("=m - If top 2 elements are equal, execute macro m. Note that instead of equals, >, <, !>, !<, != are also supported (reverse order of operands)")
				fmt.Println("HELP")
				fmt.Println("#(anything) - Comments")
				fmt.Println("clear - Clear screen")
				fmt.Println("q - Quit Program")

				// Add numbers
			} else if num, err := strconv.ParseFloat(input[i], 64); err == nil {
				stack = append(stack, num)

				// Add strings
			} else if input[i] == "[" {
				secondIndex := search(input, "]")
				if secondIndex == -1 {
					fmt.Println("godc: invalid syntax")
					break
				}
				temp := input[i+1 : secondIndex]
				stack = append(stack, temp)
				break

				// Printing
			} else if input[i] == "p" {
				if len(stack) > 0 {
					fmt.Println(stack[len(stack)-1])
				}
			} else if input[i] == "n" {
				if len(stack) > 0 {
					fmt.Println(stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
			} else if input[i] == "f" {
				for _, element := range stack {
					fmt.Println(element)
				}

				// Stack control
			} else if input[i] == "c" {
				stack = nil
			} else if input[i] == "d" {
				if len(stack) > 0 {
					stack = append(stack, stack[len(stack)-1])
				}
			} else if input[i] == "r" {
				if len(stack) > 1 {
					top := stack[len(stack)-1]
					stack[len(stack)-1] = stack[len(stack)-2]
					stack[len(stack)-2] = top
				}
			} else if input[i] == "R" {
				top := int(stack[len(stack)-1].(float64))
				stack = stack[:len(stack)-1]

				if top > 0 {
					if top > len(stack) {
						top = len(stack)
					}
					a := stack[len(stack)-top]
					for x := top; x > 1; x-- {
						stack[len(stack)-x] = stack[len(stack)-x+1]
					}
					stack[len(stack)-1] = a
				} else {
					if -top > len(stack) {
						top = len(stack)
					}
					a := stack[len(stack)-1]
					for x := 1; x < top; x++ {
						stack[len(stack)-x] = stack[len(stack)-x-1]
					}
					stack[len(stack)-top] = a
				}
			} else if input[i] == "z" {
				stack = append(stack, float64(len(stack)))

				// Registers
			} else if strings.HasPrefix(input[i], "s") {
				if len(input[i]) < 2 {
					fmt.Println("godc: invalid command")
					break
				}
				name := input[i][1]
				if len(stack) > 0 {
					temp := stack[len(stack)-1]
					switch v := temp.(type) {
					case float64:
						registers[name] = v
					case []string:
						registers[name] = v
					default:
						return
					}
					stack = stack[:len(stack)-1]
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if strings.HasPrefix(input[i], "l") {
				if len(input[i]) < 2 {
					fmt.Println("godc: invalid command")
					break
				}
				name := input[i][1]
				if registers[name] != nil {
					stack = append(stack, registers[name])
				} else {
					fmt.Println("godc: register empty")
				}

				// Execute Macros (strings)
			} else if input[i] == "x" {
				if len(stack) == 0 {
					fmt.Println("godc: stack empty")
				}
				switch v := stack[len(stack)-1].(type) {
				case []string:
					input = v
					stack = stack[:len(stack)-1]
					i = -1
					continue
				default:
					fmt.Println("godc: invalid macro")
				}

			} else if input[i] == "Q" {
				break

				// Conditionals
			} else if command := search(conditionals[:], input[i][:len(input[i])-1]); command != -1 {
				var flag bool
				if len(stack) < 2 {
					fmt.Println("godc: stack empty")
					break
				}
				name := input[i][len(input[i])-1]
				switch conditionals[command] {
				case "=":
					if stack[len(stack)-1] == stack[len(stack)-2] {
						flag = true
					}
				case ">":
					if stack[len(stack)-1].(float64) > stack[len(stack)-2].(float64) {
						flag = true
					}
				case "<":
					if stack[len(stack)-1].(float64) < stack[len(stack)-2].(float64) {
						flag = true
					}
				case "!>":
					if stack[len(stack)-1].(float64) < stack[len(stack)-2].(float64) {
						flag = true
					}
				case "!<":
					if stack[len(stack)-1].(float64) > stack[len(stack)-2].(float64) {
						flag = true
					}
				case "!=":
					if stack[len(stack)-1] != stack[len(stack)-2] {
						flag = true
					}
				default:
					return
				}
				if flag {
					stack = stack[:len(stack)-2]
					if registers[name] == nil {
						fmt.Println("godc: register empty")
					} else {
						switch v := registers[name].(type) {
						case []string:
							input = v
							i = -1
							continue
						default:
							fmt.Println("godc: invalid macro")
						}
					}
				}
				stack = stack[:len(stack)-2]

				// Arithmetic
			} else if input[i] == "+" {
				if len(stack) > 1 {
					result := stack[len(stack)-2].(float64) + stack[len(stack)-1].(float64)
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "-" {
				if len(stack) > 1 {
					result := stack[len(stack)-2].(float64) - stack[len(stack)-1].(float64)
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "*" {
				if len(stack) > 1 {
					result := stack[len(stack)-2].(float64) * stack[len(stack)-1].(float64)
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "/" {
				if len(stack) > 1 {
					result := stack[len(stack)-2].(float64) / stack[len(stack)-1].(float64)
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "%" {
				if len(stack) > 1 {
					result := modulo(stack[len(stack)-2].(float64), stack[len(stack)-1].(float64))
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "~" {
				if len(stack) > 1 {
					quotient := stack[len(stack)-2].(float64) / stack[len(stack)-1].(float64)
					remainder := modulo(stack[len(stack)-2].(float64), stack[len(stack)-1].(float64))
					stack[len(stack)-1] = remainder
					stack[len(stack)-2] = quotient
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "^" {
				if len(stack) > 1 {
					result := math.Pow(stack[len(stack)-2].(float64), stack[len(stack)-1].(float64))
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "|" {
				if len(stack) > 2 {
					result := math.Pow(stack[len(stack)-3].(float64), stack[len(stack)-2].(float64))
					result = modulo(result, stack[len(stack)-1].(float64))
					stack = stack[:len(stack)-2]
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}
			} else if input[i] == "v" {
				if len(stack) > 0 {
					result := math.Sqrt(stack[len(stack)-1].(float64))
					stack[len(stack)-1] = result
				} else {
					fmt.Println("godc: stack empty")
				}

				// Helper commmands
			} else if input[i] == "#" {
				break
			} else if input[i] == "clear" {
				c := exec.Command("clear")
				c.Stdout = os.Stdout
				c.Run()
			} else if input[i] == "q" {
				return
			} else {
				fmt.Println("godc: invalid command")
			}
		}
	}
	return
}
