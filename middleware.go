package repository

import (
	"github.com/gomodule/redigo/redis"
)

// do is a function that executes redis commands that create or add something
func (r *RedisRepository) do(commandName string, args ...interface{}) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		HandleError(conn.Close())
	}()
	_, err = conn.Do(commandName, args...)
	return err
}

// getStrings is a function that executes redis commands for lists that return something.
// Converts result in string
func (r *RedisRepository) getStrings(commandName string, args ...interface{}) ([]string, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		HandleError(conn.Close())
	}()
	res, err := redis.Strings(conn.Do(commandName, args...))
	return res, err
}

// getString is a function that executes redis commands for entities that return one argument.
// Converts result in string
func (r *RedisRepository) getString(commandName string, args ...interface{}) (string, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		HandleError(conn.Close())
	}()
	res, err := redis.String(conn.Do(commandName, args...))
	return res, err
}

// getHashMapString is a function that executes redis commands for entities that return array of strings.
// Converts result in map[string]string
func (r *RedisRepository) getHashMapString(commandName string, args ...interface{}) (map[string]string, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		HandleError(conn.Close())
	}()
	res, err := redis.StringMap(conn.Do(commandName, args...))
	return res, err
}

func (r *RedisRepository) getBytes(commandName string, args ...interface{}) ([]byte, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		HandleError(conn.Close())
	}()
	res, err := redis.Bytes(conn.Do(commandName, args...))
	return res, err
}
func (r *RedisRepository) getByteSlices(commandName string, args ...interface{}) ([][]byte, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		HandleError(conn.Close())
	}()
	res, err := redis.ByteSlices(conn.Do(commandName, args...))
	return res, err
}
