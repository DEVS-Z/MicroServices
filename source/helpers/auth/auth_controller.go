package auth

type AuthController struct {
	authService *AuthService
}

// func injection() *AuthController {
// 	return &AuthController{
// 		authService: &AuthService{},
// 	}
// }

// func SignIn(res http.ResponseWriter, req *http.Request) {
// 	controller := injection()

// 	println("SignIn called with request body:", req.Body)
// 	ctx, err := client.Init(res, req)
// 	if err != nil {
// 		_ = ctx.InternalServerError(err)
// 		return
// 	}

// 	dto, err := client.GetDTO[*AuthDTO](req.Context(), string(middleware.DtoKey))
// 	if err != nil {
// 		_ = ctx.Error("Invalid validated data", err)
// 		return
// 	}

// 	token, err := controller.authService.signIn(req.Context(), dto.Username, dto.Email, dto.Password)
// 	if err != nil {
// 		_ = ctx.Error("Authentication failed", err)
// 		return
// 	}

// 	data := map[string]any{"token": token}
// 	_ = ctx.Success("Sign in successful", []any{data})
// }

// func SignUp(res http.ResponseWriter, req *http.Request) {
// 	controller := injection()
// 	ctx, err := client.Init(res, req)
// 	if err != nil {
// 		_ = ctx.InternalServerError(err)
// 		return
// 	}

// 	dto, err := client.GetDTO[*AuthDTO](req.Context(), string(middleware.DtoKey))
// 	if err != nil {
// 		_ = ctx.InternalServerError(err)
// 		return
// 	}

// 	result, err := controller.authService.signUp(req.Context(), dto.Username, dto.Email, dto.Password)
// 	if err != nil {
// 		_ = ctx.Error("Authentication failed", err)
// 		return
// 	}

// 	_ = ctx.Success("Sign up successful", []any{result})

// }
