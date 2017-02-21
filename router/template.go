package router

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Title         string
	ShowLogin     string
	WebsocketAddr string
	RoomID        string
	Username      string
	// demo
	MyRooms      []string
	SharedRooms  []string
	Participants []string
	MyImages     []string
	SharedImages []string
}

// Prepends ClientPath + adds login information
func RenderHeaderFooterTemplate(w http.ResponseWriter, r *http.Request, data Data, templates ...string) {
	var clientPathTemplates []string
	for _, temp := range templates {
		clientPathTemplates = append(clientPathTemplates, serverConf.ClientPath+temp)
	}
	clientPathTemplates = append(clientPathTemplates, serverConf.ClientPath+"templates/header.html")
	clientPathTemplates = append(clientPathTemplates, serverConf.ClientPath+"templates/footer.html")
	clientPathTemplates = append(clientPathTemplates, serverConf.ClientPath+"templates/sidebar.html")
	t := template.Must(template.ParseFiles(clientPathTemplates...))

	user, err := serverAuth.GetCurrentUser(w, r)
	if err == nil && user != "" {
		data.Username = user
		data.MyRooms = []string{"41", "114", "153", "342", "644", "935"}
		data.SharedRooms = []string{"53", "44", "21", "15", "3"}
		data.Participants = []string{"0", "0", "0", "1", "0", "0"}
		data.MyImages = []string{"http://puu.sh/ud14z/619fe0a778.jpg", "http://puu.sh/ud11r/cbbfdb5649.jpg", "http://puu.sh/ud12c/1b41864c9b.jpg", "http://puu.sh/ud12r/68a7b11968.jpg", "http://puu.sh/ud130/340d662945.jpg", "http://puu.sh/ud164/168dfa1b94.jpg"}
		data.SharedImages = []string{"http://puu.sh/ud178/ee6837c2c1.jpg", "http://puu.sh/ud17w/843c2ea585.jpg", "http://puu.sh/ud17Z/7d3738613c.jpg", "http://puu.sh/ud18r/d02816e253.jpg", "http://puu.sh/ud18Z/df37c60d67.jpg"}
	}
	fmt.Printf("TEMPLATE DATA: %+v\n", data)
	t.Execute(w, data)
}
