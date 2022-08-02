package cmd

import (
	"log"
	"strconv"
	"strings"
)

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

/* Describe the location to the user */
func describeLocation() {
	msg := locations[game.Loc].descriptions.small
	if game.Abbrev[game.Loc]%game.AbbNum == 0 || msg == "" {
		msg = locations[game.Loc].descriptions.big
	}
	if !conditionBit(game.Loc, COND_FORCED) && isDark(game.Loc) {
		msg = messages[PITCH_DARK]
	}
	if game.Place[BEAR] == CARRIED {
		speak(true, messages[TAME_BEAR])
	}
	speak(true, msg)
	if game.Loc >= LOC_FOOF1 && game.Loc <= LOC_FOOF6 {
		settings.CRT.Show(false, "poof")
	}

	if game.Loc == LOC_Y2 && pct(25) && !game.Closng {
		speak(true, messages[SAYS_PLUGH])
		settings.CRT.Show(false, "word_plugh")
	}
}

/*  Print out descriptions of objects at this location.  If
 *  not closing and property value is negative, Tally off
 *  another treasure.  Rug is special case; once seen, its
 *  game.Prop is RUG_DRAGON (dragon on it) till dragon is killed.
 *  Similarly for chain; game.Prop is initially CHAINING_BEAR (locked to
 *  bear).  These hacks are because game.Prop=0 is needed to
 *  get full score. */
func listObjects() {
	if !isDark(game.Loc) {
		game.Abbrev[game.Loc]++
		for i := game.Atloc[game.Loc]; i != 0; i = game.Link[i] {
			objectType := i
			if objectType > NOBJECTS {
				objectType = objectType - NOBJECTS
			}
			if objectType == STEPS && isToting(NUGGET) {
				continue
			}
			if game.Prop[objectType] < 0 {
				if game.Closed {
					continue
				}
				game.Prop[objectType] = STATE_FOUND
				switch objectType {
				case RUG:
					game.Prop[RUG] = RUG_DRAGON
				case CHAIN:
					game.Prop[CHAIN] = CHAINING_BEAR
				}
				game.Tally--
				/*  Note: There used to be a test here to see whether the
				 *  player had blown it so badly that he could never ever see
				 *  the remaining treasures, and if so the lamp was zapped to
				 *  35 Turns.  But the tests were too simple-minded; things
				 *  like killing the bird before the snake was gone (can never
				 *  see jewelry), and doing it "right" was hopeless.  E.G.,
				 *  could cross troll bridge several times, using up all
				 *  available treasures, breaking vase, using coins to buy
				 *  batteries, etc., and eventually never be able to get
				 *  across again.  If bottle were left on far side, could then
				 *  never get eggs or trident, and the effects propagate.  So
				 *  the whole thing was flushed.  anyone who makes such a
				 *  gross blunder isn't likely to find everything else anyway
				 *  (so goes the rationalisation). */
			}
			kk := game.Prop[objectType]
			if objectType == STEPS {
				switch {
				case game.Fixed[STEPS] == STEPS_UP:
					kk = STEPS_UP
				case game.Fixed[STEPS] == STEPS_DOWN:
					kk = STEPS_DOWN
				}
			}
			pspeak(objectType, look, false, kk)
		}
	}
}

/*  Check if this Loc is eligible for any hints.  If been here int
 *  enough, display.  Ignore "HINTS" < 4 (special stuff, see database
 *  notes). */
