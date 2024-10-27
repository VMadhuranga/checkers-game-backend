package application

var validationErrorMessages = map[validationError]string{
	{
		"Username",
		"required",
	}: "User name is required",
	{
		"Username",
		"alpha",
	}: "Username must contain letters only",
	{
		"Password",
		"required",
	}: "Password is required",
	{
		"Password",
		"alphanum",
	}: "Password must contain letters and numbers only",
	{
		"Password",
		"min",
	}: "Password must contain at least 5 characters",
	{
		"ConfirmPassword",
		"eqfield",
	}: "Passwords do not match",
}
