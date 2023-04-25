package model

import "github.com/baptistemehat/go-leadsheet/core/song/model/song"

type Formattable interface {
	Format(f song.SongFormatter) string
}
