package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/models"
)

// Register
func RegisterHandler(ctx *gin.Context, conn *pgx.Conn) {
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	hashedPassword := models.HashPassword(input.Password)

	query := `INSERT INTO users (fullname, email, password, phone, address, roles_id) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := conn.Exec(context.Background(), query, input.Fullname, input.Email, hashedPassword, input.Phone, input.Address, 2)

	if err != nil {
		ctx.JSON(500, models.Response{
			Status:  false,
			Message: "Failed to register user to database",
		})
		return
	}

	ctx.JSON(201, models.Response{
		Status:  true,
		Message: "User registered successfully",
	})
}

// Login
func LoginHandler(ctx *gin.Context, conn *pgx.Conn) {
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{Status: false, Message: "Invalid request body"})
		return
	}

	var storedUser models.User
	query := `SELECT id_user, fullname, email, password FROM users WHERE email = $1`
	err := conn.QueryRow(context.Background(), query, input.Email).Scan(
		&storedUser.Id, &storedUser.Fullname, &storedUser.Email, &storedUser.Password,
	)

	if err != nil {
		ctx.JSON(404, models.Response{Status: false, Message: "Wrong Email or Password"})
		return
	}

	if !models.VerifyPassword(storedUser.Password, input.Password) {
		ctx.JSON(401, models.Response{Status: false, Message: "Wrong Email or Password"})
		return
	}

	ctx.JSON(200, models.Response{Status: true, Message: "Login successful", Data: storedUser})
}

// Get All Users
func GetAllUsersHandler(ctx *gin.Context, conn *pgx.Conn) {
	// Tambahkan roles_id dan password ke dalam query SELECT
	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture 
			  FROM users JOIN roles on roles.id_roles = users.roles_id`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		ctx.JSON(500, models.Response{Status: false, Message: "Gagal mengambil data database"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.Id,
			&u.RolesId,
			&u.Fullname,
			&u.Email,
			&u.Password,
			&u.Address,
			&u.Phone,
			&u.Picture,
		)
		if err != nil {
			continue
		}
		users = append(users, u)
	}

	ctx.JSON(200, models.Response{
		Status:  true,
		Message: "Success get all users",
		Data:    users,
	})
}

// Get User By ID
func GetUserByIdHandler(ctx *gin.Context, conn *pgx.Conn) {
	idParam := ctx.Param("id")
	var u models.User

	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture 
	          FROM users JOIN roles on roles.id_roles = users.roles_id
			  WHERE id_user = $1`

	err := conn.QueryRow(context.Background(), query, idParam).Scan(
		&u.Id, 
		&u.RolesId, 
		&u.Fullname, 
		&u.Email, 
		&u.Password, 
		&u.Address, 
		&u.Phone, 
		&u.Picture,
	)

	if err != nil {
		ctx.JSON(404, models.Response{Status: false, Message: "User not found!"})
		return
	}

	ctx.JSON(200, models.Response{
		Status:  true,
		Message: "User Found!",
		Data:    u,
	})
}

// Update User
func UpdateUserHandler(ctx *gin.Context, conn *pgx.Conn) {
	idParam := ctx.Param("id")
	var input models.User
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{Status: false, Message: "Invalid request body"})
		return
	}

	query := `UPDATE users SET fullname=$1, phone=$2, address=$3 WHERE id_user=$4`
	_, err := conn.Exec(context.Background(), query, input.Fullname, input.Phone, input.Address, idParam)

	if err != nil {
		ctx.JSON(500, models.Response{Status: false, Message: "Failed to update user"})
		return
	}

	ctx.JSON(200, models.Response{Status: true, Message: "User updated successfully!"})
}

// Delete User
func DeleteUserHandler(ctx *gin.Context, conn *pgx.Conn) {
	idParam := ctx.Param("id")

	query := `DELETE FROM users WHERE id_user = $1`
	result, err := conn.Exec(context.Background(), query, idParam)

	if err != nil || result.RowsAffected() == 0 {
		ctx.JSON(404, models.Response{Status: false, Message: "User not found"})
		return
	}

	ctx.JSON(200, models.Response{Status: true, Message: "User Deleted successfully!"})
}
