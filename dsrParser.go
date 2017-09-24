package main

import (
	"errors"
	"log"
	"os"
	"strings"
	"unicode"
)

type dsgConfig struct {
	serverID         string
	calls            [][]string
	roles            []string
	welcomeChannelID string
	welcomeMessage   string
	goodbyeChannelID string
	goodbyeMessage   string
	banChannelID     string
	banMessage       string
}

/*
 * WHEN MAKING THE .DSR FILES, YOU MUST FOLLOW THIS FORMAT: [PLEASE READ]
 * LINE 1: Always the server ID with a ; on the end
 * LINE 2: Always the welcome Channel ID with a ; on the end
 * LINE 3: Always the welcome Channel Message with a ; on the end,
 *      Add a %s where the username should go! [Ex: Welcome, %s!]
 * LINE 4: Always the goodbye Channel ID with a ; on the end
 * LINE 5: Always the goodbye Channel Message with a ; on the end
 *      Add a %s where the username should go! [Ex: Welcome, %s!]
 * LINE 6: Always the ban     Channel ID with a ; on the end
 * LINE 7: Always the ban     Channel Message with a ; on the end
 *      Add a %s where the username should go! [Ex: Welcome, %s!]
 *
 * LINES PROCEEDING must be done like so:
 * CALL,CALL,CALL=ROLE;
 *
 * Calls cannot have spaces due to the fact you can call
 * multiple roles with a single command, ex: $role a b c
 *
 * But; the role itself CAN have spaces. An example of a line would be:
 * role,arole=A Role;
 */

/*
 * WHEN MAKING A SERVERCONFIG.DSR FILE: [PLEASE READ]
 * It must be set up as the following for each line and each .dsr file corresponding to a server:
 * guildID=fileName.dsr;
 *
 * Example:
 * 123456789101112131=myServer.dsr
 */

func getConfigForGuildID(guildID string) (dsgConfig, error) {
	file, err := os.Open("roles/serverconfig.dsr")
	if err != nil {
		log.Fatal(err)
		return dsgConfig{}, err
	}

	data := make([]byte, 5000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
		return dsgConfig{}, err
	}

	s := string(data[:count])
	s = SpaceMap(s)
	fLines := strings.Split(s, ";")

	var fileName string

	for i := 0; i < len(fLines); i++ {
		t := strings.Split(fLines[i], "=")
		if len(t) >= 2 {
			if t[0] == guildID {
				fileName = t[1]
			}
		}
	}

	if fileName == "" {
		return dsgConfig{}, errors.New("Could not find config file for that guildID")

	}

	return handledsr(fileName), nil
}

func handledsr(filename string) dsgConfig { // Opens a dsr file and returns the role calls,
	//													and then the actual role
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 5000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data[:count])
	fLines := strings.Split(s, ";") // Separates lines
	// Removed the first line, which is the server ID
	sID := fLines[0]
	welcomeChannelID := fLines[1]
	welcomeMessage := fLines[2]
	goodbyeChannelID := fLines[3]
	goodbyeMessage := fLines[4]
	banChannelID := fLines[5]
	banMessage := fLines[6]
	fLines = fLines[7:]
	var fCalls []string
	var fRoles []string
	for i := 0; i < len(fLines); i++ {
		t := strings.Split(fLines[i], "=")
		t[0] = SpaceMap(t[0])
		if len(t) >= 2 {
			fCalls = append(fCalls, t[0]) // Should get the information before the '=' on that line
			fRoles = append(fRoles, t[1]) // Should get the information after the '=' on that line
		}
	}

	var fRolesFinal []string
	var fCallsFinal [][]string
	for i := 0; i < len(fRoles); i++ {
		fRolesFinal = append(fRolesFinal, fRoles[i])
	}
	for i := 0; i < len(fCalls); i++ {
		t := strings.Split(fCalls[i], ",")
		fCallsFinal = append(fCallsFinal, t)
	}

	return dsgConfig{sID, fCallsFinal, fRolesFinal, welcomeChannelID, welcomeMessage, goodbyeChannelID, goodbyeMessage, banChannelID, banMessage}
}

//SpaceMap ... removes all whitespace from a string efficiently
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
