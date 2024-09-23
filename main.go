/*
 * main.go
 *
 * Copyright 2024 Dariusz Sikora <ds@isangeles.dev>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either Version 2 of the License, or
 * (at your option) any later Version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 *
 *
 */

// Tool for converting flame data from XML to JSON format.
package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"github.com/isangeles/flame/data/res"
)

func main() {
	typeFlag := flag.String("type", "", "Type of data to convert")
	flag.Parse()
	input, err := readInput()
	if err != nil {
		panic(fmt.Errorf("Unable to read input: %v", err))
	}
	output, err := convData(input, dataType(*typeFlag))
	if err != nil {
		panic(fmt.Errorf("Unable to convert characters: %v", err))
	}
	fmt.Printf("%v\n", output)
}

// Reads the input from the stdin.
func readInput() ([]byte, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, fmt.Errorf("Can't read input")
	}
	var input []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Unable to read input: %v", err)
	}
	return input, nil
}

// Returns flame data type corresponding to specified type flag.
func dataType(flag string) interface{} {
	switch flag {
	case "area":
		return new(res.AreaData)
	case "characters", "chars":
		return new(res.CharactersData)
	case "dialogs":
		return new(res.DialogsData)
	case "effects":
		return new(res.EffectsData)
	case "armors":
		return new(res.ArmorsData)
	case "weapons":
		return new(res.WeaponsData)
	case "misc", "misc-items":
		return new(res.MiscItemsData)
	case "quests":
		return new(res.QuestsData)
	case "races":
		return new(res.RacesData)
	case "recipes":
		return new(res.RecipesData)
	case "skills":
		return new(res.SkillsData)
	case "trainings":
		return new(res.TrainingsData)
	default:
		return nil
	}
}

// Converts the data from XML to JSON format according to specified data type.
func convData(input []byte, data interface{}) (string, error) {
	err := xml.Unmarshal(input, data)
	if err != nil {
		return "", fmt.Errorf("Unable to unmarshal data: %v", err)
	}
	json, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Unable to marshal data: %v", err)
	}
	return string(json[:]), nil
}
