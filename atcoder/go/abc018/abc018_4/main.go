package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	io, flush := NewIO()
	defer flush()
	n, m, p, q, r := io.ScanInt(), io.ScanInt(), io.ScanInt(), io.ScanInt(), io.ScanInt()
	chocoHappiness := PrepareEmpty2DintArray(n, m)
	for i := 0; i < r; i++ {
		g, b, h := io.ScanInt()-1, io.ScanInt()-1, io.ScanInt()
		chocoHappiness[g][b] += h
	}
	maxHappiness := 0
	for i := 1<<uint(p) - 1; i < 1<<uint(n); i = NextCmb(i) {
		happiness := 0
		sum := make([]int, m)
		for g := 0; g < n; g++ {
			// skip if the girl is not included in the group
			if i>>uint(g)&1 == 0 {
				continue
			}
			for b := 0; b < m; b++ {
				sum[b] += chocoHappiness[g][b]
			}
		}
		sort.Ints(sum)
		for j := 0; j < q; j++ {
			happiness += sum[m-1-j]
		}
		if maxHappiness < happiness {
			maxHappiness = happiness
		}
	}
	io.Println(maxHappiness)
}

func NextCmb(bits int) int {
	lowest := bits & -bits
	newBits := bits + lowest
	newBits |= ((bits & ^newBits) / lowest) >> 1
	return newBits
}

func PrepareEmptyIntArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 0
	}
	return arr
}

func PrepareEmpty2DintArray(y, x int) [][]int {
	arr := make([][]int, y)
	for i := 0; i < y; i++ {
		arr[i] = PrepareEmptyIntArray(x)
	}
	return arr
}

type IO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func NewIO() (*IO, func()) {
	io := &IO{
		scanner: newScanner(),
		writer:  newWriter(),
	}
	return io, func() { io.writer.Flush() }
}

func newScanner() *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 1000005), 1000005)
	s.Split(bufio.ScanWords)
	return s
}

func newWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func (io *IO) ScanString() string {
	if !io.scanner.Scan() {
		panic("scan string failed")
	}
	return io.scanner.Text()
}

func (io *IO) ScanStrings(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = io.ScanString()
	}
	return strs
}

func (io *IO) Scan2DStrings(y, x int) [][]string {
	strings := make([][]string, y)
	for i := 0; i < y; i++ {
		strings[i] = io.ScanStrings(x)
	}
	return strings
}

func (io *IO) Scan2DGraph(y int) [][]string {
	strs := make([][]string, y)
	for i := 0; i < y; i++ {
		strs[i] = strings.Split(io.ScanString(), "")
	}
	return strs
}

func (io *IO) ScanInt() int {
	return int(io.ScanInt64())
}

func (io *IO) ScanInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = io.ScanInt()
	}
	return ints
}

func (io *IO) Scan2DInts(y, x int) [][]int {
	ints := make([][]int, y)
	for i := 0; i < y; i++ {
		ints[i] = io.ScanInts(x)
	}
	return ints
}

func (io *IO) ScanInt64() int64 {
	i, err := strconv.ParseInt(io.ScanString(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *IO) ScanFloat64() float64 {
	i, _ := strconv.ParseFloat(io.ScanString(), 64)
	return i
}

func (io *IO) ScanFloat64s(n int) []float64 {
	floats := make([]float64, n)
	for i := 0; i < n; i++ {
		floats[i] = io.ScanFloat64()
	}
	return floats
}

func (io *IO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
