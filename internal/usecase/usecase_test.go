package usecase_test

import (
	"errors"
	"testing"
	"url-shortener/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetURL(alias string) (string, error) {
	args := m.Called(alias)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) SaveURL(alias, url string) error {
	args := m.Called(alias, url)
	return args.Error(0)
}

func TestGetURL(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := usecase.New(mockRepo)

	t.Run("should return URL when alias exists", func(t *testing.T) {
		mockRepo.On("GetURL", "shortAlias").Return("https://example.com", nil)

		url, err := usecase.GetURL("shortAlias")

		assert.NoError(t, err)
		assert.Equal(t, "https://example.com", url)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when alias does not exist", func(t *testing.T) {
		mockRepo.On("GetURL", "non_existent_alias").Return("", errors.New("alias not found"))

		url, err := usecase.GetURL("non_existent_alias")

		assert.Error(t, err, "Expected an error for non-existing alias")
		assert.Equal(t, "", url, "Expected empty string for non-existing alias")

		mockRepo.AssertExpectations(t)
	})
}

func TestSaveURL(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := usecase.New(mockRepo)

	t.Run("should return alias when URL is saved successfully", func(t *testing.T) {
		mockRepo.On("SaveURL", mock.AnythingOfType("string"), "https://example.com").Return(nil)

		alias, err := usecase.SaveURL("https://example.com")

		assert.NoError(t, err)
		assert.Len(t, alias, 10)
		mockRepo.AssertExpectations(t)
	})

}
