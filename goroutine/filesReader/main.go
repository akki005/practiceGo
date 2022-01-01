package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	quit := make(chan bool, 3)

	filePaths := []string{
		"/home/akash/InterviewPrep/LanguagePractice/GO/goroutine/filesReader/file1.txt",
		"/home/akash/InterviewPrep/LanguagePractice/GO/goroutine/filesReader/file2.txt",
		"/home/akash/InterviewPrep/LanguagePractice/GO/goroutine/filesReader/file3.txt",
	}

	filesDataChs := []chan string{}
	errDataChs := []chan error{}

	for _, path := range filePaths {
		file, errCh := fileReader(path)
		filesDataChs = append(filesDataChs, file)
		errDataChs = append(errDataChs, errCh)
	}

	allFilesdata := []string{}

	readerFn := func(fileDataCh chan string, errCh chan error) {
		for {
			select {
			case data := <-fileDataCh:
				log.Println("data", data)
				allFilesdata = append(allFilesdata, data)
			case err := <-errCh:
				log.Println("err", err)
				quit <- true
				return
			}
		}
	}

	for index, filesDataCh := range filesDataChs {
		go readerFn(filesDataCh, errDataChs[index])
	}

	for i := 0; i < 3; i++ {
		<-quit
	}
	fmt.Println("done reading")
	fmt.Println(allFilesdata)
}

func fileReader(filePath string) (chan string, chan error) {

	chStr := make(chan string)
	errCh := make(chan error)

	go func() {

		defer func() {
			close(chStr)
			close(errCh)
		}()

		file, err := os.Open(filePath)

		if err != nil {
			errCh <- err
			return
		}

		defer file.Close()

		// Start reading from the file using a scanner.
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			chStr <- line
		}

		if scanner.Err() != nil {
			errCh <- scanner.Err()
			return
		}

	}()

	return chStr, errCh
}
