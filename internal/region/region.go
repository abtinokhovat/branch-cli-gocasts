package region

import "fmt"

const StoragePath = "storage/region.json"

type Region struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r *Region) String() string {
	return fmt.Sprintf("#%d: %s", r.Id, r.Name)
}
