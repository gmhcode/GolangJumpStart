package main

//Set = Set class
type Set map[string]struct{}

func getValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
