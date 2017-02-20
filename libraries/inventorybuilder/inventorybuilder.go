package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	dumpFile := flag.String("d", "/usr/share/ansible/cache_dump.db", "Full path to the cache_dump.db file")
	inventoryFile := flag.String("i", "/usr/share/ansible/hosts", "Full path of the inventory file that will be created")
	ansibleUser := flag.String("u", "ansible", "The Ansible user that will be added in the inventory file")
	ansibleSSHKey := flag.String("k", "/usr/share/ansible/.ssh/id_rsa", "The Ansible user's private key path that will be added in the inventory file")
	ansibleSSHPort := flag.String("p", "22", "The SSH port used by Ansible that will be added in the inventory file")
	flag.Parse()

	aRecords := GetARecords(dumpFile)
	inventory := BuildInventory(aRecords)
	WriteInventory(inventoryFile, ansibleUser, ansibleSSHKey, ansibleSSHPort, inventory)
}

// GetARecords returns all the A DNS records from a dumpFile
func GetARecords(dumpFile *string) [][]byte {
	dump, err := ioutil.ReadFile(*dumpFile)
	if err != nil {
		log.Fatal(err)
	}
	aRecordsReg, _ := regexp.Compile(`(?m)^[a-zA-Z0-9-. \t]*IN A[ \t]*(?:[0-9]{1,3}.){3}[0-9]{1,3}`) // match all A records
	aRecords := aRecordsReg.FindAll(dump, -1)
	return aRecords
}

// BuildInventory returns blabla
func BuildInventory(aRecords [][]byte) map[string]interface{} {
	groupedHosts := make(map[string][]string)
	hostReg, _ := regexp.Compile("[a-zA-Z0-9-.]*")

	groups := make(map[string]string)

	groups["artifactory"] = "artifactory"
	groups["jenkins"] = "jenkins"
	groups["gitlab"] = "gitlab"
	groups["web"] = "web"
	groups["app"] = "app"

	for _, record := range aRecords {
		record := string(record)
		splitRecord := strings.Split(record, "-")
		for group, substring := range groups {
			if (len(splitRecord) == 6) && (splitRecord[4] == substring) {
				groupedHosts[group] = append(groupedHosts[group], hostReg.FindString(record))
			}
		}
	}

	inventory := make(map[string]interface{})

	for group, hosts := range groupedHosts {
		zonedHosts := make(map[string][]string)
		for _, host := range hosts {
			splitHost := strings.Split(host, "-")
			zonedHosts[splitHost[3]] = append(zonedHosts[splitHost[3]], host)
		}
		inventory[group] = zonedHosts
	}
	return inventory
}

// GetGroupedHosts returns blabla
func GetGroupedHosts(aRecords [][]byte, groups map[string]string) map[string][]string {
	groupedHosts := make(map[string][]string)
	hostReg, _ := regexp.Compile("^[a-zA-Z0-9-.]*")

	for _, record := range aRecords {
		record := string(record)
		splitRecord := strings.Split(record, "-")
		for group, substring := range groups {
			if (len(splitRecord) == 6) && (splitRecord[4] == substring) {
				groupedHosts[group] = append(groupedHosts[group], hostReg.FindString(record))
			}
		}
	}

	return groupedHosts
}

// WriteInventory returns blabla
func WriteInventory(inventoryFile *string, ansibleUser *string, ansibleSSHKey *string, ansibleSSHPort *string, inventory map[string]interface{}) {
	file, err := os.Create(*inventoryFile)
	if err != nil {
		log.Fatal(err)
	}

	for group, zones := range inventory {
		for zone, hosts := range zones.(map[string][]string) {
			file.WriteString("[" + group + "-" + zone + "]\n")
			for _, host := range hosts {
				file.WriteString(host + " ansible_user=" + *ansibleUser + " ansible_ssh_private_key_file=" + *ansibleSSHKey + " ansible_port=" + *ansibleSSHPort + "\n")
			}
			file.WriteString("\n")
		}
	}
	file.Close()
	err = os.Chmod(*inventoryFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
