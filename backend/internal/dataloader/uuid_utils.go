package dataloader

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// StringToPgUUID converts a string UUID to pgtype.UUID with error handling
func StringToPgUUID(id string) (pgtype.UUID, error) {
	googleUUID, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid UUID: %w", err)
	}
	
	return pgtype.UUID{
		Bytes: googleUUID,
		Valid: true,
	}, nil
}

// PgUUIDToString converts pgtype.UUID to string safely
func PgUUIDToString(pgUUID pgtype.UUID) (string, error) {
	if !pgUUID.Valid {
		return "", fmt.Errorf("invalid pgtype.UUID")
	}
	
	googleUUID, err := uuid.FromBytes(pgUUID.Bytes[:])
	if err != nil {
		return "", fmt.Errorf("failed to convert pgtype.UUID to string: %w", err)
	}
	
	return googleUUID.String(), nil
}

// StringsToPgUUIDs converts a slice of string UUIDs to pgtype.UUIDs
// Returns valid UUIDs and a map of errors by index
func StringsToPgUUIDs(ids []string) ([]pgtype.UUID, map[int]error) {
	uuids := make([]pgtype.UUID, 0, len(ids))
	errors := make(map[int]error)
	
	for i, id := range ids {
		pgUUID, err := StringToPgUUID(id)
		if err != nil {
			errors[i] = err
			continue
		}
		uuids = append(uuids, pgUUID)
	}
	
	return uuids, errors
}

// CreateOrderedResultMap creates a map from database results for ordered reconstruction
func CreateOrderedResultMap[T any](
	results []T, 
	getID func(T) (string, error),
) (map[string]T, error) {
	resultMap := make(map[string]T, len(results))
	
	for _, result := range results {
		id, err := getID(result)
		if err != nil {
			continue // Skip invalid entries
		}
		resultMap[id] = result
	}
	
	return resultMap, nil
}