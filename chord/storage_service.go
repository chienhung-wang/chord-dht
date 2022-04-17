package chord

import (
	"errors"
)

type StorageService interface {
	Get(string) (string, error)
	Put(string, string) error
	Delete(string) error
	GetLocalTable() map[string]string
}

type hashTable struct {
	table map[string]string
}

func NewStorageService() StorageService {
	return &hashTable{
		table: make(map[string]string),
	}
}

func (kv *hashTable) GetLocalTable() map[string]string {
	return kv.table
}

func (kv *hashTable) Get(key string) (string, error) {
	if val, ok := kv.table[key]; ok {
		return val, nil
	} else {
		return "", errors.New("key not found")
	}
}

func (kv *hashTable) Put(key string, val string) error {
	kv.table[key] = val
	return nil
}

func (kv *hashTable) Delete(key string) error {
	if _, ok := kv.table[key]; ok {
		delete(kv.table, key)
		return nil
	} else {
		return errors.New("key not found")
	}
}
