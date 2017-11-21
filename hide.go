package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

var (
	messageToLong = errors.New("Message to long!")
)

// add a bit to a byte
func setLsb(n uint8, c uint8) (r uint8) {
	i, _ := strconv.ParseInt(string(c), 10, 64)
	r = n&^1 | uint8(i)
	return
}

func hideMessage(message, dst string, img image.Image, conf *Configuration) error {
	dstW, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer dstW.Close()

	message = conf.Encryption.Encrypt(message)
	bMessage := bytes2Binary([]byte(fmt.Sprintf("%d%s", len(message), message)))
	bMessageLength := len(bMessage)
	index := 0

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	if bMessageLength > (width * height * 3) {
		return messageToLong
	}

	nRGBA := toRGBA(img)
	i := 3

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := nRGBA.RGBAAt(x, y) // image color
			cr := color.RGBA{}      // result color
			cr.A = c.A

			for index+i > bMessageLength {
				i--
			}

			if index+i <= bMessageLength {
				bits := bMessage[index : index+i] // getting first `i` bits
				if i >= 1 {
					cr.R = setLsb(c.R, bits[0])
				}

				if i >= 2 {
					cr.G = setLsb(c.G, bits[1])
				}

				if i == 3 {
					cr.B = setLsb(c.B, bits[2])
				}
			}

			nRGBA.SetRGBA(x, y, cr) // set new pixel color to RGBA object
			index += i
		}
	}

	err = png.Encode(dstW, nRGBA)
	if err != nil {
		return err
	}

	return nil
}
