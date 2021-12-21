package response;

type response struct{
	Responsemsg string `json:"responsemsg"`;
	Status bool `json:"status"`;
}

func Internalservererrorresponse()response{

	return response{"Internal server error, try again",false};
}
func Getauthresponse(msg string,status bool)response{
	return response{Responsemsg:msg, Status:status};
} 

