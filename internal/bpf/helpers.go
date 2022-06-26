/*
 * Copyright © 2022 Authors of Patu
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package bpf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	progPath	= "./bpf"
)

func compileEbpfProg(debug bool) error {
	cmd := exec.Command("make", "compile")
	cmd.Dir = progPath
	cmd.Env = os.Environ()
	
	if debug {
		cmd.Env = append(cmd.Env, "DEBUG=1")
	}
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Run()
	if code := cmd.ProcessState.ExitCode(); code != 0 || err != nil {
		return fmt.Errorf("\"%s \" failed with code: %d, err: %v", strings.Join(cmd.Args, " "), code, err)
	}
	fmt.Printf("eBPF programs compiled successfully.")
	return nil
}

func loadBpfProg(debug bool) error {
	if os.Getuid() != 0 {
		return fmt.Errorf("eBPF program loading requires root privileges.")
	}
	cmd := exec.Command("make", "load")
	cmd.Dir = progPath
	cmd.Env = os.Environ()
	
	if debug {
		cmd.Env = append(cmd.Env, "DEBUG=1")
	}
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Run()
	if code := cmd.ProcessState.ExitCode(); code != 0 || err != nil {
		return fmt.Errorf("\"%s \" failed with code: %d, err: %v", strings.Join(cmd.Args, " "), code, err)
	}
	fmt.Printf("eBPF programs loaded successfully.")
	return nil
}

func attachBpfProg() error {
	if os.Getuid() != 0 {
		return fmt.Errorf("eBPF program loading requires root privileges.")
	}
	cmd := exec.Command("make", "attach")
	cmd.Dir = progPath
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if code := cmd.ProcessState.ExitCode(); code != 0 || err != nil {
		return fmt.Errorf("\"%s \" failed with code: %d, err: %v", strings.Join(cmd.Args, " "), code, err)
	}
	fmt.Printf("eBPF programs attached successfully.")
	return nil
}

func detachBpfProg() error {
	cmd := exec.Command("make", "detach")
	cmd.Dir = progPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if code := cmd.ProcessState.ExitCode(); code != 0 || err != nil {
		return fmt.Errorf("\"%s \" failed with code: %d, err: %v", strings.Join(cmd.Args, " "), code, err)
	}
	fmt.Printf("eBPF programs detached successfully.")
	return nil
}

func unloadBpfProg() error {
	cmd := exec.Command("make", "unload")
	cmd.Dir = progPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if code := cmd.ProcessState.ExitCode(); code != 0 || err != nil {
		return fmt.Errorf("\"%s \" failed with code: %d, err: %v", strings.Join(cmd.Args, " "), code, err)
	}
	fmt.Printf("eBPF programs attached successfully.")
	return nil
}
