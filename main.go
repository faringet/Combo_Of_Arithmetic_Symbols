package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strings"
)

func main() {
	// промежуток между цифрами от 9 до 0
	digits := []string{"9", "8", "7", "6", "5", "4", "3", "2", "1", "0"}

	// возможные операторы
	operators := []string{"", "+", "-"}

	// вложенные циклы, которые используются для генерации всех возможных комбинаций операторов
	for _, op1 := range operators {
		for _, op2 := range operators {
			for _, op3 := range operators {
				for _, op4 := range operators {
					for _, op5 := range operators {
						for _, op6 := range operators {
							for _, op7 := range operators {
								for _, op8 := range operators {

									// собираем комбинацию из цифр и операторов
									expression := strings.Join([]string{digits[0], op1, digits[1], op2, digits[2], op3, digits[3], op4, digits[4], op5, digits[5], op6, digits[6], op7, digits[7], op8, digits[8], digits[9]}, "")
									// присваиваем переменной финальное значение в формате float64, полученное из func eval
									val, _ := eval(expression)
									// при удовлетворении условия ("так, чтобы в результате получилось 200") - выводим комбинацию из цифр и операторов (expression)
									if val == 200 {
										fmt.Println(expression)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// передаем комбинацию в функцию, возвращаем значение в float64
func eval(expression string) (float64, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, err
	}
	return result.(float64), nil
}
