// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"cogentcore.org/core/base/exec"
	"cogentcore.org/core/base/logx"
	"cogentcore.org/core/cmd/core/config"
)

// Setup installs platform-specific dependencies for the current platform.
// It only needs to be called once per system.
func Setup(c *config.Config) error { //types:add
	vc := exec.Verbose().SetBuffer(false)
	switch runtime.GOOS {
	case "darwin":
		p, err := exec.Output("xcode-select", "-p")
		if err != nil || p == "" {
			err := vc.Run("xcode-select", "--install")
			if err != nil {
				return err
			}
		} else {
			logx.PrintlnWarn("xcode tools already installed")
		}
		return nil
	case "linux":
		// Based on https://docs.fyne.io/started
		if _, err := exec.LookPath("apt-get"); err == nil {
			err := vc.Run("sudo", "apt-get", "update")
			if err != nil {
				return err
			}
			return vc.Run("sudo", "apt-get", "install", "-f", "-y", "gcc", "libgl1-mesa-dev", "libegl1-mesa-dev", "mesa-vulkan-drivers", "xorg-dev")
		}
		if _, err := exec.LookPath("dnf"); err == nil {
			return vc.Run("sudo", "dnf", "install", "gcc", "libX11-devel", "libXcursor-devel", "libXrandr-devel", "libXinerama-devel", "mesa-libGL-devel", "libXi-devel", "libXxf86vm-devel")
		}
		if _, err := exec.LookPath("pacman"); err == nil {
			return vc.Run("sudo", "pacman", "-S", "xorg-server-devel", "libxcursor", "libxrandr", "libxinerama", "libxi")
		}
		if _, err := exec.LookPath("eopkg"); err == nil {
			return vc.Run("sudo", "eopkg", "it", "-c", "system.devel", "mesalib-devel", "libxrandr-devel", "libxcursor-devel", "libxi-devel", "libxinerama-devel")
		}
		if _, err := exec.LookPath("zypper"); err == nil {
			return vc.Run("sudo", "zypper", "install", "gcc", "libXcursor-devel", "libXrandr-devel", "Mesa-libGL-devel", "libXi-devel", "libXinerama-devel", "libXxf86vm-devel")
		}
		if _, err := exec.LookPath("xbps-install"); err == nil {
			return vc.Run("sudo", "xbps-install", "-S", "base-devel", "xorg-server-devel", "libXrandr-devel", "libXcursor-devel", "libXinerama-devel")
		}
		if _, err := exec.LookPath("apk"); err == nil {
			return vc.Run("sudo", "apk", "add", "gcc", "libxcursor-dev", "libxrandr-dev", "libxinerama-dev", "libxi-dev", "linux-headers", "mesa-dev")
		}
		if _, err := exec.LookPath("nix-shell"); err == nil {
			return vc.Run("nix-shell", "-p", "libGL", "pkg-config", "xorg.libX11.dev", "xorg.libXcursor", "xorg.libXi", "xorg.libXinerama", "xorg.libXrandr", "xorg.libXxf86vm")
		}
		return errors.New("unknown Linux distro; please file a bug report at https://github.com/cogentcore/core/issues")
	case "windows":
		if _, err := exec.LookPath("gcc"); err != nil {
			err := vc.Run("curl", "-OL", "https://github.com/skeeto/w64devkit/releases/download/v2.0.0/w64devkit-x64-2.0.0.exe")
			if err != nil {
				return err
			}
			path, err := filepath.Abs("w64devkit-x64-2.0.0.exe")
			if err != nil {
				return err
			}
			err = vc.Run(path, "x", "-oC:", "-aoa")
			if err != nil {
				return err
			}
			err = windowsRegistryAddPath(`C:\w64devkit\bin`)
			if err != nil {
				return err
			}
		} else {
			logx.PrintlnWarn("gcc already installed")
		}
		if _, err := exec.LookPath("git"); err != nil {
			err := vc.Run("curl", "-OL", "https://github.com/git-for-windows/git/releases/download/v2.45.2.windows.1/Git-2.45.2-64-bit.exe")
			if err != nil {
				return err
			}
			path, err := filepath.Abs("Git-2.45.2-64-bit.exe")
			if err != nil {
				return err
			}
			err = vc.Run(path)
			if err != nil {
				return err
			}
		} else {
			logx.PrintlnWarn("git already installed")
		}
		return nil
	}
	return fmt.Errorf("platform %q not supported for core setup", runtime.GOOS)
}
