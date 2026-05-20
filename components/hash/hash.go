package hash

func HashPassword(password string) uint64 {
	const (
		offset uint64 = 14695981039346656037
		prime  uint64 = 1099511628211
	)

	hash := offset

	for i := 0; i < len(password); i++ {
		hash ^= uint64(password[i])
		hash *= prime
	}

	hash ^= hash >> 32
	hash *= 0xd6e8feb86659fd93
	hash ^= hash >> 32
	return hash
}

func Verify(input string, expected uint64) bool {
	return HashPassword(input) == expected
}
