package api

import (
	"net/http"
	"thelastking/kingseafood/api/common"
	"thelastking/kingseafood/controller/users_bussiness"
	"thelastking/kingseafood/model"
	"thelastking/kingseafood/model/req_users"
	"thelastking/kingseafood/pkg/security"
	"thelastking/kingseafood/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SignUpHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqSignUp req_users.RequestSignUp
		if err := c.ShouldBind(&reqSignUp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't request sign up",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(reqSignUp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't validator",
			})
			return
		}

		idUsers, err := uuid.NewUUID()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "uuid fails",
			})
			return
		}
		time := time.Now().UTC()
		pwd := security.HashAndSalt([]byte(reqSignUp.Password))
		role := model.MEMBERS.String()
		users := model.Users{
			UserID:    idUsers.String(),
			FullName:  reqSignUp.FullName,
			Email:     reqSignUp.Email,
			Password:  pwd,
			Male:      reqSignUp.Male,
			Role:      role,
			CreatedAt: &time,
			UpdatedAt: &time,
		}
		biz := users_bussiness.NewSignUpController(service.NewSql(db))
		dataUser, err := biz.NewSignUp(c.Request.Context(), &users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"comment": "Invalid database users",
			})
			return
		}
		c.JSON(http.StatusOK, common.ReponseData(dataUser))
	}
}

// SIGNIN
func SignInHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqSignIn req_users.RequestSignIn
		if err := c.ShouldBind(&reqSignIn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "reqSignIn failed",
			})
			return
		}
		validate := validator.New()
		if err := validate.Struct(reqSignIn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "Can't validate",
			})
			return
		}

		users := model.Users{
			Email:    reqSignIn.Email,
			Password: reqSignIn.Password,
		}

		biz := users_bussiness.NewSignInController(service.NewSql(db))
		if err := biz.NewSignIn(c.Request.Context(), &users); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"comment": "reqSignIn database failed",
			})
			return
		}

		// Kiểm tra mật khẩu
		isValidPassword := security.ComparePasswords(users.Password, []byte(reqSignIn.Password))
		if !isValidPassword {
			c.JSON(http.StatusBadRequest, gin.H{
				"comment": "Email or password is incorrect",
			})
			return
		}

		c.JSON(http.StatusOK, common.ReponseData("Đăng nhập thành công"))
	}
}
