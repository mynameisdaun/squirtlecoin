package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/mynameisdaun/squirtlecoin/utils"
	"os"
)

const (
	dbName       = "blockchain.db"
	blocksBucket = "blocks"
	dataBucket   = "data"
	checkpoint   = "checkpoint"
)

var db *bolt.DB

func getDbname() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(getDbname(), 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func Close() {
	DB().Close()
}

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\n Data: %b\n", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockChain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}

func EmptyBlocks() {
	DB().Update(func(tx *bolt.Tx) error {
		utils.HandleErr(tx.DeleteBucket([]byte(blocksBucket)))
		_, err := tx.CreateBucket([]byte(blocksBucket))
		utils.HandleErr(err)
		return nil
	})
}
