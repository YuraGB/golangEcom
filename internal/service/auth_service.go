package service

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterInput struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	UserLoginInput
}

func LoginService(userInput UserLoginInput) User {
	// Simulate a user login process
	// In a real application, you would check the credentials against a database
	if userInput.Email == "test@email.com" && userInput.Password == "password" {
		return User{
			ID:       1,
			Name:     "Test User",
			Email:    userInput.Email,
			LastName: "User",
		}
	}
	return User{}
}

func RegisterService(inputData UserRegisterInput) User {
	// Simulate a user registration process
	// In a real application, you would save the user to a database
	return User{
		ID:       1,
		Name:     inputData.Name,
		Email:    inputData.Email,
		LastName: inputData.Lastname,
	}
}

func EditService(userData User) User {
	// Simulate a user edit process
	// In a real application, you would update the user in a database
	return User{
		ID:       userData.ID,
		Name:     userData.Name,
		Email:    userData.Email,
		LastName: userData.LastName,
	}
}

func DeleteService(id int) bool {
	// Simulate a user deletion process
	// In a real application, you would delete the user from a database
	return true
}
