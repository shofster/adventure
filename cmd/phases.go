package cmd

import (
	"strconv"
)

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

/*  Attack.  Assume target if unambiguous.  "Throw" also links here.
 *  Attackable objects fall into two categories: enemies (snake,
 *  dwarf, etc.)  and others (bird, clam, machine).  Ambiguous if 2
 *  enemies, or no enemies but 2 others. */
func attack(command Command) PhaseCodeType {
	verb := command.verb
	object := command.obj
	if object == INTRANSITIVE {
		changes := 0
		if atdwrf(game.Loc) > 0 {
			object = DWARF
			changes++
		}
		if isHere(SNAKE) {
			object = SNAKE
			changes++
		}
		if isAt(DRAGON) && game.Prop[DRAGON] == DRAGON_BARS {
			object = DRAGON
			changes++
		}
		if isAt(TROLL) {
			object = TROLL
			changes++
		}
		if isAt(OGRE) {
			object = OGRE
			changes++
		}
		if isHere(BEAR) && game.Prop[BEAR] == UNTAMED_BEAR {
			object = BEAR
			changes++
		}
		/* check for low-priority targets */
		if object == INTRANSITIVE {
			/* Can't attack bird or machine by throwing axe. */
			if isHere(BIRD) && verb != THROW {
				object = BIRD
				changes++
			}
			if isHere(VEND) && verb != THROW {
				object = VEND
				changes++
			}
			/* Clam and oyster both treated as clam for intransitive case;
			 * no harm done. */
			if isHere(CLAM) || isHere(OYSTER) {
				object = CLAM
				changes++
			}
		}
		if changes >= 2 {
			return GO_UNKNOWN
		}
	}
	if object == BIRD {
		if game.Closed {
			speak(true, messages[UNHAPPY_BIRD])
		} else {
			destroy(BIRD)
			speak(true, messages[BIRD_DEAD])
		}
		return GO_CLEAROBJ
	}
	if object == VEND {
		ifElse(game.Prop[VEND] == VEND_BLOCKS,
			func() { stateChange(VEND, VEND_UNBLOCKS) },
			func() { stateChange(VEND, VEND_BLOCKS) })
		return GO_CLEAROBJ
	}
	if object == BEAR {
		switch game.Prop[BEAR] {
		case UNTAMED_BEAR:
			speak(true, messages[BEAR_HANDS])
		case SITTING_BEAR, CONTENTED_BEAR:
			speak(true, messages[BEAR_CONFUSED])
		case BEAR_DEAD:
			speak(true, messages[ALREADY_DEAD])
		}
		return GO_CLEAROBJ
	}
	if object == DRAGON && game.Prop[DRAGON] == DRAGON_BARS {
		/*  Fun stuff for dragon.  If he insists on attacking it, win!
		 *  Set game.Prop to dead, move dragon to central Loc (still
		 *  Fixed), move rug there (not Fixed), and move him there,
		 *  too.  Then do a null motion to get new description. */
		speak(true, messages[BARE_HANDS_QUERY])
		if !silentYesNo() {
			speak(true, messages[NASTY_DRAGON])
			settings.CRT.Show(false, "dragon")
			return GO_MOVE
		}
		stateChange(DRAGON, DRAGON_DEAD)
		game.Prop[RUG] = RUG_FLOOR
		/* Hardcoding LOC_SECRET5 as the dragon's death location is ugly.
		 * The way it was computed before was worse; it depended on the
		 * two dragon locations being LOC_SECRET4 and LOC_SECRET6 and
		 * LOC_SECRET5 being right between them. */
		move(DRAGON+NOBJECTS, FIXED)
		move(RUG+NOBJECTS, FREE)
		move(DRAGON, LOC_SECRET5)
		move(RUG, LOC_SECRET5)
		drop(BLOOD, LOC_SECRET5)
		var i ObjectType
		for i = 1; i <= NOBJECTS; i++ {
			if game.Place[i] == objects[DRAGON].plac || game.Place[i] == objects[DRAGON].fixd {
				move(i, LOC_SECRET5)
			}
		}
		game.Loc = LOC_SECRET5
		return GO_MOVE
	}
	if object == OGRE {
		speak(true, messages[OGRE_DODGE])
		if atdwrf(game.Loc) == 0 {
			return GO_CLEAROBJ
		}
		settings.CRT.Show(false, "knife")
		speak(true, messages[KNIFE_THROWN])
		settings.CRT.Show(false, "knife")
		destroy(OGRE)
		dwarves := 0
		for i := 1; i < PIRATE; i++ {
			if game.Dloc[i] == game.Loc {
				dwarves++
				game.Dloc[i] = LOC_LONGWEST
				game.Dseen[i] = false
			}
		}
		ifElse(dwarves > 1,
			func() { speak(true, messages[OGRE_PANIC1]) },
			func() { speak(true, messages[OGRE_PANIC2]) })
		return GO_CLEAROBJ
	}
	switch object {
	case INTRANSITIVE:
		speak(true, messages[NO_TARGET])
	case CLAM, OYSTER:
		speak(true, messages[SHELL_IMPERVIOUS])
	case SNAKE:
		speak(true, messages[SNAKE_WARNING])
	case DWARF:
		if game.Closed {
			return GO_DWARFWAKE
		}
		speak(true, messages[BARE_HANDS_QUERY])
	case DRAGON:
		speak(true, messages[ALREADY_DEAD])
		break
	case TROLL:
		speak(true, messages[ROCKY_TROLL])
	default:
		speak(true, actions[verb].message)
	}
	return GO_CLEAROBJ
}

/* Only called on FEE FIE FOE FOO (AND FUM).  Advance to next state if given
 * in proper order. Look up foo in special section of vocab to determine which
 * word we've got. Last word zips the eggs back to the giant room (unless
 * already there). */
