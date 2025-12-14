package models

type Line struct {
	Name    string `yaml:"name"`
	Trigger struct {
		Webhook string `yaml:"webhook"`
		File    string `yaml:"file"`
	} `yaml:"trigger"`
	Knots struct {
		Groups []struct {
			Knots []string `yaml:"knots"`
		} `yaml:"groups"`
	} `yaml:"knots"`
}
