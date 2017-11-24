package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/ljesparis/glsb/encryption"

	"github.com/urfave/cli"
)

var (
	// default configuration
	conf = defaultConfig
)

func throwError(m string) {
	fmt.Println(m)
	os.Exit(1)
}

func panicIfErr(err error) {
	if err != nil {
		throwError(err.Error())
	}
}

// verify image format
// only png is supported.
func checkImageFormat(f string) error {
	ext := filepath.Ext(f)[1:]
	if ext == "png" {
		return nil
	}

	return fmt.Errorf("Image format `%s` not supported", ext)
}

// open png file if exist
func openPngFile(f string) (image.Image, error) {
	ff, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	defer ff.Close()

	i, err := png.Decode(ff)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func configureApp(c *cli.Context) {
	if c.GlobalIsSet("encryption") {
		conf.Encryption.Key = c.GlobalString("key")
		e := c.GlobalString("encryption")
		switch e {
		case "xor":
			conf.Encryption.Method = encryption.Xor
			break
		}
	}
}

func read(c *cli.Context) {
	if c.NArg() < 1 {
		cli.ShowSubcommandHelp(c)
		return
	}

	configureApp(c)
	src := c.Args().Get(0)
	err := checkImageFormat(src)
	panicIfErr(err)
	i, err := openPngFile(src)
	panicIfErr(err)
	message, err := revealMessage(i, conf)
	panicIfErr(err)
	messageLength := len(message)
	fmt.Printf(`Message Found! =).
Message Length: %d
Message: %s
`, messageLength, message)
}

func write(c *cli.Context) {
	if c.NArg() < 3 {
		cli.ShowSubcommandHelp(c)
		return
	}

	configureApp(c)
	args := c.Args()
	src, dst, m := args.Get(0), args.Get(1), args.Get(2)
	err := checkImageFormat(src)
	panicIfErr(err)
	err = checkImageFormat(dst)
	panicIfErr(err)
	i, err := openPngFile(src)
	panicIfErr(err)
	err = hideMessage(m, dst, i, conf)
	panicIfErr(err)
	fmt.Println("Message successfully hidden")
}

func methods(_ *cli.Context) {
	fmt.Println("\nEncryption Methods Supported:")
	for _, m := range []string{"xor"} {
		fmt.Println("    *", m)
	}

	fmt.Println("")
}

func main() {
	app := cli.NewApp()
	app.Author = author
	app.Email = email
	app.Version = version
	app.Description = description
	app.Usage = usage
	app.UsageText = usageText
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "encryption",
			Value: "None",
		},
		cli.StringFlag{
			Name:  "key",
			Value: "hide message with glsb!",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "write",
			Usage:     "Used to hide messages within images",
			UsageText: "glsb [options...] write [src] [dst] [message]",
			Action:    write,
		},
		{
			Name:      "read",
			Usage:     "Used to read hidden messages that are placed within images",
			UsageText: "glsb [options...] read [src]",
			Action:    read,
		},
		{
			Name:      "methods",
			Usage:     "Used to list encryption methods supported by glsb",
			UsageText: "glsb methods",
			Action:    methods,
		},
	}

	app.Run(os.Args)
}
