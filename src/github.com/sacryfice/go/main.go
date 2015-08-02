package main

import (
	//"reflect"
	"image/png"
	"image/color"
	//"fmt"
	"os"
	"log"
)

type ImageSet interface {
	Set(x, y int, c color.Color)
}

func main() {
		//открываем файл
	    file, err := os.Open("1.png")
	    if err != nil {
	        log.Fatal(err)
	    }

	    //конвертируем в картинку
	    img, err := png.Decode(file)
	    if err != nil {
	        log.Fatal(err)
	    }

	    //функция для установки пикселя
	    imgSet := img.(ImageSet)

	    //цикл по картинке
    	size := img.Bounds().Size()
    	for y := 0; y < size.Y; y++ {
	        for x := 0; x < size.X; x++ {
	        	c := img.At(x, y);
	        	_, _, _, a := c.RGBA()

	        	if (a < 65535) {
	        		pixel := color.RGBA{
	        			255, 
	        			0, 
	        			0, 
	        			255,
	        		}
	        		imgSet.Set(x, y, pixel)
	        	}
	        }
    	}

    	//создаём новый файл
		fd, err := os.Create("./_1.png")
		if err != nil {
			log.Fatal(err)
		}

		//декодируем изображение в новый файл
		err = png.Encode(fd, img)
		if err != nil {
			log.Fatal(err)
		}
}