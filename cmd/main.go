package cmd

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

func Main(s Settings) {
	settings = s
	game.Novice = askYesNo(messages[WELCOME_YOU], messages[CAVE_NEARBY], messages[NO_MESSAGE])
	if game.Novice {
		game.Limit = NOVICELIMIT
	}
	initGame()

	settings.TTY.Buttons[0].OnTapped = func() {
		score(term_score)
		settings.TTY.Focus()
	}
	settings.TTY.Buttons[1].OnTapped = func() {
		inven()
		settings.TTY.Focus()
	}

	/* interpret commands until EOF or interrupt
	* show score and exit
	* interpret commands until EOF or interrupt */
	for {
		// if we're supposed to move, move
		if !doMove() {
			continue
		}
		// get command
		if !doCommand() {
			break
		}
	}
	terminate(term_end)
}
func initGame() {
	var init = func() {
		for i := 1; i <= NOBJECTS; i++ {
			game.Place[i] = LOC_NOWHERE
		}
		for i := 1; i <= NLOCATIONS; i++ {
			l := locations[i]
			if !(l.descriptions.big == "" || keys[i] == 0) {
				k := keys[i]
				if travels[k].isTerminate() {
					conditions[i] |= COND_FORCED
				}
			}
		}
		/*  Set up the game.Atloc and game.Link arrays.
		*  We'll use the DROP subroutine, which prefaces new objects on the
		*  lists.  Since we want things in the other order, we'll run the
		*  loop backwards.  If the object is in two locs, we drop it twice.
		*  Also, since two-placed objects are typically best described
		*  last, we'll drop them first.	 */
		for i := NOBJECTS; i >= 1; i-- {
			if objects[i].fixd > 0 {
				drop(ObjectType(i+NOBJECTS), objects[i].fixd)
				drop(ObjectType(i), objects[i].plac)
			}
		}
		for i := 1; i <= NOBJECTS; i++ {
			k := NOBJECTS + 1 - i
			game.Fixed[k] = objects[k].fixd
			// fixd: -1 == FIXED, 0 == FREE
			if objects[k].plac != 0 && objects[k].fixd <= 0 {
				drop(ObjectType(k), objects[k].plac)
			}
		}
		/*  Treasure props are initially -1, and are set to 0 the first time
		 *  they are described.  game.Tally keeps track of how many are
		 *  not yet found, so we know when to close the cave. */
		for treasure := 1; treasure <= NOBJECTS; treasure++ {
			if objects[treasure].treasure {
				if objects[treasure].inventory != "" {
					game.Prop[treasure] = STATE_NOTFOUND // -1
				}
				game.Tally = game.Tally - game.Prop[treasure]
			}
		}
		game.Conds = COND_CAVE
	}
	init()
	if settings.AllowSave {
		_ = load()
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s1)
	setSeed(random.Int31())
}

/* End of game.  Let's tell him all about it. */
func terminate(mode int) {
	points, mxscor := score(mode)
	if points+game.Trnluz+1 >= mxscor && game.Trnluz != 0 {
		speak(false, messages[TOOK_LONG])
	}
	if points+game.Saved+1 >= mxscor && game.Saved != 0 {
		speak(false, messages[WITHOUT_SUSPENDS])
	}
	speak(true, messages[TOTAL_SCORE], points, mxscor, game.Turns, plural(int(game.Turns)))
	for i := 1; i <= NCLASSES; i++ {
		if classes[i].threshold >= points {
			speak(true, classes[i].message)
			i = classes[i].threshold + 1 - points
			speak(true, messages[NEXT_HIGHER], i, plural(i))
			speak(true, "Adventure Ends. Press Enter")
			getInput()
			os.Exit(mode)
		}
	}
	speak(false, messages[OFF_SCALE])
	speak(false, messages[NO_HIGHER])
	speak(true, "Adventure Ends. Press Enter")
	getInput() // insures score is readable before windows closes.
	os.Exit(mode)
}

