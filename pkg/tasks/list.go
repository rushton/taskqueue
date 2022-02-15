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
	Started     *time.Time
	Finished    *time.Time
	Description string
}

func (i Item) String() string {
	if i.Started != nil {
		timeSpent := time.Since(*i.Started).Truncate(time.Second)
		return fmt.Sprintf(
			"%s %s %s",
			i.Created.Format("2006-01-02 03:04"),
			timeSpent,
			i.Description,
		)
	}
	return fmt.Sprintf(
		"%s <Not Started> %s",
		i.Created.Format("2006-01-02 03:04"),
		i.Description,
	)
}

func List() ([]Item, error) {
	q := getQueue()
	defer q.Close()
	itms := make([]Item, q.Length())
	var offset uint64 = 0
	for i := 0; i < int(q.Length()); i++ {
		qitm, err := q.PeekByOffset(uint64(i))
		if err != nil {
			return nil, err
		}
		offset += 1
		var itm Item
		err = qitm.ToObjectFromJSON(&itm)
		if err != nil {
			return nil, err
		}
		itms[i] = itm
	}
	return itms, nil
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
	if q.Length() == 0 {
		started := time.Now()
		i.Started = &started
	}

	_, err := q.EnqueueObjectAsJSON(i)
	if err != nil {
		panic(err)
	}
}

func Done() (*Item, error) {
	itm := pop()
	q := getQueue()

	// start the next item
	nxt, err := q.Peek()

	if err == goque.ErrEmpty {
		return itm, nil
	}
	if err != nil {
		return nil, err
	}
	if nxt == nil {
		return itm, nil
	}
	var nxtItm Item
	err = nxt.ToObjectFromJSON(&nxtItm)
	if err != nil {
		return nil, err
	}
	startedTime := time.Now()
	nxtItm.Started = &startedTime
	_, err = q.UpdateObjectAsJSON(nxt.ID, nxtItm)
	if err != nil {
		return nil, err
	}
	return itm, nil
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
