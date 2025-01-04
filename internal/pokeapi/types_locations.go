package pokeapi

type respLocations struct {
	Count 		int
	Next 		*string 
	Previous 	*string
	Results 	[]struct{
		Name 	string
		URL 	string
	}
}