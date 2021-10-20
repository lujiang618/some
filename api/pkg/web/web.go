package web

var Web *Server

func StartWebServer() {
	Web = NewServer()
	Web.Start()
	Web.Destroy()
}
