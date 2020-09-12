package queries

const (
	GET_ALL_LEVEL = `SELECT * FROM tb_level WHERE level_status = 1`
	GET_BY_ID_LEVEL = `SELECT * FROM tb_level WHERE level_id = ? AND level_status = 1`
	CREATE_LEVEL = `INSERT INTO tb_level VALUES (NULL, ?, 1)`
	UPDATE_LEVEL = `UPDATE tb_level
					SET level_name = ?
					WHERE level_id = ? AND level_status = 1`
	DELETE_LEVEL = `UPDATE tb_level
					SET level_status = 0
					WHERE level_id = ?`
	GET_ALL_USER = `SELECT *
					FROM tb_user
					WHERE user_status = 1 AND
					(user_f_name LIKE ? OR user_l_name LIKE ?)
					ORDER BY 4
					LIMIT %v, %v`
	GET_TOTAL_USER = `SELECT COUNT(*) FROM tb_user WHERE user_status = 1`
	GET_BY_ID_USER = `SELECT * FROM tb_user WHERE user_id = ? AND user_status = 1`
	GET_BY_EMAIL_USER = `SELECT * FROM tb_user WHERE user_email = ? AND user_status = 1`
	CREATE_USER = `INSERT INTO tb_user VALUES (?, ?, ?, ?, ?, ?, 0, ?, 1)`
	CREATE_USER_BY_ADMIN = `INSERT INTO tb_user VALUES (?, ?, ?, ?, ?, ?, ?, ?, 1)`
	UPDATE_USER = `UPDATE tb_user
					SET user_email = ?,
						user_f_name = ?,
						user_l_name = ?,
						user_gender = ?,
						user_balance = ?,
						user_level = ?
					WHERE user_id = ? AND user_status = 1`
	DELETE_USER = `UPDATE tb_user
					SET user_status = 0
					WHERE user_id = ?`


)
