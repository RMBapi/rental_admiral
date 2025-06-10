package response

import (
	"example.com/rental/models"
)


type UserEvent struct {
	Name     string  `json:"name"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Number   string  `json:"phone"`
	UType    string  `json:"type"`
}

func SendUserResponse(user models.User, tempPassword string) UserEvent {
	uEvent := UserEvent{
		Name:   user.Name,
		Number: user.Number,
		UType: models.UserTypeMap[user.UType],
	}

	if user.UType == models.Admin || user.UType == models.SystemManager ||
		user.UType == models.Dispatcher || user.UType == models.Agent {

		if user.Email != nil {
			uEvent.Email = user.Email
		}
		if tempPassword != "" {
			uEvent.Password = &tempPassword
		}
	}

	return uEvent
}