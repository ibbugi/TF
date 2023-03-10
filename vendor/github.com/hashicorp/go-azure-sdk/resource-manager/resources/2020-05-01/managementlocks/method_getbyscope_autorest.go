package managementlocks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagementLockObject
}

// GetByScope ...
func (c ManagementLocksClient) GetByScope(ctx context.Context, id ScopedLockId) (result GetByScopeOperationResponse, err error) {
	req, err := c.preparerForGetByScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementlocks.ManagementLocksClient", "GetByScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementlocks.ManagementLocksClient", "GetByScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementlocks.ManagementLocksClient", "GetByScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByScope prepares the GetByScope request.
func (c ManagementLocksClient) preparerForGetByScope(ctx context.Context, id ScopedLockId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetByScope handles the response to the GetByScope request. The method always
// closes the http.Response Body.
func (c ManagementLocksClient) responderForGetByScope(resp *http.Response) (result GetByScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
