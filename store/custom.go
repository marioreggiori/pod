package store

import (
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type Custom struct {
	Command     string `json:"cmd"`
	Image       string `json:"img"`
	Description string `json:"dsc"`
}

func AddCustom(cm *Custom) error {
	if !isAvailable {
		return fmt.Errorf("storage not available")
	}
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("custom"))
		val, err := json.Marshal(cm)
		if err != nil {
			return err
		}
		err = b.Put([]byte(cm.Command), val)
		return err
	})
}

func RemoveCustom(cmd string) error {
	if !isAvailable {
		return fmt.Errorf("storage not available")
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("custom"))
		return b.Delete([]byte(cmd))
	})
}

func GetCustom() (res []Custom) {
	if !isAvailable {
		return
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("custom"))
		b.ForEach(func(k, v []byte) error {
			elem := Custom{}
			err := json.Unmarshal(v, &elem)
			if err != nil {
				return err
			}
			res = append(res, elem)
			return nil
		})
		return nil
	})

	return
}
