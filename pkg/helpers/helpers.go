package helpers

import (
	"fmt"
	"regexp"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func NewCounter(initialValue int64) *Counter {
	return &Counter{value: initialValue}
}

func (c *Counter) Increment() int64 {
	return atomic.AddInt64(&c.value, 1)
}

// IsValidPhone ...
func IsValidPhone(phone string) bool {
	r := regexp.MustCompile(`^\+998[0-9]{2}[0-9]{7}$`)
	return r.MatchString(phone)
}

// IsValidEmail ...
func IsValidEmail(email string) bool {
	r := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return r.MatchString(email)
}

// IsValidLogin ...
func IsValidLogin(login string) bool {
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,29}$`)
	return r.MatchString(login)
}

// IsValidUUID ...
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// FormatID форматирует число в строку с заданным префиксом и нулевым заполнением
func FormatID(prefix string, value int64) string {
	return fmt.Sprintf("%s%06d", prefix, value)
}

// OrderIDCounter глобальный счетчик для Order ID
var OrderIDCounter = NewCounter(1)

// GetNextOrderID возвращает следующее значение Order ID
func GetNextOrderID() string {
	nextValue := OrderIDCounter.Increment()
	return FormatID("O-", nextValue)
}

// ProductIDCounter глобальный счетчик для Product ID
var ProductIDCounter = NewCounter(1)

// GetNextProductID возвращает следующее значение Product ID
func GetNextProductID() string {
	nextValue := ProductIDCounter.Increment()
	return FormatID("P-", nextValue)
}
