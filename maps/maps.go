package maps

//return a map with key type string and value int
func GetMap(name string) int {
	maptTest := map[string]int{
		"Juan":   18,
		"Yohan":  24,
		"Freddy": 31,
	}
	maptTest["llave1"] = 1
	maptTest["llave12"] = 100
	delete(maptTest, "llave1")
	delete(maptTest, "llave12")

	return maptTest[name]
}
