package api

import (
	"context"
	"net/url"

	"github.com/ngrok/ngrok-api-go/v5"
)

type ListResponse[T any] interface {
	Items() []T
	NextPage() *string
}

type Iter[T any, LRT ListResponse[T]] struct {
	client *ngrok.BaseClient
	n      int
	items  []T
	err    error

	nextPage *url.URL
}

func NewIter[T any, LRT ListResponse[T]](client *ngrok.BaseClient, nextPage *url.URL) *Iter[T, LRT] {
	return &Iter[T, LRT]{
		client:   client,
		nextPage: nextPage,
		n:        -1,
	}
}

// Next returns true if there is another value available in the iterator. If it
// returs true it also advances the iterator to that next available item.
func (it *Iter[T, LRT]) Next(ctx context.Context) bool {
	// no more if there is an error
	if it.err != nil {
		return false
	}

	// advance the iterator
	it.n += 1

	// is there an available item?
	if it.n < len(it.items) {
		return true
	}

	if it.nextPage == nil {
		return false
	}

	// fetch the next page
	var resp LRT
	err := it.client.Do(ctx, "GET", it.nextPage, nil, &resp)
	if err != nil {
		it.err = err
		return false
	}

	// parse the next page URI as soon as we get it and store it
	// so we can use it on the next fetch
	if resp.NextPage() != nil {
		it.nextPage, it.err = url.Parse(*resp.NextPage())
		if it.err != nil {
			return false
		}
	} else {
		it.nextPage = nil
	}

	// page with zero items means there are no more
	if len(resp.Items()) == 0 {
		return false
	}

	it.n = -1
	it.items = resp.Items()
	return it.Next(ctx)
}

// Item returns the item currently pointed to by the iterator.
func (it *Iter[T, LRT]) Item() *T {
	return &it.items[0]
}

// If Next() returned false because an error was encountered while fetching the
// next value Err() will return that error. A caller should always check Err()
// after Next() returns false.
func (it *Iter[T, LRT]) Err() error {
	return it.err
}
