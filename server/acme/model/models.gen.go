package model

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

import (
	"github.com/cortezaproject/corteza/server/acme/types"
	"github.com/cortezaproject/corteza/server/pkg/dal"
)

var Cuser = &dal.Model{
	Ident:        "cuser",
	ResourceType: types.CuserResourceType,

	Attributes: dal.AttributeSet{
		&dal.Attribute{
			Ident: "ID",
			Type:  &dal.TypeID{},
			Store: &dal.CodecAlias{Ident: "id"},
		},

		&dal.Attribute{
			Ident: "Name", Sortable: true,
			Type:  &dal.TypeText{},
			Store: &dal.CodecAlias{Ident: "name"},
		},

		&dal.Attribute{
			Ident: "CreatedAt", Sortable: true,
			Type: &dal.TypeTimestamp{
				DefaultCurrentTimestamp: true, Timezone: true, Precision: -1,
			},
			Store: &dal.CodecAlias{Ident: "created_at"},
		},

		&dal.Attribute{
			Ident: "UpdatedAt", Sortable: true,
			Type:  &dal.TypeTimestamp{Nullable: true, Timezone: true, Precision: -1},
			Store: &dal.CodecAlias{Ident: "updated_at"},
		},

		&dal.Attribute{
			Ident: "DeletedAt", Sortable: true,
			Type:  &dal.TypeTimestamp{Nullable: true, Timezone: true, Precision: -1},
			Store: &dal.CodecAlias{Ident: "deleted_at"},
		},

		&dal.Attribute{
			Ident: "CreatedBy",
			Type: &dal.TypeRef{HasDefault: true,
				DefaultValue: 0,

				RefAttribute: "id",
				RefModel: &dal.ModelRef{
					ResourceType: "corteza::system:user",
				},
			},
			Store: &dal.CodecAlias{Ident: "created_by"},
		},

		&dal.Attribute{
			Ident: "UpdatedBy",
			Type: &dal.TypeRef{HasDefault: true,
				DefaultValue: 0,

				RefAttribute: "id",
				RefModel: &dal.ModelRef{
					ResourceType: "corteza::system:user",
				},
			},
			Store: &dal.CodecAlias{Ident: "updated_by"},
		},

		&dal.Attribute{
			Ident: "DeletedBy",
			Type: &dal.TypeRef{HasDefault: true,
				DefaultValue: 0,

				RefAttribute: "id",
				RefModel: &dal.ModelRef{
					ResourceType: "corteza::system:user",
				},
			},
			Store: &dal.CodecAlias{Ident: "deleted_by"},
		},
	},

	Indexes: dal.IndexSet{
		&dal.Index{
			Ident: "PRIMARY",
			Type:  "BTREE",

			Fields: []*dal.IndexField{
				{
					AttributeIdent: "ID",
				},
			},
		},
	},
}

func init() {
	models = append(
		models,
		Cuser,
	)
}
