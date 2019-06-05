package videogeneration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// GenerateVideoEndpoint : Generates an endpoint for the Generate Video Service
func GenerateVideoEndpoint(svc VideoGenerateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		// 
		req := request.(VideoInputObject)
		serviceID, err := svc.AddService(req.pdfID,req.audiofileID,req.mapping)
		if err != nil {
			return nil, err
		}
		return AddResponse{serviceID, err}, nil

	}
}


// Decodes the request body for the service

func DecodeService(_ context.Context, r *http.Request) (interface{}, error) {
	var request  //TODO:Error
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}



// EncodeResponse :
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}


func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	//	w.WriteHeader(Err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
