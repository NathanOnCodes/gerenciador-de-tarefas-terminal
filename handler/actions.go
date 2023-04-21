package handler

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"gerenciador-tarefas-terminal/models"
	"io"
	"os"
	"strings"
)

const (
	todoFile = "tarefas.json"
)


func AllActionsTodo(){
	add := flag.Bool("add", false, "adicionar uma nova tarefa")
	complete := flag.Int("complete", 0 , "marcar uma tarefa concluída")
	del := flag.Int("del", 0 , "deletar uma tarefa")
	list := flag.Bool("list", false, "listar tarefas")

	flag.Parse()

	todos := &models.Todos{}


	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch{
		case *add:
			task, err := getInput(os.Stdin, flag.Args()...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.AddTodo(task)
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			
		case *complete > 0:
			err := todos.CompleteTodo(*complete)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *del > 0:
			err := todos.DeleteTodo(*del)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *list:
			todos.Print()
		default:
			fmt.Fprintln(os.Stdout, "comando inválido")
			os.Exit(0)
	}
}

func getInput (r io.Reader, args ...string) (string, error) {

	if len(args) >  0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("tarefas vazias não são permetidas")
	}
	return text, nil
}