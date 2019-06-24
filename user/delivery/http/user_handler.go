package http


import (
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"qlass-native/user"
	"qlass-native/models"
)


// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type UserHandler struct {
	UserUsecase user.Usecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewUserHandler(e *echo.Echo, usecase article.Usecase) {
	handler := &UserHandler{
		AUsecase: usecase,
	}

	http.handleFunc("/articles",handler.Store).method("POST")
	// e.GET("/articles", handler.FetchArticle)
	// e.POST("/articles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}


func isRequestValid(model *models.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the user by given request body
func (a *UserHandler) Store(res http.ResponseWriter, req *http.Request) error {
	 res.Header().Set("Content-Type", "application/json")
	var user models.User
	err := req.Bind(&user)
	if err != nil {
		return http.Error(res, err.Error(), http.StatusUnprocessableEntity)		
	}

	if ok, err := isRequestValid(&user); !ok {
		return http.Error(res, err.Error(), http.StatusBadRequest)				
	}

r.Body	
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.AUsecase.Store(ctx, &article)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, article)
}