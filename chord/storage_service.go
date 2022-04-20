package chord

import (
	"errors"
	"sync"
)

type StorageService interface {
	Get(string) (string, error)
	Put(string, string) error
	Delete(string) error
	GetLocalTable() map[string]string
}

type hashTable struct {
	Table     map[string]string
	writeLock sync.Mutex
}

func NewStorageService() StorageService {
	return &hashTable{
		Table: make(map[string]string),
	}
}

func (kv *hashTable) GetLocalTable() map[string]string {
	return kv.Table
}

func (kv *hashTable) Get(key string) (string, error) {
	kv.writeLock.Lock()
	if val, ok := kv.Table[key]; ok {
		kv.writeLock.Unlock()
		return val, nil
	} else {
		kv.writeLock.Unlock()
		return "", errors.New("key not found")
	}
}

func (kv *hashTable) Put(key string, val string) error {
	kv.writeLock.Lock()
	kv.Table[key] = val
	kv.writeLock.Unlock()
	return nil
}

func (kv *hashTable) Delete(key string) error {
	kv.writeLock.Lock()
	if _, ok := kv.Table[key]; ok {
		delete(kv.Table, key)
		kv.writeLock.Unlock()
		return nil
	} else {
		kv.writeLock.Unlock()
		return errors.New("key not found")
	}
}
