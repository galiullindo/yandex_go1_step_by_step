package main

import (
	"errors"
	"fmt"
	"slices"
)

var ErrUnacceptablePosition = errors.New("unacceptable position")

type Position struct {
	level int
	name  string
}

var acceptablePositions = []Position{
	{5, "директор"},
	{4, "зам. директора"},
	{3, "начальник цеха"},
	{2, "мастер"},
	{1, "рабочий"},
}

func NewPosition(name string) (Position, error) {
	for _, position := range acceptablePositions {
		if position.name == name {
			return position, nil
		}
	}
	return Position{}, ErrUnacceptablePosition
}

type Worker struct {
	name           string
	position       Position
	salary         uint
	experienceYear uint
}

func NewWorker(name string, position string, salary uint, experienceYear uint) (*Worker, error) {
	workerPosition, err := NewPosition(position)
	if err != nil {
		return nil, err
	}
	return &Worker{name, workerPosition, salary, experienceYear}, nil
}

func (w *Worker) CalculateIncome() uint {
	return w.salary * w.experienceYear
}

func (w *Worker) Info() string {
	return fmt.Sprintf("%s — %d — %s", w.name, w.CalculateIncome(), w.position.name)
}

type CompanyInterface interface {
	AddWorkerInfo(name string, position string, salary uint, experienceYear uint) error
	SortWorkers() ([]string, error)
}

type Company struct {
	workers []*Worker
}

func (c *Company) AddWorkerInfo(name string, position string, salary uint, experienceYear uint) error {
	worker, err := NewWorker(name, position, salary, experienceYear)
	if err != nil {
		return err
	}
	c.workers = append(c.workers, worker)
	return nil
}

func (c *Company) sort() {
	slices.SortFunc(c.workers, func(w1 *Worker, w2 *Worker) int {
		if w1.CalculateIncome() > w2.CalculateIncome() {
			return -1
		}
		if w1.CalculateIncome() < w2.CalculateIncome() {
			return 1
		}
		if w1.position.level > w2.position.level {
			return -1
		}
		if w1.position.level < w2.position.level {
			return 1
		}
		return 0
	})
}

func (c *Company) SortWorkers() ([]string, error) {
	infoList := make([]string, 0, len(c.workers))
	c.sort()
	for _, worker := range c.workers {
		infoList = append(infoList, worker.Info())
	}
	return infoList, nil
}
