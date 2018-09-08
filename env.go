package main

// Env contains the environment for the shell
type Env struct {
	pwd  string
	path []string
}

func createEnv(configFile string) Env {
	env := Env{pwd: "/home/jeff", path: []string{"/bin", "/sbin", "/usr/bin", "/usr/sbin"}}
	return env
}
