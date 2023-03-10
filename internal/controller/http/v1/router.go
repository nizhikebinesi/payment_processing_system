// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /balances/{idFrom}/transfer/{idTo})
	TransferByIds(ctx echo.Context, idFrom int64, idTo int64) error

	// (GET /balances/{id})
	GetBalanceById(ctx echo.Context, id int64, params GetBalanceByIdParams) error

	// (POST /balances/{id})
	AccrueOrWriteOffBalance(ctx echo.Context, id int64) error

	// (GET /balances/{id}/transactions)
	GetBindedTransactions(ctx echo.Context, id int64, params GetBindedTransactionsParams) error

	// (POST /reservation/balances/{id})
	ReserveOnSeparateAccount(ctx echo.Context, id int64) error

	// (POST /reservation/balances/{id}/cancel)
	CancelReservation(ctx echo.Context, id int64) error

	// (POST /reservation/balances/{id}/confirm)
	ConfirmReservation(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// TransferByIds converts echo context to params.
func (w *ServerInterfaceWrapper) TransferByIds(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "idFrom" -------------
	var idFrom int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "idFrom", runtime.ParamLocationPath, ctx.Param("idFrom"), &idFrom)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter idFrom: %s", err))
	}

	// ------------- Path parameter "idTo" -------------
	var idTo int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "idTo", runtime.ParamLocationPath, ctx.Param("idTo"), &idTo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter idTo: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransferByIds(ctx, idFrom, idTo)
	return err
}

// GetBalanceById converts echo context to params.
func (w *ServerInterfaceWrapper) GetBalanceById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBalanceByIdParams
	// ------------- Optional query parameter "currency" -------------

	err = runtime.BindQueryParameter("form", true, false, "currency", ctx.QueryParams(), &params.Currency)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter currency: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBalanceById(ctx, id, params)
	return err
}

// AccrueOrWriteOffBalance converts echo context to params.
func (w *ServerInterfaceWrapper) AccrueOrWriteOffBalance(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccrueOrWriteOffBalance(ctx, id)
	return err
}

// GetBindedTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetBindedTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBindedTransactionsParams
	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBindedTransactions(ctx, id, params)
	return err
}

// ReserveOnSeparateAccount converts echo context to params.
func (w *ServerInterfaceWrapper) ReserveOnSeparateAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ReserveOnSeparateAccount(ctx, id)
	return err
}

// CancelReservation converts echo context to params.
func (w *ServerInterfaceWrapper) CancelReservation(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CancelReservation(ctx, id)
	return err
}

// ConfirmReservation converts echo context to params.
func (w *ServerInterfaceWrapper) ConfirmReservation(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ConfirmReservation(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/balances/:idFrom/transfer/:idTo", wrapper.TransferByIds)
	router.GET(baseURL+"/balances/:id", wrapper.GetBalanceById)
	router.POST(baseURL+"/balances/:id", wrapper.AccrueOrWriteOffBalance)
	router.GET(baseURL+"/balances/:id/transactions", wrapper.GetBindedTransactions)
	router.POST(baseURL+"/reservation/balances/:id", wrapper.ReserveOnSeparateAccount)
	router.POST(baseURL+"/reservation/balances/:id/cancel", wrapper.CancelReservation)
	router.POST(baseURL+"/reservation/balances/:id/confirm", wrapper.ConfirmReservation)

}
