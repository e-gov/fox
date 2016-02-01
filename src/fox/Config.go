package main

// Structure for storing configuration
type Config struct {
	Storage struct{
		Filepath string
			 }
}

// Sanitize the configuration after loading
func (c *Config)Sanitize(){
	// Make sure the db path ends with a forwardslash
	s := c.Storage.Filepath
	if len(s) > 0{
		if string(s[len(s) - 1]) != "/" {
			c.Storage.Filepath = s + "/"
		}
	}
}