func bigWords(id VocabType) PhaseCodeType {
	var fum = func() {
		if game.Loc == LOC_GIANTROOM {
			speak(true, messages[START_OVER])
		} else {
			/* This is new behavior in Open Adventure - sounds better when
			 * player isn't in the Giant Room. */
			speak(true, messages[NOTHING_HAPPENS])
		}
		game.Foobar = WORD_EMPTY
	}
	if (game.Foobar == WORD_EMPTY && id == FEE) ||
		(game.Foobar == FEE && id == FIE) ||
		(game.Foobar == FIE && id == FOE) ||
		(game.Foobar == FOE && id == FOO) ||
		(game.Foobar == FOE && id == FUM) {
		game.Foobar = int(id)
		if (id != FOO) && (id != FUM) {
			speak(true, messages[OK_MAN])
			return GO_CLEAROBJ
		}
		game.Foobar = WORD_EMPTY
		if game.Place[EGGS] == objects[EGGS].plac || (isToting(EGGS) && game.Loc == objects[EGGS].plac) {
			speak(false, messages[NOTHING_HAPPENS])
			return GO_CLEAROBJ
		} else if id == FUM {
			fum()
			return GO_CLEAROBJ
		} else {
			/*  Bring back troll if we steal the eggs back from him before
			 *  crossing. */
			if game.Place[EGGS] == LOC_NOWHERE && game.Place[TROLL] == LOC_NOWHERE && game.Prop[TROLL] == TROLL_UNPAID {
				game.Prop[TROLL] = TROLL_PAIDONCE
			}
			if isHere(EGGS) {
				pspeak(EGGS, look, true, EGGS_VANISHED)
			} else if game.Loc == objects[EGGS].plac {
				pspeak(EGGS, look, true, EGGS_HERE)
			} else {
				pspeak(EGGS, look, true, EGGS_DONE)
			}
			move(EGGS, objects[EGGS].plac)
			return GO_CLEAROBJ
		}
	}
	fum()
	return GO_CLEAROBJ
}

/*  Blast.  No effect unless you've got dynamite, which is a neat trick! */
func blast() {
	if game.Prop[ROD2] == STATE_NOTFOUND || !game.Closed {
		speak(true, messages[REQUIRES_DYNAMITE])
	} else {
		if isHere(ROD2) {
			game.Bonus = bonus_splatter
			speak(true, messages[SPLATTER_MESSAGE])
		} else if game.Loc == LOC_NE {
			game.Bonus = bonus_defeat
			speak(true, messages[DEFEAT_MESSAGE])
		} else {
			game.Bonus = bonus_victory
			speak(true, messages[VICTORY_MESSAGE])
		}
		terminate(ENDGAME_SIGN)
	}
}

/*  Break.  Only works for mirror in repository and, of course, the vase. */
func vbreak(verb VerbType, object ObjectType) PhaseCodeType {
	switch object {
	case MIRROR:
		if game.Closed {
			stateChange(MIRROR, MIRROR_BROKEN)
			return GO_DWARFWAKE
		}
		speak(true, messages[TOO_FAR])
	case VASE:
		if game.Prop[VASE] == VASE_WHOLE {
			if isToting(VASE) {
				drop(VASE, game.Loc)
			}
			stateChange(VASE, VASE_BROKEN)
			game.Fixed[VASE] = FIXED
			break
		}
		fallthrough
	default:
		speak(true, actions[verb].message)
	}
	return GO_CLEAROBJ
}

/*  Brief.  Intransitive only.  Suppress full descriptions after first time. */
func brief() PhaseCodeType {
	game.AbbNum = 10000
	game.Detail = 3
	speak(true, messages[BRIEF_CONFIRM])
	return GO_CLEAROBJ
}

/*  Carry an object.  Special cases for bird and cage (if bird in cage, can't
 *  take one without the other).  Liquids also special, since they depend on
 *  status of bottle.  Also, various side effects, etc. */
