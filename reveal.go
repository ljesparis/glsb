package main

import (
	"errors"
	"image"
	"strconv"
)

var (
	messageNotFound = errors.New("Message Not Found =)")
)

// Extract message from image's pixels.
// by default utf8 encoding is used.
func revealMessage(img image.Image, conf *Configuration) (buff string, err error) {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	rgbaObject := toRGBA(img)

	var msgLen int = 0        // message length
	var msgLenAsString string // message length as string(need to be parsed to int)

	var cbuff int = 0 // buffer that host a char extracted bit by bit from pixels
	var n uint8 = 0   // used to decode `cbuff`

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cRGBA := rgbaObject.RGBAAt(x, y) // getting rgba color from `x` and `y` position
			for _, c := range []uint8{cRGBA.R, cRGBA.G, cRGBA.B} {
				lsb := c & 1                 // getting least significant bit
				cbuff += int(lsb << (7 - n)) // decoding `lsb` bit using utf8 encoding as default
				if n == 8 {
					b := byte(cbuff) // casting integer to byte

					// finding message length
					if b > 47 && b < 58 && len(buff) == 0 {
						msgLenAsString += string(b)
					} else {
						msgLen, _ = strconv.Atoi(msgLenAsString) // parsing message length to integer

						// if buffer already reached message length
						// let's go to end
						if len(buff) == msgLen {
							goto end
						}

						buff += string(b)
					}

					// let's restart `n` and `cbuff` to start
					// another char finding and decoding.
					n, cbuff = 0, 0
				}

				n++
			}
		}
	}

end:
	if msgLen > 0 {
		buff = conf.Encryption.Decrypt(buff)
	} else {
		buff = ""
		err = messageNotFound
	}

	return
}