// C: checkhints
func checkHints() {
	if conditions[game.Loc] >= game.Conds {
		return
	}
	for hint := 0; hint < NHINTS; hint++ {
		if game.Hinted[hint] {
			continue
		}
		if !conditionBit(game.Loc, uint32(hint+1+COND_HBASE)) {
			game.Hintlc[hint] = -1
		}
		game.Hintlc[hint]++
		/*  Come here if he's been at required Loc(s) for some unused hint. */
		if game.Hintlc[hint] >= hints[hint].turns {
			switch hint { // a "break" will ask if want hint
			case 0: // cave
				if game.Prop[GRATE] == GRATE_CLOSED && !isHere(KEYS) {
					break
				}
				game.Hintlc[hint] = 0
				return
			case 1: // bird
				if game.Place[BIRD] == game.Loc && isToting(ROD) && game.Oldobj == BIRD {
					break
				}
				return
			case 2: // snake
				if isHere(SNAKE) && !isHere(BIRD) {
					break
				}
				game.Hintlc[hint] = 0
				return
			case 3: // maze
				if game.Atloc[game.Loc] == NO_OBJECT &&
					game.Atloc[game.Oldloc] == NO_OBJECT &&
					game.Atloc[game.Oldlc2] == NO_OBJECT &&
					game.Holdng > 1 {
					break
				}
				game.Hintlc[hint] = 0
				return
			case 4: // dark
				if game.Prop[EMERALD] != STATE_NOTFOUND && game.Prop[PYRAMID] == STATE_NOTFOUND {
					break
				}
				game.Hintlc[hint] = 0
				return
			case 5: // witt
				break
			case 6: // urn
				if game.Dflag == 0 {
					break
				}
				game.Hintlc[hint] = 0
				return
			case 7: // woods
				if game.Atloc[game.Loc] == NO_OBJECT &&
					game.Atloc[game.Oldloc] == NO_OBJECT &&
					game.Atloc[game.Oldlc2] == NO_OBJECT {
					break
				}
				return
			case 8: // ogre
				i := atdwrf(game.Loc)
				if i < 0 {
					game.Hintlc[hint] = 0
					return
				}
				if isHere(OGRE) && i == 0 {
					break
				}
				return
			case 9: // jade
				if game.Tally == 1 && game.Prop[JADE] < 0 {
					break
				}
				game.Hintlc[hint] = 0
				return
			default:
				// Should never hap[pen
				panic("checkHints: hint exceeds list size")
			}
			/* Fall through to hint display */
			game.Hintlc[hint] = 0
			if !askYesNo(hints[hint].question, messages[NO_MESSAGE], messages[OK_MAN]) {
				return
			}
			speak(true, messages[HINT_COST], hints[hint].penalty, plural(hints[hint].penalty))
			game.Hinted[hint] = askYesNo(messages[WANT_HINT], hints[hint].hint, messages[OK_MAN])
			if game.Hinted[hint] && game.Limit > WARNTIME {
				game.Limit += TurnType(WARNTIME * hints[hint].penalty)
			}
		}
	}
}

/* Pre-processes a command input to see if we need to tease out a few specific cases:
 * - "enter water" or "enter stream":
 *   wierd specific case that gets the user wet, and then kicks us back to get another command
 * - <object> <verb>:
 *   Irregular form of input, but should be allowed. We switch back to <verb> <object> form for
 *   furtherprocessing.
 * - "grate":
 *   If in location with grate, we move to that grate. If we're in a number of other places,
 *   we move to the entrance.
 * - "water plant", "oil plant", "water door", "oil door":
 *   Change to "pour water" or "pour oil" based on context
 * - "cage bird":
 * - "cage bird":
 *   If bird is present, we change to "carry bird"
 *
 * Returns true if pre-processing is complete, and we're ready to move to the primary command
 * processing, false otherwise. */
// return value was ignored. Sets state to PREPROCESSED if OK.
// C: preprocess_command
func preprocessCommand(command *Command) {
	if command.word[0].typ == MOTION && command.word[0].id == ENTER &&
		(command.word[1].id == STREAM || command.word[1].id == WATER) {
		ifElse(liquidLoc(game.Loc) == WATER,
			func() { speak(true, messages[FEET_WET]) },
			func() { speak(true, messages[WHERE_QUERY]) })
	} else {
		if command.word[0].typ == OBJECT {
			/* From OV to VO form */
			if command.word[1].typ == ACTION {
				w := command.word[0]
				command.word[0] = command.word[1]
				command.word[1] = w
			}
			if command.word[0].id == GRATE {
				command.word[0].typ = MOTION
				settings.CRT.Show(true, "grate")
				switch game.Loc {
				case LOC_START, LOC_VALLEY, LOC_SLIT:
					command.word[0].id = DEPRESSION
				case LOC_COBBLE, LOC_DEBRIS, LOC_AWKWARD, LOC_BIRDCHAMBER, LOC_PITTOP:
					command.word[0].id = ENTRANCE
				}
			}
			if (command.word[0].id == WATER || command.word[0].id == OIL) &&
				(command.word[1].id == PLANT || command.word[1].id == DOOR) {
				if isAt(ObjectType(command.word[1].id)) {
					command.word[1] = command.word[0]
					command.word[0].id = POUR
					command.word[0].typ = ACTION
					command.word[0].raw = "pour"
				}
			}
			if command.word[0].id == CAGE && command.word[1].id == BIRD && isHere(CAGE) && isHere(BIRD) {
				command.word[0].id = CARRY
				command.word[0].typ = ACTION
			}
		}
	}
	/* If no word type is given for the first word, we assume it's a motion. */
	if command.word[0].typ == NO_WORD_TYPE {
		command.word[0].typ = MOTION
	}
	command.state = PREPROCESSED
}

