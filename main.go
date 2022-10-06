package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"os"
)

const (
	server     = "YOUR_SERVER:PORT"
	user       = "YOUR_USER"
	pass       = "YOUR_PASS"
	serverPath = "/YOUR_DIRECTORY_TO_PB_SS/svss/"
)

func main() {

	filename := fileList()
	fmt.Print("Found", len(filename), "files to download!\n")

	for xpng := 0; xpng < len(filename); xpng++ {
		c, err := ftp.Dial(server)
		if err != nil {
			log.Panic(err)
		}
		err = c.Login(user, pass)
		if err != nil {
			log.Panic(err)
		}
		defer c.Quit()

		filename := fileList()
		filepath := serverPath + filename[xpng]
		localPath := "downloads/" + filename[xpng]
		log.Println("Downloading: " + filename[xpng])

		// Create local file
		file, err := os.Create(localPath)
		if err != nil {
			log.Println("[MAIN] Can't create a local path:", err)
		}
		defer file.Close() // As the DEFER is inside a FOR LOOP, it will be executed after the end of the FUNCTION!
		res, err := c.Retr(filepath)
		if err != nil {
			log.Println("[MAIN] Error while retrieving file from server:", err)
			// Delete the local file on error
			err = os.Remove(localPath)
			if err != nil {
				log.Println("[MAIN] Error while deleting file:", err)
			}
			continue
		}
		// Copy the file
		defer res.Close() // As the DEFER is inside a FOR LOOP, it will be executed after the end of the FUNCTION!
		_, err = io.Copy(file, res)
		if err != nil {
			log.Println("[MAIN] Error while copying the file: ", err)
			continue
		}
		// (un)Comment this line to keep the file on server side
		//c.Delete(filepath) // PS: Looks like it always throws an error, but the file is deleted anyway.
	}
	verifyLocalFiles()
	defer DisgordMain()

}
func verifyLocalFiles() { // Verify the integrity of the files
	dir, err := os.Open("downloads/")
	if err != nil {
		log.Panic(err)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Panic(err)
	}
	for _, file := range files { // Delete files with 0 bits (corrupted)
		if file.Size() == 0 {
			fmt.Println("Deleting file: ", file.Name())
			err := os.Remove("downloads/" + file.Name())
			if err != nil {
				log.Panic("[MAIN] ", err)
			}
		}
	}
}

// Func return the file list from the server
func fileList() []string {
	c, err := ftp.Dial(server)
	if err != nil {
		log.Panic(err)
	}
	err = c.Login(user, pass)
	if err != nil {
		log.Panic(err)
	}
	defer c.Quit()

	ftpFileList, err := c.List(serverPath)
	if err != nil {
		log.Panic("[MAIN] Error on List: ", err)
	}
	if len(ftpFileList) == 0 {
		log.Panic("[MAIN] No files found")
	}
	var fileList []string
	for _, file := range ftpFileList {
		if file.Size > 1000 && file.Name != "pbsvss.htm" {
			fileList = append(fileList, file.Name)
		}
	}
	return fileList
}
