package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	data, err := db.Get([]byte("key"), nil)

	if err != nil {
		fmt.Println(err.Error(), "----")
	} else {
		fmt.Println(data)
	}

	err = db.Put([]byte("key"), []byte("value"), nil)

	data, err = db.Get([]byte("key"), nil)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(data)
	}

	err = db.Delete([]byte("key"), nil)

	// iter := db.NewIterator(nil, nil)
	// for iter.Next() {
	// 	// Remember that the contents of the returned slice should not be modified, and
	// 	// only valid until the next call to Next.
	// 	key := iter.Key()
	// 	value := iter.Value()
	// 	...
	// }
	// iter.Release()
	// err = iter.Error()

	// iter := db.NewIterator(nil, nil)
	// for ok := iter.Seek(key); ok; ok = iter.Next() {
	// 	// Use key/value.
	// 	...
	// }
	// iter.Release()
	// err = iter.Error()

	// iter := db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("xoo")}, nil)
	// for iter.Next() {
	// 	// Use key/value.
	// 	...
	// }
	// iter.Release()
	// err = iter.Error()

	// iter := db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
	// for iter.Next() {
	// 	// Use key/value.
	// 	...
	// }
	// iter.Release()
	// err = iter.Error()

	// batch := new(leveldb.Batch)
	// batch.Put([]byte("foo"), []byte("value"))
	// batch.Put([]byte("bar"), []byte("another value"))
	// batch.Delete([]byte("baz"))
	// err = db.Write(batch, nil)
	// ...

	// o := &opt.Options{
	// 	Filter: filter.NewBloomFilter(10),
	// }
	// db, err := leveldb.OpenFile("path/to/db", o)
	// ...
	// defer db.Close()
}
