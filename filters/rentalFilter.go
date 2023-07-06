package filters

/*
Supported Query Filters:

- price_min (number)
- price_max (number)
- limit (number)
- offset (number)
- ids (comma separated list of rental ids)
- near (comma separated pair [lat,lng])
- sort (string)
*/

//Making them all strings to start for ease of coding when converting the parameter values
type RentalFilter struct {
	PriceMin string   `json:"price_min"`
	PriceMax string   `json:"price_max"`
	Limit    int      `json:"limit"`
	Offset   int      `json:"offset"`
	Ids      []string `json:"ids"`
	Near     string   `json:"near"`
	Sort     string   `json:"sort"`
}
