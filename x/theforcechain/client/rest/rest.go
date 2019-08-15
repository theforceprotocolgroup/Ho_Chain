package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/theforceprotocolgroup/theforcechain/x/theforcechain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/gorilla/mux"
)

const (
	restName = "id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeID string) {
	r.HandleFunc(fmt.Sprintf("/%s/theforcechain", storeID), idsHandler(cliCtx, storeID)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/theforcechain", storeID), setOrderHandler(cliCtx)).Methods("PUT")
}

// --------------------------------------------------------------------------------------
// Tx Handler

type setOrderReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Id        string       `json:"id"`
	Borrower  string       `json:"borrower"`
	Lender    string       `json:"lender"`
	TokenGet  sdk.Coin     `json:"tokenGet"`
	TokenGive sdk.Coin     `json:"tokenGive"`
	Owner     string       `json:"owner"`
}

func setOrderHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setOrderReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgSetOrder(req.Id, req.Borrower, req.Lender, req.TokenGet, req.TokenGive, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

//--------------------------------------------------------------------------------------
// Query Handlers

func idsHandler(cliCtx context.CLIContext, storeID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/ids", storeID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
