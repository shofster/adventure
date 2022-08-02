package cmd

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

/*
 * I/O and support riutines.
 *
 */

/* Find the state message from msg and print it.  Modes are:
 * touch = for inventory, what you can touch
 * look = the full description for the state the object is in
 * listen = the sound for the state the object is in
 * study = text on the object.
 * change = state change */

func pspeak(obj ObjectType, mode SpeakType, blank bool, state int, optional ...any) {
	switch mode {
	case touch:
		if settings.Logger {
			log.Printf("inventory object: %s\n", objects[obj].inventory)
		}
		speak(blank, objects[obj].inventory, optional)
	case look:
		if objects[obj].descriptions != nil {
			if settings.Logger {
				log.Printf("look object: %s @ %s state %d\n", ObjectText(obj),
					LocationText(game.Loc), state)
			}
			speak(blank, objects[obj].descriptions[state], optional)
			if len(objects[obj].descriptions) > 0 {
				switch obj {
				case FISSURE:
					switch state {
					case UNBRIDGED:
						objects[FISSURE].pic = ""
					case BRIDGED:
						objects[FISSURE].pic = "bridge_crystal"
					}
				case BIRD:
					switch state {
					case BIRD_CAGED:
						objects[CAGE].pic = "cage_bird"
					default:
						objects[CAGE].pic = "cage"
					}
				case BEAR:
					switch state {
					case 1:
						objects[obj].pic = "bear_sitting"
					case 2:
						objects[obj].pic = "bear_wandering"
					default:
						objects[obj].pic = "bear"
					}
				}
				settings.CRT.Show(false, objects[obj].pic)
			}
		}
	case hear:
		if objects[obj].sounds != nil {
			if settings.Logger {
				log.Printf("hear object: %d %s\n", state, objects[obj].sounds[state])
			}
			speak(blank, objects[obj].sounds[state], optional)
			settings.CRT.Show(false, objects[obj].pic)
		}
	case study:
		if objects[obj].texts != nil {
			if settings.Logger {
				log.Printf("study object: %d %s\n", state, objects[obj].texts[state])
			}
			speak(blank, objects[obj].texts[state], optional)
		}
	case change:
		if objects[obj].states != nil {
			speak(blank, objects[obj].states[state], optional)
			switch obj {
			case FISSURE:
				if settings.Logger {
					log.Printf("change object: %d %s\n", state, ObjectText(obj))
				}
				switch state {
				case UNBRIDGED:
					objects[FISSURE].pic = ""
				case BRIDGED:
					objects[FISSURE].pic = "bridge_crystal"
				}
			case DRAGON:
				if settings.Logger {
					log.Printf("change object: %d %s\n", state, ObjectText(obj))
				}
				objects[DRAGON].pic = "dragon"
			case BIRD:
				if settings.Logger {
					log.Printf("change object: %d %s\n", state, ObjectText(obj))
				}
				switch state {
				case BIRD_CAGED:
					objects[CAGE].pic = "cage_bird"
				default:
					objects[CAGE].pic = "cage"
				}
			case LAMP:
				if settings.Logger {
					log.Printf("change object: %d %s\n", state, ObjectText(obj))
				}
				switch state {
				case LAMP_BRIGHT:
					objects[obj].pic = "lamp_on"
				case LAMP_DARK:
					objects[obj].pic = "lamp_off"
				}
			case BEAR:
				if settings.Logger {
					log.Printf("change object: %d %s\n", state, ObjectText(obj))
				}
				switch state {
				case 1:
					objects[obj].pic = "bear_sitting"
				case 2:
					objects[obj].pic = "bear_wandering"
				default:
					objects[obj].pic = "bear"
				}
			case GRATE:
				if settings.Logger {
					log.Printf("change object: %d %s\n", state, ObjectText(obj))
				}
				switch state {
				case GRATE_OPEN:
					objects[obj].pic = "chamber_grate"
				}
			default:
				if settings.Logger {
					log.Printf("unknown change object: %d %s\n", state, ObjectText(obj))
				}
				return
			}
			settings.CRT.Show(false, objects[obj].pic)
		}
	default:
		if settings.Logger {
			log.Printf("  pspeak: object %d invalid  mode %d\n", obj, mode)
		}
	}
}

/* Get user input on stdin, parse and map to command */
func getCommandInput(command *Command) {
	var words = make([]string, 0)
	for {
		input := strings.Trim(getInput(), " ")
		if input == "" {
			speak(false, "I can't hear you. Please speak louder.")
			continue
		}
		ws := strings.Split(input, " ")
		for _, w := range ws {
			if len(w) > 0 {
				words = append(words, w)
			}
		}
		switch len(words) {
		case 1:
			words = append(words, "")
			fallthrough
		case 2:
			break
		default:
			words = nil
			speak(false, messages[TWO_WORDS], nil)
		}
		if len(words) == 2 {
			break
		}
	}
	/* populate command with parsed vocabulary metadata */
	command.word[0].raw = words[0]
	id, typ := getVocabMetadata(command.word[0].raw)
	command.word[0].id = id
	command.word[0].typ = typ
	command.word[1].raw = words[1]
	id, typ = getVocabMetadata(command.word[1].raw)
	command.word[1].id = id
	command.word[1].typ = typ
	command.state = TOKENIZED
	if settings.Logger {
		log.Printf("getCommandInput: %s", command)
	}
	command.state = GIVEN
}
func speak(_ bool, msg string, optional ...any) {
	p := func(s string) {
		if settings.Logger {
			if len(s) > 20 {
				log.Printf("     '%s' ...\n", s[0:20])
			} else {
				log.Printf("     '%s'\n", s)
			}
		}
		if msg == "NULL" {
			return
		}
		settings.TTY.Speak(s)
	}
	if len(optional) > 0 {
		switch reflect.ValueOf(optional[0]).Kind().String() {
		case "int", "string":
			p(fmt.Sprintf(msg, optional...))
			return
		}
	}
	p(msg)
}
func getInput() string {
	return settings.TTY.Ask("")
}
func askYesNo(ask, yes, no string) bool {
	speak(false, ask)
	yn := settings.TTY.AskYesNo("Please answer (Y)es or (N)o.")
	if yn {
		if yes != "" {
			speak(false, yes)
		}
	} else {
		if no != "" {
			speak(false, no)
		}
	}
	return yn
}
func silentYesNo() bool {
	return settings.TTY.AskYesNo("Please answer (Y)es or (N)o.")
}
