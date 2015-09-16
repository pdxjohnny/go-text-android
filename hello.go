package hello

import (
	"fmt"
	"net/http"
	"html"

	// "github.com/pdxjohnny/microsocket/client"

	"github.com/pdxjohnny/go-text-android/web"
)

type SmsManager interface {
	Send(to, message string)
}

func Greetings(name string) string {
	return fmt.Sprintf("Hello again %s!", name)
}

func StartWeb(sms SmsManager) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/send", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		sendTo := "9716781233"
		sendMessage := "Hello from go server"
		sms.Send(sendTo, sendMessage)
	})
	go web.Run(mux)
}

//
// func SendUpdate(data interface{}) {
// 	conn := client.NewClient()
// 	connUrl := fmt.Sprintf("http://%s:%d/ws", "localhost", 14000)
// 	err := conn.Connect(connUrl)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// go conn.Read()
// 	conn.SendInterface(map[string]interface{}{
// 		"Method": "Incomming",
// 		"Data": data,
// 	})
// }
