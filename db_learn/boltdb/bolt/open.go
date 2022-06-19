package main

import (
	"github.com/boltdb/bolt"
  "fmt"
	"log"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

  err = db.Update(func (tx *bolt.Tx) error  {
    // 创建 Black 表
    b, err := tx.CreateBucket([]byte("Black"))
    if err != nil {
      return fmt.Errorf("---%s---", err)
    }

    if b != nil {
      err := b.Put([]byte("l"), []byte("t"))
      if err != nil{
        log.Panic("err")
      }
    }

    return nil
  })
  if err != nil {
    log.Panic(err)
  }
}
