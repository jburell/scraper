package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
//	"sync"
)

func main(){
	//var wg sync.WaitGroup

	startNum := 18
	endNum := 3405
	//numURLs := endNum - startNum

	//wg.Add(numURLs)
	for i := startNum; i < endNum; i += 1 {
		/*go*/ get_image(i)
	}
	//wg.Wait()
}

func get_image(nr int){
	
	imgURL := fmt.Sprintf("%s%05d%s", "http://2015.revision-party.net/media/photowall/photobooth_display/thumbs/big_Capture", nr, ".jpg")
	fmt.Println(imgURL)

	resp, err := http.Get(imgURL)
	if err != nil {
		fmt.Println(err)
	}else{
		// open output file
		outDir := "imgs" 
		err := os.MkdirAll(outDir, 0644)
		
		if err != nil {
			panic(err)
		}

	    fo, err := os.Create(fmt.Sprintf("%s%05d%s", fmt.Sprintf("%s%s",outDir, "/photowall_"), nr,".jpg"))
	    if err != nil {
	        panic(err)
	    }

	    // close fo on exit and check for its returned error
	    defer func() {
	        if err := fo.Close(); err != nil {
	            panic(err)
	        }
	    }()

		buf := make([]byte, 1024)
		for {
			// read a chunk
			n, err := resp.Body.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}

			// write a chunk
			if _, err := fo.Write(buf[:n]); err != nil {
				panic(err)
			}
		}
	}
}
