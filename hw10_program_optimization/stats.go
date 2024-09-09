package hw10programoptimization

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Email string
}

type DomainStat map[string]int

var (
	ErrDomainDoesNotMatch = errors.New("domain does not match")
	ErrIncorrectEmail     = errors.New("incorrect email")
)

func extractSiteName(email string, stats DomainStat) {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		log.Println(ErrIncorrectEmail.Error())
		return
	}
	stats[strings.ToLower(email[atIndex+1:])]++
}

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	stats := make(DomainStat)
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	scanner := bufio.NewScanner(r)
	var user User

	for scanner.Scan() {
		err := json.Unmarshal(scanner.Bytes(), &user)
		if err != nil {
			return nil, fmt.Errorf("error with unmarshalling the string %v", scanner.Text())
		}

		if !strings.HasSuffix(user.Email, domain) {
			continue
		}

		extractSiteName(user.Email, stats)
	}
	return stats, nil
}

func GetDomainStatOld(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var user User
		if err = json.Unmarshal([]byte(line), &user); err != nil {
			return
		}
		result[i] = user
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		matched, err := regexp.Match("\\."+domain, []byte(user.Email))
		if err != nil {
			return nil, err
		}

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
