package services

type User struct {
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

func (c *Services) Helloworld() error {
	data := User{
		Name: "nana",
		Age:  20,
	}
	err := c.ctx.JSON(data)
	if err != nil {
		return err
	}

	return nil
}
