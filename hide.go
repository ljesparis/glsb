package main

import (
	"errors"
	"fmt"
	"image"
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

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := nRGBA.RGBAAt(x, y) // image color

			if index+1 <= bMessageLength {
				c.R = setLsb(c.R, bMessage[index])
			}

			if index+2 <= bMessageLength {
				c.G = setLsb(c.G, bMessage[index+1])
			}

			if index+3 <= bMessageLength {
				c.B = setLsb(c.B, bMessage[index+2])
			}

			nRGBA.SetRGBA(x, y, c) // set new pixel color to RGBA object
			index += 3
		}
	}

	err = png.Encode(dstW, nRGBA)
	if err != nil {
		return err
	}

	return nil
}
