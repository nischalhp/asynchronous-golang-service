package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"

	"github.com/labstack/echo"
)

/*
Method to handle the post requests
*/
func HandlePostRequest(context echo.Context) error {
	fmt.Println("Processing post request")
	requestToken, err := handleFormData(context)
	fmt.Println("finished processing form data method")
	if err != nil {
		return handleErrorFunctionEchoContext(context, err)
	}
	return context.String(http.StatusOK, requestToken)
}

func handleErrorFunctionEchoContext(context echo.Context, err error) error {
	return context.String(http.StatusInternalServerError, err.Error())
}

/*
DownloadFiles
This method is used to download files from the echo request
*/
func handleFormData(context echo.Context) (string, error) {
	fmt.Println("Handling form data")
	forms, err := context.MultipartForm()
	if err != nil {
		return "", err
	}
	request_token := "token"
	go downloadFiles(context, forms, request_token)
	fmt.Println("sending response back")
	return request_token, nil
}

func downloadFiles(context echo.Context, forms *multipart.Form, request_token string) {
	fmt.Println("Handling files")
	var wg sync.WaitGroup
	for file := range forms.File {
		fmt.Println("Adding to waitgroup")
		wg.Add(1)
		go handleFiles(file, context, &wg)
	}
	wg.Wait()
	// update request token table with status
}

func handleFiles(file string, context echo.Context, wg *sync.WaitGroup) {

	// waiting for the go routine to finish
	defer wg.Done()

	content, err := context.FormFile(file)
	var er = handleError(err, context)
	if er != nil {
		handleError(er, context)
	}
	src, err := content.Open()
	var openEr = handleError(err, context)
	if openEr != nil {
		handleError(openEr, context)
	}
	defer src.Close()

	dest, err := os.Create("/Users/nischal/code/go/test-files/" + content.Filename)
	var createEr = handleError(err, context)
	if createEr != nil {
		handleError(createEr, context)
	}

	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		handleError(err, context)
	}
}
