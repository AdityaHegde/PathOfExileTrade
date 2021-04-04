package service

import (
	"context"
	"fmt"
	"github.com/AdityaHegde/PathOfExileTrade/account"
	"strconv"

	servicemodel "github.com/AdityaHegde/PathOfExileTrade/model/service"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// AddToQueue is exported
func AddToQueue(rdb *redis.Client, listingType *servicemodel.ListingType, user *account.User) {
	if userAlreadyQueued(rdb, listingType, user) {
		return
	}

	// List is used to maintain a queue
	_, rPushErr := rdb.RPush(ctx, listingType.Name, user.Name).Result()
	// Set is used to make sure we do not duplicate
	_, sAddErr := rdb.SAdd(ctx, listingType.Name, user.Name).Result()

	if rPushErr != nil {
		fmt.Println("Right push failed")
		fmt.Println(rPushErr)
	}
	if sAddErr != nil {
		fmt.Println("Set add failed")
		fmt.Println(sAddErr)
	}
}

// RemoveFromQueue is exported
func RemoveFromQueue(rdb *redis.Client, listingType *servicemodel.ListingType, user *account.User) {
	_, lRemErr := rdb.LRem(ctx, listingType.Name, 1, user.Name).Result()
	_, sRemErr := rdb.SRem(ctx, listingType.Name, user.Name).Result()

	if lRemErr != nil {
		fmt.Println("Left remove failed")
		fmt.Println(lRemErr)
	}
	if sRemErr != nil {
		fmt.Println("Set remove failed")
		fmt.Println(sRemErr)
	}
}

// GetUsersForListing is exported
func GetUsersForListing(rdb *redis.Client, listing *servicemodel.Listing) []string {
	users := getNUsersFromQueue(rdb, &listing.ListingType, listing.MaxParticipants)
	claimUsersForListing(rdb, listing, users)
	return users
}

// ClearListing is exported
func ClearListing(rdb *redis.Client, listing *servicemodel.Listing) {
	listingStrID := strconv.FormatUint(listing.ID, 10)
	rdb.LTrim(ctx, listingStrID, 1, 0)
}

func userAlreadyQueued(rdb *redis.Client, listingType *servicemodel.ListingType, user *account.User) bool {
	exists, err := rdb.SIsMember(ctx, listingType.Name, user.Name).Result()

	return err == redis.Nil || exists
}

func getNUsersFromQueue(rdb *redis.Client, listingType *servicemodel.ListingType, num uint) []string {
	var users []string = make([]string, num)

	for i := 0; i < int(num); i++ {
		users[i] = getUserFromQueue(rdb, listingType)
	}

	return users
}

func getUserFromQueue(rdb *redis.Client, listingType *servicemodel.ListingType) string {
	user, lPopErr := rdb.LPop(ctx, listingType.Name).Result()

	if lPopErr != nil {
		_, sRemErr := rdb.SRem(ctx, listingType.Name, user).Result()
		if sRemErr != nil {
			fmt.Println("Set remove failed")
			fmt.Println(sRemErr)
		}
	} else {
		fmt.Println("Left pop failed")
		fmt.Println(lPopErr)
	}

	return user
}

func claimUsersForListing(rdb *redis.Client, listing *servicemodel.Listing, users []string) {
	listingStrID := strconv.FormatUint(listing.ID, 10)

	len, lLenErr := rdb.LLen(ctx, listingStrID).Result()

	if len > 0 || (lLenErr != nil && lLenErr != redis.Nil) {
		ClearListing(rdb, listing)
	}

	for user := range users {
		rdb.RPush(ctx, listingStrID, user)
	}
}
