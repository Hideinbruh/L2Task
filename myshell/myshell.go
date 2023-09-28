package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Необходимо реализовать свой собственный UNIX-шелл-утилиту с
//поддержкой ряда простейших команд:
//- cd <args> - смена директории (в качестве аргумента могут
//быть то-то и то)
//- pwd - показать путь до текущего каталога
//- echo <args> - вывод аргумента в STDOUT
//- kill <args> - "убить" процесс, переданный в качестве
//аргумента (пример: такой-то пример)
//- ps - выводит общую информацию по запущенным процессам в
//формате *такой-то формат*
//Так же требуется поддерживать функционал fork/exec-команд
//Дополнительно необходимо поддерживать конвейер на пайпах
//(linux pipes, пример cmd1 | cmd2 | .... | cmdN).
//*Шелл — это обычная консольная программа, которая будучи
//запущенной, в интерактивном сеансе выводит некое приглашение
//в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись
//ввода, обрабатывает команду согласно своей логике
//и при необходимости выводит результат на экран. Интерактивный
//сеанс поддерживается до тех пор, пока не будет введена
//команда выхода (например \quit)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
		input = strings.TrimSuffix(input, "\n")
		parts := strings.Split(input, " ")
		fmt.Println(parts)
		switch parts[0] {
		case "cd":
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "cd: пропущен аргумент команды")
				continue
			}
			errCd := os.Chdir(parts[1])
			if errCd != nil {
				fmt.Fprintln(os.Stderr, errCd)
			}

		case "pwd":
			pwd, errPwd := os.Getwd()
			if errPwd != nil {
				fmt.Fprintln(os.Stderr, errPwd)
			}
			fmt.Println(pwd)

		case "kill":
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "kill: пропущен аргумент команды")
				continue
			}
			cmd := exec.Command("kill", "-9", parts[1])
			errCmd := cmd.Run()
			if errCmd != nil {
				fmt.Fprintln(os.Stderr, errCmd)
			}
		case "echo":
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "echo: пропущен аргумент команды")
			}
			fmt.Println(parts[2])

		case "ps":
			cmd := exec.Command("ps")
			if err := cmd.Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		default:
			fmt.Fprintln(os.Stderr, "Неизвестная команда")
		}
	}
}
