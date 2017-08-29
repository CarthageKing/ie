// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": risk_assessment TestHelpers
//
// Command:
// $ goagen
// --design=github.com/intervention-engine/ie/design
// --out=$(GOPATH)/src/github.com/intervention-engine/ie
// --version=v1.2.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/intervention-engine/ie/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

// ListRiskAssessmentBadRequest runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListRiskAssessmentBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RiskAssessmentController, id string, endDate *time.Time, riskServiceID string, startDate *time.Time) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		query["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		query["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		query["start_date"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/patients/%v/risk_assessments", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		prms["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		prms["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		prms["start_date"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RiskAssessmentTest"), rw, req, prms)
	listCtx, _err := app.NewListRiskAssessmentContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// ListRiskAssessmentInternalServerError runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListRiskAssessmentInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RiskAssessmentController, id string, endDate *time.Time, riskServiceID string, startDate *time.Time) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		query["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		query["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		query["start_date"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/patients/%v/risk_assessments", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		prms["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		prms["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		prms["start_date"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RiskAssessmentTest"), rw, req, prms)
	listCtx, _err := app.NewListRiskAssessmentContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// ListRiskAssessmentOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListRiskAssessmentOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RiskAssessmentController, id string, endDate *time.Time, riskServiceID string, startDate *time.Time) (http.ResponseWriter, app.RiskAssessmentCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		query["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		query["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		query["start_date"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/patients/%v/risk_assessments", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		prms["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		prms["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		prms["start_date"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RiskAssessmentTest"), rw, req, prms)
	listCtx, _err := app.NewListRiskAssessmentContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.RiskAssessmentCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.RiskAssessmentCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.RiskAssessmentCollection", resp)
		}
	}

	// Return results
	return rw, mt
}

// ListRiskAssessmentOKList runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListRiskAssessmentOKList(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.RiskAssessmentController, id string, endDate *time.Time, riskServiceID string, startDate *time.Time) (http.ResponseWriter, app.RiskAssessmentListCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		query["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		query["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		query["start_date"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/patients/%v/risk_assessments", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if endDate != nil {
		sliceVal := []string{(*endDate).Format(time.RFC3339)}
		prms["end_date"] = sliceVal
	}
	{
		sliceVal := []string{riskServiceID}
		prms["risk_service_id"] = sliceVal
	}
	if startDate != nil {
		sliceVal := []string{(*startDate).Format(time.RFC3339)}
		prms["start_date"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "RiskAssessmentTest"), rw, req, prms)
	listCtx, _err := app.NewListRiskAssessmentContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.RiskAssessmentListCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.RiskAssessmentListCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.RiskAssessmentListCollection", resp)
		}
	}

	// Return results
	return rw, mt
}
