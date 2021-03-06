package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spaceuptech/helpers"

	"github.com/spaceuptech/space-cloud/gateway/config"
	"github.com/spaceuptech/space-cloud/gateway/managers/admin"
	"github.com/spaceuptech/space-cloud/gateway/managers/syncman"
	"github.com/spaceuptech/space-cloud/gateway/model"
	"github.com/spaceuptech/space-cloud/gateway/utils"
)

// HandlePostIntegration handles the post integration request
func HandlePostIntegration(adminMan *admin.Manager, syncMan *syncman.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request
		token := utils.GetTokenFromHeader(r)

		// Get the body of the request
		req := new(config.IntegrationConfig)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusBadRequest, err)
			return
		}
		defer utils.CloseTheCloser(r.Body)

		// Validate the token
		reqParams, err := adminMan.IsTokenValid(r.Context(), token, "integration", "modify", map[string]string{"integration": req.ID})
		if err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusForbidden, err)
			return
		}

		// Create a context object
		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		// Save the integration hook
		reqParams.Method = r.Method
		reqParams.Path = r.URL.Path
		reqParams.Headers = r.Header
		reqParams.Payload = req

		// Enable the integration
		status, err := syncMan.EnableIntegration(ctx, req, reqParams)
		if err != nil {
			_ = helpers.Response.SendErrorResponse(ctx, w, status, err)
			return
		}

		_ = helpers.Response.SendOkayResponse(ctx, status, w)
	}
}

// HandleDeleteIntegration handles the delete integration request
func HandleDeleteIntegration(adminMan *admin.Manager, syncMan *syncman.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request
		token := utils.GetTokenFromHeader(r)

		vars := mux.Vars(r)
		name := vars["name"]

		// Lets close the body since we wont be needing it
		defer utils.CloseTheCloser(r.Body)

		// Validate the token
		reqParams, err := adminMan.IsTokenValid(r.Context(), token, "integration", "modify", map[string]string{"integration": name})
		if err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusForbidden, err)
			return
		}

		// Create a context object
		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		reqParams.Method = r.Method
		reqParams.Path = r.URL.Path
		reqParams.Headers = r.Header

		// remove the integration
		status, err := syncMan.RemoveIntegration(ctx, name, reqParams)
		if err != nil {
			_ = helpers.Response.SendErrorResponse(ctx, w, status, err)
			return
		}

		_ = helpers.Response.SendOkayResponse(ctx, status, w)
	}
}

// HandleGetIntegrations handles the get integration hook request
func HandleGetIntegrations(adminMan *admin.Manager, syncMan *syncman.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request
		token := utils.GetTokenFromHeader(r)

		// Get the path parameters
		integrationID := "*"
		if id := r.URL.Query().Get("id"); id != "" {
			integrationID = id
		}

		// Close the body
		defer utils.CloseTheCloser(r.Body)

		// Validate the token
		reqParams, err := adminMan.IsTokenValid(r.Context(), token, "integration", "read", map[string]string{"integration": integrationID})
		if err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusForbidden, err)
			return
		}

		// Create a context object
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		// Remove the integration hook
		reqParams.Method = r.Method
		reqParams.Path = r.URL.Path
		reqParams.Headers = r.Header
		status, integrations, err := syncMan.GetIntegrations(ctx, integrationID, reqParams)
		if err != nil {
			_ = helpers.Response.SendErrorResponse(ctx, w, status, err)
			return
		}

		_ = helpers.Response.SendResponse(ctx, w, status, model.Response{Result: integrations})
	}
}

