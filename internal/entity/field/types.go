package field

import "database/sql"

func New[K sql.Scanner](t K, v any) K {
	if err := t.Scan(v); err != nil {
		panic(err)
	}

	return t
}
