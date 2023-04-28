package model

type Formattable interface {
	Format(f Formatter) (string, error)
}
