package router

import (
  "html/template"
  "net/http"
)

type TemplateData struct {
  Title string
}

// Prepends ClientPath + adds login information
func RenderHeaderFooterTemplate(w http.ResponseWriter, r *http.Request, data map[string]string, templates ...string) {
  var clientPathTemplates []string
  for _, temp := range templates {
    clientPathTemplates = append(clientPathTemplates, serverConf.ClientPath + temp)
  }
  clientPathTemplates = append(clientPathTemplates, serverConf.ClientPath + "templates/header.html")
  clientPathTemplates = append(clientPathTemplates, serverConf.ClientPath + "templates/footer.html")
  t := template.Must(template.ParseFiles(clientPathTemplates...))

  data["Username"] = serverAuth.CurrentUser(w, r)
  t.Execute(w, data)
}