/* Actually execute the move to the new location and dwarf movement */
func doMove() bool {
	if settings.Logger {
		l := game.Loc
		n := game.Newloc
		lc := ConditionText(uint32(l))
		nc := ConditionText(uint32(n))
		log.Printf("doMove: from %s (%s) to %s (%s)", LocationText(l), lc, LocationText(n), nc)
	}
	//  Can't leave cave once it's closing (except by main office).
	if isOutside(game.Newloc) && game.Newloc != 0 && game.Closng {
		speak(false, messages[EXIT_CLOSED])
		game.Newloc = game.Loc
		if !game.Panic {
			game.Clock2 = 15
		}
		game.Panic = true
	}
	/*  See if a dwarf has seen him and has come from where he
	 *  wants to go.  If so, the dwarf's blocking his way.  If
	 *  coming from Place forbidden to pirate (dwarves rooted in
	 *  Place) let him get out (and attacked). */
	if game.Newloc != game.Loc && !isForced(game.Loc) && !isNoArrr(game.Loc) {
		for i := 1; i <= NDWARVES-1; i++ {
			if game.Odloc[i] == game.Newloc && game.Dseen[i] {
				game.Newloc = game.Loc
				speak(false, messages[DWARF_BLOCK])
				settings.CRT.Show(false, "dwarf")
				break
			}
		}
	}
	game.Loc = game.Newloc // show movement
	if !dwarfMove() {
		croak()
	}
	if game.Loc == LOC_NOWHERE {
		croak()
	}
	/* The easiest way to get killed is to fall into a pit in
	 * pitch darkness. */
	if !isForced(game.Loc) && isDark(game.Loc) && game.Wzdark && pct(35) {
		speak(true, messages[PIT_FALL])
		game.Oldlc2 = game.Loc
		croak()
		return false
	}
	settings.CRT.Show(true, locations[game.Loc].pic)
	return true
}

/*  "You're dead, Jim."
 *
 *  If the current Loc is zero, it means the clown got himself killed.
 *  We'll allow this maxdie times.  NDEATHS is automatically set based
 *  on the number of snide messages available.  Each death results in
 *  a message (obituaries[n]) which offers reincarnation; if accepted,
 *  this results in message obituaries[0], obituaries[2], etc.  The
 *  last time, if he wants another chance, he gets a snide remark as
 *  we exit.  When reincarnated, all objects being carried get dropped
 *  at game.Oldlc2 (presumably the last Place prior to being killed)
 *  without change of props.  The loop runs backwards to assure that
 *  the bird is dropped before the cage.  (This kluge could be changed
 *  once we're sure all references to bird and cage are done by
 *  keywords.)  The lamp is a special case (it wouldn't do to leave it
 *  in the cave). It is turned off and left outside the building (only
 *  if he was carrying it, of course).  He himself is left inside the
 *  building (and heaven help him if he tries to xyzzy back into the
 *  cave without the lamp!).  game.Oldloc is zapped so, he can't just
 *  "retreat". */

/*  Okay, he's dead.  Let's get on with it. */
func croak() {
	query := obituaries[game.Numdie].query
	response := obituaries[game.Numdie].response
	game.Numdie++
	if game.Closng {
		//  He died during closing time.  No resurrection.  Tally up a
		//*  death and exit.
		speak(true, messages[DEATH_CLOSING])
		terminate(term_end)
	} else if !askYesNo(query, response, messages[OK_MAN]) || game.Numdie == NDEATHS {
		// Player is asked if he wants to try again. If not, or if
		// he's already used all of his lives, we end the game
		terminate(term_end)
	} else {
		/* If player wishes to continue, we empty the liquids in the
		* user's inventory, turn off the lamp, and drop all items
		* where he died. */
		game.Place[WATER] = LOC_NOWHERE
		game.Place[OIL] = LOC_NOWHERE
		if isToting(LAMP) {
			game.Prop[LAMP] = LAMP_DARK
		}
		var i ObjectType
		var j ObjectType
		for j = 1; j <= NOBJECTS; j++ {
			i = NOBJECTS + 1 - j
			if isToting(i) {
				// Always leave lamp where it's accessible aboveground
				if i == LAMP {
					drop(LAMP, LOC_START)
				} else {
					drop(i, game.Oldlc2)
				}
			}
		}
		game.Oldloc = LOC_BUILDING
		game.Loc = LOC_BUILDING
		game.Newloc = LOC_BUILDING
	}
}

