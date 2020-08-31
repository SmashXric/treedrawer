package tree

import (
	"fmt"
	"log"
	"strconv"

	"github.com/m1gwings/treedrawer/drawer"
)

// NodeValue is the interface that wraps the Draw method.
//
// The Draw method allows to convert data into its unicode canvas representation.
// With the Draw method you can control how your data is going to be appear on the tree.
type NodeValue interface {
	Draw() *drawer.Drawer
}

// NodeInt64 is the default type for drawing int64s on the tree.
type NodeInt64 int64

// Draw satisfies the NodeValue interface.
func (i NodeInt64) Draw() *drawer.Drawer {
	return NodeString(strconv.Itoa(int(i))).Draw()
}

// NodeString is the default type for drawing strings on the tree.
type NodeString string

// Draw satisfies the NodeValue interface.
func (s NodeString) Draw() *drawer.Drawer {
	d := drawer.NewDrawer(len(s), 1)
	for i, r := range s {
		err := d.DrawRune(r, i, 0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing %d th rune of %s string in NodeString.Draw() method: %v", i, s, err))
		}
	}
	return d
}