package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func Users(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		j, _ := json.Marshal(tom)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

type Store struct {
	redis *redis.Client
}

func (s *Store) GetUsers() {
	keys := s.redis.Keys(context.Background(), "users*")

	var users = make(map[string]string)
	for _, userKey := range keys.Val() {
		res := s.redis.Get(context.Background(), userKey)
		address := res.Val()
		users[userKey] = address
	}

}

func NewStore() Store {
	return Store{
		redis,
	}
}
func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	http.HandleFunc("/users", Users)

	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}
