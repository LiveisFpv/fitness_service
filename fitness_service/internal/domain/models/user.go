package models

type User struct {
	User_id             int
	User_firstName      string
	User_lastName       string
	User_middleName     *string
	User_birthday       string
	User_height         int
	User_weight         float64
	User_fitness_target string
	User_sex            bool
	User_hypertain      bool
	User_diabet         bool
	User_level          *int
}

// Error implements error.
func (u *User) Error() string {
	panic("unimplemented")
}
