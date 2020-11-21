package bitemporal

import (
	"sort"
	"time"

	"github.com/guregu/null"
	"github.com/lovung/gotemporal"
)

type Collection []Model

func (c *Collection) ToSlice() []Model {
	return *c
}

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Verify() bool {
	list := c.ToSlice()

	c.SortByValidTime()
	for i := range list {
		if list[i].ValidFrom.Equal(list[i].ValidTo.Time) {
			return false
		}
		if i == len(list)-1 {
			break
		}
		if list[i].ValidTo.IsZero() {
			return false
		}
		if !list[i].ValidTo.Time.Equal(list[i+1].ValidFrom) {
			return false
		}
	}
	return true
}

// SortByValidTime the collection by the valid time
func (c *Collection) SortByValidTime() {
	list := c.ToSlice()

	sort.SliceStable(
		list,
		func(i int, j int) bool {
			return list[i].ValidFrom.Before(list[j].ValidFrom)
		},
	)
}

func (c Collection) GetByID(id gotemporal.IDer) (Model, error) {
	list := c.ToSlice()

	for i := range list {
		if list[i].ID == id.GetID() {
			return list[i], nil
		}
	}
	return Model{}, gotemporal.ErrNotFound
}

func (c Collection) GetByValidAt(at time.Time) (Model, error) {
	list := c.ToSlice()

	c.SortByValidTime()
	for i := range list {
		if list[i].ValidFrom.After(at) {
			continue
		}
		if list[i].ValidTo.IsZero() || list[i].ValidTo.Time.After(at) {
			return list[i], nil
		}
	}
	return Model{}, gotemporal.ErrNotFound
}

func (c *Collection) DeleteByID(id gotemporal.IDer, modifyTime null.Time) (Collection, error) {
	var (
		pIndex  *int
		now     = time.Now()
		newList Collection
	)

	list := c.ToSlice()

	c.SortByValidTime()

	for i := range list {
		if list[i].ID == id.GetID() {
			pIndex = &i
			break
		}
	}

	if pIndex == nil {
		return nil, gotemporal.ErrNotFound
	}

	index := *pIndex
	// Delete first item
	if index == 0 {
		list[index].SysTo = null.NewTime(now, true)
		newModel := list[1]
		if !modifyTime.IsZero() && newList[0].ValidFrom.After(modifyTime.Time) {
			newModel.Clean()
			newModel.ValidFrom = modifyTime.Time
			newList = append(newList, newModel)
		}
		return newList, nil
	}

	// Delete the last item
	if index == c.Len() {
		list[index].SysTo = null.NewTime(now, true)
		newModel := list[index-1]
		if modifyTime.IsZero() || newList[len(newList)-1].ValidFrom.Before(modifyTime.Time) {
			newModel.Clean()
			newModel.ValidTo = modifyTime
			newList = append(newList, newModel)
		}
		return newList, nil
	}

	list[index].SysTo = null.NewTime(now, true)
	prevModel := list[index-1]
	nextModel := list[index+1]
	if !modifyTime.IsZero() &&
		newList[index-1].ValidFrom.Before(modifyTime.Time) &&
		newList[index].ValidTo.Time.After(modifyTime.Time) {
		prevModel.Clean()
		nextModel.Clean()

		prevModel.ValidTo = modifyTime
		nextModel.ValidFrom = modifyTime.Time

		newList = append(newList, prevModel)
		newList = append(newList, nextModel)
	}

	return newList, nil
}
