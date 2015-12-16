// Copyright 2015 Tamás Gulácsi
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package oerr

import (
	"errors"

	"github.com/boltdb/bolt"
)

var ErrNotFound = errors.New("not found")

// Open the DB, read-only mode.
func Open(dbPath string) (GetCloser, error) {
	db, err := bolt.Open(dbPath, 0664, &bolt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	tx, err := db.Begin(false)
	if err != nil {
		return nil, err
	}
	return dbS{tx.Bucket([]byte(bucketName))}, nil
}

type dbS struct {
	*bolt.Bucket
}

func (db dbS) Get(id MsgID) (data MsgData, err error) {
	if db.Bucket == nil {
		return data, errors.New("db is closed")
	}
	key, err := id.MarshalBinary()
	if err != nil {
		return data, err
	}
	val := db.Bucket.Get(key)
	if len(val) == 0 {
		return data, ErrNotFound
	}
	err = data.UnmarshalBinary(val)
	return data, err
}

func (db dbS) Close() error {
	if db.Bucket == nil {
		return nil
	}
	tx := db.Bucket.Tx()
	if tx == nil {
		return nil
	}
	realDB := tx.DB()
	tx.Rollback()
	if realDB == nil {
		return nil
	}
	return realDB.Close()
}

type GetCloser interface {
	Get(MsgID) (MsgData, error)
	Close() error
}
