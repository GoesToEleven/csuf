package main

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/scrypt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type User struct {
	Email        string
	PasswordSalt []byte
	PasswordHash []byte
}

const (
	// scrypt is used for strong keys
	// these are the recommended scrypt parameters
	scryptN      = 16384
	scryptR      = 8
	scryptP      = 1
	scryptKeyLen = 32
)

// AuthenticateUser checks the email & password against what's in the database
func AuthenticateUser(ctx context.Context, email, password string) (*User, error) {
	key := datastore.NewKey(ctx, "User", email, 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	hash, err := scrypt.Key([]byte(password), user.PasswordSalt, scryptN, scryptR, scryptP, scryptKeyLen)
	if err != nil {
		return nil, err
	}

	if subtle.ConstantTimeCompare(hash, user.PasswordHash) != 1 {
		return nil, fmt.Errorf("invalid username or password")
	}

	return &user, nil
}

// CreateUser creates a new user record
func CreateUser(ctx context.Context, email, password string) error {
	if email == "" {
		return fmt.Errorf("invalid email")
	}
	if password == "" {
		return fmt.Errorf("invalid password")
	}
	key := datastore.NewKey(ctx, "User", email, 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	salt := make([]byte, 16)
	rand.Read(salt)
	hash, err := scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, scryptKeyLen)
	if err != nil {
		return err
	}

	user = User{
		Email:        email,
		PasswordSalt: salt,
		PasswordHash: hash,
	}

	_, err = datastore.Put(ctx, key, &user)
	return err
}

// UpdateUser updates a user record
func UpdateUser(ctx context.Context, user *User) error {
	key := datastore.NewKey(ctx, "User", user.Email, 0, nil)
	_, err := datastore.Put(ctx, key, user)
	return err
}
