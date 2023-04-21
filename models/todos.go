package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
	"github.com/alexeyco/simpletable"
)

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) AddTodo(task string){
	todo := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) CompleteTodo(index int) error {
	ls := *t

	if index <= 0 || index > len(ls){
		return errors.New("sintaxe invalida")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) DeleteTodo(index int) error {
	ls := *t

	if index <= 0 || index > len(ls){
		return errors.New("sintaxe invalida")
	}

	*t = append(ls[:index-1], ls[index:]...)
	
	return nil
}

func (t *Todos) Load(filename string) error {
	file, fileErr := ioutil.ReadFile(filename)
	if fileErr != nil {
		if errors.Is(fileErr, os.ErrNotExist){
			return nil
		}
		log.Println("erro: ", fileErr)
	}

	if len(file) == 0 {
		return fileErr
	}

	fileErr = json.Unmarshal(file,t)
	if fileErr != nil {
		return fileErr
	}
	return nil
}

func (t *Todos) Store(filename string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}



func (t *Todos) Print(){
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Tarefa"},
			{Align: simpletable.AlignCenter, Text: "Concluida?"},
			{Align: simpletable.AlignRight, Text: "Criado em"},
			{Align: simpletable.AlignRight, Text: "Completado em"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("nao")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("sim")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("Tarefas pendentes: %d", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}