package auth

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" bson:"email"`
	Password string `json:"password" validate:"required" bson:"password"`
}

type UserRegister struct {
	Name     string `json:"name" validate:"required" bson:"name"`
	Gender   string `json:"gender" validate:"required" bson:"gender"`
	Email    string `json:"email" validate:"required,email" bson:"email"`
	Password string `json:"password" validate:"required" bson:"password"`
}

type User struct {
	Name      string `json:"name" validate:"required" bson:"name"`
	Gender    string `json:"gender" validate:"required" bson:"gender"`
	Email     string `json:"email" validate:"required,email" bson:"email"`
	Password  string `json:"password" validate:"required" bson:"password"`
	Role      Role   `json:"role" validate:"required" bson:"role"`
	IsActive  bool   `json:"isActive" validate:"required" bson:"isActive"`
	IsDeleted bool   `json:"isDeleted" validate:"required" bson:"isDeleted"`
	CreatedBy string `json:"createdBy" validate:"required" bson:"createdBy"`
	UpdatedBy string `json:"updatedBy" validate:"required" bson:"updatedBy"`
	Version   string `json:"version" validate:"required" bson:"version"`
}

type UserUpdate struct {
	Name      *string `json:"name,omitempty" bson:"name,omitempty"`
	Gender    *string `json:"gender,omitempty" bson:"gender,omitempty"`
	Email     *string `json:"email,omitempty" bson:"email,omitempty"`
	Password  *string `json:"password,omitempty" bson:"password,omitempty"`
	Role      *Role   `json:"role,omitempty" bson:"role,omitempty"`
	IsActive  *bool   `json:"isActive,omitempty" bson:"isActive,omitempty"`
	IsDeleted *bool   `json:"isDeleted,omitempty" bson:"isDeleted,omitempty"`
	UpdatedBy string  `json:"updatedBy" validate:"required" bson:"updatedBy"`
	Version   string  `json:"version" validate:"required" bson:"version"`
}

type LoginResponse struct {
	Name         string `json:"name" validate:"required" bson:"name"`
	Email        string `json:"email" validate:"required" bson:"email"`
	Role         Role   `json:"role" validate:"required" bson:"role"`
	AccessToken  string `json:"accessToken" validate:"required" bson:"accessToken"`
	RefreshToken string `json:"refreshToken" validate:"required" bson:"refreshToken"`
}
