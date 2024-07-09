package utils

type Queue[TItem any] struct {
	items []TItem

	IsEmpty func() bool

	Enqueue func(item TItem)
	Dequeue func() TItem
}

func CreateQueue[TItem any]() Queue[TItem] {

	data := make([]TItem, 0)

	enqueue := func(item TItem) {
		data = append(data, item)
	}

	dequeue := func() TItem {
		item := data[0]
		data = data[1:]
		return item
	}

	isEmpty := func() bool {
		return len(data) == 0
	}

	return Queue[TItem]{
		items:   data,
		IsEmpty: isEmpty,
		Enqueue: enqueue,
		Dequeue: dequeue,
	}
}
