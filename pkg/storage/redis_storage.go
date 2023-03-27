package storage

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type RedisStorage struct {
	DBName     string
	DBPort     int
	DBUser     string
	DBPassword string
}

func (r *RedisStorage) connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     r.DBName + ":" + fmt.Sprint(r.DBPort),
		Password: r.DBPassword,
		DB:       0,
	})
	return client
}

func (r *RedisStorage) InsertShortenedLinks(link string, shortlink string) error {
	conn_redis := r.connect()
	defer conn_redis.Close()

	if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
		link = "http://" + link
	}

	err := conn_redis.Set("/links/"+shortlink, link, 0).Err()
	if err != nil {
		logrus.Info(err)
	}
	val, err := conn_redis.Get("/links/" + shortlink).Result()
	if err != nil {
		logrus.Info(err)
	}

	if val == "" {
		return fmt.Errorf("unable to shorten url")
	}
	return nil
}

func (r *RedisStorage) UpdateDomainCount(domainName string) error {
	logrus.Info("Updating domain count")
	conn_redis := r.connect()
	defer conn_redis.Close()
	// Get value from database
	key := domainName
	res := conn_redis.ZScore("tags", key)
	val := res.Val()

	// intVal, err := strconv.Atoi(val)
	intVal := val + 1

	// Update value to database
	_, err := conn_redis.ZAdd("tags", redis.Z{intVal, key}).Result()
	if err != nil {
		log.Fatalf("Error adding %s", key)
	}
	return nil
}

func (r *RedisStorage) GetLink(shortUrl string) (string, error) {
	conn_redis := r.connect()
	defer conn_redis.Close()
	logrus.Info("Short url:", shortUrl)
	val, err := conn_redis.Get("/links/" + shortUrl).Result()
	if err != nil {
		logrus.Info(err)
	}

	return val, nil
}

func (r *RedisStorage) GetAllLinks() map[string]string {
	logrus.Info("Get all links")
	return MapShortUrls
}

// func GetDomainCount() map[string]int {
func (r *RedisStorage) GetDomainCount() map[string]int {

	logrus.Info("Get all links")
	MapDomainCountOrdered := make(map[string]int)
	conn_redis := r.connect()
	// Get value from database
	result, err := conn_redis.ZRevRangeWithScores("tags", 0, 2).Result()
	if err != nil {
		log.Fatalf("Error retrieving top 5 keys: %v", err)
	}
	for _, zItem := range result {
		log.Printf("%v\n", zItem)
		log.Printf("%v\n", zItem)
		val, _ := strconv.Atoi(fmt.Sprint(zItem.Score))
		MapDomainCountOrdered[fmt.Sprint(zItem.Member)] = val
	}

	return MapDomainCountOrdered
}
