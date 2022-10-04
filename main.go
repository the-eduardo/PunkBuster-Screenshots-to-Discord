package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jlaffaye/ftp"
)

const (
	// FILE PATHS TO DOWNLOAD:
	//ftp://
	server     = "YOUR_SERVER:PORT"
	user       = "YOUR_USER"
	pass       = "YOUR_PASS"
	serverPath = "/YOUR_DIRECTORY_TO_PB_SS/svss/"
)

func main() {
	for xpng := 0; xpng <= 1000; xpng++ {
		c, err := ftp.Dial(server)
		if err != nil {
			log.Fatal(err)
		}
		err = c.Login(user, pass)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Quit()

		filename := fileList(c)
		filepath := serverPath + filename[xpng]
		localPath := "downloads/" + filename[xpng]
		log.Println("Downloading: " + filepath)

		// Create local file
		file, err := os.Create(localPath)
		if err != nil {
			log.Println("Can't create a local path:", err)
		}
		defer file.Close() // As the DEFER is inside a FOR LOOP, it will be executed after the end of the FUNCTION!
		res, err := c.Retr(filepath)
		if err != nil {
			log.Println("Error while retrieving file from server:", err)
			// Delete the file
			os.Remove(localPath)
			continue
		}
		// Copy the file
		defer res.Close() // As the DEFER is inside a FOR LOOP, it will be executed after the end of the FUNCTION!
		_, err = io.Copy(file, res)
		if err != nil {
			log.Println("Error while copying the file: ", err)
			continue
		}
		// Delete the file TODO: Comment this line to keep the file on server side
		/*err = c.Delete(filepath)
		if err != nil {
			log.Println("Error on delete: ", err)
			continue
		}*/
	}

}

// func file list that return the list
func fileList(c *ftp.ServerConn) []string {
	ftpFileList, err := c.List(serverPath)
	if err != nil {
		log.Fatal("Error on List: ", err)
	}
	if len(ftpFileList) == 0 {
		panic("No files found")
	}
	var fileList []string
	for _, file := range ftpFileList {
		if file.Size > 1000 && file.Name != "pbsvss.htm" {
			fileList = append(fileList, file.Name)

		}
	}
	fmt.Printf("\t Found %v files! \n\n", len(fileList))
	return fileList
}

// TODO - Delete file after download completed and delete the local file on error
