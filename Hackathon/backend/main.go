package main

import "backend/app/controllers"

//"log"

func main() {
	//models.ReadFoodfiles()
	controllers.StartMainServer()
	/*
		user2, _ := models.GetUser(1)
		heights, _ := user2.GetHeightsByUser()
		for _, v := range heights {
			fmt.Println(v)
		}

		height, _ := user2.GetLatestHeightByUser()
		fmt.Println(height)*/

	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "changed2"
		u.Item1 = 1
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*fmt.Println(models.Db)
	user, _ := models.GetUserByName("test1")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	valid, _ := session.ChaeckSession()
	fmt.Println(valid)
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/
	/*
		u := &models.User{}
		u.Name = "test"
		u.PassWord = "testtest"
		u.Maxscore = 0
		fmt.Println(u)

		u.CreateUser()
		user, _ := models.GetUser(1)
		user.CreateHeight(30.2)
		fmt.Println(user)
	*/
	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Maxscore = 10
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
		u.Name = "test2"
		u.Maxscore=10
		u.UpdateUser()

		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)*/
	/*
		h, _ := models.GetHeight(1)
		fmt.Println(h)
	*/
	/*
		user, _ := models.GetUser(3)
		user.CreateHeight(50.2)
		fmt.Println(user)*/
	/*
		heights, _ := models.GetHeights()
		for _, v := range heights {
			fmt.Println(v)
		}*/
	/*
		user2, _ := models.GetUser(3)
		heights, _ := user2.GetHeightsByUser()
		for _, v := range heights {
			fmt.Println(v)
		}*/
	/*
		h, _ := models.GetHeight(3)
		h.DeleteHeight()
	*/
}
