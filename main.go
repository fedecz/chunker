package main

import (
	"fmt"
	qr "github.com/skip2/go-qrcode"
	"io"
	"log"
	"os"
	filepath "path/filepath"
)

const chunkSize = 1 * 1001 //1024

func generateQRCode(chunkBytes []byte, filename string) error {
	// Generate a random seed for the QR code
	//rand.Seed(time.Now().UnixNano())

	err := qr.WriteFile(string(chunkBytes), qr.Highest, 1024, filename)
	if err != nil {
		return err
	}

	// Encode the QR code in hex format
	//encodedQRCode := hex.EncodeToString(qrCode)

	// Save the QR code to a file
	//f, err := os.Create(filepath.Join("qr-codes", filename))
	//if err != nil {
	//	return err
	//}
	//defer f.Close()
	//
	//_, err = f.Write(encodedQRCode)
	//if err != nil {
	//	return err
	//}

	log.Println("QR code saved to", filename)

	return nil
}

func splitFileIntoChunks(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var i = 0
	for {
		chunk := make([]byte, chunkSize)
		var shouldBreak = false
		read, err := file.Read(chunk)
		if err == io.EOF {
			shouldBreak = true
		}
		//fmt.Printf("read %d\n", read)
		if read == 0 {
			break
		}

		//_, err := file.Readn(int(float64(len(chunk)) - int64(len(chunk)/chunkSize*chunkSize)))
		//if err != nil {
		//		return nil, err
		//	}
		//chunk = append(chunk[:len(chunk)-int(n)], []byte{0}...)

		//yieldChunk := chunk[:n]
		// Generate QR code and save it to disk
		err = generateQRCode(chunk, fmt.Sprintf("%s-%02d.png", filepath.Base(filePath), i))
		if err != nil {
			return err
		}
		i++
		if shouldBreak {
			break
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("missing argument")
	}
	filePath := os.Args[1]
	err := splitFileIntoChunks(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//f, err := os.Open("image.base64-0.png")
	//f, err := os.Open("/home/fede/Downloads/test.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//result, err := readqr.Decode(f)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)

	fmt.Println("File has been successfully split into chunks and QR codes generated.")
}
