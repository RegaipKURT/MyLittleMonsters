package main

// Bu program tldr ve cheat.sh kullanarak aranan bir kod örneğini bulmaya yarar.
// orjinal programlar
// cheat.sh: https://github.com/chubin/cheat.sh
// tldr: https://github.com/tldr-pages/tldr

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) <= 1 {
		cmd := exec.Command("curl", "cht.sh")
		res, err := cmd.Output()
		if err != nil {
			fmt.Println("cheat.sh error: ", err.Error())
		}
		fmt.Println(string(res))

		cmd2 := exec.Command("tldr", "-h")
		res, err = cmd2.Output()
		if err != nil {
			fmt.Println("tldr error:", err.Error())
		}
		fmt.Println(string(res))

		fmt.Println("BU PROGRAM YUKARIDA GÖRÜLEN cht.sh ve tldr PROGRAMLARI ÜZERİNDEN ARAMA YAPAR.\ntldr komutunun yüklenmiş olması gerekir.")
		fmt.Println("En az bir programlama dili veya parametre belirtin!")
		fmt.Println("Kullanım (1): kodbul <programlama dili> <aranacak ifadeler>")
		fmt.Println("Kullanım (2): kodbul <bash girdisi>")

		os.Exit(0)
	}

	var args string
	var url string

	if len(os.Args) == 2 {
		if os.Args[1] != "-h" {
			cmd := exec.Command("tldr", os.Args[1])
			res, err := cmd.Output()
			if err != nil {
				fmt.Println("tldr error: ", err.Error())
			}
			fmt.Println(string(res))
		} else {
			fmt.Println("BU PROGRAM cht.sh ve tldr PROGRAMLARI ÜZERİNDEN ARAMA YAPAR.\ntldr komutunun yüklenmiş olması gerekir.")
			fmt.Println("\nKullanım (1): kodbul <programlama dili> <aranacak ifadeler>")
			fmt.Println("Kullanım (2): kodbul <bash girdisi>")
			os.Exit(0)
		}
	}

	for _, w := range os.Args[2:] {
		args += w + " "
	}

	args, url = "", ""

	for _, w := range os.Args[2:] {
		args += w + "+"
	}

	if len(args) > 0 {
		url = "cht.sh/" + os.Args[1] + "/" + args[0:len(args)-1]
	} else {
		url = "cht.sh/" + os.Args[1]
	}

	cmd := exec.Command("curl", url)
	res, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