/* mode is 'scoregame' (term_score) if scoring,
* 'quitgame' (term_quit) if quitting,
* 'endgame' (term_end) if died or won */
func score(mode int) (int, int) {
	s := 0
	/*  The present scoring algorithm is as follows:
	 *     Objective:          Points:        Present total possible:
	 *  Getting well into cave   25                    25
	 *  Each treasure < chest    12                    60
	 *  Treasure chest itself    14                    14
	 *  Each treasure > chest    16                   224
	 *  Surviving             (MAX-NUM)*10             30
	 *  Not quitting              4                     4
	 *  Reaching "game.Closng"   25                    25
	 *  "Closed": Quit/Killed    10
	 *            Klutzed        25
	 *            Wrong way      30
	 *            Success        45                    45
	 *  Came to Witt's End        1                     1
	 *  Round out the total       2                     2
	 *                                       TOTAL:   430
	 *  Points can also be deducted for using hints or too many Turns, or for
	 *  saving intermediate positions. */
	/*  First Tally up the treasures.  Must be in building and not broken.
	 *  Give the poor guy 2 points just for finding each treasure. */
	mxscor := 0
	for i := 1; i <= NOBJECTS; i++ {
		if !objects[i].treasure {
			continue
		}
		if objects[i].inventory != "" {
			k := 12
			if i == CHEST {
				k = 14
			}
			if i > CHEST {
				k = 16
			}
			if game.Prop[i] > STATE_NOTFOUND {
				s += 2
			}
			if game.Place[i] == LOC_BUILDING && game.Prop[i] == STATE_FOUND {
				s += k - 2
			}
			mxscor += k
		}
	}
	/*  Now look at how he finished and how far he got.  NDEATHS and
	 *  game.Numdie tell us how well he survived.  game.Dflag will tell us
	 *  if he ever got suitably deep into the cave.  game.Closng still
	 *  indicates whether he reached the endgame.  And if he got as far as
	 *  "cave Closed" (indicated by "game.Closed"), then Bonus is zero for
	 *  mundane exits or 133, 134, 135 if he blew it (so to speak). */
	s += (NDEATHS - int(game.Numdie)) * 10
	mxscor += NDEATHS * 10
	if mode == term_end {
		s += 4
	}
	mxscor += 4
	if game.Dflag != 0 {
		s += 25
	}
	mxscor += 25
	if game.Closng {
		s += 25
	}
	mxscor += 25
	if game.Closed {
		switch game.Bonus {
		case bonus_none:
			s += 10
		case bonus_splatter:
			s += 25
		case bonus_defeat:
			s += 30
		case bonus_victory:
			s += 45

		}
	}
	/* Did he come to Witt's End as he should? */
	if game.Place[MAGAZINE] == LOC_WITTSEND {
		s += 1
	}
	mxscor += 1
	/* Round it off. */
	s += 2
	mxscor += 2
	/* Deduct for hints/Turns/saves. Hints < 4 are special; see database desc. */
	for i := 0; i < NHINTS; i++ {
		if game.Hinted[i] {
			s = s - hints[i].penalty
		}
	}
	if game.Novice {
		s -= 5
	}
	if game.Clshnt {
		s -= 10
	}
	s = s - game.Trnluz - game.Saved
	/* Return to score command if that's where we came from. */
	if mode == term_score {
		speak(true, messages[GARNERED_POINTS], score, mxscor, game.Turns, plural(int(game.Turns)))
	}
	return s, mxscor
}