// HandleGetIntegrationTokens handles the get integration tokens request
func HandleGetIntegrationTokens(syncMan *syncman.Manager) http.HandlerFunc {
	type request struct {
		ID  string `json:"id"`
		Key string `json:"key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the body of the request
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusBadRequest, err)
			return
		}
		defer utils.CloseTheCloser(r.Body)

		// Get tokens for integration
		status, tokens, err := syncMan.GetIntegrationTokens(r.Context(), req.ID, req.Key)
		if err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, status, err)
			return
		}

		_ = helpers.Response.SendResponse(r.Context(), w, status, model.Response{Result: tokens})
	}
}

// HandleAddIntegrationHook handles the add integration hook request
func HandleAddIntegrationHook(adminMan *admin.Manager, syncMan *syncman.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request
		token := utils.GetTokenFromHeader(r)

		vars := mux.Vars(r)
		name := vars["name"]

		// Get the body of the request
		req := new(config.IntegrationHook)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusBadRequest, err)
			return
		}
		defer utils.CloseTheCloser(r.Body)

		// Validate the token
		reqParams, err := adminMan.IsTokenValid(r.Context(), token, "integration-hook", "modify", map[string]string{"integration": name, "hook": req.ID})
		if err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusForbidden, err)
			return
		}

		// Create a context object
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		// Save the integration hook
		reqParams.Method = r.Method
		reqParams.Path = r.URL.Path
		reqParams.Headers = r.Header
		reqParams.Payload = req

		status, err := syncMan.AddIntegrationHook(ctx, name, req, reqParams)
		if err != nil {
			_ = helpers.Response.SendErrorResponse(ctx, w, status, err)
			return
		}

		_ = helpers.Response.SendOkayResponse(ctx, status, w)
	}
}

// HandleDeleteIntegrationHook handles the delete integration hook request
func HandleDeleteIntegrationHook(adminMan *admin.Manager, syncMan *syncman.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request
		token := utils.GetTokenFromHeader(r)

		vars := mux.Vars(r)
		name := vars["name"]
		hookID := vars["id"]

		// Close the body
		defer utils.CloseTheCloser(r.Body)

		// Validate the token
		reqParams, err := adminMan.IsTokenValid(r.Context(), token, "integration-hook", "modify", map[string]string{"integration": name, "hook": hookID})
		if err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusForbidden, err)
			return
		}

		// Create a context object
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		// Remove the integration hook- User Management (Auth Providers): auth-provider
		// - Database: db-config, db-schema, db-prepared-query, db-rule
		// - Eventing: eventing-trigger, eventing-config, eventing-schema, eventing-rule
		// - Filestore: filestore-config,  filestore-rule
		// - Project: letsencrypt, project, ingress-global (this can go in ingress too)
		// -  Ingress: ingress-route
		// - Remote services: remote-service
		// - Deployments: service, service-route
		// - Secret: secret
		// - Integration: integration, integration-hook
		reqParams.Method = r.Method
		reqParams.Path = r.URL.Path
		reqParams.Headers = r.Header

		status, err := syncMan.RemoveIntegrationHook(ctx, name, hookID, reqParams)
		if err != nil {
			_ = helpers.Response.SendErrorResponse(ctx, w, status, err)
			return
		}

		_ = helpers.Response.SendOkayResponse(ctx, status, w)
	}
}

// HandleGetIntegrationHooks handles the get integration hook request
func HandleGetIntegrationHooks(adminMan *admin.Manager, syncMan *syncman.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request
		token := utils.GetTokenFromHeader(r)

		vars := mux.Vars(r)
		name := vars["name"]

		// Get the path parameters
		hookID := "*"
		if id := r.URL.Query().Get("id"); id != "" {
			hookID = id
		}

		// Close the body
		defer utils.CloseTheCloser(r.Body)

		// Validate the token
		reqParams, err := adminMan.IsTokenValid(r.Context(), token, "integration-hook", "read", map[string]string{"integration": name, "hook": hookID})
		if err != nil {
			_ = helpers.Response.SendErrorResponse(r.Context(), w, http.StatusForbidden, err)
			return
		}

		// Create a context object
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		// Remove the integration hook
		reqParams.Method = r.Method
		reqParams.Path = r.URL.Path
		reqParams.Headers = r.Header
		status, hooks, err := syncMan.GetIntegrationHooks(ctx, name, hookID, reqParams)
		if err != nil {
			_ = helpers.Response.SendErrorResponse(ctx, w, status, err)
			return
		}

		_ = helpers.Response.SendResponse(ctx, w, status, model.Response{Result: hooks})
	}
}
