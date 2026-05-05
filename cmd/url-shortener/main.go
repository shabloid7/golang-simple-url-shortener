package main

import (
	"log"
	"github.com/shabloid7/golang-simple-url-shortener/internal/config"
	"github.com/shabloid7/golang-simple-url-shortener/internal/handler"
	"github.com/shabloid7/golang-simple-url-shortener/internal/repository"
	"github.com/shabloid7/golang-simple-url-shortener/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)



func main() {
	cfg := config.Load()

	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	repo := repository.NewRedisRepository(redisClient)
	serv := service.NewURLService(repo, cfg.BaseURL)
	h := handler.NewHandler(serv)

	r := gin.Default()
	h.RegisterRoutes(r)

	log.Printf("starting server on port %s", cfg.ServerPort)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}