package service

import (
	"fmt"
	"math/rand"
	"regexp"
	"unicode"

	"github.com/google/uuid"
)

func checkLogin(login string) bool {
	regularNum := regexp.MustCompile(`[0-9]`)
	regularStr := regexp.MustCompile(`[a-zA-Z]`)
	if len(login) < 8 || !regularNum.MatchString(login) || !regularStr.MatchString(login) {
		return false
	}
	return true
}

func checkPassword(password string) bool {
	regularNum := regexp.MustCompile(`[0-9]`)
	regularLower := regexp.MustCompile(`[a-z]`)
	regularUpper := regexp.MustCompile(`[A-Z]`)
	if len(password) < 8 ||
		!regularLower.MatchString(password) ||
		!regularUpper.MatchString(password) ||
		!regularNum.MatchString(password) {
		return false
	}
	for _, value := range password {
		if !unicode.IsLetter(value) && !unicode.IsDigit(value) {
			return true
		}
	}
	return false
}

func generateToken() string {
	var token string
	for len(token) < 20 {
		randomInt := rand.Intn(122-48) + 48
		if randomInt > 47 && randomInt < 58 {
			token += string(rune(randomInt))
		} else {
			if randomInt > 64 && randomInt < 91 {
				token += string(rune(randomInt))
			} else {
				if randomInt > 96 && randomInt < 123 {
					token += string(rune(randomInt))
				}
			}
		}
	}
	return token
}

func prepareGetDocumentKey(documentID uuid.UUID) string {
	return fmt.Sprintf("document_ID:%s", documentID.String())
}

func prepareGetUserIDKey(token string) string {
	return fmt.Sprintf("get_user_id_by_token:%s", token)
}

func prepareGetUserKey(id uuid.UUID) string {
	return fmt.Sprintf("get_user_id:%s", id.String())
}

func prepareCheckGrantKey(documentID uuid.UUID, login string) string {
	return fmt.Sprintf("document_ID:%s:login:%s", documentID.String(), login)
}
