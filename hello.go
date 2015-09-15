package hello

import (
	"fmt"

	"github.com/pdxjohnny/go-text-android/web"
)

func Greetings(name string) string {
	return fmt.Sprintf("Hello again %s!", name)
}

func StartWeb() {
	go web.Run()
}
