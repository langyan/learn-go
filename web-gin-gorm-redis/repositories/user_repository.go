package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/langyan/web-gin-gorm-redis/cache"
	"github.com/langyan/web-gin-gorm-redis/config"
	"github.com/langyan/web-gin-gorm-redis/models"
)

func GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	cacheKey := fmt.Sprintf("user:%d", id)

	// Try to get user from cache
	cachedUserStr, err := cache.GetCache(ctx, cacheKey)
	if err == nil {
		fmt.Println("Cache hit")
		var cachedUser models.User
		if err := json.Unmarshal([]byte(cachedUserStr), &cachedUser); err != nil {
			return nil, err
		}
		return &cachedUser, nil
	}
	fmt.Println("Cache miss")

	// If not in cache, fetch from DB
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Store in cache for 10 minutes
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	cache.SetCache(ctx, cacheKey, string(userJSON), 10*time.Minute)
	return &user, nil
}
