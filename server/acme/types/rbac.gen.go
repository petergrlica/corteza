package types

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

import (
	"fmt"
	"strconv"
)

type (
	// Component struct serves as a virtual resource type for the acme component
	//
	// This struct is auto-generated
	Component struct{}
)

var (
	_ = fmt.Printf
	_ = strconv.FormatUint
)

// RbacResource returns string representation of RBAC resource for Cuser by calling CuserRbacResource fn
//
// RBAC resource is in the corteza::acme:cuser/... format
//
// This function is auto-generated
func (r Cuser) RbacResource() string {
	return CuserRbacResource(r.ID)
}

// CuserRbacResource returns string representation of RBAC resource for Cuser
//
// RBAC resource is in the corteza::acme:cuser/... format
//
// This function is auto-generated
func CuserRbacResource(id uint64) string {
	cpts := []interface{}{CuserResourceType}
	if id != 0 {
		cpts = append(cpts, strconv.FormatUint(id, 10))
	} else {
		cpts = append(cpts, "*")
	}

	return fmt.Sprintf(CuserRbacResourceTpl(), cpts...)

}

func CuserRbacResourceTpl() string {
	return "%s/%s"
}

// RbacResource returns string representation of RBAC resource for Component by calling ComponentRbacResource fn
//
// RBAC resource is in the corteza::acme/... format
//
// This function is auto-generated
func (r Component) RbacResource() string {
	return ComponentRbacResource()
}

// ComponentRbacResource returns string representation of RBAC resource for Component
//
// RBAC resource is in the corteza::acme/ format
//
// This function is auto-generated
func ComponentRbacResource() string {
	return ComponentResourceType + "/"

}

func ComponentRbacResourceTpl() string {
	return "%s"
}
