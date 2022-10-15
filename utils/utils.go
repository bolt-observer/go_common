package utils

import (
	"encoding/base64"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func SafeBase64Decode(s string) ([]byte, error) {
	s = strings.TrimRight(s, "=")
	addendum := strings.Repeat("=", (4-(len(s)%4))%4)

	return base64.StdEncoding.DecodeString(s + addendum)
}

func SetDiffGeneric[T comparable, V any](orig, changed map[T]V) map[T]V {
	result := make(map[T]V)

	for k := range changed {
		if _, exists := orig[k]; !exists {
			result[k] = changed[k]
		}
	}

	return result
}

func SetDiff[T comparable](orig, changed map[T]struct{}) map[T]struct{} {
	result := make(map[T]struct{})

	for k, _ := range changed {
		if _, exists := orig[k]; !exists {
			result[k] = struct{}{}
		}
	}

	return result
}

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func AreElementsUnique[T comparable](elems []T) bool {
	return len(elems) == len(GetUnique(elems))
}

func GetUnique[T comparable](arr []T) []T {
	occurred := map[T]bool{}
	result := []T{}
	for e := range arr {
		if !occurred[arr[e]] {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func GetEnvWithDefault(key, def string) string {
	result := os.Getenv(key)
	if result == "" {
		return def
	}
	return result
}

func GetEnv(key string) string {
	result := os.Getenv(key)
	if result == "" {
		log.Fatalf("Environment variable %s not found", key)
		return ""
	}
	return result
}

func Max[T constraints.Ordered](s ...T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func Min[T constraints.Ordered](s ...T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}
