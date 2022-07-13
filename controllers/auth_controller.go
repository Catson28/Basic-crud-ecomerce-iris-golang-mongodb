package controllers

import (
	"fmt"
	"os"

	"tentativa/datamodels"
	"tentativa/datamodels/request"

	"tentativa/services"
	"tentativa/util"

	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type AuthController struct {
	Ctx     iris.Context
	Service services.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Service: services.NewUserService(),
	}
}

const defaultSecretKey = "sercrethatmaycontainch@r$32chars"

func getSecretKey() string {
	secret := os.Getenv(util.AppName + "_SECRET")
	if secret == "" {
		return defaultSecretKey
	}

	return secret
}

// UserClaims represents the user token claims.
type UserClaims struct {
	UserID string            `json:"user_id"`
	Roles  []datamodels.Role `json:"roles"`
}

// Verificar permite apenas clientes autorizados.
func Verify() iris.Handler {
	secret := getSecretKey()

	verifier := jwt.NewVerifier(jwt.HS256, []byte(secret), jwt.Expected{Issuer: util.AppName})
	// Ativar o recurso de bloco de token do lado do servidor (mesmo antes de seu tempo de expiração):
	verifier.Extractors = []jwt.TokenExtractor{jwt.FromHeader} 
	// extrai o token apenas da Authorization: Bearer $token
	return verifier.Verify(func() interface{} {
		return new(UserClaims)
	})
}

// AllowAdmin permite apenas clientes autorizados com função de acesso "admin".
// Deve ser registrado após Verify.
func (a *AuthController) AllowAdmin(ctx iris.Context) {
	if !IsAdmin(ctx) {
		a.Ctx.StopWithText(iris.StatusForbidden, "admin access required")
		return
	}

	a.Ctx.Next()
}

// SignIn accepts the user form data and returns a token to authorize a client.

func (a *AuthController) PostSignIn() (response datamodels.Response) {
	//fmt.Println("err")
	/*
	 */
	sign := request.SignRequest{}
	//usir := a.Ctx.ReadQuery(&sign)
	dSign := a.Ctx.ReadJSON(&sign)
	if dSign != nil {

		if dSign, ok := dSign.(validator.ValidationErrors); ok {
			fmt.Println("erro grande")
			validationErrors := wrapValidationErrors(dSign)

			response.Code = 40001
			response.Msg = fmt.Sprintf("参数解析失败：%v", dSign)
			response.Data = iris.NewProblem().Title("Validation error").Detail("One or more fields failed to be validated").Type("/api/v1/games/validation-errors").Key("errors", validationErrors)
			//response.Data = usir
			//response.Data = &sign
			return
		}

	}

	secret := getSecretKey()
	signer := jwt.NewSigner(jwt.HS256, []byte(secret), 15*time.Minute)

	//fmt.Println("erro 11111")
	//user, ok := a.Service.GetByUsernameAndPassword(username, password)
	user, ok := a.Service.GetSinger(sign)
	if !ok {
		//fmt.Println("erro de cabetula")
		//fmt.Println(ok)
		//a.Ctx.StopWithText(iris.StatusBadRequest, "wrong username or password")
		response.Code = 40002
		response.Msg = fmt.Sprintf("参数解析失败：%v", user)
		response.Data = "erro de insercao, usuario ou palavra passe incorreta"
		return
	}

	//fmt.Println(firstname)
	//fmt.Println(password)
	//fmt.Println(username)
	//fmt.Println("username")

	claims := UserClaims{
		UserID: user.ID.Hex(),
		Roles:  user.Roles,
	}

	// Optionally, generate a JWT ID.
	jti, err := util.GenerateUUID()
	if err != nil {
		a.Ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	token, err := signer.Sign(claims, jwt.Claims{
		ID:     jti,
		Issuer: util.AppName,
	})

	if err != nil {
		a.Ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	//a.Ctx.Write(token)

	//return a.Service.SignIn(username, password)

	response.Code = 20000
	response.Msg = "success"
	response.Data = token

	return
}

func (a *AuthController) PostRegister() (response datamodels.Response) {
	userRequest := request.UserRequest{}

	user := datamodels.User{}
	err := a.Ctx.ReadJSON(&userRequest)

	if err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := wrapValidationErrors(errs)

			response.Code = 40001
			response.Msg = fmt.Sprintf("参数解析失败：%v", err)
			response.Data = iris.NewProblem().Title("Validation error").Detail("One or user fields failed to be validated").Type("/api/v1/users/validation-errors").Key("errors", validationErrors)

			return
		}

		fmt.Println("errrooooooo")
		a.Ctx.StopWithStatus(iris.StatusInternalServerError)
	}
	/*

	 */
	hashedPassword, err := util.GeneratePassword(userRequest.Password)
	if err != nil {
	}
	user.Password = hashedPassword
	user.Firstname = userRequest.Firstname
	user.Username = userRequest.Username
	user.Roles = userRequest.Roles

	return a.Service.Create(user)
}


func (a *AuthController) GetSignOut() (response datamodels.Response) {
	a.Ctx.Logout()

	response.Code = 20000
	response.Msg = "success"
	response.Data = "Logout"
	return
}

// GetClaims retorna as declarações de cliente autorizadas atuais.
func GetClaims(ctx iris.Context) *UserClaims {
	claims := jwt.Get(ctx).(*UserClaims)
	return claims
}

// GetUserID retorna o ID de usuário do cliente autorizado atual extraído de declarações.
func GetUserID(ctx iris.Context) string {
	return GetClaims(ctx).UserID
}

// IsAdmin relata se o cliente atual tem acesso de administrador.
func IsAdmin(ctx iris.Context) bool {
	for _, role := range GetClaims(ctx).Roles {
		if role == datamodels.Admin {
			return true
		}
	}
	return false
}
