// Copyright (c) 2018 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package docker

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func selectCaps(selectOption string) TableEntry {
	return Entry(fmt.Sprintf("cap_%s", selectOption), selectOption)
}

var _ = Describe("capabilities", func() {
	var (
		args      []string
		id        string
		anotherID string
		stdout    string
		exitCode  int
	)

	BeforeEach(func() {
		id = randomDockerName()
		anotherID = randomDockerName()
	})

	AfterEach(func() {
		Expect(ExistDockerContainer(id)).NotTo(BeTrue())
		Expect(ExistDockerContainer(anotherID)).NotTo(BeTrue())
	})

	DescribeTable("drop and add capabilities",
		func(selectOption string) {
			Skip("Issue https://github.com/kata-containers/agent/issues/190")
			args = []string{"--name", id, "--rm", "--cap-drop", selectOption, FedoraImage, "capsh --print"}
			stdout, _, exitCode = dockerRun(args...)
			Expect(exitCode).To(Equal(0))
			Expect(stdout).NotTo(ContainSubstring("cap_" + selectOption))

			args = []string{"--name", anotherID, "--rm", "--cap-add", selectOption, FedoraImage, "capsh --print"}
			stdout, _, exitCode = dockerRun(args...)
			Expect(exitCode).To(Equal(0))
			Expect(stdout).To(ContainSubstring("cap_" + selectOption))
		},
		selectCaps("audit_control"),
		selectCaps("audit_read"),
		selectCaps("audit_write"),
		selectCaps("block_suspend"),
		selectCaps("chown"),
		selectCaps("dac_override"),
		selectCaps("dac_read_search"),
		selectCaps("fowner"),
		selectCaps("fsetid"),
		selectCaps("ipc_lock"),
		selectCaps("ipc_owner"),
		selectCaps("kill"),
		selectCaps("lease"),
		selectCaps("linux_immutable"),
		selectCaps("mac_admin"),
		selectCaps("mac_override"),
		selectCaps("mknod"),
		selectCaps("net_admin"),
		selectCaps("net_bind_service"),
		selectCaps("net_broadcast"),
		selectCaps("net_raw"),
		selectCaps("setgid"),
		selectCaps("setfcap"),
		selectCaps("setuid"),
		selectCaps("setpcap"),
		selectCaps("sys_admin"),
		selectCaps("sys_boot"),
		selectCaps("sys_chroot"),
		selectCaps("sys_nice"),
		selectCaps("sys_pacct"),
		selectCaps("sys_rawio"),
		selectCaps("sys_resource"),
		selectCaps("sys_time"),
		selectCaps("syslog"),
		selectCaps("wake_alarm"),
	)
})