// C: vcarry
func vcarry(verb VerbType, object ObjectType) PhaseCodeType {
	if object == INTRANSITIVE {
		/*  Carry, no object given yet.  OK if only one object present. */
		if game.Atloc[game.Loc] == NO_OBJECT ||
			game.Link[game.Atloc[game.Loc]] != 0 ||
			atdwrf(game.Loc) > 0 {
			return GO_UNKNOWN
		}
		object = game.Atloc[game.Loc]
	}
	if isToting(object) {
		speak(true, actions[verb].message)
		return GO_CLEAROBJ
	}
	if object == MESSAG {
		speak(true, messages[REMOVE_MESSAGE])
		destroy(MESSAG)
		return GO_CLEAROBJ
	}
	if game.Fixed[object] != FREE {
		switch object {
		case PLANT:
			/* Next guard tests whether plant is tiny or stashed */
			ifElse(game.Prop[PLANT] <= PLANT_THIRSTY,
				func() { speak(true, messages[DEEP_ROOTS]) },
				func() { speak(true, messages[YOU_JOKING]) })
		case BEAR:
			ifElse(game.Prop[BEAR] == SITTING_BEAR,
				func() { speak(true, messages[BEAR_CHAINED]) },
				func() { speak(true, messages[YOU_JOKING]) })
		case CHAIN:
			ifElse(game.Prop[BEAR] != UNTAMED_BEAR,
				func() { speak(true, messages[STILL_LOCKED]) },
				func() { speak(true, messages[YOU_JOKING]) })
		case RUG:
			ifElse(game.Prop[RUG] == RUG_HOVER,
				func() { speak(true, messages[RUG_HOVERS]) },
				func() { speak(true, messages[YOU_JOKING]) })
		case URN:
			speak(true, messages[URN_NOBUDGE])
		case CAVITY:
			speak(true, messages[DOUGHNUT_HOLES])
		case BLOOD:
			speak(true, messages[FEW_DROPS])
		case SIGN:
			speak(true, messages[HAND_PASSTHROUGH])
		default:
			speak(true, messages[YOU_JOKING])
		}
		return GO_CLEAROBJ
	}
	if object == WATER || object == OIL {
		if !isHere(BOTTLE) || liquid() != object {
			if !isToting(BOTTLE) {
				speak(true, messages[NO_CONTAINER])
				return GO_CLEAROBJ
			}
			if game.Prop[BOTTLE] == EMPTY_BOTTLE {
				return fill(verb, BOTTLE)
			} else {
				speak(true, messages[BOTTLE_FULL])
			}
			return GO_CLEAROBJ
		}
		object = BOTTLE
	}
	if game.Holdng >= INVLIMIT { // 7 items maximum
		speak(true, messages[CARRY_LIMIT])
		return GO_CLEAROBJ
	}
	if object == BIRD && game.Prop[BIRD] != BIRD_CAGED && stashed(BIRD) != BIRD_CAGED {
		if game.Prop[BIRD] == BIRD_FOREST_UNCAGED {
			destroy(BIRD)
			speak(true, messages[BIRD_CRAP])
			return GO_CLEAROBJ
		}
		if !isToting(CAGE) {
			speak(true, messages[CANNOT_CARRY])
			return GO_CLEAROBJ
		}
		if isToting(ROD) {
			speak(true, messages[BIRD_EVADES])
			return GO_CLEAROBJ
		}
		game.Prop[BIRD] = BIRD_CAGED
	}
	if (object == BIRD || object == CAGE) &&
		(game.Prop[BIRD] == BIRD_CAGED || stashed(BIRD) == BIRD_CAGED) {
		/* expression maps BIRD to CAGE and CAGE to BIRD */
		carry(BIRD+CAGE-object, game.Loc)
		settings.CRT.Show(false, "cage_bird")
		objects[CAGE].pic = "cage_bird"
	}
	carry(object, game.Loc)
	if object == BOTTLE && liquid() != NO_OBJECT {
		game.Place[liquid()] = CARRIED
	}
	if isGemstone(object) && game.Prop[object] != STATE_FOUND {
		game.Prop[object] = STATE_FOUND
		game.Prop[CAVITY] = CAVITY_EMPTY
	}
	speak(false, messages[OK_MAN])
	return GO_CLEAROBJ
}

/* Do something to the bear's chain */
// C: chain
func chain(verb VerbType) PhaseCodeType {
	if verb != LOCK {
		if game.Prop[BEAR] == UNTAMED_BEAR {
			speak(true, messages[BEAR_BLOCKS])
			return GO_CLEAROBJ
		}
		if game.Prop[CHAIN] == CHAIN_HEAP {
			speak(true, messages[ALREADY_UNLOCKED])
			return GO_CLEAROBJ
		}
		game.Prop[CHAIN] = CHAIN_HEAP
		game.Fixed[CHAIN] = FREE
		if game.Prop[BEAR] != BEAR_DEAD {
			game.Prop[BEAR] = CONTENTED_BEAR
		}
		switch game.Prop[BEAR] {
		case BEAR_DEAD:
			/* Can't be reached until the bear can die in some way other
			 * than a bridge collapse. Leave in incase this states, but
			 * exclude from coverage testing. */
			game.Fixed[BEAR] = FIXED
		default:
			game.Fixed[BEAR] = FREE
		}
		speak(true, messages[CHAIN_UNLOCKED])
		return GO_CLEAROBJ
	}
	if game.Prop[CHAIN] != CHAIN_HEAP {
		speak(true, messages[ALREADY_LOCKED])
		return GO_CLEAROBJ
	}
	if game.Loc != objects[CHAIN].plac {
		speak(true, messages[NO_LOCKSITE])
		return GO_CLEAROBJ
	}
	game.Prop[CHAIN] = CHAIN_FIXED
	if isToting(CHAIN) {
		drop(CHAIN, game.Loc)
	}
	game.Fixed[CHAIN] = FIXED
	speak(true, messages[CHAIN_LOCKED])
	return GO_CLEAROBJ
}

/*  Discard object.  "Throw" also comes here for most objects.  Special cases for
 *  bird (might attack snake or dragon) and cage (might contain bird) and vase.
 *  Drop coins at vending machine for extra batteries. */
