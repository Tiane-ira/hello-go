package etcd

import (
	"context"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"
)

func Lock(index int, key string) {
	session, err := concurrency.NewSession(client)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	mux := concurrency.NewMutex(session, key)
	log.Printf("acquired lock on %s %d", key, index)
	err = mux.Lock(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("lock on %d", index)
	//i := rand.Intn(5) + 1
	time.Sleep(time.Second * time.Duration(70))
	err = mux.Unlock(context.TODO())
	log.Printf("relese lock on %s %d", key, index)
}

func TryLock(index int, key string) {
	session, err := concurrency.NewSession(client)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	mux := concurrency.NewMutex(session, key)
	log.Printf("acquired lock on %s %d", key, index)
	err = mux.TryLock(context.TODO())
	if err != nil {
		log.Printf("获取锁失败：%v", err)
		return
	}
	log.Printf("lock on %d", index)
	//i := rand.Intn(5) + 1
	time.Sleep(time.Second * time.Duration(70))
	err = mux.Unlock(context.TODO())
	log.Printf("relese lock on %s %d", key, index)
}
