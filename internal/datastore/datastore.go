package datastore

import (
	"fmt"
	"go-blog-api/ent"
	"go-blog-api/internal/config"

	"entgo.io/ent/dialect"

	_ "github.com/go-sql-driver/mysql"
)

func NewClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", 
										config.EnvVar.DBUser,
										config.EnvVar.DBPass,
										config.EnvVar.DBHost,
										config.EnvVar.DBPort,
										config.EnvVar.DBSchema)

	return ent.Open(dialect.MySQL, dsn, entOptions...)
}