// C: discard
func discard(verb VerbType, object ObjectType) PhaseCodeType {
	if object == ROD && !isToting(ROD) && isToting(ROD2) {
		object = ROD2
	}
	if !isToting(object) {
		speak(true, actions[verb].message)
		return GO_CLEAROBJ
	}
	if isGemstone(object) && isAt(CAVITY) && game.Prop[CAVITY] != CAVITY_FULL {
		speak(true, messages[GEM_FITS])
		game.Prop[object] = STATE_IN_CAVITY
		game.Prop[CAVITY] = CAVITY_FULL
		if isHere(RUG) && ((object == EMERALD && game.Prop[RUG] != RUG_HOVER) ||
			(object == RUBY && game.Prop[RUG] == RUG_HOVER)) {
			switch {
			case object == RUBY:
				speak(true, messages[RUG_SETTLES])
			case isToting(RUG):
				speak(true, messages[RUG_WIGGLES])
			default:
				speak(true, messages[RUG_RISES])
			}
			if !isToting(RUG) || object == RUBY {
				var loc LocationType = RUG_FLOOR
				if game.Prop[RUG] == RUG_HOVER {
					loc = RUG_HOVER
				}
				game.Prop[RUG] = int(loc)
				if loc == RUG_HOVER {
					loc = objects[SAPPH].plac
				}
				move(RUG+NOBJECTS, loc)
			}
		}
		drop(object, game.Loc)
		return GO_CLEAROBJ
	}
	if object == COINS && isHere(VEND) {
		destroy(COINS)
		drop(BATTERY, game.Loc)
		pspeak(BATTERY, look, true, FRESH_BATTERIES)
		return GO_CLEAROBJ
	}
	if liquid() == object {
		object = BOTTLE
	}
	if object == BOTTLE && liquid() != NO_OBJECT {
		game.Place[liquid()] = LOC_NOWHERE
	}
	if object == BEAR && isAt(TROLL) {
		stateChange(TROLL, TROLL_GONE)
		move(TROLL, LOC_NOWHERE)
		move(TROLL+NOBJECTS, FREE)
		move(TROLL2, objects[TROLL].plac)
		move(TROLL2+NOBJECTS, objects[TROLL].fixd)
		juggle(CHASM)
		drop(object, game.Loc)
		return GO_CLEAROBJ
	}
	if object == VASE {
		if game.Loc != objects[PILLOW].plac {
			ifElse(isAt(PILLOW),
				func() { stateChange(VASE, VASE_WHOLE) },
				func() { stateChange(VASE, VASE_DROPPED) })
			if game.Prop[VASE] != VASE_WHOLE {
				game.Fixed[VASE] = FIXED
			}
			drop(object, game.Loc)
			return GO_CLEAROBJ
		}
	}
	if object == CAGE && game.Prop[BIRD] == BIRD_CAGED {
		drop(BIRD, game.Loc)
	}
	if object == BIRD {
		settings.CRT.Show(false, "bird")
		if isAt(DRAGON) && game.Prop[DRAGON] == DRAGON_BARS {
			speak(true, messages[BIRD_BURNT])
			destroy(BIRD)
			return GO_CLEAROBJ
		}
		if isHere(SNAKE) {
			speak(true, messages[BIRD_ATTACKS])
			if game.Closed {
				return GO_DWARFWAKE
			}
			destroy(SNAKE)
			/* Set game.Prop for use by travel options */
			game.Prop[SNAKE] = SNAKE_CHASED
		} else {
			speak(true, messages[OK_MAN])
		}
		ifElse(isForest(game.Loc),
			func() { game.Prop[BIRD] = BIRD_FOREST_UNCAGED },
			func() { game.Prop[BIRD] = BIRD_UNCAGED })
		drop(object, game.Loc)
		return GO_CLEAROBJ
	}
	speak(false, messages[OK_MAN])
	drop(object, game.Loc)
	return GO_CLEAROBJ
}

/*  Drink.  If no object, assume water and look for it here.  If water is in
 *  the bottle, drink that, else must be at a water Loc, so drink stream. */
// C: drink
func drink(verb VerbType, object ObjectType) PhaseCodeType {
	if object == INTRANSITIVE && liquidLoc(game.Loc) != WATER &&
		(liquid() != WATER || !isHere(BOTTLE)) {
		return GO_UNKNOWN
	}
	if object == BLOOD {
		destroy(BLOOD)
		stateChange(DRAGON, DRAGON_BLOODLESS)
		game.Blooded = true
		return GO_CLEAROBJ
	}
	if object != INTRANSITIVE && object != WATER {
		speak(true, messages[RIDICULOUS_ATTEMPT])
		return GO_CLEAROBJ
	}
	if liquid() == WATER && isHere(BOTTLE) {
		game.Place[WATER] = LOC_NOWHERE
		stateChange(BOTTLE, EMPTY_BOTTLE)
		return GO_CLEAROBJ
	}
	speak(true, actions[verb].message)
	return GO_CLEAROBJ
}

/*  Eat.  Intransitive: assume food if present, else ask what.  Transitive: food
 *  ok, some things lose appetite, rest are ridiculous. */
// C: eat
func eat(verb VerbType, object ObjectType) PhaseCodeType {
	if object == INTRANSITIVE && !isHere(FOOD) { // INTRANSITIVE has special value == -1
		return GO_UNKNOWN
	}
	switch object {
	case FOOD:
		destroy(FOOD)
		speak(true, messages[THANKS_DELICIOUS])
	case BIRD, SNAKE, CLAM, OYSTER, DWARF, DRAGON, TROLL, BEAR, OGRE:
		speak(true, messages[LOST_APPETITE])
	default:
		speak(true, actions[verb].message)
	}
	return GO_CLEAROBJ
}

/* Extinguish.  Lamp, urn, dragon/volcano (nice try). */
// C: extinguish
func extinguish(verb VerbType, object ObjectType) PhaseCodeType {
	if object == INTRANSITIVE { // deduce a known extinguishable object
		if isHere(LAMP) && game.Prop[LAMP] == LAMP_BRIGHT {
			object = LAMP
		}
		if isHere(URN) && game.Prop[URN] == URN_LIT {
			object = URN
		}
		if object == INTRANSITIVE {
			return GO_UNKNOWN
		}
	}
	switch object {
	case URN:
		if game.Prop[URN] != URN_EMPTY {
			stateChange(URN, URN_DARK)
		} else {
			pspeak(URN, change, true, URN_DARK)
		}
	case LAMP:
		stateChange(LAMP, LAMP_DARK)
		ifElse(isDark(game.Loc),
			func() { speak(true, messages[PITCH_DARK]) },
			func() { speak(true, messages[NO_MESSAGE]) })
	case DRAGON, VOLCANO:
		speak(true, messages[BEYOND_POWER])
	default:
		speak(true, actions[verb].message)
	}
	return GO_CLEAROBJ
}

/*  Feed.  If bird, no seed.  Snake, dragon, troll: quip.  If dwarf, make him
 *  mad.  Bear, special. */
