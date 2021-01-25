package data

import "github.com/jmoiron/sqlx"

type Tasks struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	DueDate    string `json:"dueDate" db:"due_date"`
	CompltDate string `json:"compltDate" db:"complt_date"`
	Status     int    `json:"status" db:"status"`
}

// TaskGet get the task detail by id
func TaskGet(db *sqlx.DB, in *Tasks) (*Tasks, error) {
	q := "SELECT id, `name`, `due_date`, `complt_date`, `status` FROM tasks WHERE id = ?"

	task := &Tasks{}
	err := db.Get(task, q, in.ID) //return error if not exist
	if err != nil {
		return nil, err
	}

	return task, nil
}

// TaskCreate insert the task detail.
func TaskCreate(db *sqlx.DB, in *Tasks) error {
	q := "INSERT INTO tasks (`name`, `due_date`, `complt_date`, `status`) VALUES (:name, :due_date, :complt_date, :status)"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

// TaskUpdate update the task detail by id, id will not be update.
func TaskUpdate(db *sqlx.DB, in *Tasks) error {
	_, err := TaskGet(db, in)
	if err != nil {
		return err
	}

	q := "UPDATE tasks SET `name`=:name, `due_date`=:due_date, `complt_date`=:complt_date, `status`=:status WHERE id=:id"

	_, err = db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

// TaskDelete delete the task record
func TaskDelete(db *sqlx.DB, in *Tasks) error {
	_, err := TaskGet(db, in)
	if err != nil {
		return err
	}

	q := "DELETE FROM tasks WHERE `id` = ?"

	_, err = db.Exec(q, in.ID)
	if err != nil {
		return err
	}

	return nil
}
