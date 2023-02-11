package registration


type Registration struct{
	name string
	surname string 
	login string 
	password string
}

var Registered []Registration

func(r *Registration) Registration(newName string,newSurname string,newLogin string,newPassword string){
	newRegistration := Registration{name:newName,surname:newSurname,login:newLogin,password:newPassword}
	Registered = append(Registered, newRegistration)
}