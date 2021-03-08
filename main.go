package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
    var files []string

    root := "./"
	os.RemoveAll("backup/")

	filterDate := time.Now().AddDate(0, 0, -90).Format("2006-01-02")

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {

		fileStatus, _ := os.Stat(file)

		modFile := fileStatus.ModTime().Format("2006-01-02")

		if filepath.Ext(file) == ".xml" && filterDate <= modFile{


			bytesRead, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
		
			folderDst := "backup/" + filepath.Dir(file)
			if _, err := os.Stat(folderDst); os.IsNotExist(err) {

				err = os.MkdirAll(folderDst, 0755)
				if err != nil {
					log.Fatal(err)
				}

			}

			dst := "backup/" + file
			err = ioutil.WriteFile(dst, bytesRead, 0755)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Copy " + file )

		}

    }
}