/*  Handle the closing of the cave.  The cave closes "Clock1" Turns
 *  after the last treasure has been located (including the pirate's
 *  chest, which may of course never show up).  Note that the
 *  treasures need not have been taken yet, just located.  Hence,
 *  Clock1 must be large enough to get out of the cave (it only ticks
 *  while inside the cave).  When it hits zero, we branch to 10000 to
 *  start closing the cave, and then sit back and wait for him to try
 *  to get out.  If he doesn't within Clock2 Turns, we close the cave;
 *  if he does try, we assume he panics, and give him a few additional
 *  Turns to get frantic before we close.  When Clock2 hits zero, we
 *  transport him into the final puzzle.  Note that the puzzle depends
 *  upon all sorts of random things.  For instance, there must be no
 *  water or oil, since there are beanstalks which we don't want to be
 *  able to water, since the code can't handle it.  Also, we can have
 *  no keys, since there is a grate (having moved the Fixed object!)
 *  there separating him from all the treasures.  Most of these
 *  problems arise from the use of negative Prop numbers to suppress
 *  the object descriptions until he's actually moved the objects. */
// C: closecheck
func closecheck() bool {
	/* If a turn threshold has been met, apply penalties and tell
	 * the player about it. */
	for i := 0; i < NTHRESHOLDS; i++ {
		if game.Turns == thresholds[i].threshold+1 {
			game.Trnluz += thresholds[i].loss
			speak(true, thresholds[i].message)
		}
	}
	/*  Don't tick game.Clock1 unless well into cave (and not at Y2). */
	if game.Tally == 0 && inDeep(game.Loc) && game.Loc != LOC_Y2 {
		game.Clock1--
	}
	/*  When the first warning comes, we lock the grate, destroy
	 *  the bridge, kill all the dwarves (and the pirate), remove
	 *  the troll and bear (unless dead), and set "Closng" to
	 *  true.  Leave the dragon; too much trouble to move it.
	 *  from now until Clock2 runs out, he cannot unlock the
	 *  grate, move to any location outside the cave, or create
	 *  the bridge.  Nor can he be resurrected if he dies.  Note
	 *  that the snake is already gone, since he got to the
	 *  treasure accessible only via the hall of the mountain
	 *  king. Also, he's been in giant room (to get eggs), so we
	 *  can refer to it.  Also, also, he's gotten the pearl, so we
	 *  know the bivalve is an oyster.  *And*, the dwarves must
	 *  have been activated, since we've found chest. */
	if game.Clock1 == 0 {
		game.Prop[GRATE] = GRATE_CLOSED
		game.Prop[FISSURE] = UNBRIDGED
		for i := 1; i <= NDWARVES; i++ {
			game.Dseen[i] = false
			game.Dloc[i] = LOC_NOWHERE
		}
		destroy(TROLL)
		move(TROLL+NOBJECTS, FREE)
		move(TROLL2, objects[TROLL].plac)
		move(TROLL2+NOBJECTS, objects[TROLL].fixd)
		juggle(CHASM)
		if game.Prop[BEAR] != BEAR_DEAD {
			destroy(BEAR)
		}
		game.Prop[CHAIN] = CHAIN_HEAP
		game.Fixed[CHAIN] = FREE
		game.Prop[AXE] = AXE_HERE
		game.Fixed[AXE] = FREE
		speak(true, messages[CAVE_CLOSING])
		game.Clock1 = -1
		game.Closng = true
		return game.Closed
	} else if game.Clock1 < 0 {
		game.Clock2--
	}
	if game.Clock2 == 0 {
		/*  Once he's panicked, and Clock2 has run out, we come here
		 *  to set up the storage room.  The room has two locs,
		 *  hardwired as LOC_NE and LOC_SW.  At the ne end, we
		 *  Place empty bottles, a nursery of plants, a bed of
		 *  oysters, a pile of lamps, rods with stars, sleeping
		 *  dwarves, and him.  At the sw end we Place grate over
		 *  treasures, snake pit, covey of caged birds, more rods, and
		 *  pillows.  A mirror stretches across one wall.  Many of the
		 *  objects come from known locations and/or states (e.g. the
		 *  snake is known to have been destroyed and needn't be
		 *  carried away from its old "Place"), making the various
		 *  objects be handled differently.  We also drop all other
		 *  objects he might be carrying (lest he have some which
		 *  could cause trouble, such as the keys).  We describe the
		 *  flash of light and trundle back. */
		game.Prop[BOTTLE] = put(ObjectType(BOTTLE), LocationType(LOC_NE), EMPTY_BOTTLE)
		game.Prop[PLANT] = put(ObjectType(PLANT), LocationType(LOC_NE), PLANT_THIRSTY)
		game.Prop[OYSTER] = put(ObjectType(OYSTER), LocationType(LOC_NE), STATE_FOUND)
		game.Prop[LAMP] = put(ObjectType(LAMP), LocationType(LOC_NE), LAMP_DARK)
		game.Prop[ROD] = put(ObjectType(ROD), LocationType(LOC_NE), STATE_FOUND)
		game.Prop[DWARF] = put(ObjectType(DWARF), LocationType(LOC_NE), 0)
		game.Loc = LOC_NE
		game.Oldloc = LOC_NE
		game.Newloc = LOC_NE
		/*  Leave the grate with normal (non-negative) property.
		 *  Reuse sign. */
		put(GRATE, LOC_SW, 0)
		put(SIGN, LOC_SW, 0)
		game.Prop[SIGN] = ENDGAME_SIGN
		game.Prop[SNAKE] = put(SNAKE, LOC_SW, SNAKE_CHASED)
		game.Prop[BIRD] = put(ObjectType(BIRD), LocationType(LOC_SW), BIRD_CAGED)
		game.Prop[CAGE] = put(ObjectType(CAGE), LocationType(LOC_SW), STATE_FOUND)
		game.Prop[ROD2] = put(ObjectType(ROD2), LocationType(LOC_SW), STATE_FOUND)
		game.Prop[PILLOW] = put(ObjectType(PILLOW), LocationType(LOC_SW), STATE_FOUND)

		game.Prop[MIRROR] = put(ObjectType(MIRROR), LocationType(LOC_NE), STATE_FOUND)
		game.Fixed[MIRROR] = LOC_SW

		var o ObjectType
		for o = 1; o <= NOBJECTS; o++ {
			if isToting(o) {
				destroy(o)
			}
		}
		speak(true, messages[CAVE_CLOSED])
		game.Closed = true
		return game.Closed
	}
	lampCheck()
	return false
}

