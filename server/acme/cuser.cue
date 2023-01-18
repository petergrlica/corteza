package acme

import (
	"github.com/cortezaproject/corteza/server/codegen/schema"
)

cuser: {
	features: {
		labels:  false
		paging:  true
		sorting: true
	}

	model: {
		ident: "cuser"
		attributes: {
			id: schema.IdField
			name: {
				sortable: true
				ident:    "name"
				goType:   "string"
				dal: {}
			}
			created_at: schema.SortableTimestampNowField
			updated_at: schema.SortableTimestampNilField
			deleted_at: schema.SortableTimestampNilField
			created_by: schema.AttributeUserRef
			updated_by: schema.AttributeUserRef
			deleted_by: schema.AttributeUserRef
		}

		indexes: {
			"primary": {attribute: "id"}
		}
	}

	filter: {
		struct: {
			name: {goType: "string"}
			deleted: {goType: "filter.State", storeIdent: "deleted_at"}
		}

		query: ["name"]
		byNilState: ["deleted"]
	}

	rbac: {
		operations: {
			"cuser.create": description: "Create custom user"
		}
	}

	store: {
		ident: "acmeCuser"

		api: {
			lookups: [
				{
					fields: ["id"]
					description: """
						searches for custom user by ID

						Returns custom user
						"""
				}, {
					fields: ["name"]
					description: """
						searches for custom user by name
						"""
				},
			]
		}
	}
}
