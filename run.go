package main

import "os/exec"

func safe(someVar string) {
    exec.Command("echo", someVar).Run()
}
