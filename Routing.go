package main

import "log"

func (s *ApiServer) Run() {
	// GET /api/register
	s.App.Get("/", s.GetInfo)

	// GET /flights/LAX-SFO
	log.Fatal(s.App.Listen(s.Port))
}
