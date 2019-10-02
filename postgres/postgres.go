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
	"github.com/raedahgroup/dcrextdata/postgres/models"
	"time"
)

type PgDb struct {
	db           *sql.DB
	queryTimeout time.Duration
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
		Fetcher: pg.chartMempool,
		Appender: appendChartMempool,
	})
}

func (pg *PgDb) chartMempool(ctx context.Context, charts *cache.ChartData) (interface{}, func(), error) {
	ctx, cancel := context.WithTimeout(ctx, pg.queryTimeout)

	charts.Height()
	mempoolSlice, err := models.Mempools(models.MempoolWhere.Time.GT(time.Unix(int64(charts.MempoolTime()), 0))).All(ctx, pg.db)
	if err != nil {
		return nil, cancel, fmt.Errorf("chartBlocks: %s", err.Error())
	}
	return mempoolSlice, cancel, nil
}

// Append the results from retrieveChartBlocks to the provided ChartData.
// This is the Appender half of a pair that make up a cache.ChartUpdater.
func appendChartMempool(charts *cache.ChartData, mempoolSliceInt interface{}) error {
	mempoolSlice := mempoolSliceInt.(models.MempoolSlice)
	chartsMempool := charts.Mempool
	for _, mempoolData := range mempoolSlice {
		chartsMempool.Time = append(chartsMempool.Time, uint64(mempoolData.Time.Unix()))
		chartsMempool.Fees = append(chartsMempool.Fees, mempoolData.TotalFee.Float64)
		chartsMempool.TxCount = append(chartsMempool.TxCount, uint64(mempoolData.NumberOfTransactions.Int))
		chartsMempool.Size = append(chartsMempool.Size, uint64(mempoolData.Size.Int))
	}
	return nil
}
