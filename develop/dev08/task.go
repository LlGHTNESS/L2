package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"main.go/Downloads/L2/develop/dev08/internal/shell"
	"main.go/Downloads/L2/develop/dev08/internal/shell/shellcommands"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	myShell := shell.NewShell()

	// Регистрируем команды
	myShell.RegisterComand("pwd", shellcommands.PWDCommand{})
	myShell.RegisterComand("cd", shellcommands.CDCommand{})
	myShell.RegisterComand("echo", shellcommands.EchoCommand{})
	myShell.RegisterComand("ps", shellcommands.PSCommand{})
	myShell.RegisterComand("kill", shellcommands.KillCommand{})
	myShell.RegisterComand("exec", shellcommands.ExecCommand{})
	myShell.RegisterComand("exit", shellcommands.ExitCommand{})

	// Обрабатываем ввод
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, "|") // Делим на команды
		for _, rawCommand := range commands {
			command := strings.Fields(rawCommand)       // Делим на аргументы
			myShell.ExecComand(command[0], command[1:]) // Вызывае команду
			fmt.Print(">")
		}
	}
}
