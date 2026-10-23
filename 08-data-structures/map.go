package main

func findInMap(items map[string]struct{}, target string) bool {
	_, ok := items[target]
	return ok
}
