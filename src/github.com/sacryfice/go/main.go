package main

import (
	"github.com/disintegration/imaging"
	"image/jpeg"
	"io/ioutil"
	"strings"
	"fmt"
	"os"
	"log"
)

func findJpeg() (jpgArr []string) {

	//читаем файлы в директории
	files, _ := ioutil.ReadDir("./")

	//файлы, имеющие расширение jpg выносим в отдельный массив
    for _, f := range files {
    		split := strings.Split(f.Name(), ".")
    		ext := split[len(split) - 1];

    		if (ext == "jpg") {
            	jpgArr = append(jpgArr, f.Name())
    		}
    }

    return
}

func main() {
	jpgArr := findJpeg()
	for _, fileName := range jpgArr {
	    file, err := os.Open(fileName)
	    if err != nil {
	        log.Fatal(err)
	    }

	    fmt.Println(file)

	    // decode jpeg into image.Image
	    img, err := jpeg.Decode(file)
	    if err != nil {
	        log.Fatal(err)
	    }
	    file.Close()

	    // and preserve aspect ratio
	    //m := resize.Resize(100, 0, img, resize.Lanczos3)
        m := imaging.FlipH(img)

	    out, err := os.Create("_" + fileName)
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer out.Close()

	    // write new image to file
	    jpeg.Encode(out, m, nil)
	}
}

/*package main

import (
    "github.com/nfnt/resize"
    "image/jpeg"
    "log"
    "os"
    "fmt"
)

func main() {
    // open "test.jpg"
    file, err := os.Open("1.jpg")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(file)

    // decode jpeg into image.Image
    img, err := jpeg.Decode(file)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    // and preserve aspect ratio
    m := resize.Resize(100, 0, img, resize.Lanczos3)

    out, err := os.Create("test_resized.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // write new image to file
    jpeg.Encode(out, m, nil)
}*/