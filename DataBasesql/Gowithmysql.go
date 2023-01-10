package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cover struct {
	Id   int
	Name string
}

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("mysql", "root:42085344720062546@(127.0.0.1:3306)/coursedb")
	if err != nil {
		panic(err)
	}
	// Add and Update
	// cover := Cover{9, "Cranberry"}
	// // err = AddCover(cover)
	// err = UpdateCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	// Delete
	// err = DeleteCover(9)
	// if err != nil {
	// 	panic(err)
	// }

	// Read
	// covers, err := GetCovers()
	// if err != nil {
	// 	panic(err)
	// }

	// 	cover, err := GetCover(1)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(cover)

	// get coversx
	// covers, err := GetcoversX()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//get coverx
	// cover, err := GetcoverX(1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cover)

	// forloop value in covers {}
	// for _, v := range covers {
	// 	fmt.Println(v)
	// }

}
func GetcoversX() ([]Cover, error) {
	query := "SELECT id,name FROM new_table"
	covers := []Cover{}
	err := db.Select(&covers, query)
	if err != nil {
		return nil, err
	}
	return covers, nil
}
func GetcoverX(id int) (*Cover, error) {
	query := "SELECT id ,name FROM new_table WHERE id = ?"
	cover := Cover{}
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	query := "SELECT id,name FROM new_table"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}
		covers = append(covers, cover)
	}
	return covers, nil
}
func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	query := "SELECT id,name FROM new_table WHERE id=?"
	row := db.QueryRow(query, id)
	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}
func AddCover(cover Cover) error {
	query := "INSERT INTO new_table (id,name) VALUES (?,?)"
	result, err := db.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("can not insert")
	}
	return nil
}
func UpdateCover(cover Cover) error {
	query := "UPDATE new_table SET name = ? WHERE id = ?"
	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("can not update")
	}
	return nil
}
func DeleteCover(id int) error {
	query := "DELETE FROM new_table WHERE id=?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("can not delete")
	}
	return nil
}

//transection

func AddCoverTransection(cover Cover) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	query := "INSERT INTO new_table (id,name) VALUES (?,?)"
	result, err := tx.Exec(query, cover.Id, cover.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affect <= 0 {
		return errors.New("can not insert")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func UpdateCoverTransection(cover Cover) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE new_table SET name = ? WHERE id = ?"
	result, err := tx.Exec(query, cover.Name, cover.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affect <= 0 {
		return errors.New("can not update")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func DeleteCoverTransection(id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM new_table WHERE id=?"
	result, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affect <= 0 {
		return errors.New("can not delete")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
