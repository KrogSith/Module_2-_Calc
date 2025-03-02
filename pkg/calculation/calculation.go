package calculation

import (
	"calculator/pkg/stack"
	"fmt"
	"strconv"
	"unicode"
)

// Интерфейс с методами стека
type PushPopper[T comparable] interface {
	Push(n T)
	Pop() T
	Len() int
	GetArray() []T
}

// InfixExprToPostfixString() преобразует выражение инфиксной записи в выражение постфиксной записи и
// заполняет стек математических операций
func InfixExprToPostfixString(infixExpr string, operationsStack PushPopper[string], postfixExpr PushPopper[string]) (PushPopper[string], error) {
	for i := 0; i < len(infixExpr); i++ {
		fmt.Println("===========================")
		fmt.Println("operationsStack:", operationsStack)
		fmt.Println("postfixExpr    :", postfixExpr)
		s := string(infixExpr[i])
		fmt.Println(s)
		// if s == " " {
		// 	continue
		// }
		s_rune := []rune(s)[0]
		if unicode.IsDigit(s_rune) {
			number := ""
			for j := i; j < len(infixExpr); j++ {
				s = string(infixExpr[j])
				s_rune := []rune(s)[0]
				if unicode.IsDigit(s_rune) != true {
					postfixExpr.Push(number)
					fmt.Printf("1 postfixExpr.Push(%v)\n", number)
					i = j - 1
					break
				}
				if j == len(infixExpr)-1 {
					number += s
					postfixExpr.Push(number)
					fmt.Printf("2 postfixExpr.Push(%v)\n", number)
					i = j
					break
				}
				number += s
			}
			continue
		}

		if s == "(" {
			operationsStack.Push(s)
			fmt.Printf("2.1 operationsStack.Push(%v)\n", s)
			continue
		}
		if s == "-" || s == "+" {
			if operationsStack.Len() == 0 {
				operationsStack.Push(s)
				fmt.Printf("3 operationsStack.Push(%v)\n", s)
				continue
			}
			element := operationsStack.Pop()
			fmt.Println("3.1 operationsStack.Pop() = ", element)

			if element == "*" || element == "/" {
				postfixExpr.Push(element)
				fmt.Printf("3.2 postfixExpr.Push(%v)\n", element)

				if operationsStack.Len() != 0 {
					for i := operationsStack.Len() - 1; i >= 0; i-- {
						element := operationsStack.Pop()
						fmt.Println("3.2 operationsStack.Pop() = ", element)

						if element == "(" || element == "+" || element == "-" {
							postfixExpr.Push(element)
							fmt.Printf("4 postfixExpr.Push(%v)\n", element)
							break
						}
						if element == "*" || element == "/" {
							postfixExpr.Push(element)
							fmt.Printf("4.1 postfixExpr.Push(%v)\n", element)
							continue
						}
					}
				}
			} else if element == "+" || element == "-" {
				//operationsStack.Push(element)
				postfixExpr.Push(element)
				fmt.Printf("5 postfixExpr.Push(%v)\n", element)
			}
			operationsStack.Push(s)
			fmt.Printf("6 operationsStack.Push(%v)\n", s)
			continue
		}
		if s == "*" || s == "/" {
			if operationsStack.Len() == 0 {
				operationsStack.Push(s)
				fmt.Printf("6.1 postfixExpr.Push(%v)\n", s)
				continue
			}
			element := operationsStack.Pop()
			fmt.Println("6.2 operationsStack.Pop() = ", element)

			if element == "+" || element == "-" {
				operationsStack.Push(element)
				operationsStack.Push(s)
				fmt.Printf("7 operationsStack.Push(%v)\n", element)
				fmt.Printf("8 operationsStack.Push(%v)\n", s)
				continue
			}
			operationsStack.Push(element)
			operationsStack.Push(s)
			fmt.Printf("9 operationsStack.Push(%v)\n", element)
			fmt.Printf("10 operationsStack.Push(%v)\n", s)

			continue
		}
		if s == ")" {
			for operation := 0; operation < operationsStack.Len(); operation++ {
				element := operationsStack.Pop()
				if element != "(" {
					postfixExpr.Push(element)
				} else {
					break
				}
			}
			// element := operationsStack.Pop()
			// fmt.Println("10.1 operationsStack.Pop() = ", element)
			// if element != "+" && element != "-" && element != "/" && element != "*" {
			// 	return stack.NewStack[string](), fmt.Errorf("500")
			// }
			// postfixExpr.Push(element)
			// fmt.Printf("10.2 postfixExpr.Push(%v)\n", element)
			// operationsStack.Pop()
			// fmt.Println("10.3 operationsStack.Pop()")
			// continue
		}
	}

	for i := operationsStack.Len() - 1; i >= 0; i-- {
		postfixExpr.Push(operationsStack.GetArray()[i])
		fmt.Printf("11 postfixExpr.Push(%v)\n", operationsStack.GetArray()[i])
		//fmt.Println("777777777777777777777")
	}
	//fmt.Println(operationsStack)
	return postfixExpr, nil
}

// Проводит математические операции с полученным на вход выражением в постфиксной записи,
// попутно заполняя и убирая элементы из стека с числами
func StackCalc(postfixExpr PushPopper[string], numbersStack PushPopper[float64]) (float64, error) {
	fmt.Println(postfixExpr)
	for i := 0; i < postfixExpr.Len(); i++ {
		element := postfixExpr.GetArray()[i]
		// fmt.Println(element)
		n, err := strconv.Atoi(element)
		// fmt.Println(numbersStack)
		if err == nil {
			numbersStack.Push(float64(n))
			continue
		}

		if numbersStack.Len() < 2 {
			return 0, fmt.Errorf("Invalid expression")
		}

		if element == "+" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 + n1
			numbersStack.Push(oper)
			//fmt.Println(oper)
		}
		if element == "-" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 - n1
			numbersStack.Push(oper)
			//fmt.Println(oper)
		}
		if element == "/" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 / n1
			numbersStack.Push(oper)
		}
		if element == "*" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 * n1
			numbersStack.Push(oper)
		}
	}
	if numbersStack.Len() != 1 {
		return 0, fmt.Errorf("Invalid expression")
	} else {
		return numbersStack.Pop(), nil
	}
}

// IsBracketsRight() Проверяет правильность расстановки скобок
func IsBracketsRight(str string) bool {
	num := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			num += 1
		} else if string(str[i]) == ")" {
			num -= 1
		}
		if num < 0 {
			return false
		}
	}
	if num == 0 {
		return true
	} else {
		return false
	}
}

// Calc() вызывает проверочные и вычислительные функции
func Calc(infixExpr string) (float64, error) {
	for i := 0; i < len(infixExpr); i++ {
		if infixExpr[i] == '+' || infixExpr[i] == '-' || infixExpr[i] == '/' || infixExpr[i] == '*' ||
			infixExpr[i] == '(' || infixExpr[i] == ')' {
			continue
		} else {
			_, err := strconv.Atoi(string(infixExpr[i]))
			if err != nil {
				return 0, fmt.Errorf("Invalid expression")
			}
		}
	}

	if !IsBracketsRight(infixExpr) {
		return 0, fmt.Errorf("Invalid expression")
	}

	str, err := InfixExprToPostfixString(infixExpr, stack.NewStack[string](), stack.NewStack[string]())
	if err != nil {
		return 0, err
	}

	return StackCalc(str, stack.NewStack[float64]())
}