// C: feed
func feed(verb VerbType, object ObjectType) PhaseCodeType {
	switch object {
	case BIRD:
		speak(true, messages[BIRD_PINING])
	case DRAGON:
		ifElse(game.Prop[DRAGON] != DRAGON_BARS,
			func() { speak(true, messages[RIDICULOUS_ATTEMPT]) },
			func() { speak(true, messages[NOTHING_EDIBLE]) })
	case SNAKE:
		ifElse(!game.Closed && isHere(BIRD),
			func() {
				destroy(BIRD)
				speak(true, messages[BIRD_DEVOURED])
			},
			func() { speak(true, messages[NOTHING_EDIBLE]) })
	case TROLL:
		speak(true, messages[TROLL_VICES])
	case DWARF:
		ifElse(isHere(FOOD),
			func() {
				game.Dflag += 2
				speak(true, messages[REALLY_MAD])
			},
			func() { speak(true, actions[verb].message) })
	case BEAR:
		if game.Prop[BEAR] == BEAR_DEAD {
			speak(true, messages[RIDICULOUS_ATTEMPT])
		} else if game.Prop[BEAR] == UNTAMED_BEAR {
			if isHere(FOOD) {
				destroy(FOOD)
				game.Fixed[AXE] = FREE
				game.Prop[AXE] = AXE_HERE
				stateChange(BEAR, SITTING_BEAR)
			} else {
				speak(true, messages[NOTHING_EDIBLE])
			}
		} else {
			speak(true, actions[verb].message)
		}
	case OGRE:
		ifElse(isHere(FOOD),
			func() { speak(true, messages[OGRE_FULL]) },
			func() { speak(true, actions[verb].message) })
	default:
		speak(true, messages[AM_GAME])
	}
	return GO_CLEAROBJ
}

/*  Fill.  Bottle or urn must be empty, and liquid available.  (Vase
 *  is nasty.) */
// C: fill
func fill(verb VerbType, object ObjectType) PhaseCodeType {
	if object == VASE {
		if liquidLoc(game.Loc) == NO_OBJECT {
			speak(true, messages[FILL_INVALID])
			return GO_CLEAROBJ
		}
		if !isToting(VASE) {
			speak(true, messages[ARENT_CARRYING])
			return GO_CLEAROBJ
		}
		speak(true, messages[SHATTER_VASE])
		game.Prop[VASE] = VASE_BROKEN
		game.Fixed[VASE] = FIXED
		drop(VASE, game.Loc)
		return GO_CLEAROBJ
	}
	if object == URN {
		if game.Prop[URN] != URN_EMPTY {
			speak(true, messages[FULL_URN])
			return GO_CLEAROBJ
		}
		if !isHere(BOTTLE) {
			speak(true, messages[FILL_INVALID])
			return GO_CLEAROBJ
		}
		k := liquid()
		switch k {
		case WATER:
			game.Prop[BOTTLE] = EMPTY_BOTTLE
			speak(true, messages[WATER_URN])
		case OIL:
			game.Prop[URN] = URN_DARK
			game.Prop[BOTTLE] = EMPTY_BOTTLE
			speak(true, messages[OIL_URN])
		case NO_OBJECT:
			fallthrough
		default:
			speak(true, messages[FILL_INVALID])
			return GO_CLEAROBJ
		}
		game.Place[k] = LOC_NOWHERE
		return GO_CLEAROBJ
	}
	if object != INTRANSITIVE && object != BOTTLE {
		speak(true, actions[verb].message)
		return GO_CLEAROBJ
	}
	if object == INTRANSITIVE && !isHere(BOTTLE) {
		return GO_UNKNOWN
	}
	if isHere(URN) && game.Prop[URN] != URN_EMPTY {
		speak(true, messages[URN_NOPOUR])
		return GO_CLEAROBJ
	}
	if liquid() != NO_OBJECT {
		speak(true, messages[BOTTLE_FULL])
		return GO_CLEAROBJ
	}
	if liquidLoc(game.Loc) == NO_OBJECT {
		speak(true, messages[NO_LIQUID])
		return GO_CLEAROBJ
	}
	ifElse(liquidLoc(game.Loc) == OIL,
		func() { stateChange(BOTTLE, OIL_BOTTLE) },
		func() { stateChange(BOTTLE, WATER_BOTTLE) })
	return GO_CLEAROBJ
}

/* Find.  Might be carrying it, or it might be here.  Else give caveat. */
// C: find
func find(verb VerbType, object ObjectType) PhaseCodeType {
	if isToting(object) {
		speak(true, messages[ALREADY_CARRYING])
		return GO_CLEAROBJ
	}
	if game.Closed {
		speak(true, messages[NEEDED_NEARBY])
		return GO_CLEAROBJ
	}
	if isAt(object) || (liquid() == object && isAt(BOTTLE)) ||
		object == liquidLoc(game.Loc) || (object == DWARF && atdwrf(game.Loc) > 0) {
		speak(true, messages[YOU_HAVEIT])
		return GO_CLEAROBJ
	}
	speak(true, actions[verb].message)
	return GO_CLEAROBJ
}

