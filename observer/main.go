package main

import (
	"fmt"
	"time"
)

type Data struct {
	data int64
}

type Observers interface {
	OnNotify(data Data)
}

type Notifier interface {
	Register(observers Observers)
	Deregister(observers Observers)
	Notify(data Data)
}

type DataObserver struct {
	ID string
}
type DataNotifier struct {
	observers map[Observers]struct{}
}

func (notifier DataNotifier) Register(l Observers) {
	notifier.observers[l] = struct{}{}
}
func (notifier DataNotifier) Deregister(l Observers) {
	delete(notifier.observers, l)
}
func (notifier *DataNotifier) Notify(e Data) {
	for o := range notifier.observers {
		o.OnNotify(e)
	}
}

func (observer DataObserver) OnNotify(data Data) {
	fmt.Print("--------------------------------\n")
	fmt.Printf("Income data %d \n", data.data)
	fmt.Printf("ID %s \n", observer.ID)

}

func main() {
	notifier := DataNotifier{
		observers: map[Observers]struct{}{},
	}
	notifier.Register(&DataObserver{ID: "robot_01"})
	notifier.Register(&DataObserver{ID: "robot_02"})
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			notifier.Notify(Data{data: t.UnixNano()})
		}
	}
}
