package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

//DBLogger I Dont Know What this mean
type DBLogger struct{}

//BeforeQuery Function to handle Data before Query Insert
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

//AfterQuery Func to Handle data after query whic mean just show you data after you insert query
func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

//New To Handle connection to Database
func New(opts *pg.Options) *pg.DB {

	return pg.Connect(opts)
}
