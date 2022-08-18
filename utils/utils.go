package utils

import (
	"log"
	"os"

	"golang.org/x/exp/constraints"
)

func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
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
	if len(elems) == len(GetUnique(elems)) {
		return true
	}

	return false
}

func GetUnique[T comparable](arr []T) []T {
	occurred := map[T]bool{}
	result := []T{}
	for e := range arr {
		if occurred[arr[e]] != true {
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
