package options

import (
	"github.com/cortezaproject/corteza/server/codegen/schema"
)

acme: schema.#optionsGroup & {
	handle: "acme"
	options: {
		enabled: {
			type:          "bool"
			defaultGoExpr: "true"
			description:   "Enable custom user"
		}
	}
	title: "Custom ACME user"
}
