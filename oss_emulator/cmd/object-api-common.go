package emulator

var globalObjectAPI ObjectLayer

func newStorageAPI(disk string) (storage StorageAPI, err error) {
	metaDir := pathJoin(disk, "meta")
	dataDir := pathJoin(disk, "data")
    InitMeta(metaDir)
	storage, err = newStorage(dataDir)
	return
}