/* Check game Limit and lamp timers */
// C: lampcheck
func lampCheck() {
	if game.Prop[LAMP] == LAMP_BRIGHT {
		game.Limit--
	}
	/*  Another way we can force an end to things is by having the
	 *  lamp give out.  When it gets close, we come here to warn him.
	 *  First following arm checks if the lamp and fresh batteries are
	 *  here, in which case we replace the batteries and continue.
	 *  Second is for other cases of lamp dying.  Even after it goes
	 *  out, he can explore outside for a while if desired. */
	if game.Limit <= WARNTIME {
		if isHere(BATTERY) && game.Prop[BATTERY] == FRESH_BATTERIES && isHere(LAMP) {
			speak(true, messages[REPLACE_BATTERIES])
			game.Prop[BATTERY] = DEAD_BATTERIES
			game.Limit += BATTERYLIFE
			game.Lmwarn = false
		} else if !game.Lmwarn && isHere(LAMP) {
			game.Lmwarn = true
			if game.Prop[BATTERY] == DEAD_BATTERIES {
				speak(true, messages[MISSING_BATTERIES])
			} else if game.Place[BATTERY] == LOC_NOWHERE {
				speak(true, messages[LAMP_DIM])
			} else {
				speak(true, messages[GET_BATTERIES])
			}
		}
	}
	if game.Limit == 0 {
		game.Limit = -1
		game.Prop[LAMP] = LAMP_DARK
		if isHere(LAMP) {
			speak(true, messages[LAMP_OUT])
		}
	}
}

/*  Analyse a verb.  Remember what it was, go back for object if second word
 *  unless verb is "say", which snarfs arbitrary second word.
 */
