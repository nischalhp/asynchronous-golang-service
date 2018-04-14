package main

import (
	"net/http"

	"github.com/labstack/echo"
)

//var wg sync.WaitGroup

/*
func main() {
	//go echo(os.Stdin, os.Stdout)
	//time.Sleep(5 * time.Second)
	isProducerDone := make(chan bool)
	buffer := make(chan int)

	//wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Producing ... ", i)
			buffer <- i
			time.Sleep(1 * time.Second)
			//wg.Done()
		}

		isProducerDone <- true
	}()
	//wg.Wait()

	go func() {
		for product := range buffer {
			fmt.Println("Consuming ..", product)
		}
	}()

	<-isProducerDone
}

func echo(out io.Writer, in io.Reader) {
	io.Copy(out, in)
}
*/
//mentioning handlers

func main() {
	e := echo.New()
	e.GET("/", welcomePage)
	e.POST("/upload", DownloadFiles)
	e.Logger.Fatal(e.Start(":1323"))
}

func welcomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the matrix!")
}

func handleError(err error, context echo.Context) error {
	if err != nil {
		return context.String(http.StatusNotFound, err.Error())
	}
	return nil
}