/* Fly.  Snide remarks unless hovering rug is here. */
// C: fly
func fly(verb VerbType, object ObjectType) PhaseCodeType {
	if object == INTRANSITIVE {
		if !isHere(RUG) {
			speak(true, messages[FLAP_ARMS])
			return GO_CLEAROBJ
		}
		if game.Prop[RUG] != RUG_HOVER {
			speak(true, messages[RUG_NOTHING2])
			return GO_CLEAROBJ
		}
		object = RUG
	}
	if object != RUG {
		speak(true, actions[verb].message)
		return GO_CLEAROBJ
	}
	if game.Prop[RUG] != RUG_HOVER {
		speak(true, messages[RUG_NOTHING1])
		return GO_CLEAROBJ
	}
	if game.Loc == LOC_CLIFF {
		game.Oldlc2 = game.Oldloc
		game.Oldloc = game.Loc
		game.Newloc = LOC_LEDGE
		speak(true, messages[RUG_GOES])
	} else if game.Loc == LOC_LEDGE {
		game.Oldlc2 = game.Oldloc
		game.Oldloc = game.Loc
		game.Newloc = LOC_CLIFF
		speak(true, messages[RUG_RETURNS])
	} else {
		/* should never happen */
		speak(true, messages[NOTHING_HAPPENS])
	}
	return GO_TERMINATE
}

/* Inventory. If object treat same as find.  Else report on current burden. */
// C: inven
func inven() PhaseCodeType {
	empty := true
	var i ObjectType
	for i = 1; i <= NOBJECTS; i++ {
		if i == BEAR || !isToting(i) {
			continue
		}
		if empty {
			speak(true, messages[NOW_HOLDING])
			empty = false
		}
		pspeak(i, touch, false, -1)
	}
	if isToting(BEAR) {
		speak(false, messages[TAME_BEAR])
	}
	if empty {
		speak(true, messages[NO_CARRY])
	}
	return GO_CLEAROBJ
}

/*  Light.  Applicable only to lamp and urn. */
// C: light
func light(verb VerbType, object ObjectType) PhaseCodeType {
	if object == INTRANSITIVE {
		selects := 0
		if isHere(LAMP) && game.Prop[LAMP] == LAMP_DARK && game.Limit >= 0 {
			object = LAMP
			selects++
		}
		if isHere(URN) && game.Prop[URN] == URN_DARK {
			object = URN
			selects++
		}
		if selects != 1 {
			return GO_UNKNOWN
		}
	}
	switch object {
	case URN:
		ifElse(game.Prop[URN] == URN_EMPTY,
			func() { stateChange(URN, URN_EMPTY) },
			func() { stateChange(URN, URN_LIT) })
	case LAMP:
		if game.Limit < 0 {
			speak(true, messages[LAMP_OUT])
			break
		}
		stateChange(LAMP, LAMP_BRIGHT)
		if game.Wzdark {
			return GO_TOP
		}
	default:
		speak(true, actions[verb].message)
	}
	return GO_CLEAROBJ
}

/*  Listen.  Intransitive only.  Print stuff based on object sound proprties. */
// C: listen
func listen() PhaseCodeType {
	soundlatch := false
	var sound = locations[game.Loc].sound
	if sound != SILENT {
		speak(true, messages[sound])
		if !locations[game.Loc].loud {
			speak(true, messages[NO_MESSAGE])
		}
		soundlatch = true
	}
	var i ObjectType
	for i = 1; i <= NOBJECTS; i++ {
		if !isHere(i) || objects[i].sounds[0] == "" || game.Prop[i] < 0 {
			continue
		}
		mi := game.Prop[i]
		/* (ESR) Some unpleasant magic on object states here. Ideally
		 * we'd have liked the bird to be a normal object that we can
		 * use state_change() on; can't do it, because there are
		 * actually two different series of per-state birdsounds
		 * depending on whether player has drunk dragon's blood. */
		if i == BIRD && game.Blooded {
			mi += 3
		}
		pspeak(i, hear, true, mi, game.ZZword)
		speak(true, messages[NO_MESSAGE])
		if i == BIRD && mi == BIRD_ENDSTATE {
			destroy(BIRD)
		}
		soundlatch = true
	}
	if !soundlatch {
		speak(true, messages[ALL_SILENT])
	}
	return GO_CLEAROBJ
}

/* Lock, unlock, no object given.  Assume various things if present. */
// C: lock
func lock(verb VerbType, object ObjectType) PhaseCodeType {
	if object == INTRANSITIVE { // try to find which locakable Object
		if isHere(CLAM) {
			object = CLAM
		}
		if isHere(OYSTER) {
			object = OYSTER
		}
		if isAt(DOOR) {
			object = DOOR
		}
		if isAt(GRATE) {
			object = GRATE
		}
		if isHere(CHAIN) {
			object = CHAIN
		}
		if object == INTRANSITIVE {
			speak(true, messages[NOTHING_LOCKED])
			return GO_CLEAROBJ
		}
	}
	switch object {
	case CHAIN:
		if isHere(KEYS) {
			return chain(verb)
		} else {
			speak(true, messages[NO_KEYS])
		}
	case GRATE:
		if isHere(KEYS) {
			if game.Closng {
				speak(true, messages[EXIT_CLOSED])
				if !game.Panic {
					game.Clock2 = PANICTIME
				}
				game.Panic = true
			} else {
				ifElse(verb == LOCK,
					func() { stateChange(GRATE, GRATE_CLOSED) },
					func() { stateChange(GRATE, GRATE_OPEN) })
			}
		} else {
			speak(true, messages[NO_KEYS])
		}
	case CLAM:
		if verb == LOCK {
			speak(true, messages[HUH_MAN])
		} else if isToting(CLAM) {
			speak(true, messages[DROP_CLAM])
		} else if !isToting(TRIDENT) {
			speak(true, messages[CLAM_OPENER])
		} else {
			destroy(CLAM)
			drop(OYSTER, game.Loc)
			drop(PEARL, LOC_CULDESAC)
			speak(true, messages[PEARL_FALLS])
		}
	case OYSTER:
		if verb == LOCK {
			speak(true, messages[HUH_MAN])
		} else if isToting(OYSTER) {
			speak(true, messages[DROP_OYSTER])
		} else if !isToting(TRIDENT) {
			speak(true, messages[OYSTER_OPENER])
		} else {
			speak(true, messages[OYSTER_OPENS])
		}
	case DOOR:
		ifElse(game.Prop[DOOR] == DOOR_UNRUSTED,
			func() { speak(true, messages[OK_MAN]) },
			func() { speak(true, messages[RUSTY_DOOR]) })
	case CAGE:
		speak(true, messages[NO_LOCK])
	case KEYS:
		speak(true, messages[CANNOT_UNLOCK])
	default:
		speak(true, actions[verb].message)
	}
	return GO_CLEAROBJ
}

