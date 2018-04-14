package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

/*
DownloadFiles
This method is used to download files from the echo request
*/
func DownloadFiles(context echo.Context) error {
	print("hell")
	forms, err := context.MultipartForm()
	if err != nil {
		return context.String(http.StatusNotFound, err.Error())
	}

	for file := range forms.File {
		processedFile := make(chan bool)
		println(file)
		go handleFiles(file, context, processedFile)
		<-processedFile
	}
	return context.String(http.StatusOK, "All processed")
}

func handleFiles(file string, context echo.Context, processedFileStatus chan bool) error {
	content, err := context.FormFile(file)
	var er = handleError(err, context)
	if er != nil {
		return er
	}

	src, err := content.Open()
	var openEr = handleError(err, context)
	if openEr != nil {
		return openEr
	}
	defer src.Close()

	dest, err := os.Create("/Users/nischal/code/go/" + content.Filename)
	var createEr = handleError(err, context)
	if createEr != nil {
		return createEr
	}

	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		return handleError(err, context)
	}

	processedFileStatus <- true
	return nil

}
