package data

import (
	"fmt"
)

const resTable = "resolutionTypes"
type ResolutionType struct {
	Id int
	Name string
}

func NewResolutionType(id int, name string) (* ResolutionType) {
	return &ResolutionType{
		Id: id,
		Name: name,
	}
}

func getResolutionTypeById(id int) (* ResolutionType, error) {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE id = ?;", resTable);
	rows, err := runQuery(tableq, id)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err)

	r := make([]ResolutionType, 0)
	for rows.Next() {
		cR := ResolutionType{}
		err = rows.Scan(&cR.Id, &cR.Name) 
		if err != nil {
			logerr(err)
			return nil, err
		}
		r = append(r, cR);
	}

	 return NewResolutionType(r[0].Id, r[0].Name), nil;	

}