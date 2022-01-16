package repository

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuranRepo_ListSuratL10N(t *testing.T) {
	ctx := context.TODO()
	repo := NewQuran(theDb)
	result, err := repo.ListSuratL10N(ctx, 1, 0, 10)
	log.Printf("%+v\n", result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, len(result), 10)
}

func TestQuranRepo_ListAyatL10N(t *testing.T) {
	ctx := context.TODO()
	repo := NewQuran(theDb)
	result, err := repo.ListAyatL10N(ctx, 1, 2, 6, 10)
	log.Printf("%+v\n", result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, len(result), 10)

}
