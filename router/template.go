package router

import (
	"html/template"
	"net/http"
)

type Data struct {
  Title string
  ShowLogin string
  WebsocketAddr string
  RoomId string
  Username string
  Rooms []string
  Participants []string
  VideoId string
}
// Prepends ClientPath + adds login information
func (s ServerConfig) RenderHeaderFooterTemplate(
  w http.ResponseWriter,
  r *http.Request,
  data Data,
  templates ...string,
) {
	var clientPathTemplates []string
	for _, temp := range templates {
		clientPathTemplates = append(clientPathTemplates, s.config.ClientPath+temp)
	}
	clientPathTemplates = append(clientPathTemplates, s.config.ClientPath+"templates/header.html")
	clientPathTemplates = append(clientPathTemplates, s.config.ClientPath+"templates/footer.html")
	clientPathTemplates = append(clientPathTemplates, s.config.ClientPath+"templates/sidebar.html")
	t := template.Must(template.ParseFiles(clientPathTemplates...))

	user, err := s.auth.GetCurrentUser(w, r)
	if err == nil && user != "" {
    data.Username = user
    data.Rooms = []string{"121", "240", "320", "520", "555", "677"}
    data.Participants = []string{"0", "0", "0", "1", "0", "0"}
	}
	t.Execute(w, data)
}
