package disjoint_set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSet(t *testing.T) {
	djs := NewDisjointSet()

	// good sets tests
	set1, ok := djs.MakeSet(1)
	assert.True(t, ok, "set1 not ok")
	assert.NotNil(t, set1, "set1 should not be nil")
	assert.Len(t, set1.Items, 1)
	assert.Equal(t, set1.Items[0], 1)

	set2, ok := djs.MakeSet(2)
	assert.True(t, ok, "set2 not ok")
	assert.NotNil(t, set2, "set1 should not be nil")
	assert.Len(t, set2.Items, 1)
	assert.Equal(t, set2.Items[0], 2)

	// test duplicate
	dupSet, ok := djs.MakeSet(2)
	assert.False(t, ok, "dupSet is ok")
	assert.Nil(t, dupSet, "dupSet should be nil")

}

func TestFindSet(t *testing.T) {
	djs := NewDisjointSet()

	set1, _ := djs.MakeSet(1)
	set2, _ := djs.MakeSet(2)
	setAsdf, _ := djs.MakeSet("asdf")

	found1, ok := djs.FindSet(1)
	assert.True(t, ok)
	assert.Equal(t, set1.Items[0], found1)

	found2, ok := djs.FindSet(2)
	assert.True(t, ok)
	assert.Equal(t, set2.Items[0], found2)

	foundAsdf, ok := djs.FindSet("asdf")
	assert.True(t, ok)
	assert.Equal(t, setAsdf.Items[0], foundAsdf)

	foundBad, ok := djs.FindSet("bad")
	assert.False(t, ok)
	assert.Nil(t, foundBad)
}

func TestUnion(t *testing.T) {
	djs := NewDisjointSet()

	set1, _ := djs.MakeSet(1)
	set2, _ := djs.MakeSet(2)
	set3, _ := djs.MakeSet(3)
	set4, _ := djs.MakeSet(4)
	set5, _ := djs.MakeSet(5)
	djs.Union(set1, set2)
	djs.Union(set3, set4)
	assert.Len(t, djs.Sets, 3)
	assert.Equal(t, djs.Items[1], set1)
	assert.Equal(t, djs.Items[2], set1)
	assert.Equal(t, djs.Items[3], set3)
	assert.Equal(t, djs.Items[4], set3)
	assert.Equal(t, djs.Items[5], set5)

	assert.Equal(t, djs.Sets[1], set1)
	assert.Equal(t, djs.Sets[3], set3)
	assert.Equal(t, djs.Sets[5], set5)
}
