package input

import (
	"bufio"
	"errors"
	"io"

	"find-bin-width/pkg/vpdb"
	"find-bin-width/pkg/xtype"
)

type LineGroup struct {
	AttributeStableID vpdb.AttributeStableID
	EntityTypeID      vpdb.EntityTypeID
	DataType          xtype.DataType
	Values            any
	IsNA              bool
}

type LineGroupIterator interface {
	HasNext() bool

	Next() (LineGroup, error)
}

type inMemoryLineGroupIterator struct {
	index  int
	sorted []lineKey
	groups map[lineKey]lineEntries
}

func (i *inMemoryLineGroupIterator) HasNext() bool {
	return i.index+1 < len(i.sorted)
}

func (i *inMemoryLineGroupIterator) Next() (LineGroup, error) {
	if !i.HasNext() {
		return LineGroup{}, errors.New("no such element")
	}

	key := i.sorted[i.index]
	i.index++

	if val, ok := i.groups[key]; ok {
		return LineGroup{
			AttributeStableID: key.attributeSourceID,
			EntityTypeID:      key.entityTypeID,
			DataType:          val.dataType,
			Values:            val.entries,
			IsNA:              val.na,
		}, nil
	}

	return LineGroup{}, errors.New("illegal state: " + key.String() + " could not be found in line group map")
}

func StreamGroups(stream io.Reader, rmNa bool, lineSplitter LineSplitFunction) (LineGroupIterator, error) {
	return &collectingLineGroupIterator{
		it: inputIterator{
			scanner: bufio.NewScanner(stream),
			splitFn: lineSplitter,
		},
		rmNa: rmNa,
	}, nil
}

type collectingLineGroupIterator struct {
	it     inputIterator
	buf    LineSplitResult
	hasBuf bool
	rmNa   bool
}

func (c *collectingLineGroupIterator) HasNext() bool {
	return c.it.hasNext() || c.hasBuf
}

func (c *collectingLineGroupIterator) Next() (LineGroup, error) {
	var key1, key2, value string
	values := lineEntries{}

	if c.hasBuf {
		key1, key2, value = c.buf.split()
		c.hasBuf = false
	} else {
		if tmp, err := c.it.next(); err != nil {
			return LineGroup{}, c.wrapError(err)
		} else {
			key1, key2, value = tmp.split()
		}
	}

	key, err := makeLineKey(key1, key2)
	if err != nil {
		return LineGroup{}, c.wrapError(err)
	}

	err = values.append(&key, value, c.rmNa)
	if err != nil {
		return LineGroup{}, c.wrapError(err)
	}

	for c.it.hasNext() {
		next, err := c.it.next()
		if err != nil {
			return LineGroup{}, c.wrapError(err)
		}

		if next[0] != key1 || next[1] != key2 {
			c.buf = next
			c.hasBuf = true
			break
		}

		err = values.append(&key, next[2], c.rmNa)
		if err != nil {
			return LineGroup{}, c.wrapError(err)
		}
	}

	return LineGroup{
		AttributeStableID: key.attributeSourceID,
		EntityTypeID:      key.entityTypeID,
		DataType:          values.dataType,
		Values:            values.entries,
		IsNA:              values.na,
	}, nil
}

func (c *collectingLineGroupIterator) wrapError(e error) error {
	if v, ok := e.(inputError); ok {
		return v
	}

	return inputError{c.it.line, e}
}
