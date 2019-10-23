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

