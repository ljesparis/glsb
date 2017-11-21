glsb
-

It's a tool to hide and read messages to/from images,
applying least significant bit(lsb) technique.

Installation
-

First of all clone the tool's source from the
master repository
```bash
git clone https://github.com/ljesparis/glsb.git
```

Lastly go to glsb's folder and compile.
```bash
cd $GOPATH/src/github.com/ljesparis/glsb
go build .
```

Done!!

Usage
-

#### Without enccryption
Write a message within an image
```bash
glsb write /path/to/image/src.png dst.png "Hello world"
```

Read hiden message
```bash
glsb read /path/to/image/dst.png
```

#### With enccryption
Write a message within an image
```bash
glsb --encryption xor --key "this is my key" write /path/to/image/src.png dst.png "Hello world"
```

Read hiden message
```bash
glsb --encryption xor --key "this is my key" read /path/to/image/dst.png
```

Formats Supported
-
 ### Image
 - [x] [PNG](https://en.wikipedia.org/wiki/Portable_Network_Graphics)
 - [ ] [JPEG](https://en.wikipedia.org/wiki/JPEG)
 - [ ] [GIF](https://en.wikipedia.org/wiki/GIF)
 - [ ] [BMP](https://en.wikipedia.org/wiki/BMP_file_format)

### Sound
 - [ ] [MP3](https://en.wikipedia.org/wiki/MP3)
 - [ ] [WAV](https://en.wikipedia.org/wiki/WAV)
 
Encryption methods supported
-
 - [x] [XOR](https://en.wikipedia.org/wiki/XOR_cipher)
 - [ ] [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)
 - [ ] [BLOWFISH](https://en.wikipedia.org/wiki/Blowfish_%28cipher%29)
 - [ ] [TWOFISH](https://en.wikipedia.org/wiki/Twofish)

License
-
[MIT](./LICENSE) license