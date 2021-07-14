package main

import (
	"C"
	"sync"
)
import (
	"bufio"
	"imgManager/imagesClassifier"
	"imgManager/unzipper"
	"log"
	"os"
)

func main() {
	scanners := bufio.NewScanner(os.Stdin)
	scanners.Scan()
	input := scanners.Text()

	log.Println("The path is " + input)

	var wg sync.WaitGroup
	wg.Add(1)

	go unzipper.SortZipFile(input, &wg)

	wg.Wait()

	wg.Add(1)

	imagesClassifier.FilesClassifier(input, &wg)

	wg.Wait()


}
