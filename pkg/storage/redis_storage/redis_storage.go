package redis_storage

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// ShortUrls := make(map[string]string)
var MapShortUrls map[string]string
var MapDomainCount map[string]int

func connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func StoreShortenedLinks(link string, shortlink string) error {
	conn_redis := connect()
	logrus.Println("Shortened Links Storage Redis")

	if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
		link = "http://" + link
	}

	err := conn_redis.Set("/links/"+shortlink, link, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := conn_redis.Get("/links/" + shortlink).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	if val == "" {
		return fmt.Errorf("unable to shorten url")
	}
	return nil
}

func UpdateDomainCount(domainName string) error {
	logrus.Info("Updating domain count")
	conn_redis := connect()
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

func GetLink(shortUrl string) (string, error) {
	conn_redis := connect()
	logrus.Info("Short url:", shortUrl)
	val, err := conn_redis.Get("/links/" + shortUrl).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val, nil
}

func GetAllLinks() map[string]string {
	logrus.Info("Get all links")
	return MapShortUrls
}

// func GetDomainCount() map[string]int {
func GetDomainCount() map[string]int {

	logrus.Info("Get all links")
	MapDomainCountOrdered := make(map[string]int)
	conn_redis := connect()
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
