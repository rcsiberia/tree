package tree

import (
	"testing"
	"reflect"
)

func btree123() *Btree {
	btree := New()
	btree.Insert(1).Insert(2).Insert(3)
	return btree
}

func TestBtree(t *testing.T) {
	btree := btree123()

	if btree.Size() != 3 {
		t.Error("size of tree should be 3")
	}
	if btree.Get(2) != 2 {
		t.Error("get value should be 2")
	}
	if btree.Head().Value != 1 {
		t.Error("beggining of tree should be 1")
	}
	if btree.Root.Value != 2 {
		t.Error("Root of tree should be 2")
	}
}

func TestIterateBtree(t *testing.T) {
	const capacity = 3
	btree := btree123()

	var b [capacity]int
	for i, n := 0, btree.Head(); n != nil; i, n = i + 1, n.Next() {
		b[i] = n.Value.(int)
	}
	c := [capacity]int{1, 2, 3}
	if !reflect.DeepEqual(c , b) {
		t.Error(c, "should equal", b)
	}

	for i, n := 0, btree.Tail(); n != nil; i, n = i + 1, n.Prev() {
		b[i] = n.Value.(int)
	}
	d := [capacity]int{3, 2, 1}
	if !reflect.DeepEqual(b , d) {
		t.Error(b, "should equal", d)
	}

	slice := btree.Slice()
	for i, v := range slice {
		if c[i] != v {
			t.Error(c[i], "should equal", v)
		}
	}
}

func TestRemoveBtree(t *testing.T) {
	btree := btree123()
	btree.Remove(1)

	size := btree.Size()
	if size != 2 {
		t.Error(size, "should be size", 2)
	}
}