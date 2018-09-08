package main

type prompt struct {
  ps1 string
}

func (p prompt) format() string {
  return p.ps1
}
