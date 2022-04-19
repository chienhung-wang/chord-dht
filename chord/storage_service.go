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
	Table map[string]string
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
	if val, ok := kv.Table[key]; ok {
		return val, nil
	} else {
		return "", errors.New("key not found")
	}
}

func (kv *hashTable) Put(key string, val string) error {
	kv.Table[key] = val
	return nil
}

func (kv *hashTable) Delete(key string) error {
	if _, ok := kv.Table[key]; ok {
		delete(kv.Table, key)
		return nil
	} else {
		return errors.New("key not found")
	}
}
