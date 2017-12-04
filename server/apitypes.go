package main

type Version struct {
	Version string
}

type ApiError struct {
	Error string `json:"error"`
}

type LstInterfaces struct {
	Intrefaces []string `json:"interfaces"`
}

type DetailInterface struct {
	Name      string   `json:"name"`
	Hw_addr   string   `json:"hw_addr"`
	Inet_addr []string `json:"inet_addr"`
	MTU       int      `json:"MTU"`
}
