package response

import "net/http"

type Code struct {
	HttpStatusCode int
	ErrorCode      string
	Message        string
}

var Ok Code = Code{http.StatusOK, "", "Request performed successfully"}
var Created Code = Code{http.StatusCreated, "", "Resource created successfully"}

var ValidationError Code = Code{http.StatusBadRequest, "VALIDATION_ERROR", "Request validation error"}
var UnsupportedFeature Code = Code{http.StatusBadRequest, "SELECTED_TYPE_NOT_SUPPORTED", "Selected type not supported"}
var GenericResourceNotFound Code = Code{http.StatusNotFound, "RESOURCE_NOT_FOUND", "Requested resource not found"}
var DuplicateError Code = Code{http.StatusConflict, "DUPLICATE_ERROR", "Duplicate resource, already exists"}
var InvalidStatePrecondition Code = Code{http.StatusConflict, "INVALID_STATE_PRECONDITION", "Invalid resource state for this action"}
var ReferenceIdAlreadyExist Code = Code{http.StatusConflict, "REFERENCE_ID_ALREADY_EXISTS", "Reference ID already exists"}
var GenericServerError Code = Code{http.StatusInternalServerError, "SERVER_ERROR", "There is a problem processing your request. Error has been logged."}
