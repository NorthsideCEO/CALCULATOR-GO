package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)
	
	var firstNum, secondNum int
	var sign string
	fmt.Println("Введите HELP для ознакомления со списком операций или введите знак для двух целых чисел")
	for {
		for {
			sign, _ = scanner.ReadString('\n')
			sign = strings.TrimSpace(sign)
			_ = strings.ToLower(sign)

			switch sign {
			case "HELP":
				fmt.Println("\"+\" - операция сложения \n \"-\" - операция вычитания \n \"*\" - операция умножения \n \"/\"- операция деления \n \"**\" - " +
					"операция возведения в степень \n \"sqrt\" - операция возведения в квадрат \n \"%\" - остаток от деления ")
				continue
			}
			break
		}
		switch sign {
		case "sqrt":
			for {
				_, err := fmt.Scan(&firstNum)
				if err != nil {
					fmt.Println("Введенным значением должно быть число,повторите попытку")
					continue
				}
				for {
					_, err := fmt.Scan(&secondNum)
					if err != nil {
						fmt.Println("Введенным значением должно быть число,повторите попытку")
						continue
					}
					fmt.Println(float64(firstNum) * float64(firstNum))
					fmt.Println("Вычисление закончено,программа закроется через 10 секунд...")
					time.Sleep(10 * time.Second)
					break
				}
				break
			}
		case "+":
			for {
				_, err := fmt.Scan(&firstNum)
				if err != nil {
					fmt.Println("Введенным значением должно быть число,повторите попытку")
					continue
				}
				for {
					_, err := fmt.Scan(&secondNum)
					if err != nil {
						fmt.Println("Введенным значением должно быть число,повторите попытку")
						continue
					}
					fmt.Println(float64(firstNum) + float64(secondNum))
					fmt.Println("Вычисление закончено,программа закроется через 10 секунд...")
					time.Sleep(10 * time.Second)
					break
				}
				break
			}
		case "-":
			for {
				_, err := fmt.Scan(&firstNum)
				if err != nil {
					fmt.Println("Введенным значением должно быть число,повторите попытку")
					continue
				}
				for {
					_, err := fmt.Scan(&secondNum)
					if err != nil {
						fmt.Println("Введенным значением должно быть число,повторите попытку")
						continue
					}
					fmt.Println(float64(firstNum) - float64(secondNum))
					fmt.Println("Вычисление закончено,программа закроется через 10 секунд...")
					time.Sleep(10 * time.Second)
					break
				}
				break
			}
		case "*":
			for {
				_, err := fmt.Scan(&firstNum)
				if err != nil {
					fmt.Println("Введенным значением должно быть число,повторите попытку")
					continue
				}
				for {
					_, err := fmt.Scan(&secondNum)
					if err != nil {
						fmt.Println("Введенным значением должно быть число,повторите попытку")
						continue
					}
					fmt.Println(float64(firstNum) * float64(secondNum))
					fmt.Println("Вычисление закончено,программа закроется через 10 секунд...")
					time.Sleep(10 * time.Second)
					break
				}
				break
			}
		case "/":
			for {
				_, err := fmt.Scan(&firstNum)
				if err != nil {
					fmt.Println("Введенным значением должно быть число,повторите попытку")
					continue
				}
				for {
					_, err := fmt.Scan(&secondNum)
					if err != nil {
						fmt.Println("Введенным значением должно быть число,повторите попытку")
						continue
					}
					fmt.Println(float64(firstNum) / float64(secondNum))
					fmt.Println("Вычисление закончено,программа закроется через 10 секунд...")
					time.Sleep(10 * time.Second)
					break
				}
				break
			}

		case "**":
			for {
				_, err := fmt.Scan(&firstNum)
				if err != nil {
					fmt.Println("Введенным значением должно быть число,повторите попытку")
					continue
				}
				for {
					_, err := fmt.Scan(&secondNum)
					if err != nil {
						fmt.Println("Введенным значением должно быть число,повторите попытку")
						continue
					}
					fmt.Println(math.Pow(float64(firstNum), float64(secondNum)))
					fmt.Println("Вычисление закончено,программа закроется через 10 секунд...")
					time.Sleep(10 * time.Second)
					break
				}
				break
			}
		case "%":
			for {
				_, err := fmt.Scan(&firstNum)
				if err != nil {
					fmt.Println("Введенным значением должно быть число,повторите попытку")
					continue
				}
				for {
					_, err := fmt.Scan(&secondNum)
					if err != nil {
						fmt.Println("Введенным значением должно быть число,повторите попытку")
						continue
					}
					fmt.Println(firstNum % secondNum)
					fmt.Println("Вычисление закончено,программа закроется через 10 секунд...")
					break
				}
				break
			}
		default:
			fmt.Println("Выбранной команды не существует,введите HELP для ознакомления со списком операций и повторите попытку")
			continue
		}
		break
	}
}
