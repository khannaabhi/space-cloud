package configman

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spacecloud-io/space-cloud/model"
)

func extractPathParams(urlPath, method string) (op, module, typeName, resourceName string, err error) {
	// Set the default operation to single
	op = "list"
	if method == http.MethodPost {
		op = "single"
	}

	// Check if url has proper length
	arr := strings.Split(urlPath[1:], "/")
	if len(arr) > 5 || len(arr) < 4 {
		err = fmt.Errorf("invalid config url provided - %s", urlPath)
		return
	}

	// Check the operation type
	if len(arr) == 5 {
		op = "single"
		resourceName = arr[4]
	}

	// Set the other parameters
	module = arr[2]
	typeName = arr[3]
	return
}

func prepareErrorResponseBody(err error, schemaErrors []string) interface{} {
	return model.ErrorResponse{
		Error:        err.Error(),
		SchemaErrors: schemaErrors,
	}
}
