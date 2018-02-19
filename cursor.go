package dbflex

import (
	"github.com/eaciit/toolkit"
)

type ICursor interface {
	Reset() error
	Fetch(interface{}) error
	Fetchs(interface{}, int) error
	Count() int
	CountAsync() <-chan int
	Close()
	Error() error
	SetCloseAfterFetch() ICursor
	CloseAfterFetch() bool
	SetCountQuery(IQuery)
	CountQuery() IQuery
}

type CursorBase struct {
	err             error
	closeafterfetch bool

	self       ICursor
	countQuery IQuery
}

func (b *CursorBase) SetError(err error) {
	b.err = err
}

func (b *CursorBase) Error() error {
	return b.err
}

func (b *CursorBase) this() ICursor {
	if b.self == nil {
		return b
	} else {
		return b.self
	}
}

func (b *CursorBase) SetThis(o ICursor) ICursor {
	b.self = o
	return o
}

func (b *CursorBase) Reset() error {
	panic("not implemented")
}

func (b *CursorBase) Fetch(interface{}) error {
	panic("not implemented")
}

func (b *CursorBase) Fetchs(interface{}, int) error {
	panic("not implemented")
}

func (b *CursorBase) Count() int {
	if b.countQuery == nil {
		b.SetError(toolkit.Errorf("cursor has no countquery"))
		return 0
	}

	recordcount := struct {
		Count int
	}{}
	err := b.countQuery.Cursor(nil).Fetch(&recordcount)
	if err != nil {
		b.SetError(toolkit.Errorf("unable to get count. %s", err.Error()))
		return 0
	}

	return recordcount.Count
}

func (b *CursorBase) CountAsync() <-chan int {
	out := make(chan int)
	go func(o chan int) {
		o <- b.Count()
	}(out)
	return out
}

func (b *CursorBase) SetCountQuery(q IQuery) {
	b.countQuery = q
}

func (b *CursorBase) CountQuery() IQuery {
	return b.countQuery
}

func (b *CursorBase) SetCloseAfterFetch() ICursor {
	b.closeafterfetch = true
	return b.this()
}

func (b *CursorBase) CloseAfterFetch() bool {
	return b.closeafterfetch
}

func (b *CursorBase) Close() {
}

func (b *CursorBase) Serialize(dest interface{}) error {
	return toolkit.Error("Serialize is not yet implemented")
}
