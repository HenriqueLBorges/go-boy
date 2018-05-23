package main

import (
	"bufio"
	"fmt"
	"os"
)

type gameboy struct {
	memory         [4096]byte
	a              byte
	b              byte
	c              byte
	d              byte
	e              byte
	h              byte
	l              byte
	programCounter uint16  //Program counter
	stackPointer   uint16  //Stack pointer
	flags          [4]byte //flag register | 0 - Zero Flag | 1 - Subtract Flag | 2 - Half Carry Flag | 3 - Carry Flag
}

type rom struct {
	memory []byte
}

//Rom functions
func (r rom) load(path string) {
	rom, err := os.Open(path)
	defer rom.Close()

	if err != nil {
		fmt.Println("Can't load ROM. Err =", err)
	} else {
		fileInfo, err := rom.Stat()
		if err != nil {
			fmt.Println("Can't descript file. Err =", err)
		} else {
			r.memory = make([]byte, fileInfo.Size())
			buffer := bufio.NewReader(rom)
			_, err = buffer.Read(r.memory)
			if err != nil {
				fmt.Printf("Can't read ROM data. err =", err)
			} else {
				fmt.Println(r.memory)
			}
		}
	}
}

//Helpers
//Get file size
func descriptFile(file *os.File) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Can't descript file. Err =", err)
	} else {
		fmt.Println("File Description")
		fmt.Println("Name =", fileInfo.Name())
		fmt.Println("Size =", fileInfo.Size(), "bytes")
		fmt.Println("Mode =", fileInfo.Size())
		fmt.Println("ModTime =", fileInfo.ModTime())
	}
}

func main() {
	path := os.Args[1]
	//var gb gameboy
	var game rom
	game.load(path)
	//fmt.Println("Memory = ", gb.memory)
}
