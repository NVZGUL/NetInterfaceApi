package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func InternalServerError(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ApiError{Error: "error message"})
}

func NotFound(w http.ResponseWriter, req *http.Request, name string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ApiError{Error: fmt.Sprintf("interface %s was not found", name)})
}

const V = "v1"

func GetVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Version{Version: V})
}

func GetInterfaces(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if params["api_version"] != V {
		InternalServerError(w, req)
		return
	}
	l, err := net.Interfaces()
	if err != nil {
		panic(err)

	}
	var netInterfaces []string
	for _, f := range l {
		netInterfaces = append(netInterfaces, f.Name)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LstInterfaces{Intrefaces: netInterfaces})
}

func GetDetailInterface(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if params["api_version"] != V {
		InternalServerError(w, req)
		return
	}
	l, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	var intr net.Interface
	var addrLst []string
	exist := false
	for _, i := range l {
		if params["i_name"] == i.Name {
			intr = i
			addr, err := intr.Addrs()
			if err != nil {
				panic(err)
			}
			for _, a := range addr {
				fmt.Println(a)
				addrLst = append(addrLst, a.String())
			}
			exist = true
			break
		}
	}

	if exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(DetailInterface{
			Name:      intr.Name,
			Hw_addr:   intr.HardwareAddr.String(),
			Inet_addr: addrLst,
			MTU:       intr.MTU,
		})
	} else {
		NotFound(w, req, params["i_name"])
	}

}

func Handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/service/version", GetVersion).Methods("GET")
	router.HandleFunc("/service/{api_version}/interfaces", GetInterfaces).Methods("GET")
	router.HandleFunc("/service/{api_version}/interfaces/{i_name}", GetDetailInterface).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(InternalServerError)

	return router
}
