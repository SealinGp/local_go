package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

//url 映射表 短url => 长url
type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}
type record struct {
	Key,URL string
}


func NewURLStore(fileName string) *URLStore {
	store := &URLStore{
		urls: make(map[string]string),
		save: make(chan record,saveQueueLength),
	}
	if err := store.load(fileName); err != nil {
		log.Println("Error loading URLStore:", err)
	}
	go store.saveLoop(fileName)
	return store
}


func (u *URLStore)Get(key string) string {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.urls[key]
}

func (u *URLStore)Set(key,url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _,present := u.urls[key];present {
		return false
	}
	u.urls[key] = url
	return true
}

func (u *URLStore)Count() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return len(u.urls)
}

func (u *URLStore)Put(url string) string {
	for {
		key := genKey(u.Count()) //生成短url的key
		if ok := u.Set(key,url);ok {
			u.save <- record{key,url}
			return key
		}
	}
	return ""
}

func (u *URLStore)load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			u.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (u *URLStore) saveLoop(filename string) error {
	f,err := os.OpenFile(filename,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	e := gob.NewEncoder(f)

	for {
		r := <-u.save
		if err = e.Encode(r); err != nil {
			log.Println("error saving to gob")
		}
	}
}