package database

func AccessTokenIsValid(token string) bool {

	var userId string

	err := DB.QueryRow("SELECT userId FROM UserToken WHERE token=?", token).Scan(&userId)

	if err != nil {
		return false
	}

	return true

}
