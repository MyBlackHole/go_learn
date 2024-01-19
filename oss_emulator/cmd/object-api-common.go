package emulator

var globalObjectAPI ObjectLayer

func newStorageAPI(disk string) (storage StorageAPI, err error) {
	storage, err = newStorage(disk)
	return
}
