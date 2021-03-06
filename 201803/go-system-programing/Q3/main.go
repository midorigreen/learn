package main

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	endian()
	// Q3.1
	copyFile()

	// Q3.2
	randRead()

	// Q3.3
	zipW()
}

func endian() {
	// 10000 x ビッグエンディアン
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}

func copyFile() {
	input, err := os.Open("./old.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	output, err := os.Create("./new.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	written, err := io.Copy(output, input)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("written count: %v \n", written)
}

func randRead() {
	reader := io.LimitReader(rand.Reader, 1024)
	buf := &bytes.Buffer{}
	io.Copy(buf, reader)

	// 1024 byte
	fmt.Println(buf.Len())
}

func zipW() {
	arch, err := os.Create("archive.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer arch.Close()

	zipWriter := zip.NewWriter(arch)
	defer zipWriter.Close()

	w, err := zipWriter.Create("tmp.txt")
	if err != nil {
		log.Fatalln(err)
	}
	r := strings.NewReader("Goならわかるシステムプログラミング")

	io.Copy(w, r)
}
