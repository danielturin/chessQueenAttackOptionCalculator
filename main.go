package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {

	// construct obstacles matrix for O(1) obstacles check
	o_m := buildObstaclesMatrix(obstacles, k, n)

	attackOptsCounter := int32(0)
	// HANDLE ROW TO THE RIGHT
	// if queen is not on the right-most column
	if c_q < n {
		for i := c_q + 1; i <= n; i++ {
			if o_m[r_q-1][i-1] != 1 {
				attackOptsCounter++
			} else {
				break
			}
		}
	}

	// HANDLE ROW TO THE LEFT
	if c_q > 1 {
		for i := c_q - 1; i >= 1; i-- {
			if o_m[r_q-1][i-1] != 1 {
				attackOptsCounter++
			} else {
				break
			}
		}
	}

	// HANDLE COLUMN DOWNWARDS
	if r_q > 1 {
		for i := r_q - 1; i >= 1; i-- {
			if o_m[i-1][c_q-1] != 1 {
				attackOptsCounter++
			} else {
				break
			}
		}
	}

	// HANDLE COLUMN UPWARDS
	if r_q < n {
		for i := r_q + 1; i <= n; i++ {
			if o_m[i-1][c_q-1] != 1 {
				attackOptsCounter++
			} else {
				break
			}
		}
	}

	// HANDLE BOTTOM LEFT DIAGONAL OPTIONS
	if r_q > 1 && c_q > 1 {
		r_i := r_q - 1
		c_i := c_q - 1
		for r_i >= 1 && c_i >= 1 {
			if o_m[r_i-1][c_i-1] != 1 {
				attackOptsCounter++
				r_i--
				c_i--
			} else {
				break
			}
		}
	}

	// HANDLE TOP RIGHT DIAGONAL OPTIONS
	if r_q < n && c_q < n {
		r_i := r_q + 1
		c_i := c_q + 1
		for r_i <= n && c_i <= n {
			if o_m[r_i-1][c_i-1] != 1 {
				attackOptsCounter++
				r_i++
				c_i++
			} else {
				break
			}
		}
	}

	// HANDLE TOP LEFT DIAGONAL OPTIONS
	if r_q < n && c_q > 1 {
		r_i := r_q + 1
		c_i := c_q - 1
		for r_i <= n && c_i >= 1 {
			if o_m[r_i-1][c_i-1] != 1 {
				attackOptsCounter++
				r_i++
				c_i--
			} else {
				break
			}
		}
	}

	// HANDLE BOTTOM RIGHT DIAGONAL OPTIONS
	if r_q > 1 && c_q < n {
		r_i := r_q - 1
		c_i := c_q + 1
		for r_i >= 1 && c_i <= n {
			if o_m[r_i-1][c_i-1] != 1 {
				attackOptsCounter++
				r_i--
				c_i++
			} else {
				break
			}
		}
	}

	fmt.Printf("Attack Options Counter is %d", attackOptsCounter)
	return attackOptsCounter
}

func buildObstaclesMatrix(obstacles [][]int32, k int32, n int32) [][]int32 {
	o_m := make([][]int32, n)
	for i := range o_m {
		o_m[i] = make([]int32, n)
	}

	// init matrix values to 0
	for r := int32(0); r < n; r++ {
		for c := int32(0); c < n; c++ {
			o_m[r][c] = 0
		}
	}

	temp_row_val := int32(0)
	temp_col_val := int32(0)

	for i := int32(0); i < k; i++ {
		temp_row_val = obstacles[i][0]
		temp_col_val = obstacles[i][1]
		o_m[temp_row_val-1][temp_col_val-1] = 1
	}

	return o_m
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != 2 {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
