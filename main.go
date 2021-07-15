package main

import (
	"bufio"
	imggather "imgManager/imgGather"
	"log"
	"os"
)

func main() {
	scanners := bufio.NewScanner(os.Stdin)
	scanners.Scan()
	input := scanners.Text()

	log.Println("The path is " + input)


	imggather.ImgGather(input)


}
