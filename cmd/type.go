package cmd

type Service struct {
	name        string
	host        string
	image       string
	port        int
	entrypoints string
	tls         bool
	network     string
}

func newService() *Service {
	return &Service{
		name:        "service",
		host:        "Host(`example.com`)",
		image:       "example/example:latest",
		port:        80,
		entrypoints: "websecure",
		tls:         true,
		network:     "traefik-global-proxy",
	}
}
