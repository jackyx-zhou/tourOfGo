// Exercise: Errors

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return x, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	for {
		diff := (z*z - x) / (2*z)
		if math.Abs(diff) < 0.000000000001 {
			return z, nil
		} else {
			z -= diff
			fmt.Println(z)
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	val, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}

// --------------------------------------------------------
// Exercise: rot13Reader

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rotr.r.Read(b)
	for i:=0; i < n; i++ {
		c := &b[i]
		switch {
		case *c >= 'A' && *c <= 'M':
			fallthrough
		case *c >= 'a' && *c <= 'm':
			*c += 13
		case *c >= 'N' && *c <= 'Z':
			fallthrough
		case *c >= 'n' && *c <= 'z':
			*c -= 13
		}
	}
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// ---------------------------------------------------------
// Exercise: Equivalent Binary Trees
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Since we use range with channel, 
// it needs to be closed to finish the blocking and exit loop
func WalkTree (t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// Since it's a sorted tree we can determine equivalence by inspecting each element
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	var arr1, arr2 []int
	go WalkTree(t1, ch1)
	
	ch2 := make(chan int)
	go WalkTree(t2, ch2)
	
	for i1 := range ch1 {
		arr1 = append(arr1, i1)
	}
	
	for i2 := range ch2 {
		arr2 = append(arr2, i2)
	}
	
	if len(arr1) != len(arr2) {
        return false
    }
	// sorted tree
    for i, v := range arr1 {
        if v != arr2[i] {
            return false
        }
    }
    return true
}

func main() {
	ch1 := make(chan int)
	go WalkTree(tree.New(2), ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
	fmt.Println("---------------------------")
	fmt.Println("Should return true:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Should return false:", Same(tree.New(1), tree.New(2)))
}
