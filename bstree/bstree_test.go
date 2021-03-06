package bstree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpsert(t *testing.T) {
	bst := BSTree{}

	bst.Upsert("foo", "bar")
	assert.Equal(t, "foo", bst.root.key)
	assert.Equal(t, "bar", bst.root.val)

	// left
	bst.Upsert("bbb", "bar-b")
	assert.Equal(t, "bbb", bst.root.left.key)
	assert.Equal(t, "bar-b", bst.root.left.val)
	bst.Upsert("aaa", "bar-a")
	assert.Equal(t, "aaa", bst.root.left.left.key)
	assert.Equal(t, "bar-a", bst.root.left.left.val)

	// right
	bst.Upsert("yyy", "bar-y")
	assert.Equal(t, "yyy", bst.root.right.key)
	assert.Equal(t, "bar-y", bst.root.right.val)
	bst.Upsert("zzz", "bar-z")
	assert.Equal(t, "zzz", bst.root.right.right.key)
	assert.Equal(t, "bar-z", bst.root.right.right.val)

	// overwrite
	bst.Upsert("foo", "1337")
	assert.Equal(t, "foo", bst.root.key)
	assert.Equal(t, "1337", bst.root.val)
}

func TestValue(t *testing.T) {
	var err error
	bst := BSTree{}
	bst.Upsert("foo", "bar")

	// root
	val, err := bst.Value("foo")
	assert.Equal(t, "bar", val)
	assert.Equal(t, nil, err)

	// left
	_, err = bst.Value("aaa")
	assert.Equal(t, ErrorNotFound, err)
	bst.Upsert("aaa", "bar-a")
	val, err = bst.Value("aaa")
	assert.Equal(t, "bar-a", val)
	assert.Equal(t, nil, err)

	// right
	_, err = bst.Value("zzz")
	assert.Equal(t, ErrorNotFound, err)
	bst.Upsert("zzz", "bar-z")
	val, err = bst.Value("zzz")
	assert.Equal(t, "bar-z", val)
	assert.Equal(t, nil, err)
}

func TestIsLeafHasLeftHasRight(t *testing.T) {
	var nd *node

	nd = &node{}
	assert.Equal(t, true, nd.isLeaf())
	assert.Equal(t, false, nd.hasLeft())
	assert.Equal(t, false, nd.hasRight())

	nd = &node{
		left: &node{},
	}
	assert.Equal(t, false, nd.isLeaf())
	assert.Equal(t, true, nd.hasLeft())
	assert.Equal(t, false, nd.hasRight())

	nd = &node{
		right: &node{},
	}
	assert.Equal(t, false, nd.isLeaf())
	assert.Equal(t, false, nd.hasLeft())
	assert.Equal(t, true, nd.hasRight())

	nd = &node{
		left:  &node{},
		right: &node{},
	}
	assert.Equal(t, false, nd.isLeaf())
	assert.Equal(t, true, nd.hasLeft())
	assert.Equal(t, true, nd.hasRight())
}

func TestDelete(t *testing.T) {
	// leaf node
	{
		var err error
		bst := BSTree{}
		bst.Upsert("foo", "bar")
		bst.Upsert("aaa", "bar-a")
		bst.Upsert("zzz", "bar-z")

		// left
		err = bst.Delete("aaa")
		assert.Equal(t, nil, err)
		err = bst.Delete("aaa")
		assert.Equal(t, ErrorNotFound, err)

		// right
		err = bst.Delete("zzz")
		assert.Equal(t, nil, err)
		err = bst.Delete("zzz")
		assert.Equal(t, ErrorNotFound, err)
	}
	// left child
	{
		var err error
		bst := BSTree{}
		bst.Upsert("foo", "bar")
		bst.Upsert("aaa", "bar-a")

		err = bst.Delete("foo")
		assert.Equal(t, nil, err)
		err = bst.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	// right child
	{
		var err error
		bst := BSTree{}
		bst.Upsert("foo", "bar")
		bst.Upsert("zzz", "bar-z")

		err = bst.Delete("foo")
		assert.Equal(t, nil, err)
		err = bst.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	// two children
	{
		var err error
		bst := BSTree{}
		bst.Upsert("foo", "bar")
		bst.Upsert("aaa", "bar-a")
		bst.Upsert("zzz", "bar-z")

		err = bst.Delete("foo")
		assert.Equal(t, nil, err)
		err = bst.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
}

func TestMin(t *testing.T) {
	bst := BSTree{}
	bst.Upsert("foo", "bar")
	bst.Upsert("bbb", "bar-b")
	bst.Upsert("aaa", "bar-a")
	assert.Equal(t, "bar-a", bst.root.min().val)
}

func TestHeight(t *testing.T) {
	bst := BSTree{}
	assert.Equal(t, 0, bst.Height())

	bst.Upsert("foo", "bar")
	assert.Equal(t, 1, bst.Height())

	bst.Upsert("aaa", "bar-a")
	assert.Equal(t, 2, bst.Height())
	bst.Upsert("zzz", "bar-z")
	assert.Equal(t, 2, bst.Height())
}

func TestIter(t *testing.T) {
	bst := BSTree{}
	bst.Upsert("foo", "bar")
	bst.Upsert("aaa", "bar-a")
	bst.Upsert("zzz", "bar-z")
	bst.Upsert("bbb", "bar-b")
	bst.Upsert("rrr", "bar-r")
	var n int
	for i := range bst.Iter() {
		switch n {
		case 0:
			assert.Equal(t, "aaa", i.Key)
			assert.Equal(t, "bar-a", i.Val)
		case 1:
			assert.Equal(t, "bbb", i.Key)
			assert.Equal(t, "bar-b", i.Val)
		case 2:
			assert.Equal(t, "foo", i.Key)
			assert.Equal(t, "bar", i.Val)
		case 3:
			assert.Equal(t, "rrr", i.Key)
			assert.Equal(t, "bar-r", i.Val)
		case 4:
			assert.Equal(t, "zzz", i.Key)
			assert.Equal(t, "bar-z", i.Val)
		}
		n++
	}
	assert.Equal(t, 5, n)
}
