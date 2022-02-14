package tasks

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/beeker1121/goque"
)

type Item struct {
	Created     time.Time
	Description string
}

func (i Item) String() string {
	timeSpent := time.Since(i.Created).Truncate(time.Second)
	return fmt.Sprintf(
		"%s %s %s",
		i.Created.Format("2006-01-02 03:04"),
		timeSpent,
		i.Description,
	)
}

func Head() *Item {
	q := getQueue()
	defer q.Close()
	r, err := q.Peek()
	if err == goque.ErrEmpty {
		return nil
	}
	if err != nil {
		panic(err)
	}
	var itm Item
	err = r.ToObjectFromJSON(&itm)
	if err != nil {
		panic(err)
	}
	return &itm
}

func Put(i Item) {
	q := getQueue()
	defer q.Close()

	_, err := q.EnqueueObjectAsJSON(i)
	if err != nil {
		panic(err)
	}
}

func pop() *Item {
	q := getQueue()
	defer q.Close()
	r, err := q.Dequeue()
	if err == goque.ErrEmpty {
		return nil
	}
	if err != nil {
		panic(err)
	}
	var itm Item
	err = r.ToObjectFromJSON(&itm)
	if err != nil {
		panic(err)
	}
	return &itm
}

func Done() {
	itm := pop()
	fmt.Println(itm)
}

func getQueue() *goque.Queue {
	q, err := goque.OpenQueue(getDir())
	if err != nil {
		panic(err)
	}
	return q
}

func getDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(home, ".taskqueue")
}