/*  Pour.  If no object, or object is bottle, assume contents of bottle.
 *  special tests for pouring water or oil on plant or rusty door. */
// C: pour
func pour(verb VerbType, object ObjectType) PhaseCodeType {
	if object == BOTTLE || object == INTRANSITIVE {
		object = liquid()
	}
	if object == NO_OBJECT {
		return GO_UNKNOWN
	}
	if !isToting(object) {
		speak(true, actions[verb].message)
		return GO_CLEAROBJ
	}
	if object != OIL && object != WATER {
		speak(true, messages[CANT_POUR])
		return GO_CLEAROBJ
	}
	if isHere(URN) && game.Prop[URN] == URN_EMPTY {
		return fill(verb, URN)
	}
	game.Prop[BOTTLE] = EMPTY_BOTTLE
	game.Place[object] = LOC_NOWHERE
	if !(isAt(PLANT) || isAt(DOOR)) {
		speak(true, messages[GROUND_WET])
		return GO_CLEAROBJ
	}
	if !isAt(DOOR) {
		if object == WATER {
			/* cycle through the three plant states */
			n := (game.Prop[PLANT] + 1) % 3
			stateChange(PLANT, n)
			game.Prop[PLANT2] = game.Prop[PLANT]
			return GO_MOVE
		} else {
			speak(true, messages[SHAKING_LEAVES])
			return GO_CLEAROBJ
		}
	} else {
		ifElse(object == OIL,
			func() { stateChange(DOOR, DOOR_UNRUSTED) },
			func() { stateChange(DOOR, DOOR_RUSTED) })
		return GO_CLEAROBJ
	}
}

/*  Quit.  Intransitive only.  Verify intent and exit if that's what he wants. */
func quit() PhaseCodeType {
	if askYesNo(messages[REALLY_QUIT], messages[OK_MAN], messages[OK_MAN]) {
		terminate(term_quit)
	}
	return GO_CLEAROBJ
}

/*  Read.  Print stuff based on objtxt.  Oyster (?) is special case. */
func read(command Command) PhaseCodeType {
	if command.obj == INTRANSITIVE {
		command.obj = NO_OBJECT
		var i ObjectType
		for i = 1; i <= NOBJECTS; i++ {
			if isHere(i) && objects[i].texts[0] != "" && game.Prop[i] >= 0 {
				command.obj = command.obj*NOBJECTS + i
			}
		}
		if command.obj > NOBJECTS || command.obj == NO_OBJECT || isDark(game.Loc) {
			return GO_UNKNOWN
		}
	}
	if isDark(game.Loc) {
		speak(true, messages[NO_SEE], command.word[0].raw)
	} else if command.obj == OYSTER {
		if !isToting(OYSTER) || !game.Closed {
			speak(true, messages[DONT_UNDERSTAND])
		} else if !game.Clshnt {
			game.Clshnt = askYesNo(messages[CLUE_QUERY], messages[WAYOUT_CLUE], messages[OK_MAN])
		} else {
			pspeak(OYSTER, hear, true, 1) // Not really a sound, but oh well.
		}
	} else if objects[command.obj].texts[0] == "" || game.Prop[command.obj] == STATE_NOTFOUND {
		speak(true, actions[command.verb].message)
	} else {
		pspeak(command.obj, study, true, game.Prop[command.obj])
	}
	return GO_CLEAROBJ
}

/*  Z'ZZZ (word gets recomputed at startup; different each game). */
// C: reservoir
func reservoir() PhaseCodeType {
	if !isAt(RESER) && game.Loc != LOC_RESBOTTOM {
		speak(true, messages[NOTHING_HAPPENS])
		return GO_CLEAROBJ
	}
	ifElse(game.Prop[RESER] == WATERS_PARTED,
		func() { stateChange(RESER, WATERS_UNPARTED) },
		func() { stateChange(RESER, WATERS_PARTED) })
	if isAt(RESER) {
		return GO_CLEAROBJ
	}
	game.Oldlc2 = game.Loc
	game.Newloc = LOC_NOWHERE
	speak(true, messages[NOT_BRIGHT])
	return GO_TERMINATE
}

/* Rub.  Yields various snide remarks except for lit urn. */
// C: rub
func rub(verb VerbType, object ObjectType) PhaseCodeType {
	if object == URN && game.Prop[URN] == URN_LIT {
		destroy(URN)
		drop(AMBER, game.Loc)
		game.Prop[AMBER] = AMBER_IN_ROCK
		game.Tally--
		drop(CAVITY, game.Loc)
		speak(true, messages[URN_GENIES])
	} else if object != LAMP {
		speak(true, messages[PECULIAR_NOTHING])
	} else {
		speak(true, actions[verb].message)
	}
	return GO_CLEAROBJ
}

/* Say.  Echo WD2. Magic words override. */
// C: say
func say(command Command) PhaseCodeType {
	if command.word[1].typ == MOTION &&
		(command.word[1].id == XYZZY || command.word[1].id == PLUGH || command.word[1].id == PLOVER) {
		return GO_WORD2
	}
	if command.word[1].typ == ACTION && command.word[1].id == PART {
		return reservoir()
	}
	if command.word[1].typ == ACTION &&
		(command.word[1].id == FEE || command.word[1].id == FIE || command.word[1].id == FOE ||
			command.word[1].id == FOO || command.word[1].id == FUM || command.word[1].id == PART) {
		return bigWords(command.word[1].id)
	}
	speak(true, messages[OKEY_DOKEY], command.word[1].raw)
	return GO_CLEAROBJ
}

