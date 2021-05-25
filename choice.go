package main

type Choice struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Beats []int  `json:"-"`
}
