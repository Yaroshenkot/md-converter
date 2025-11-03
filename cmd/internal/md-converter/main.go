package main

import (
	"converter/converter"
	"fmt"
	"io"
	"os"
)

func ConvertFile(inputPath, outputPath string) error {
	fileBefore, err := os.Open(inputPath)

	if err != nil {
		fmt.Println("err", err)

	}
	defer fileBefore.Close()

	var user string
	data := make([]byte, 1024)

	for {
		n, err := fileBefore.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ошибка", err)
		}
		user += string(data[:n])

	}
	for _, value := range data {
		user = user + string(value)
	}
	//fmt.Print(user)
	new := converter.ConvertMDToHTML(user)

	fileAfter, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("!!!", err)

	}
	defer fileAfter.Close()

	new2 := []byte(new)

	fileAfter.Write(new2)
	if err != nil {
		fmt.Println("ошибка", err)
	}

	return nil
}
func main() {
	inputPath := "markdown.md"
	outputPath := "HTML.html"

	err := ConvertFile(inputPath, outputPath)
	if err != nil {
		fmt.Printf("ошибка!!!")
		os.Exit(1)
	}
	fmt.Println("Конвертация завершена")
}