/* Dwarves move.  Return true if player survives, false if he dies. */
func dwarfMove() bool {
	//int kk, stick, attack;
	var tk [21]LocationType

	/*  Dwarf stuff.  See earlier comments for description of
	 *  variables.  Remember sixth dwarf is pirate and is thus
	 *  very different except for motion rules. */

	/*  First off, don't let the dwarves follow him into a pit or a
	 *  wall.  Activate the whole mess the first time he gets as far
	 *  as the Hall of Mists (what INDEEP() tests).  If game.Newloc
	 *  is forbidden to pirate (in particular, if it's beyond the
	 *  troll bridge), bypass dwarf stuff.  That way pirate can't
	 *  steal return toll, and dwarves can't meet the bear.  Also,
	 *  means dwarves won't follow him into dead end in maze, but
	 *  c'est la vie.  They'll wait for him outside the dead end. */
	if game.Loc == LOC_NOWHERE || isForced(game.Loc) || isNoArrr(game.Newloc) {
		return true
	}
	/* Dwarf activity level ratchets up */
	if game.Dflag == 0 {
		if inDeep(game.Loc) {
			game.Dflag = 1
		}
		return true
	}
	/*  When we encounter the first dwarf, we kill 0, 1, or 2 of
	 *  the 5 dwarves.  If any of the survivors is at game.Loc,
	 *  replace him with the alternate. */
	if game.Dflag == 1 {
		if !inDeep(game.Loc) || (pct(95) && (!isNoBack(game.Loc) || pct(85))) {
			return true
		}
		game.Dflag = 2
		for i := 1; i <= 2; i++ {
			j := 1 + randrange(NDWARVES-1)
			if pct(50) {
				game.Dloc[j] = 0
			}
		}
		/* Alternate initial Loc for dwarf, in case one of them
		 *  starts out on top of the adventurer. */
		for i := 1; i <= NDWARVES-1; i++ {
			if game.Dloc[i] == game.Loc {
				game.Dloc[i] = DALTLC
			}
			game.Odloc[i] = game.Dloc[i]
		}
		speak(true, messages[DWARF_RAN])
		settings.CRT.Show(false, "dwarf")
		drop(AXE, game.Loc)
		return true
	}
	/*  Things are in full swing.  Move each dwarf at random,
	 *  except if he's seen us he sticks with us.  Dwarves stay
	 *  deep inside.  If wandering at random, they don't back up
	 *  unless there's no alternative.  If they don't have to
	 *  move, they attack.  And, of course, dead dwarves don't do
	 *  much of anything. */
	game.Dtotal = 0
	att := 0
	stick := 0
	for i := 1; i <= NDWARVES; i++ {
		if game.Dloc[i] == 0 {
			continue
		}
		/*  Fill tk array with all the places this dwarf might go. */
		j := 1
		kk := keys[game.Dloc[i]]
		if kk != 0 {
			for {
				if travels[kk].stop {
					break
				}
				game.Newloc = travels[kk].destval
				switch { // Have we avoided a dwarf encounter?
				case travels[kk].desttype != dest_goto,
					!inDeep(game.Newloc),
					game.Odloc[i] == game.Newloc,
					j > 1 && game.Newloc == tk[j-1],
					j >= len(tk)-1,
					game.Newloc == game.Dloc[i],
					isForced(game.Newloc),
					isNoArrr(game.Newloc) && PIRATE == i,
					travels[kk].nodwarves:
				default:
					tk[j] = game.Newloc
					j++
				}
				kk++
			}
		}
		tk[j] = game.Odloc[i]
		if j >= 2 {
			j--
		}
		j = int(1 + randrange(int32(j)))
		game.Odloc[i] = game.Dloc[i]
		game.Dloc[i] = tk[j]
		game.Dseen[i] = (game.Dseen[i] && inDeep(game.Loc)) ||
			(game.Dloc[i] == game.Loc || game.Odloc[i] == game.Loc)
		if !game.Dseen[i] {
			continue
		}
		game.Dloc[i] = game.Loc
		if spottedByPirate(i) {
			continue
		}
		/* This threatening little dwarf is in the room with him! */
		game.Dtotal++
		if game.Odloc[i] == game.Dloc[i] {
			att++
			if game.Knfloc >= 0 {
				game.Knfloc = game.Loc
			}
			if int(randrange(1000)) < 95*(game.Dflag-2) {
				stick++
			}
		}
	}
	/*  Now we know what's happening.  Let's tell the poor sucker about it. */
	if game.Dtotal == 0 {
		return true
	}
	if game.Dtotal == 1 {
		speak(true, messages[DWARF_SINGLE])
	} else {
		speak(true, messages[DWARF_PACK], game.Dtotal)
	}
	if att == 0 {
		return true
	}
	if game.Dflag == 2 {
		game.Dflag = 3
	}
	if att > 1 {
		speak(true, messages[THROWN_KNIVES], attack)
		switch {
		case stick == 0:
			speak(true, messages[NONE_HIT])
		case stick == 1:
			speak(true, messages[ONE_HIT])
		default:
			speak(true, messages[MULTIPLE_HITS], stick)
		}
	} else {
		speak(true, messages[KNIFE_THROWN])
		switch {
		case stick == 0:
			speak(true, messages[MISSES_YOU])
		default:
			speak(true, messages[GETS_YOU])
		}
	}
	if stick == 0 {
		return true
	}
	game.Oldlc2 = game.Loc
	return false
}

