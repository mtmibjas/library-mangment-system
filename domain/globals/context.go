package globals

import (
	"context"
	"crypto/rand"
	"fmt"
)

// Context key type to be used with contexts.
type contextKey string

// UUIDKey is the universally unique identifier key to be used with context.
const UUIDKey contextKey = "UUID"

// PrefixKey is the key to add an additional prefix value to the context.
const PrefixKey contextKey = "Prefix"

// TxKey is the key to attach a database transaction to the context.
const TxKey contextKey = "Tx"

// AppendToContextPrefix appends the given prefix string to the globals.PrefixKey.
func AppendToContextPrefix(ctx context.Context, prefix string) context.Context {

	pfx := ctx.Value(PrefixKey)

	if pfx == nil {
		return context.WithValue(ctx, PrefixKey, prefix)
	}

	return context.WithValue(ctx, PrefixKey, pfx.(string)+"."+prefix)
}
func GenerateAPIKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}