// C: action
func action(command Command) PhaseCodeType {
	/* Previously, actions that result in a message, but don't do anything
	 * further were called "specials". Now they're handled here as normal
	 * actions. If noaction is true, then we spit out the message and return */
	if actions[command.verb].noaction {
		speak(true, actions[command.verb].message)
		return GO_CLEAROBJ
	}
	if command.part == unknown {
		/*  Analyse an object word.  See if the thing is here, whether
		 *  we've got a verb yet, and so on.  Object must be here
		 *  unless verb is "find" or "invent(ory)" (and no new verb
		 *  yet to be analysed).  Water and oil are also funny, since
		 *  they are never actually dropped at any location, but might
		 *  be here inside the bottle or urn or as a feature of the
		 *  location. */
		if isHere(command.obj) {
			/* FALL THROUGH */
		} else if command.obj == DWARF && atdwrf(game.Loc) > 0 {
			/* FALL THROUGH */
		} else if !game.Closed && ((liquid() == command.obj && isHere(BOTTLE)) || command.obj == liquidLoc(game.Loc)) {
			/* FALL THROUGH */
		} else if command.obj == OIL && isHere(URN) && game.Prop[URN] != URN_EMPTY {
			command.obj = URN
		} else if command.obj == PLANT && isAt(PLANT2) && game.Prop[PLANT2] != PLANT_THIRSTY {
			command.obj = PLANT2
		} else if command.obj == KNIFE && game.Knfloc == game.Loc {
			game.Knfloc = -1
			speak(true, messages[KNIVES_VANISH])
			return GO_CLEAROBJ
		} else if command.obj == ROD && isHere(ROD2) {
			command.obj = ROD2
		} else if (command.verb == FIND || command.verb == INVENTORY) && (command.word[1].id == WORD_EMPTY || command.word[1].id == WORD_NOT_FOUND) {
			/* FALL THROUGH */
		} else {
			speak(true, messages[NO_SEE], command.word[0].raw)
			return GO_CLEAROBJ
		}
		// FALL THROUGH is here.
		if command.verb != 0 {
			command.part = transitive
		}
	}
	switch command.part {
	case intransitive:
		if command.word[1].raw != "" && command.verb != SAY {
			// have object in word 2. process it.
			return GO_WORD2
		}
		if command.verb == SAY {
			/* KEYS is not special, anything not NO_OBJECT or INTRANSITIVE
			 * will do here. We're preventing interpretation as an intransitive
			 * verb when the word is unknown. */
			ifElse(command.word[1].raw != "",
				func() { command.obj = KEYS },
				func() { command.obj = NO_OBJECT })
		}
		if command.obj == NO_OBJECT || command.obj == INTRANSITIVE {
			/*  Analyse an intransitive verb (ie, no object given yet). */
			switch command.verb {
			case CARRY:
				return vcarry(command.verb, INTRANSITIVE)
			case DROP:
				return GO_UNKNOWN
			case SAY:
				return GO_UNKNOWN
			case UNLOCK:
				return lock(command.verb, INTRANSITIVE)
			case NOTHING:
				speak(false, messages[OK_MAN])
				return GO_CLEAROBJ
			case LOCK:
				return lock(command.verb, INTRANSITIVE)
			case LIGHT:
				return light(command.verb, INTRANSITIVE)
			case EXTINGUISH:
				return extinguish(command.verb, INTRANSITIVE)
			case WAVE:
				return GO_UNKNOWN
			case TAME:
				return GO_UNKNOWN
			case GO:
				speak(true, actions[command.verb].message)
				return GO_CLEAROBJ
			case ATTACK:
				command.obj = INTRANSITIVE
				return attack(command)
			case POUR:
				return pour(command.verb, INTRANSITIVE)
			case EAT:
				return eat(command.verb, INTRANSITIVE)
			case DRINK:
				return drink(command.verb, INTRANSITIVE)
			case RUB:
				return GO_UNKNOWN
			case THROW:
				return GO_UNKNOWN
			case QUIT:
				return quit()
			case FIND:
				return GO_UNKNOWN
			case INVENTORY:
				return inven()
			case FEED:
				return GO_UNKNOWN
			case FILL:
				return fill(command.verb, INTRANSITIVE)
			case BLAST:
				blast()
				return GO_CLEAROBJ
			case SCORE:
				score(term_score)
				return GO_CLEAROBJ
			case FEE, FIE, FOE, FOO, FUM:
				return bigWords(command.word[0].id)
			case BRIEF:
				return brief()
			case READ:
				command.obj = INTRANSITIVE // could have been NO_OBJECT
				return read(command)
			case BREAK:
				return GO_UNKNOWN
			case WAKE:
				return GO_UNKNOWN
			case SAVE:
				return suspend()
			case RESUME:
				resume()
				return GO_TOP
			case FLY:
				return fly(command.verb, INTRANSITIVE)
			case LISTEN:
				return listen()
			case PART:
				return reservoir()
			case SEED, WASTE:
				speak(true, messages[NUMERIC_REQUIRED])
				return GO_TOP
			default:
				panic("action: intransitive verb exceeds list")
			}
		}
		/* FALLTHRU */
		fallthrough
	case transitive:
		/*  Analyse a transitive verb. */
		switch command.verb {
		case CARRY:
			return vcarry(command.verb, command.obj)
		case DROP:
			return discard(command.verb, command.obj)
		case SAY:
			return say(command)
		case UNLOCK:
			return lock(command.verb, command.obj)
		case NOTHING:
			speak(false, messages[OK_MAN])
			return GO_CLEAROBJ
		case LOCK:
			return lock(command.verb, command.obj)
		case LIGHT:
			return light(command.verb, command.obj)
		case EXTINGUISH:
			return extinguish(command.verb, command.obj)
		case WAVE:
			return wave(command.verb, command.obj)
		case TAME:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case GO:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case ATTACK:
			return attack(command)
		case POUR:
			return pour(command.verb, command.obj)
		case EAT:
			return eat(command.verb, command.obj)
		case DRINK:
			return drink(command.verb, command.obj)
		case RUB:
			return rub(command.verb, command.obj)
		case THROW:
			return throwit(command)
		case QUIT:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case FIND:
			return find(command.verb, command.obj)
		case INVENTORY:
			return find(command.verb, command.obj)
		case FEED:
			return feed(command.verb, command.obj)
		case FILL:
			return fill(command.verb, command.obj)
		case BLAST:
			blast()
			return GO_CLEAROBJ
		case SCORE:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case FEE, FIE, FOE, FOO, FUM:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case BRIEF:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case READ:
			return read(command)
		case BREAK:
			return vbreak(command.verb, command.obj)
		case WAKE:
			return wake(command.verb, command.obj)
		case SAVE:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case RESUME:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		case FLY:
			return fly(command.verb, command.obj)
		case LISTEN:
			speak(true, actions[command.verb].message)
			return GO_CLEAROBJ
		// This case should never happen - here only as placeholder
		case PART:
			return reservoir()
		case SEED:
			return seed(command.verb, command.word[1].raw)
		case WASTE:
			t, err := strconv.Atoi(command.word[1].raw)
			if err != nil {
				t = 0
			}
			return waste(command.verb, TurnType(t))
		default:
			panic("action: transitive verb exceeds list")
		}
	case unknown:
		/* Unknown verb, couldn't deduce object - might need hint */
		speak(true, messages[WHAT_DO], command.word[0].raw)
		return GO_CHECKHINT
	default:
		panic("action: speechpart not transitive, intransitive, or unonown")
	}
	return GO_CLEAROBJ
}

