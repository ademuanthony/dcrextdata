// Copyright (c) 2018-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package postgres

//go:generate sqlboiler --wipe psql --no-hooks --no-auto-timestamps

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/raedahgroup/dcrextdata/cache"
)

type PgDb struct {
	db *sql.DB
}

func NewPgDb(host, port, user, pass, dbname string) (*PgDb, error) {
	db, err := Connect(host, port, user, pass, dbname)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	return &PgDb{
		db: db,
	}, nil
}

func (pg *PgDb) Close() error {
	log.Trace("Closing postgresql connection")
	return pg.db.Close()
}

func (pg *PgDb) RegisterCharts(charts *cache.ChartData) {
	charts.AddUpdater(cache.ChartUpdater{
		Tag:      "mempool chart",
		Fetcher: func(data *cache.ChartData) (rows *sql.Rows, i func(), e error) {

		},
		Appender: nil,
	})
}

func (pgb *PgDb) chartBlocks(charts *cache.ChartData) (*sql.Rows, func(), error) {
	ctx, cancel := context.WithTimeout(pgb.ctx, pgb.queryTimeout)

	rows, err := retrieveChartBlocks(ctx, pgb.db, charts)
	if err != nil {
		return nil, cancel, fmt.Errorf("chartBlocks: %v", pgb.replaceCancelError(err))
	}
	return rows, cancel, nil
}
