// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
)

const dateTemplate = "2006-01-02 15:04"
const dateMiliTemplate = "2006-01-02 15:04:05.99"

type insertable interface {
	Insert(context.Context, boil.ContextExecutor, boil.Columns) error
}

type upsertable interface {
	Upsert(ctx context.Context, db boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error
}

func UnixTimeToString(t int64) string {
	return time.Unix(t, 0).UTC().Format(dateTemplate)
}

func RoundValue(input float64) string {
	value := input * 100
	return strconv.FormatFloat(value, 'f', 3, 64)
}

// func RoundValue(input float64, places int) (newVal float64) {
// 	input := input * 100
//  	var round float64
//  	pow := math.Pow(10, float64(places))
//  	digit := pow * input
//  	round = math.Ceil(digit)
//  	newVal = round / pow
//  	return
//  }

func int64ToTime(t int64) time.Time {
	return time.Unix(t, 0)
}

func (pg *PgDb) tryInsert(ctx context.Context, txr boil.Transactor, data insertable) error {
	err := data.Insert(ctx, pg.db, boil.Infer())
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return err
		}
		errT := txr.Rollback()
		if errT != nil {
			return errT
		}
		return err
	}
	return nil
}

func (pg *PgDb) tryUpsert(ctx context.Context, txr boil.Transactor, data upsertable) error {
	err := data.Upsert(ctx, pg.db, true, nil, boil.Infer(), boil.Infer())
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return err
		}
		errT := txr.Rollback()
		if errT != nil {
			return errT
		}
		return err
	}
	return nil
}
