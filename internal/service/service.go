package service

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/Alina9496/documents/internal/domain"
	"github.com/Alina9496/documents/internal/repo"
	"github.com/Alina9496/documents/internal/service/dto"
	"github.com/Alina9496/tool/pkg/logger"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

type Service struct {
	repo  Repository
	cache Cache
	log   *logger.Logger
}

func New(
	r Repository,
	cache Cache,
	log *logger.Logger,
) *Service {
	return &Service{
		repo:  r,
		cache: cache,
		log:   log,
	}
}

func (s *Service) isUserExists(ctx context.Context, user *domain.User) bool {
	id, err := s.repo.CheckUser(ctx, user)
	if err != nil || id == uuid.Nil {
		return false
	}
	user.ID = id
	return true
}

func (s *Service) Registration(ctx context.Context, user *domain.User) (string, error) {
	l := s.log.WithField("service_method", "Registration")
	if user == nil {
		l.Warn(ErrUserIsNil.Error())
		return "", ErrUserIsNil
	}

	if !checkLogin(user.Login) {
		l.Warn(ErrUserLoginIncorected.Error())
		return "", ErrUserLoginIncorected
	}

	if !checkPassword(user.Password) {
		l.Warn(ErrUserPasswordIncorected.Error())
		return "", ErrUserPasswordIncorected
	}

	if s.isUserExists(ctx, user) {
		l.WithError(ErrUserExists).Error("error when check user")
		return "", fmt.Errorf("error when check user: %w", ErrUserExists)
	}

	err := s.repo.Registration(ctx, user)
	if err != nil {
		l.WithError(err).Error("error when registration user")
		return "", fmt.Errorf("error when registration user: %w", ErrRegistrationUser)
	}

	return user.Login, nil
}

func (s *Service) Authentication(ctx context.Context, user *domain.User) (string, error) {
	l := s.log.WithField("service_method", "Authentication")
	if user == nil {
		l.Warn(ErrUserIsNil.Error())
		return "", ErrUserIsNil
	}

	if !checkLogin(user.Login) {
		l.Warn(ErrUserLoginIncorected.Error())
		return "", ErrUserLoginIncorected
	}

	if !checkPassword(user.Password) {
		l.Warn(ErrUserPasswordIncorected.Error())
		return "", ErrUserPasswordIncorected
	}

	if !s.isUserExists(ctx, user) {
		l.WithError(ErrUserNotFound).Error("error when check user")
		return "", fmt.Errorf("error when check user: %w", ErrUserNotFound)
	}

	user.Token = generateToken()
	err := s.repo.Authentication(ctx, user)
	if err != nil {
		l.WithError(err).Error("error when authentication user")
		return "", fmt.Errorf("error when authentication user: %w", ErrAuthenticationUser)
	}

	return user.Token, nil
}

func (s *Service) LogOut(ctx context.Context, token string) error {
	l := s.log.WithField("service_method", "LogOut")

	err := s.repo.LogOut(ctx, token)
	if err != nil {
		if errors.Is(err, repo.ErrTokenNotFound) {
			return ErrTokenNotFound
		}
		l.WithError(err).Error("error when logout user")
		return fmt.Errorf("error when logout user: %w", ErrLogOutUser)
	}
	return nil
}

func (s *Service) Upload(ctx context.Context, document *dto.Document) (name string, err error) {
	l := s.log.WithField("service_method", "Upload")

	userID, err := s.getUserID(ctx, document.Token)
	if err != nil {
		l.WithError(err).Error("error get user id")
		return "", ErrTokenNotFound
	}

	err = s.repo.ExecTx(ctx, func(ctx context.Context) error {
		documentID, err := s.repo.Save(ctx, toDocument(userID, document))
		if err != nil {
			l.WithError(err).Error("error save document")
			return err
		}

		for _, login := range document.Grant {
			err = s.repo.AddGrant(ctx, toGrant(login, userID, documentID))
			if err != nil {
				l.WithError(err).Error("error add grant")
				return err
			}
		}

		return nil
	})
	if err != nil {
		return "", err
	}
	return document.Name, nil
}

