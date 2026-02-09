package main

import "fmt"

type fahrenheit float64
type celcius float64

func (c celcius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32)
}

func (f fahrenheit) celcius() celcius {
	return celcius(f-32) / (9 / 5)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

func drawTable(hdr1 string, hdr2 string, rows int, calculation func(int) (string, string)) {
	fmt.Printf("%s\n", line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Printf("%s\n", line)
	for row := range rows {
		s1, s2 := calculation(row)
		fmt.Printf(rowFormat, s1, s2)
	}
	fmt.Printf("%s\n", line)
}

func ctof(row int) (string, string) {
	c := celcius(row*5 - 40)
	f := c.fahrenheit()
	s1 := fmt.Sprintf("%.1f", c)
	s2 := fmt.Sprintf("%.1f", f)
	return s1, s2
}

func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celcius()
	s1 := fmt.Sprintf("%.1f", f)
	s2 := fmt.Sprintf("%.1f", c)
	return s1, s2
}

func main() {
	rows := 29
	drawTable("°C", "°F", rows, ctof)
	drawTable("°F", "°C", rows, ftoc)
}
