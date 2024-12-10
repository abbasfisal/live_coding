package database

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"live_coding/internal/user/entity"

	"os"

	"sync"
)

func GenerateData() {
	Migrate()

	file, err := os.Open("users_data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	userChannel := make(chan entity.User, 100)
	processedChannel := make(chan bool, 100)

	workerLimit := 10
	sem := make(chan struct{}, workerLimit)

	var wg sync.WaitGroup

	decoder := json.NewDecoder(file)
	if token, err := decoder.Token(); err != nil || token != json.Delim('[') {
		fmt.Println("Expected start of JSON array:", err)
		return
	}

	for decoder.More() {
		var user entity.User
		if err := decoder.Decode(&user); err != nil {
			fmt.Println("Error decoding JSON object:", err)
			continue
		}

		wg.Add(1)

		go func(user entity.User) {
			defer wg.Done()
			userChannel <- user
		}(user)
	}

	if token, err := decoder.Token(); err != nil || token != json.Delim(']') {
		fmt.Println("Expected end of JSON array:", err)
	}

	go func() {
		batchSize := 100
		var batchUser []entity.User

		for user := range userChannel {
			sem <- struct{}{}

			batchUser = append(batchUser, user)

			if len(batchUser) >= batchSize {
				if err := insertUserAndAddresses(batchUser); err != nil {
					fmt.Println("Error inserting user:", err)
				}
				batchUser = nil
			}
			<-sem
		}

		if len(batchUser) > 0 {
			if err := insertUserAndAddresses(batchUser); err != nil {
				fmt.Println("Error inserting user:", err)
			}
		}

		processedChannel <- true
	}()

	wg.Wait()
	close(userChannel)
	<-processedChannel
	close(processedChannel)

	fmt.Println("Processing completed!")
}

func insertUserAndAddresses(users []entity.User) error {

	db := Get()
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.CreateInBatches(&users, 100).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting users: %v", err)
	}

	var addresses []entity.Address
	for _, user := range users {
		for _, address := range user.Addresses {
			address.UserID = user.ID
			if address.ID == uuid.Nil {
				address.ID = uuid.New()
			}
			addresses = append(addresses, address)
		}
	}

	if err := tx.CreateInBatches(&addresses, 100).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting addresses: %v", err)
	}

	tx.Commit()
	return nil
}
