package input

import (
	"bufio"
	"io"
	"slices"

	"find-bin-width/pkg/vpdb"
)

func CollectAndSortGroups(stream io.Reader, rmNA bool, lineSplitter LineSplitFunction) (LineGroupIterator, error) {
	collector := sortingCollector{
		input: inputIterator{
			scanner: bufio.NewScanner(stream),
			splitFn: lineSplitter,
		},
		rmNA: rmNA,
	}

	if err := collector.collect(); err != nil {
		return nil, err
	}

	return collector.iterator(), nil
}

type sortingCollector struct {
	rmNA   bool
	input  inputIterator
	groups map[lineKey]lineEntries
}

func (g *sortingCollector) collect() error {
	for g.input.hasNext() {
		if res, err := g.input.next(); err != nil {
			return err
		} else if err = g.add(res[0], res[1], res[2]); err != nil {
			return err
		}
	}

	return nil
}

func (g *sortingCollector) add(attributeSourceID, entityTypeID, value string) error {
	asi, err := vpdb.ParseAttributeStableID(attributeSourceID)
	if err != nil {
		return err
	}

	eti, err := vpdb.ParseEntityTypeID(entityTypeID)
	if err != nil {
		return err
	}

	if g.groups == nil {
		g.groups = make(map[lineKey]lineEntries, 1024)
	}

	key := lineKey{asi, eti}

	if val, ok := g.groups[key]; ok {
		return val.append(&key, value, g.rmNA)
	} else {
		val = lineEntries{}
		g.groups[key] = val
		return val.append(&key, value, g.rmNA)
	}
}

func (g *sortingCollector) iterator() LineGroupIterator {
	keys := make([]lineKey, len(g.groups))

	i := 0
	for k := range g.groups {
		keys[i] = k
		i++
	}

	slices.SortFunc(keys, lineKeySort)

	return &inMemoryLineGroupIterator{
		sorted: keys,
		groups: g.groups,
	}
}

func lineKeySort(a, b lineKey) int {
	if cmp := a.attributeSourceID.Compare(b.attributeSourceID); cmp != 0 {
		return cmp
	}

	return a.entityTypeID.Compare(b.entityTypeID)
}