/* Set the LCG seed */
// C: set_seed
func setSeed(seedval int32) {
	game.LcgX = seedval % LCG_M
	if game.LcgX < 0 {
		game.LcgX = LCG_M + game.LcgX
	}
	// once seed is set, we need to generate the Z`ZZZ word
	for i := 0; i < 5; i++ {
		game.ZZword[i] = byte('A' + randrange(26))
	}
	game.ZZword[1] = '\'' // force second char to apostrophe
	s := string(game.ZZword[:])
	if settings.Logger {
		log.Println("new seeded word is", s)
	}
}
func randrange(r int32) int32 {
	nextlcg := func() int32 {
		// Return the LCG's current value, and then iterate it.
		oldX := game.LcgX
		game.LcgX = (LCG_A*game.LcgX + LCG_C) % LCG_M
		return oldX
	}
	// Return a random integer from [0, range]
	return r * nextlcg() / LCG_M
}

/*  Return the index of first dwarf at the given location, zero if no dwarf is
 *  there (or if dwarves not active yet), -1 if all dwarves are dead.  Ignore
 *  the pirate (6th dwarf). */
// C: atdrf
func atdwrf(where LocationType) int {
	if game.Dflag < 2 {
		return 0
	}
	at := -1
	for i := 1; i <= NDWARVES-1; i++ {
		if game.Dloc[i] == where {
			return i // found n'th dwarf @ this location
		}
		if game.Dloc[i] != 0 {
			at = 0
		}
	}
	return at
}

