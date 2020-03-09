package controllers

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func NewUserController() User {
	return User{}
}

func (u *User) Create(name string) *User {
	u.FirstName = name
	u.LastName = "takahashi"
	u.ID = 829
	return u
}

// type UserController struct {
//     Interactor usecase.UserInteractor
// }

// func NewUserController(SqlHandler database.SqlHandler) *UserController {
//     return &UserController{
//         Interactor: usecase.UserInteractor{
//             UserRepository: &database.UserRepository{
//                 SqlHandler: SqlHandler
//             }
//         }
//     }
// }

// func (controller *UserController) Create() {
// 	u := domain.User{}
// }

// func (controller *UserController) Index(c Context) {
// 	users, err := controller.Interactor.Users()
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, users)
// }

// func (controller *UserController) Show(c Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user, err := controller.Interactor.UserById(id)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// }
