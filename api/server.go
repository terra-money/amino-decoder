package api

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/terra-project/core/app"
)

var cdc *codec.Codec

func init() {
	cdc = app.MakeCodec()
}

// Server represents the API server
type Server struct {
	Port      int    `json:"port"`

	Version string `yaml:"version,omitempty"`
	Commit  string `yaml:"commit,omitempty"`
	Branch  string `yaml:"branch,omitempty"`
}

// Router returns the router
func (s *Server) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/version", s.VersionHandler).Methods("GET")
	router.HandleFunc("/decode/tx", s.DecodeTxHandler).Methods("POST")

	return router
}
