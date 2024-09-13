package stockproducts

// Product represents any item in the company inventory
type Product struct {
	Uuid        string
	Name        string
	Reference   string
	Line        string
	Supplier    string
	Description string
	ImageUrl    string
	Variants    []ProductVariant
}

// ProductVariant represents subcategory of item such as for
// example an item that could be either in red or blue color.
type ProductVariant struct {
	// Size correspond to item size such as S, M, L, XL, etc...
	Size  *string
	Color *string
	// Length corresponds to the original length of the item
	// (e.g 50m joint spool)
	Length   *int
	Quantity int
}

type ProductSize struct {
}
