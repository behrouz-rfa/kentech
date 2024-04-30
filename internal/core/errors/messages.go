package errors

const (
	ErrCanNotCreateUser              = "could not create user"
	ErrCanNotGetUser                 = "could not get user"
	ErrCanNotGetUserAfterCreate      = "could not get user after create"
	ErrCanNotGetUserAfterUpdate      = "could not get user after update"
	ErrCanNotUpdateUser              = "could not  update user"
	ErrRoleIsNotValid                = "role is not valid"
	ErrUserNotFoundForAddClassMember = "user not found for add section member"
	ErrClassRoomNotFound             = "section not found"
	ErrCCouldNotRemoveClassroom      = "could not remove section"
	ErrCouldNotAddMember             = "could not add member"
	ErrUserAlreadyExistInClassroom   = "user already exists in section"
	ErrSchoolNotFound                = "school not found"
	ErrUpdateSchool                  = "failed to update school"
	ErrCouldNotCreateClassroom       = "could not create section"
	ErrCouldNotUpdateClassroom       = "could not update section"
	ErrNotAuthorized                 = "you are not authorized to perform this action"
)
