package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"io/ioutil"
	"net/http"
)

type DecodeRequestReq struct {
	AminoEncodedTx string `json:"amino_encoded_tx"`
}

// Marshal - nolint
func (sb DecodeRequestReq) Marshal() []byte {
	out, err := json.Marshal(sb)
	if err != nil {
		panic(err)
	}
	return out
}

// DecodeHandler handles the /decode route
func (s *Server) DecodeTxHandler(w http.ResponseWriter, r *http.Request) {
	var req DecodeRequestReq

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cdc.UnmarshalJSON(body, &req)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	txBytes, err := base64.StdEncoding.DecodeString(req.AminoEncodedTx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var stdTx auth.StdTx
	err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bz, err := cdc.MarshalJSON(stdTx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bz)
}
