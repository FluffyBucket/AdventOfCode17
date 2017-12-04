package code

import "io/ioutil"

func loadFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(dat)
}