func spottedByPirate(i int) bool {
	if i != PIRATE {
		return false
	}
	/*  The pirate's spotted him.  Pirate leaves him alone once we've
	 *  found chest.  K counts if a treasure is here.  If not, and
	 *  Tally=1 for an unseen chest, let the pirate be spotted.  Note
	 *  that game.Place[CHEST] = LOC_NOWHERE might mean that he's thrown
	 *  it to the troll, but in that case he's seen the chest
	 *  (game.Prop[CHEST] == STATE_FOUND). */
	if game.Loc == game.Chloc || game.Prop[CHEST] != STATE_NOTFOUND {
		return true
	}
	snarfed := 0
	movechest := false
	robplayer := false
	for treasure := 1; treasure <= NOBJECTS; treasure++ {
		if !objects[treasure].treasure {
			continue
		}
		/*  Pirate won't take pyramid from plover room or dark
		 *  room (too easy!). */
		if treasure == PYRAMID && (game.Loc == objects[PYRAMID].plac || game.Loc == objects[EMERALD].plac) {
			continue
		}
		if isToting(ObjectType(treasure)) || isHere(ObjectType(treasure)) {
			snarfed++
		}
		if isToting(ObjectType(treasure)) {
			movechest = true
			robplayer = true
		}
	}
	/* Force chest placement before player finds last treasure */
	if game.Tally == 1 && snarfed == 0 && game.Place[CHEST] == LOC_NOWHERE &&
		isHere(LAMP) && game.Prop[LAMP] == LAMP_BRIGHT {
		speak(true, messages[PIRATE_SPOTTED])
		settings.CRT.Show(false, "pirate")
		movechest = true
	}
	/* Do things in this order (chest move before robbery) so chest is listed
	 * last at the maze location. */
	if movechest {
		move(CHEST, game.Chloc)
		move(MESSAG, game.Chloc2)
		game.Dloc[PIRATE] = game.Chloc
		game.Odloc[PIRATE] = game.Chloc
		game.Dseen[PIRATE] = false
	} else {
		/* You might get a hint of the pirate's presence even if the
		 * chest doesn't move... */
		if game.Odloc[PIRATE] != game.Dloc[PIRATE] && pct(20) {
			speak(true, messages[PIRATE_RUSTLES])
			settings.CRT.Show(false, "pirate")
		}
	}
	if robplayer {
		speak(true, messages[PIRATE_POUNCES])
		settings.CRT.Show(false, "pirate")
		for treasure := 1; treasure <= NOBJECTS; treasure++ {
			if !objects[treasure].treasure {
				continue
			}
			if !(treasure == PYRAMID && (game.Loc == objects[PYRAMID].plac ||
				game.Loc == objects[EMERALD].plac)) {
				if isAt(ObjectType(treasure)) && game.Fixed[treasure] == FREE {
					carry(ObjectType(treasure), game.Loc)
				}
				if isToting(ObjectType(treasure)) {
					drop(ObjectType(treasure), game.Chloc)
				}
			}
		}
	}
	return false
}
func getVocabMetadata(word string) (id VocabType, typ WordType) {
	// Check for an empty string
	if word == "" {
		return WORD_NOT_FOUND, NO_WORD_TYPE
	}
	if len(word) > 5 {
		word = word[0:5]
	}
	word = strings.ToLower(word)
	ix := func() int {
		for i := 0; i < NMOTIONS; i++ {
			if motions[i].words != nil {
				for j := 0; j < len(motions[i].words); j++ {
					if word == motions[i].words[j] {
						return i
					}
				}
			}
		}
		return WORD_NOT_FOUND
	}()
	if ix != WORD_NOT_FOUND {
		typ = MOTION
		id = VocabType(ix)
		return
	}
	ix = func() int {
		for i := 0; i < NOBJECTS+1; i++ {
			if objects[i].words != nil {
				for j := 0; j < len(objects[i].words); j++ {
					if word == objects[i].words[j] {
						return i
					}
				}
			}
		}
		return WORD_NOT_FOUND
	}()
	if ix != WORD_NOT_FOUND {
		typ = OBJECT
		id = VocabType(ix)
		return
	}
	ix = func() int {
		for i := 0; i < NACTIONS+1; i++ {
			if actions[i].words != nil {
				for j := 0; j < len(actions[i].words); j++ {
					if word == actions[i].words[j] {
						return i
					}
				}
			}
		}
		return WORD_NOT_FOUND
	}()
	if ix != WORD_NOT_FOUND {
		typ = ACTION
		id = VocabType(ix)
		return
	}
	// Check for the reservoir magic word.
	s := string(game.ZZword[0:])
	if strings.EqualFold(word, s) {
		id = PART
		typ = ACTION
		return
	}
	// Check words that are actually numbers.
	_, err := strconv.Atoi(word)
	if err == nil {
		return WORD_EMPTY, NUMERIC
	}
	typ = NO_WORD_TYPE
	id = WORD_NOT_FOUND
	return
}

/* Are two travel entries equal for purposes of skip after failed condition? */
func traveleq(a, b int) bool {
	return (travels[a].condtype == travels[b].condtype) &&
		(travels[a].condarg1 == travels[b].condarg1) &&
		(travels[a].condarg2 == travels[b].condarg2) &&
		(travels[a].desttype == travels[b].desttype) &&
		(travels[a].destval == travels[b].destval)
}

/* Object must have a change-message list for this to be useful; only some do */
// C: state_change
func stateChange(object ObjectType, state int) {
	game.Prop[object] = state
	pspeak(object, change, true, state)
}

// Resets the state of the command to empty
func clearCommand(command *Command) {
	command.verb = ACT_NULL
	command.part = unknown
	game.Oldobj = command.obj
	command.obj = NO_OBJECT
	command.state = EMPTY
	command.word[0] = EmptyCommandWord
	command.word[1] = EmptyCommandWord
}

// C: move
func move(object ObjectType, where LocationType) {
	if settings.Logger {
		log.Printf(" MOVE: %s to %s\n", ObjectText(object), LocationText(where))
	}
	var from LocationType
	if object > NOBJECTS {
		from = game.Fixed[object-NOBJECTS]
	} else {
		from = game.Place[object]
	}
	if from != LOC_NOWHERE && from != CARRIED {
		carry(object, from)
	}
	drop(object, where)
}

/*  put() is the same as move(), except it returns a value used to set up the
 *  negated game.Prop values for the repository objects.
 */
// C: put
func put(object ObjectType, where LocationType, pval int) int {
	move(object, where)
	return stashed(ObjectType(pval))
}

/*  Place an object at a given Loc, prefixing it onto the game.Atloc list.  Decrement
 *  game.Holdng if the object was being toted. */
