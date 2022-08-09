package dal

type (
	IndexType string

	// Index @todo: again name, for now to avoid confusion
	Index struct {
		Ident string
		Label string

		// Keep it either slice or
		// 2 separate fields; Attribute & AttributeSet
		// Why? No reason but keeping those apart gives a bit better handling for generic/variadic function arguments
		AttributeSet AttributeSet

		// Type describes what the value represents and how it should be
		// encoded/decoded
		// IndexType have a file for it
		Type IndexType
	}
)

// Prep struct for each indexType
// attribute index <- single atr
// unique_* index <- multi atr
//func (t receiver) Type() IndexType        { return "corteza::dal:index-type:{name}" }
