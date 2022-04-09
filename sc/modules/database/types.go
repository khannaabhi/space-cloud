package database

import (
	"github.com/spacecloud-io/space-cloud/config"
	"github.com/spacecloud-io/space-cloud/model"
)

// Config describes the configuration required by a single database
type Config struct {
	Connector       *config.DatabaseConfig         `json:"connector"`
	Schemas         config.DatabaseSchemas         `json:"schemas"`
	PreparedQueries config.DatabasePreparedQueries `json:"preparedQueries"`
}

func getTypeDefinitions() model.Types {
	return model.Types{
		"config": &model.TypeDefinition{
			Schema: m{
				"type": "object",
				"properties": m{
					"dbAlias": m{
						"type": "string",
					},
					"type": m{
						"type": "string",
					},
					"name": m{
						"type": "string",
					},
					"conn": m{
						"type": "string",
					},
					"isPrimary": m{
						"type": "boolean",
					},
					"enabled": m{
						"type": "boolean",
					},
					"batchTime": m{
						"type": "integer",
					},
					"batchRecords": m{
						"type": "integer",
					},
					"limit": m{
						"type": "integer",
					},
					"driverConf": m{
						"type": "object",
						"properties": m{
							"maxConn": m{
								"type": "integer",
							},
							"maxIdleTimeout": m{
								"type": "integer",
							},
							"minConn": m{
								"type": "integer",
							},
							"maxIdleConn": m{
								"type": "integer",
							},
						},
						"required": t{"maxConn", "maxIdleTimeout", "minConn", "maxIdleConn"},
					},
				},
				"required": t{"type", "name", "conn"},
			},
			Hooks:           model.Hooks{model.PhasePreApply: struct{}{}},
			RequiredParents: []string{"project"},
		},
		"schema": &model.TypeDefinition{
			Schema: m{
				"type": "object",
				"properties": m{
					"col": m{
						"type": "string",
					},
					"dbAlias": m{
						"type": "string",
					},
					"schema": m{
						"type": "string",
					},
				},
				"required": t{"schema"},
			},
			Hooks:           model.Hooks{model.PhasePreApply: struct{}{}},
			RequiredParents: []string{"project", "db"},
		},
		"prepared-query": &model.TypeDefinition{
			Schema: m{
				"type": "object",
				"properties": m{
					"id": m{
						"type": "string",
					},
					"sql": m{
						"type": "string",
					},
					"rule": m{
						"type":                 "object",
						"additionalProperties": true,
					},
					"dbAlias": m{
						"type": "string",
					},
					"args": m{
						"type": "array",
						"items": m{
							"type": "string",
						},
					},
				},
				"required": t{"sql"},
			},
			Hooks:           model.Hooks{model.PhasePreApply: struct{}{}},
			RequiredParents: []string{"project", "db"},
		},
	}
}

type (
	m map[string]interface{}
	t []interface{}
)
