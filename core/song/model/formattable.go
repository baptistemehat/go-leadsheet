package model

// Formattable
type Formattable interface {
	Format(f Formatter) (string, error)
}
