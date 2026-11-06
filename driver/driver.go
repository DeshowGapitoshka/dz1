package driver

import (
	"bufio"
	"bytes"
	bst "dz1/BST"
	rbt "dz1/RBT"
	"dz1/hash"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func validValue(s string) bool {
	if s == "" {
		return false
	}
	i := 0
	if s[0] == '-' {
		i = 1
		if len(s) == 1 {
			return false
		}
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func DriverBst(data *os.File, out_file *os.File, flag int) time.Duration {
	var times []int64
	content, err := io.ReadAll(data)
	if err != nil {
		return 0
	}
	for i := 0; i < flag; i++ {
		if _, err := data.Seek(0, io.SeekStart); err != nil {
			return 0
		}
		start := time.Now()
		bst := &bst.BST{}
		reader := bytes.NewReader(content)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)

			if len(parts) == 0 {
				continue
			}

			switch {
			case parts[0] == "INSERT" && len(parts) == 3:
				if len(parts[1]) > 256 {
					fmt.Fprintln(out_file, "BAD")
					continue
				}
				s := parts[2]
				if !validValue(s) {
					fmt.Fprintln(out_file, "BAD")
					continue
				}
				temp_numb, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					fmt.Fprintf(out_file, "BAD\n")
					continue
				}
				out := bst.Insert(parts[1], temp_numb)
				fmt.Fprintf(out_file, "%s\n", out)

			case parts[0] == "FIND" && len(parts) == 2:
				out, flag := bst.Find(parts[1])
				if flag == true {
					fmt.Fprintf(out_file, "%d\n", out)
				} else {
					fmt.Fprintf(out_file, "NOT_FOUND\n")
				}

			case parts[0] == "ERASE" && len(parts) == 2:
				out := bst.Erase(parts[1])
				fmt.Fprintf(out_file, "%s\n", out)

			case parts[0] == "SIZE" && len(parts) == 1:
				out := bst.Scale()
				fmt.Fprintf(out_file, "%d\n", out)

			case parts[0] == "DUMP" && len(parts) == 1:
				out := bst.Dump()
				fmt.Fprintf(out_file, "DUMP %d\n", len(out))
				for _, p := range out {
					fmt.Fprintf(out_file, "%s %d\n", p.Key, p.Value)
				}
			default:
				fmt.Fprintf(out_file, "BAD\n")
			}

		}
		elapsed := int64(time.Since(start))
		times = append(times, elapsed)
	}
	if flag > 1 {
		return calcMediana(times)
	}
	return 0
}

func DriverHash(data *os.File, out_file *os.File, flag int) time.Duration {
	var times []int64
	for i := 0; i < flag; i++ {
		if _, err := data.Seek(0, io.SeekStart); err != nil {
			return 0
		}
		start := time.Now()
		hash := hash.New()
		scanner := bufio.NewScanner(data)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) == 0 {
				continue
			}

			switch {
			case parts[0] == "INSERT" && len(parts) == 3:
				if len(parts[1]) > 256 {
					fmt.Fprintln(out_file, "BAD")
					continue
				}
				s := parts[2]
				if !validValue(s) {
					fmt.Fprintln(out_file, "BAD")
					continue
				}
				temp_numb, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					fmt.Fprintf(out_file, "BAD\n")
					continue
				}
				out := hash.Insert(parts[1], temp_numb)
				fmt.Fprintf(out_file, "%s\n", out)

			case parts[0] == "FIND" && len(parts) == 2:
				out, flag := hash.Find(parts[1])
				if flag == true {
					fmt.Fprintf(out_file, "%d\n", out)
				} else {
					fmt.Fprintf(out_file, "NOT_FOUND\n")
				}

			case parts[0] == "ERASE" && len(parts) == 2:
				out := hash.Erase(parts[1])
				fmt.Fprintf(out_file, "%s\n", out)

			case parts[0] == "SIZE" && len(parts) == 1:
				out := hash.Size()
				fmt.Fprintf(out_file, "%d\n", out)

			case parts[0] == "DUMP" && len(parts) == 1:
				out := hash.Dump()
				fmt.Fprintf(out_file, "DUMP %d\n", len(out))
				for _, p := range out {
					fmt.Fprintf(out_file, "%s %d\n", p.Key, p.Value)
				}
			default:
				fmt.Fprintf(out_file, "BAD\n")
			}

		}
		elapsed := time.Since(start)
		times = append(times, int64(elapsed))
	}
	if flag > 1 {
		return calcMediana(times)
	}
	return 0
}

func DriverRbt(data *os.File, out_file *os.File, flag int) time.Duration {
	var times []int64
	for i := 0; i < flag; i++ {
		if _, err := data.Seek(0, io.SeekStart); err != nil {
			return 0
		}
		start := time.Now()
		bst := &rbt.RBT{}
		scanner := bufio.NewScanner(data)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) == 0 {
				continue
			}

			switch {
			case parts[0] == "INSERT" && len(parts) == 3:
				if len(parts[1]) > 256 {
					fmt.Fprintln(out_file, "BAD")
					continue
				}
				s := parts[2]
				if !validValue(s) {
					fmt.Fprintln(out_file, "BAD")
					continue
				}
				temp_numb, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					fmt.Fprintf(out_file, "BAD\n")
					continue
				}
				out := bst.Insert(parts[1], temp_numb)
				fmt.Fprintf(out_file, "%s\n", out)

			case parts[0] == "FIND" && len(parts) == 2:
				out, flag := bst.Find(parts[1])
				if flag == true {
					fmt.Fprintf(out_file, "%d\n", out)
				} else {
					fmt.Fprintf(out_file, "NOT_FOUND\n")
				}

			case parts[0] == "ERASE" && len(parts) == 2:
				out := bst.Erase(parts[1])
				fmt.Fprintf(out_file, "%s\n", out)

			case parts[0] == "SIZE" && len(parts) == 1:
				out := bst.Scale()
				fmt.Fprintf(out_file, "%d\n", out)

			case parts[0] == "DUMP" && len(parts) == 1:
				out := bst.Dump()
				fmt.Fprintf(out_file, "DUMP %d\n", len(out))
				for _, p := range out {
					fmt.Fprintf(out_file, "%s %d\n", p.Key, p.Value)
				}
			default:
				fmt.Fprintf(out_file, "BAD\n")
			}

		}
		elapsed := time.Since(start)
		times = append(times, int64(elapsed))
	}
	if flag > 1 {
		return calcMediana(times)
	}
	return 0
}

func Benchmark(data *os.File, out_file *os.File) {
	var time1, time2, time3 time.Duration
	// time2 = DriverRbt(data, out_file, 3)
	// time3 = DriverHash(data, out_file, 3)

	time1 = DriverBst(data, out_file, 5)
	time2 = DriverRbt(data, out_file, 5)
	time3 = DriverHash(data, out_file, 5)
	fmt.Printf("Медиана BST %.3f мс\nМедиана RBT %.3f мс\nМедиана Hash %.3f мс\n",
		time1.Seconds()*1_000,
		time2.Seconds()*1_000,
		time3.Seconds()*1_000)
	out_file.Close()
}

func calcMediana(massiv []int64) time.Duration {
	sort.Slice(massiv, func(i, j int) bool {
		return massiv[i] < massiv[j]
	})
	n := len(massiv) / 2
	if len(massiv)%2 == 0 {
		mediana := (massiv[n-1] + massiv[n]) / 2
		return time.Duration(mediana)
	}
	mediana := massiv[n]
	return time.Duration(mediana)
}
