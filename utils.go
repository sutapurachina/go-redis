package repository

func (r *RedisRepository) SaveKey(keyName string, thing interface{}) error {
	command := "SET"
	err := r.do(command, keyName, thing)
	return err
}

func (r *RedisRepository) GetKey(keyName string) ([]byte, error) {
	command := "GET"
	res, err := r.getBytes(command, keyName)
	return res, err
}

// SaveInList executes "rpush" or "lpush" to given list
func (r *RedisRepository) SaveInList(listName string, thing interface{}, end bool) error {
	var command string
	if end == true {
		command = "RPUSH"
	} else {
		command = "LPUSH"
	}
	err := r.do(command, listName, thing)
	return err
}

// SaveInHash executes "HSET ..." - creates hashmap and/or adds field with given value
func (r *RedisRepository) SaveInHash(hashTableName, fieldName string, thing interface{}) error {
	command := "HSET"
	err := r.do(command, hashTableName, fieldName, thing)
	return err
}

// GetList execute "LRANGE 0 -1" to given list
func (r *RedisRepository) GetList(listName string) ([][]byte, error) {
	command := "LRANGE"
	res, err := r.getByteSlices(command, listName, 0, -1)
	return res, err
}

func (r *RedisRepository) GetHashMapField(hashTableName string, fieldName interface{}) (string, error) {
	command := "HGET"
	res, err := r.getString(command, hashTableName, fieldName)
	return res, err
}

func (r *RedisRepository) GetHashMap(hashTableName string) (map[string]string, error) {
	command := "HGETALL"
	res, err := r.getHashMapString(command, hashTableName)
	return res, err
}

func (r *RedisRepository) DeleteByValue(listName string, value interface{}, count int) error {
	command := "LREM"
	err := r.do(command, listName, count, value)
	return err
}

func (r *RedisRepository) DeleteKey(key string) error {
	command := "DEL"
	err := r.do(command, key)
	return err
}
