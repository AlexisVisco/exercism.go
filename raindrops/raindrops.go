package raindrops

import (
	"bytes"
	"strconv"
)

func IsFactorOf(number float64, factor float64) bool {
	a := factor / number
	b := number / factor
	return b == float64(int64(b)) || a == float64(int64(a))
}

func AppendBuffer(floatVal float64, buffer *bytes.Buffer) {
	if IsFactorOf(float64(floatVal), 3.0) {
		buffer.WriteString("Pling")
	}
	if IsFactorOf(floatVal, 5.0) {
		buffer.WriteString("Plang")
	}
	if IsFactorOf(floatVal, 7.0) {
		buffer.WriteString("Plong")
	}
}

func Convert(number int) string {
	var buffer bytes.Buffer
	floatVal := float64(number)
	if number == 1 {
		return "1"
	}
	AppendBuffer(floatVal, &buffer)
	if buffer.Len() > 0 {
		return buffer.String()
	}
	return strconv.Itoa(number)
}
