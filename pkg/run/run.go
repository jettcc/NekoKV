package run

import (
	"encoding/json"
	"whiteDB/pkg/store"
	"whiteDB/pkg/warn"
)

func ExecSet(key, value string) {
	set := &store.Set{
		Key:   key,
		Value: value,
	}
	kv := store.Kvstore
	pos, err := kv.Seek()
	if err != nil {
		return
	}

	data, err := json.MarshalIndent(set, "", "\t")
	if err != nil {
		warn.ERRORF(err.Error())
		return
	}
	kv.WriterAt(pos, string(data))
	// 记录索引
	npos, _ := kv.Seek()
	kv.Index[key] = store.CommandPos{
		Pos: uint64(pos),
		Len: uint64(npos - pos),
	}
}