/*  Given the current location in "game.Loc", and a motion verb number in
 *  "motion", put the new location in "game.Newloc".  The current Loc is Saved
 *  in "game.Oldloc" in case he wants to retreat.  The current
 *  game.Oldloc is Saved in game.Oldlc2, in case he dies.  (if he
 *  does, game.Newloc will be limbo, and game.Oldloc will be what killed
 *  him, so we need game.Oldlc2, which is the last Place he was
 *  safe.) */
func playerMove(motion int) {
	travelEntry := keys[game.Loc]
	game.Newloc = game.Loc
	if travelEntry == LOC_NOWHERE {
		panic("playerMove: No travel entries")
	}
	switch motion {
	case NUL:
		return
	case BACK:
		/*  Handle "go back".  Look for verb which goes from game.Loc to
		 *  game.Oldloc, or to game.Oldlc2 If game.Oldloc has forced-motion.
		 *  te_tmp saves entry -> forced Loc -> previous Loc. */
		previous := game.Oldloc
		if isForced(previous) {
			previous = game.Oldlc2
		}
		game.Oldlc2 = game.Oldloc
		game.Oldloc = game.Loc
		if isNoBack(game.Loc) {
			speak(true, messages[TWIST_TURN])
			return
		}
		if previous == game.Loc {
			speak(true, messages[FORGOT_PATH])
			return
		}
		travelEntryTemp := 0
		for { // find appropriate motion to go BACK to previous "loc"
			desttype := travels[travelEntry].desttype
			dest := travels[travelEntry].destval // LocationType
			if desttype != dest_goto || dest != previous {
				if desttype == dest_goto {
					if isForced(dest) && travels[keys[dest]].destval == previous {
						travelEntryTemp = travelEntry
					}
				}
				if !travels[travelEntry].stop {
					travelEntry++ /* go to next travel entry for this location */
					continue
				}
				/* we've reached the end of travel entries for game.Loc */
				travelEntry = travelEntryTemp
				if travelEntry == 0 {
					speak(true, messages[NOT_CONNECTED])
					return
				}
			}
			// get to previous from current
			motion = travels[travelEntry].motion
			travelEntry = keys[game.Loc]
			break
		}
	case LOOK:
		/*  Look.  Can't give more Detail.  Pretend it wasn't dark
		 *  (though it may now be dark) so he won't fall into a
		 *  pit while staring into the gloom. */
		if game.Detail < 3 {
			speak(true, messages[NO_MORE_DETAIL])
		}
		game.Detail++
		game.Wzdark = false
		game.Abbrev[game.Loc] = 0
		return
	case CAVE:
		ifElse(isOutside(game.Loc) && game.Loc != LOC_GRATE,
			func() { speak(true, messages[FOLLOW_STREAM]) },
			func() { speak(true, messages[NEED_DETAIL]) })
		return
	default: // none of the specials
		game.Oldlc2 = game.Oldloc
		game.Oldloc = game.Loc
	}
	/* Look for a way to fulfil the motion verb passed in - travel_entry indexes
	 * the beginning of the motion entries for here (game.Loc). */
	for {
		//if settings.Logger {
		//	log.Printf("travel entry: Loc %s New %s\n", LocationText(game.Loc), LocationText(game.Newloc))
		//	log.Printf("travel entry: %s\n", &travels[travelEntry])
		//}
		if travels[travelEntry].isTerminate() || travels[travelEntry].motion == motion {
			break
		}
		if travels[travelEntry].stop {
			/*  Couldn't find an entry matching the motion word passed
			 *  in.  Various messages depending on word given. */
			switch motion {
			case EAST, WEST, SOUTH, NORTH, NE, NW, SE, SW, UP, DOWN:
				speak(true, messages[BAD_DIRECTION])
			case FORWARD, LEFT, RIGHT:
				speak(true, messages[UNSURE_FACING])
			case OUTSIDE, INSIDE:
				speak(true, messages[NO_INOUT_HERE])
			case XYZZY:
				speak(true, messages[NOTHING_HAPPENS])
				settings.CRT.Show(false, "word_xyzzy")
			case PLUGH:
				speak(true, messages[NOTHING_HAPPENS])
				settings.CRT.Show(false, "word_plugh")
			case CRAWL:
				speak(true, messages[WHICH_WAY])
			default:
				speak(true, messages[CANT_APPLY])
			}
			return
		}
		travelEntry++
	}
	/* (ESR) We've found a destination that goes with the motion verb.
	 * Next we need to check any conditional(s) on this destination, and
	 * possibly on following entries. */
	for { // L12 loop
		for {
			condtype := travels[travelEntry].condtype
			condarg1 := travels[travelEntry].condarg1
			condarg2 := travels[travelEntry].condarg2
			if condtype < cond_not {
				if condtype == cond_goto || condtype == cond_pct {
					if condarg1 == 0 || pct(int32(condarg1)) {
						break
					}
				} else if isToting(ObjectType(condarg1)) || (condtype == cond_with && isAt(ObjectType(condarg1))) {
					break
				}
			} else if game.Prop[condarg1] != condarg2 {
				break
			}
			/* We arrive here on conditional failure.
			 * Skip to next non-matching destination */
			travelEntrytemp := travelEntry
			for {
				if travels[travelEntrytemp].stop {
					panic("playermove: travel with NO alternation")
				}
				travelEntrytemp++
				if !(traveleq(travelEntry, travelEntrytemp)) {
					break
				}
			}
			travelEntry = travelEntrytemp
		}
		/* Found an eligible rule, now execute it */
		desttype := travels[travelEntry].desttype
		game.Newloc = travels[travelEntry].destval
		if desttype == dest_goto {
			settings.CRT.Show(true, locations[game.Newloc].pic)
			return
		}
		if desttype == dest_speak {
			/* Execute a speak rule */
			speak(true, messages[game.Newloc])
			game.Newloc = game.Loc // and leave in place
			return
		}
		switch game.Newloc {
		case 1:
			/* Special travel 1.  Plover-alcove passage.  Can carry only
			 * emerald.  Note: travel table must include "useless"
			 * entries going through passage, which can never be used
			 * for actual motion, but can be spotted by "go back". */
			ifElse(game.Loc == LOC_PLOVER,
				func() { game.Newloc = LOC_ALCOVE },
				func() { game.Newloc = LOC_PLOVER })
			if game.Holdng > 1 ||
				(game.Holdng == 1 && !isToting(EMERALD)) {
				game.Newloc = game.Loc
				speak(true, messages[MUST_DROP])
			}
			return
		case 2:
			/* Special travel 2.  Plover transport.  Drop the
			 * emerald (only use special travel if toting
			 * it), so he's forced to use the plover-passage
			 * to get it out.  Having dropped it, go back and
			 * pretend he wasn't carrying it after all. */
			drop(EMERALD, game.Loc)
			travelEntrytemp := travelEntry
			for {
				if travels[travelEntrytemp].stop {
					panic("playermove: travel with NO alternation")
				}
				travelEntrytemp++
				if !(traveleq(travelEntry, travelEntrytemp)) {
					break
				}
			}
			travelEntry = travelEntrytemp
		case 3:
			/* Special travel 3.  Troll bridge.  Must be done
			 * only as special motion so that dwarves won't
			 * wander across and encounter the bear.  (They
			 * won't follow the player there because that
			 * region is forbidden to the pirate.)  If
			 * game.Prop[TROLL]=TROLL_PAIDONCE, he's crossed
			 * since paying, so step out and block him.
			 * (standard travel entries check for
			 * game.Prop[TROLL]=TROLL_UNPAID.)  Special stuff
			 * for bear. */
			if game.Prop[TROLL] == TROLL_PAIDONCE {
				pspeak(TROLL, look, true, TROLL_PAIDONCE)
				game.Prop[TROLL] = TROLL_UNPAID
				destroy(TROLL2)
				move(TROLL2, LOC_NOWHERE)
				move(TROLL2+NOBJECTS, FREE)
				move(TROLL, objects[TROLL].plac)
				move(TROLL+NOBJECTS, objects[TROLL].fixd)
				juggle(CHASM)
				game.Newloc = game.Loc
				return
			}
			game.Newloc = objects[TROLL].plac + objects[TROLL].fixd - game.Loc
			if game.Prop[TROLL] == TROLL_UNPAID {
				game.Prop[TROLL] = TROLL_PAIDONCE
			}
			if !isToting(BEAR) {
				return
			}
			stateChange(CHASM, BRIDGE_WRECKED)
			game.Prop[TROLL] = TROLL_GONE
			drop(BEAR, game.Newloc)
			game.Fixed[BEAR] = FIXED
			game.Prop[BEAR] = BEAR_DEAD
			game.Oldlc2 = game.Newloc
			croak()
			return
		default:
			panic("playermove: travel exceeds 300")
		}
	} // L12
}

