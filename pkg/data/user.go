package data

import "github.com/jmoiron/sqlx"

type Usrs struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	DueDate    string `json:"dueDate" db:"due_date"`
	CompltDate string `json:"compltDate" db:"complt_date"`
	Status     int    `json:"status" db:"status"`
}

// UserGet get the user detail by id
func UserGet(db *sqlx.DB, in *Usrs) (*Usrs, error) {
	q := "SELECT id, name, dueDate, compltDate, status FROM users WHERE id = ?"

	usr := &Usrs{}
	err := db.Get(*usr, q, in.ID) //return error if not exist
	if err != nil {
		return nil, err
	}

	return usr, nil
}

// UserCreate insert the user detail.
func UserCreate(db *sqlx.DB, in *Usrs) error {
	q := "INSERT INTO users (name, due_date, complt_date, status) VALUES (:name, :due_date, :complt_date, :status)"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

// UserUpdate update the user detail by id, id will not be update.
func UserUpdate(db *sqlx.DB, in *Usrs) error {
	_, err := UserGet(db, in)
	if err != nil {
		return err
	}

	q := "UPDATE user SET name=:name, due_date=:due_date, complt_date=:complt_date, status=:status WHERE id=:id"

	_, err = db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

// UserDelete delete the user record
func UserDelete(db *sqlx.DB, in *Usrs) error {
	_, err := UserGet(db, in)
	if err != nil {
		return err
	}

	q := "DELETE FROM user WHERE id = ? LIMITE 1"

	_, err = db.Exec(q, in.ID)
	if err != nil {
		return err
	}

	return nil
}
