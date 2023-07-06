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

type RentalFilter struct {
	PriceMin int                 `json:"price_min"`
	PriceMax int                 `json:"price_max"`
	Limit    int                 `json:"limit"`
	Offset   int                 `json:"offset"`
	Ids      []int               `json:"ids"`
	Near     []map[string]string `json:"near"`
	Sort     string              `json:"sort"`
}
