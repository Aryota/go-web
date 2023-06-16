package main

type NumberData struct {
	Number int
	First  int
	Second int
	Third  int
}

type TableData struct {
	Numbers []NumberData
}

func (td *TableData) Reset() {
	td.Numbers = []NumberData{}
}
