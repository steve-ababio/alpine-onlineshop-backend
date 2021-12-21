package misc

func Contains(arr []interface{},searchvalue interface{})bool{
	for _,i := range arr{
		if(i == searchvalue){
			return true;
		}
	}
		return false;
}