/* Get and execute a command */
// C: do_command
func doCommand() bool {
	var command = Command{}
	clearCommand(&command)
	// Describe the current location and (maybe) get next command.
	for { // loop until Command is EXECUTED. the return bool.
		if command.state == EXECUTED {
			break
		}
		describeLocation()
		settings.CRT.Show(true, locations[game.Loc].pic)
		if isForced(game.Loc) { // PLUGH, PLOVER, XYZZY redirection
			playerMove(HERE)
			return true
		}
		listObjects()
		clearCommand(&command)
		for { // get a Command loop.  until "GIVEN"
			// loop while EMPTY, RAW, TOKENIZED, GIVEN
			// ignore loop (break) when PREPROCESSED, PROCESS, EXECUTED
			if command.state > GIVEN {
				break
			}
			/*  If closing time, check for any objects being toted with
			*  game.Prop < 0 and stash them.  This way objects won't be
			*  described until they've been picked up and put down
			*  separate from their respective piles. */
			if game.Closed {
				if game.Prop[OYSTER] < 0 && isToting(OYSTER) {
					pspeak(OYSTER, look, true, 1)
				}
				var i ObjectType
				for i = 1; i <= NOBJECTS; i++ {
					if isToting(i) && game.Prop[i] < 0 {
						game.Prop[i] = stashed(i)
					}
				}
			}
			// Check to see if the room is dark. If the knife is here,
			// and it's dark, the knife permanently disappears
			game.Wzdark = isDark(game.Loc)
			if game.Knfloc != LOC_NOWHERE && game.Knfloc != game.Loc {
				game.Knfloc = LOC_NOWHERE
			}
			// Check some for hints, get input from user, increment
			// turn, and pre-process commands. Keep going until
			// pre-processing is done.
			for { // Preprocessing loop.
				// loop while EMPTY, RAW, TOKENIZED, GIVEN, PREPROCESSED
				// ignore loop (break) when PROCESS, EXECUTED
				if command.state >= PREPROCESSED {
					break
				}
				checkHints()
				/* Get command input from user */
				getCommandInput(&command)
				game.Turns++
				preprocessCommand(&command) // sets PREPROCESSED, if OK
				if settings.Logger {
					log.Println("\n**********  new preprocessCommand  **********")
					log.Printf(" At %s  : %s\n", LocationText(game.Loc), &command)
				}
			}
			/* check if game is Closed, and exit if it is */
			if closecheck() {
				return true
			}
			for { // loop until all words in command are procesed
				if command.state != PREPROCESSED {
					break
				}
				command.state = PROCESSING
				if command.word[0].id == WORD_NOT_FOUND {
					/* Gee, I don't understand. */
					speak(true, messages[DONT_KNOW], command.word[0].raw)
					clearCommand(&command)
					continue
				}
				// Give user hints of shortcuts
				if strings.ToLower(command.word[0].raw) == "west" {
					game.Iwest++
					if game.Iwest == 3 {
						speak(true, messages[W_IS_WEST])
					}
				}
				if strings.ToLower(command.word[0].raw) == "go" && command.word[1].id != VocabType(WORD_EMPTY) {
					game.Igo++
					if game.Igo == 3 {
						speak(true, messages[GO_UNNEEDED])
					}
				}
				switch command.word[0].typ {
				case MOTION:
					// fmt.Printf("playerMove - from %s  : %s\n", LocationText(game.Loc), &command)
					if settings.Logger {
						log.Printf("playerMove - At %s  : %s\n", LocationText(game.Loc), &command)
					}
					playerMove(int(command.word[0].id))
					// fmt.Printf("playerMoved - to %s  : %s\n", LocationText(game.Loc), &command)
					command.state = EXECUTED
					//					settings.CRT.Show(true, locations[game.Loc].pic)
					if settings.Logger {
						log.Printf("  playerMove. EXECUTED - Now At %s", LocationText(game.Loc))
					}
					continue
				case OBJECT:
					command.part = unknown
					command.obj = ObjectType(command.word[0].id)
				case ACTION:
					if command.word[1].typ == NUMERIC {
						command.part = transitive
					} else {
						command.part = intransitive
					}
					command.verb = VerbType(command.word[0].id)
				case NUMERIC:
					speak(true, messages[DONT_KNOW], command.word[0].raw)
					clearCommand(&command)
					continue
				case NO_WORD_TYPE:
					fallthrough
				default:
					panic("doCommand: VocabType > 1000 and not 1 to 4")
				}
				// OBJECT & ACTION make it to here
				if settings.Logger {
					log.Printf("action - At %s  : %s\n", LocationText(game.Loc), &command)
				}
				phase := action(command)
				if settings.Logger {
					log.Printf("  PHASE - %s", PhaseCodeText(phase))
				}
				switch phase {
				case GO_TERMINATE:
					command.state = EXECUTED
				case GO_MOVE:
					playerMove(NUL)
					command.state = EXECUTED
				case GO_WORD2:
					/* Get second word for analysis. */
					command.word[0] = command.word[1]
					command.word[1] = EmptyCommandWord
					command.state = PREPROCESSED // forces return to this loop.
				case GO_UNKNOWN:
					/*  Random intransitive verbs come here.  Clear obj just in case
					 *  (see attack()). */
					command.word[0].raw = strings.ToUpper(command.word[0].raw)
					speak(true, messages[DO_WHAT], command.word[0].raw)
					command.obj = NO_OBJECT
					/* object cleared; we need to go back to the preprocessing step */
					command.state = GIVEN
				case GO_CHECKHINT:
					command.state = GIVEN
				case GO_DWARFWAKE:
					//  Oh dear, he's disturbed the dwarves.
					speak(true, messages[DWARVES_AWAKEN])
					terminate(term_end)
				case GO_CLEAROBJ:
					clearCommand(&command)
				case GO_TOP:
				default:
					panic("doCommand: invalid phase code")
				}
			} /* while command has not been fully processed */
		} /* loop until command is GIVEN */
	} /* loop until command is EXECUTER */
	/* command completely executed; we return true. */
	return true
}