func (s *Service) GetDocument(ctx context.Context, documentID uuid.UUID, token string) (*domain.Document, error) {
	l := s.log.WithField("service_method", "GetDocument")

	document, err := s.getDocument(ctx, documentID)
	if err != nil {
		l.WithError(err).Error("error get document")
		return nil, ErrDocumentNotFound
	}

	if document.Public {
		return document, nil
	}

	userID, err := s.getUserID(ctx, token)
	if err != nil {
		l.WithError(err).Error("error get user id")
		return nil, ErrUserNotFound
	}

	if userID == document.UserID {
		return document, nil
	}

	user, err := s.getUserByID(ctx, userID)
	if err != nil {
		l.WithError(err).Error("error get user")
		return nil, ErrUserNotFound
	}

	isAccess, err := s.checkGrant(ctx, documentID, user.Login)
	if err != nil {
		l.WithError(err).Error("error check grant")
		return nil, err
	}

	if isAccess {
		return document, nil
	}

	return nil, ErrNoAccess
}

func (s *Service) GetDocuments(ctx context.Context, filter *dto.GetDocumentsRequest) ([]domain.Document, error) {
	l := s.log.WithField("service_method", "GetDocuments")

	userID, err := s.getUserID(ctx, filter.Token)
	if err != nil {
		l.WithError(err).Error("error get user id")
		return nil, ErrUserNotFound
	}

	documents, err := s.repo.GetDocuments(ctx, toGetDocuments(userID, filter))
	if err != nil {
		l.WithError(err).Error("error get documents")
		return nil, ErrDocumentsNotFound
	}

	sort.SliceStable(documents, func(i, j int) bool {
		if documents[i].Name < documents[j].Name {
			return true
		} else if documents[i].Name > documents[j].Name {
			return false
		}

		return documents[i].CreatedAt.Before(documents[j].CreatedAt)
	})

	return documents, nil
}

func (s *Service) DeleteDocument(ctx context.Context, id uuid.UUID, token string) (uuid.UUID, error) {
	l := s.log.WithField("service_method", "DeleteDocument")

	userID, err := s.getUserID(ctx, token)
	if err != nil {
		l.WithError(err).Error("error get user id")
		return uuid.Nil, ErrUserNotFound
	}

	id, err = s.repo.DeleteDocument(ctx, id, userID)
	if err != nil {
		l.WithError(err).Error("error get document")
		return uuid.Nil, ErrDocumentNotFound
	}

	return id, nil
}

func (s *Service) getDocument(ctx context.Context, documentID uuid.UUID) (*domain.Document, error) {
	key := prepareGetDocumentKey(documentID)
	doc, exist := s.cache.Get(key)
	document, ok := doc.(*domain.Document)
	if !exist || !ok {
		document, err := s.repo.GetDocument(ctx, documentID)
		if err != nil {
			return nil, ErrDocumentNotFound
		}

		s.cache.Set(key, document, cache.DefaultExpiration)
		return document, nil
	}
	return document, nil
}

func (s *Service) getUserID(ctx context.Context, token string) (uuid.UUID, error) {
	key := prepareGetUserIDKey(token)
	userCash, exist := s.cache.Get(key)
	userID, ok := userCash.(uuid.UUID)
	if !exist || !ok {
		userID, err := s.repo.GetUserID(ctx, token)
		if err != nil {
			return uuid.Nil, ErrUserNotFound
		}
		s.cache.Set(key, userID, cache.DefaultExpiration)
		return userID, nil
	}

	return userID, nil
}

func (s *Service) getUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	key := prepareGetUserKey(id)
	userCash, exist := s.cache.Get(key)
	user, ok := userCash.(*domain.User)
	if !exist || !ok {
		user, err := s.repo.GetUser(ctx, id)
		if err != nil {
			return nil, ErrUserNotFound
		}
		s.cache.Set(key, user, cache.DefaultExpiration)
		return user, nil
	}

	return user, nil
}

func (s *Service) checkGrant(ctx context.Context, documentID uuid.UUID, login string) (bool, error) {
	key := prepareCheckGrantKey(documentID, login)
	isAccessCash, exist := s.cache.Get(key)
	isAccess, ok := isAccessCash.(bool)
	if !exist || !ok {
		isAccess, err := s.repo.CheckGrant(ctx, documentID, login)
		if err != nil {
			return false, err
		}
		s.cache.Set(key, isAccess, cache.DefaultExpiration)
		return isAccess, nil
	}

	return isAccess, nil
}