func throwSupport(spk VocabType) PhaseCodeType {
	speak(true, messages[spk])
	drop(AXE, game.Loc)
	return GO_MOVE
}

/*  Throw.  Same as discard unless axe.  Then same as attack except
 *  ignore bird, and if dwarf is present then one might be killed.
 *  (Only way to do so!)  Axe also special for dragon, bear, and
 *  troll.  Treasures special for troll. */
// C: throwit
func throwit(command Command) PhaseCodeType {
	if !isToting(command.obj) {
		speak(true, actions[command.verb].message)
		return GO_CLEAROBJ
	}
	if objects[command.obj].treasure && isAt(TROLL) {
		/*  Snarf a treasure for the troll. */
		drop(command.obj, LOC_NOWHERE)
		move(TROLL, LOC_NOWHERE)
		move(TROLL+NOBJECTS, FREE)
		drop(TROLL2, objects[TROLL].plac)
		drop(TROLL2+NOBJECTS, objects[TROLL].fixd)
		juggle(CHASM)
		speak(true, messages[TROLL_SATISFIED])
		return GO_CLEAROBJ
	}
	if command.obj == FOOD && isHere(BEAR) {
		/* But throwing food is another story. */
		command.obj = BEAR
		return feed(command.verb, command.obj)
	}
	if command.obj != AXE {
		return discard(command.verb, command.obj)
	}
	if atdwrf(game.Loc) <= 0 {
		if isAt(DRAGON) && game.Prop[DRAGON] == DRAGON_BARS {
			return throwSupport(DRAGON_SCALES)
		}
		if isAt(TROLL) {
			return throwSupport(TROLL_RETURNS)
		}
		if isAt(OGRE) {
			return throwSupport(OGRE_DODGE)
		}
		if isHere(BEAR) && game.Prop[BEAR] == UNTAMED_BEAR {
			/* This'll teach him to throw the axe at the bear! */
			drop(AXE, game.Loc)
			game.Fixed[AXE] = FIXED
			juggle(BEAR)
			stateChange(AXE, AXE_LOST)
			return GO_CLEAROBJ
		}
		command.obj = INTRANSITIVE
		return attack(command)
	}
	if randrange(NDWARVES+1) < int32(game.Dflag) {
		return throwSupport(DWARF_DODGES)
	}
	loc := atdwrf(game.Loc)
	game.Dseen[loc] = false
	game.Dloc[loc] = LOC_NOWHERE
	game.Dkill++
	var p PhaseCodeType
	ifElse(game.Dkill == 1,
		func() { p = throwSupport(DWARF_SMOKE) },
		func() { p = throwSupport(KILLED_DWARF) })
	return p
}

/* Wake.  Only use is to disturb the dwarves. */
// C: wake
func wake(verb VerbType, object ObjectType) PhaseCodeType {
	if object != DWARF || !game.Closed {
		speak(true, actions[verb].message)
		return GO_CLEAROBJ
	}
	speak(true, messages[PROD_DWARF])
	return GO_DWARFWAKE

}

/* Set seed */
func seed(verb VerbType, arg string) PhaseCodeType {
	seed, err := strconv.Atoi(arg)
	if err == nil {
		speak(true, actions[verb].message, seed)
		setSeed(int32(seed))
		game.Turns++
	}
	return GO_TOP
}

/* Burn Turns */
func waste(verb VerbType, turns TurnType) PhaseCodeType {
	game.Limit -= turns
	speak(true, actions[verb].message, game.Limit)
	return GO_TOP
}

/* Wave.  No effect unless waving rod at fissure or at bird. */
// C: wave
func wave(verb VerbType, object ObjectType) PhaseCodeType {
	if object != ROD || !isToting(object) || (!isHere(BIRD) && (game.Closng || !isAt(FISSURE))) {
		ifElse(!isToting(object) && (object != ROD || !isToting(ROD2)),
			func() { speak(true, messages[ARENT_CARRYING]) },
			func() { speak(true, actions[verb].message) })
		return GO_CLEAROBJ
	}
	if game.Prop[BIRD] == BIRD_UNCAGED && game.Loc == game.Place[STEPS] && game.Prop[JADE] == STATE_NOTFOUND {
		drop(COND_JADE, game.Loc)
		game.Prop[JADE] = STATE_FOUND
		game.Tally--
		speak(true, messages[NECKLACE_FLY])
		return GO_CLEAROBJ
	} else {
		if game.Closed {
			ifElse(game.Prop[BIRD] == BIRD_CAGED,
				func() { speak(true, messages[CAGE_FLY]) },
				func() { speak(true, messages[FREE_FLY]) })
			return GO_DWARFWAKE
		}
		if game.Closng || !isAt(FISSURE) {
			ifElse(game.Prop[BIRD] == BIRD_CAGED,
				func() { speak(true, messages[CAGE_FLY]) },
				func() { speak(true, messages[FREE_FLY]) })
			return GO_CLEAROBJ
		}
		if isHere(BIRD) {
			ifElse(game.Prop[BIRD] == BIRD_CAGED,
				func() { speak(true, messages[CAGE_FLY]) },
				func() { speak(true, messages[FREE_FLY]) })
		}
		ifElse(game.Prop[FISSURE] == BRIDGED,
			func() { stateChange(FISSURE, UNBRIDGED) },
			func() { stateChange(FISSURE, BRIDGED) })

		return GO_CLEAROBJ
	}
}