/*  Suspend.  Offer to save things in a file, but charging
 *  some points (so can't win by using Saved games to retry
 *  battles or to start over after learning ZZword).
 *  If ADVENT_NOSAVE is defined, do nothing instead. */
func suspend() PhaseCodeType {
	if !settings.AllowSave {
		speak(true, "This version does NOT allow SAVEing\n Play on.")
		return GO_CLEAROBJ
	}
	speak(true, messages[SUSPEND_WARNING])
	if !askYesNo(messages[THIS_ACCEPTABLE], messages[OK_MAN], messages[OK_MAN]) {
		return GO_CLEAROBJ
	}
	game.Saved = game.Saved + 5
	_ = game.save()
	return GO_TOP
}

/*  Resume.  Read a suspended game back from a file.
*  If ADVENT_NOSAVE is defined, do nothing instead.
*  !! Not necesssary. resumes when start from ~/adventure.json */
func resume() {
	if !settings.AllowSave {
		speak(true, "This version does NOT allow SAVEing\n Play on.")
	} else {
		speak(true, "You were automatically resumed\n Play on.")
		/*		ifElse(resumed, func() {
						speak(true, "You're already resumed.\nSee home/adventure.json\n Play on.")
					}, func() {
						speak(true, "You must first do a 'SAVE'. (creates 'home/adventure.json') ")
					})
				}
		*/
	}
}
