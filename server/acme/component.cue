package acme

import (
	"github.com/cortezaproject/corteza/server/codegen/schema"
)

component: schema.#component & {
	handle: "acme"

	resources: {
		"cuser": cuser
	}

	rbac: operations: {
		"user.search": description:    "Search or filter users"
	}
}
