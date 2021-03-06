package routing

import (
	"strings"

	"github.com/spaceuptech/space-cloud/gateway/config"
)

// SetProjectRoutes adds a project's routes to the global list of routes
func (r *Routing) SetProjectRoutes(project string, routesConfig config.IngressRoutes) error {
	routes := make(config.Routes, 0)
	for _, route := range routesConfig {
		routes = append(routes, route)
	}
	r.lock.Lock()
	defer r.lock.Unlock()

	// Delete all templates for this project
	for k := range r.goTemplates {
		if strings.HasPrefix(k, project) {
			delete(r.goTemplates, k)
		}
	}

	// Add projects to the routes object and generate go templates
	for _, route := range routes {
		route.Project = project
		route.Modify.Tmpl = config.TemplatingEngineGo

		// Parse request template
		if route.Modify.ReqTmpl != "" {
			if err := r.createGoTemplate("request", project, route.ID, route.Modify.ReqTmpl); err != nil {
				return err
			}
		}

		// Parse response template
		if route.Modify.ResTmpl != "" {
			if err := r.createGoTemplate("response", project, route.ID, route.Modify.ResTmpl); err != nil {
				return err
			}
		}
	}

	r.addProjectRoutes(project, routes)
	return nil
}

// DeleteProjectRoutes deletes a project's routes from the global list or routes
func (r *Routing) DeleteProjectRoutes(project string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.deleteProjectRoutes(project)
}

// SetGlobalConfig sets the project level config of the routing module
func (r *Routing) SetGlobalConfig(globalConfig *config.GlobalRoutesConfig) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if globalConfig != nil {
		r.globalConfig = globalConfig
	}
}