// C: drop
func drop(object ObjectType, where LocationType) {
	//if settings.Logger != nil {
	//	log.Printf(" DROP: %s in %s\n", ObjectText(object), LocationText(where))
	//}
	if object > NOBJECTS {
		game.Fixed[object-NOBJECTS] = where
	} else {
		/* The bird has to be weightless.  This ugly hack (and the
		 * corresponding code in the drop function) brought to you
		 * by the fact that when the bird is caged, we need to be able
		 * to either 'take bird' or 'take cage' and have the right thing
		 * happen. */
		if game.Place[object] == CARRIED && object != BIRD {
			game.Holdng--
		}
		game.Place[object] = where
	}
	if where == LOC_NOWHERE || where == CARRIED {
		return
	}
	game.Link[object] = game.Atloc[where]
	game.Atloc[where] = object

}

/*  Juggle an object by picking it up and putting it down again, the purpose
 *  being to get the object to the front of the chain of things at its Loc.
 */
// C: juggle
func juggle(object ObjectType) {
	move(object, game.Place[object])
	move(object+NOBJECTS, game.Fixed[object])
}

/*  Start toting an object, removing it from the list of things at its former
 *  location.  Incr Holdng unless it was already being toted.  If object>NOBJECTS
 *  (moving "Fixed" second Loc), don't change game.Place or game.Holdng. */
// C: carry
func carry(object ObjectType, where LocationType) {
	if object <= NOBJECTS {
		if game.Place[object] == CARRIED {
			return
		}
		game.Place[object] = CARRIED
		if object != BIRD {
			game.Holdng++
		}
	}
	if object == BIRD { // shows in cage
		settings.CRT.Show(false, objects[CAGE].pic)
	} else {
		o := object
		if o > NOBJECTS {
			o -= NOBJECTS
		}
		settings.CRT.Show(false, objects[o].pic)
	}
	if game.Atloc[where] == object {
		game.Atloc[where] = game.Link[object]
		return
	}
	loc := game.Atloc[where]
	for {
		if game.Link[loc] == object {
			break
		}
		loc = game.Link[loc]
	}
	game.Link[loc] = game.Link[object]
}
func destroy(object ObjectType) {
	move(object, LOC_NOWHERE)
}

/* Map a state property value to a negative range, where the object cannot be
 * picked up but the value can be recovered later.  Avoid colliding with -1,
 * which has its own meaning. */
func stashed(obj ObjectType) int {
	return -1 - game.Prop[obj]
}
func (t Travel) isTerminate() bool {
	return t.motion == 1
}

// pct. Returns true if rand[1-100] < limit.  a random %
func pct(limit int32) bool {
	return randrange(100) < limit
}
func isToting(obj ObjectType) bool {
	return game.Place[obj] == CARRIED
}

// true if an Object is at the current location (fixed or carried there).
func isAt(obj ObjectType) bool {
	return game.Place[obj] == game.Loc || game.Fixed[obj] == game.Loc
}

// true if an Object is at the current location (fixed or carried there).
// OR currently being carried (toted).
func isHere(obj ObjectType) bool {
	return isAt(obj) || isToting(obj)
}
func liquid() ObjectType {
	switch game.Prop[BOTTLE] {
	case WATER_BOTTLE:
		return WATER
	case OIL_BOTTLE:
		return OIL
	}
	return NO_OBJECT
}
func conditionBit(l LocationType, mask uint32) bool {
	//	fmt.Println("conditions:", ConditionText(mask))
	return (conditions[l] & mask) != 0
}
func liquidLoc(loc LocationType) ObjectType {
	if conditionBit(loc, COND_FLUID) {
		if conditionBit(loc, COND_OILY) {
			return OIL
		}
		return WATER
	}
	return NO_OBJECT
}
func isForced(loc LocationType) bool {
	return conditionBit(loc, COND_FORCED)
}
func isDark(_ LocationType) bool {
	return !conditionBit(game.Loc, COND_LIT) &&
		(game.Prop[LAMP] == LAMP_DARK || !isHere(LAMP))
}
func isGemstone(obj ObjectType) bool { // ?? is JADE a gemstone
	return obj == EMERALD || obj == RUBY || obj == AMBER || obj == SAPPH
}
func isAbove(loc LocationType) bool {
	return conditionBit(loc, COND_ABOVE)
}
func isForest(loc LocationType) bool {
	return conditionBit(loc, COND_FOREST)
}
func isOutside(loc LocationType) bool {
	return isAbove(loc) || isForest(loc)
}
func isNoArrr(loc LocationType) bool {
	return conditionBit(loc, COND_NOARRR)
}
func isNoBack(loc LocationType) bool {
	return conditionBit(loc, BACK)
}
func inDeep(loc LocationType) bool {
	return conditionBit(loc, COND_DEEP)
}
