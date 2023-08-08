package middleware

// using decorator pattern to redefine the AUTH logic inside the app 

func auth(fn fiber.Handler) fiber.Handler{ 

  return func(c *fiber.Ctx) error{
    // inside here you chek for the authorisation rules that you want  ex
  user = getUserFromLdap();
    if user == nil{
      // here we could also render a Unauth View using the engine !
      // Example: 
      // c.Render("errors/401" , fiber.Map{Title: "Unautherized"}, "layout/base")
      return c.SendStatus(http.StatusUnautorized)
    }

    return fn(c)
  }

}
