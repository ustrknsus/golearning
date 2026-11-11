// Structs, pointers, and methods
// Exercise: Add a Stop method and a port field.
package main

import "fmt"

type Server struct {
	Name    string
	Running bool
}

func (s *Server) Start()        { s.Running = true }
func (s Server) Status() string { return fmt.Sprintf("%s running=%t", s.Name, s.Running) }
func main() {
	server := Server{Name: "lab-01"}
	server.Start()
	fmt.Println(server.Status())
}
