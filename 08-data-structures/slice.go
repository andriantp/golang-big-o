package main

func findInSlice(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}