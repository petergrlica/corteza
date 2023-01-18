package types

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// acme/types/types.yaml

type (

	// CuserSet slice of Cuser
	//
	// This type is auto-generated.
	CuserSet []*Cuser
)

// Walk iterates through every slice item and calls w(Cuser) err
//
// This function is auto-generated.
func (set CuserSet) Walk(w func(*Cuser) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Cuser) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set CuserSet) Filter(f func(*Cuser) (bool, error)) (out CuserSet, err error) {
	var ok bool
	out = CuserSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set CuserSet) FindByID(ID uint64) *Cuser {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set CuserSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}
