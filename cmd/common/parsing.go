package common

import (
	"strings"
	"strconv"
	"bufio"
	"io"
)

func ReadInts(r io.Reader) (integers []int) {
	scanner := bufio.NewScanner(r)

	// Does Scan() take care of trim() or Atoi?
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		PanicIf(err)

		integers = append(integers, val)
	}

	return
}

func CommaSeparatedToInt(str string) (integers []int) {
	splitUp := strings.Split(str, ",")
	for _, v := range(splitUp) {
		val, err := strconv.Atoi(v)
		PanicIf(err)

		integers = append(integers, val)
	}

	return
}