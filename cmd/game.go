package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

var game = Game{
	// dwarf locations. 6th is special "pirate"
	Dloc:  [7]LocationType{LOC_NOWHERE, LOC_KINGHALL, LOC_WESTBANK, LOC_Y2, LOC_ALIKE3, LOC_COMPLEX, LOC_MAZEEND12},
	Chloc: LOC_MAZEEND12, Chloc2: LOC_MAZEEND12, AbbNum: 5, Clock1: WARNTIME, Clock2: FLASHTIME,
	Limit: GAMELIMIT, Newloc: LOC_START, Loc: LOC_START, Foobar: WORD_EMPTY,
}

type Game struct {
	LcgX    int32
	AbbNum  int          // How often to print int descriptions
	Bonus   ScoreType    // What kind of finishing Bonus we are getting
	Chloc   LocationType // pirate chest location
	Chloc2  LocationType // pirate chest alternate location
	Clock1  TurnType     // # Turns from finding last treasure to close
	Clock2  TurnType     // # Turns from warning till blinding flash
	Clshnt  bool         // has player read the clue in the endgame?
	Closed  bool         // whether we're all the way Closed
	Closng  bool         // whether it's closing time yet
	Lmwarn  bool         // has player been warned about lamp going dim?
	Novice  bool         // asked for instructions at start-up?
	Panic   bool         // has player found out he's trapped?
	Wzdark  bool         // whether the Loc he's leaving was dark
	Blooded bool         // has player drunk of dragon's blood?
	Conds   uint32       // min value for cond[Loc] if Loc has any hints
	Detail  int          // level of Detail in description
	/*  Dflag controls the level of activation of dwarves:
	 *	0	No dwarf stuff yet (wait until reaches Hall Of Mists)
	 *	1	Reached Hall Of Mists, but hasn't met first dwarf
	 *	2	Met first dwarf, others start moving, no knives thrown yet
	 *	3	A knife has been thrown (first set always misses)
	 *	3+	Dwarves are mad (increases their accuracy) */
	Dflag  int
	Dkill  int                        // dwarves killed
	Dtotal int                        // total dwarves (including pirate) in Loc
	Foobar int                        // progress in saying "FEE FIE FOE FOO".
	Holdng int                        // number of objects being carried
	Igo    int                        // # uses of "go" instead of a direction
	Iwest  int                        // # times he's said "west" instead of "w"
	Knfloc LocationType               // knife location; 0 if none, -1 after caveat
	Limit  TurnType                   // lifetime of lamp
	Loc    LocationType               // where player is now
	Newloc LocationType               // where player is going
	Numdie TurnType                   // number of times killed so far
	Oldloc LocationType               // where player was
	Oldlc2 LocationType               // where player was two moves ago
	Oldobj ObjectType                 // last object player handled
	Saved  int                        // point penalty for saves
	Tally  int                        // count of treasures gained
	Thresh int                        // current threshold for endgame scoring tier
	Trnluz int                        // # points lost so far due to Turns used
	Turns  TurnType                   // counts commands given (ignores yes/no)
	ZZword [TOKLEN]byte               // randomly generated magic word from bird
	Abbrev [NLOCATIONS + 1]int        // has location been seen?
	Atloc  [NLOCATIONS + 1]ObjectType // head of object linked list per location
	Fixed  [NOBJECTS + 1]LocationType // Fixed location of object (if  not IS_FREE)
	Link   [NOBJECTS*2 + 1]ObjectType // object-list links
	Place  [NOBJECTS + 1]LocationType // location of object
	Hinted [NHINTS]bool               // Hinted[i] = true iff hint i has been used.
	Hintlc [NHINTS]int                // Hintlc[i] = how int at LOC with cond bit i
	Prop   [NOBJECTS + 1]int          // each object's state
	Dseen  [NDWARVES + 1]bool         // true if dwarf has seen him
	Dloc   [NDWARVES + 1]LocationType // location of dwarves, initially hard-wired in
	Odloc  [NDWARVES + 1]LocationType // prior Loc of each dwarf, initially garbage
}

func (g *Game) save() (err error) {
	d, err := os.UserHomeDir()
	log.Printf("** saving: cage %s, bird %s\n",
		LocationText(LocationType(game.Prop[CAGE])),
		LocationText(LocationType(game.Prop[BIRD])))
	if err != nil {
		return err
	}
	f := path.Join(d, "adventure.json")
	bytes, err := json.MarshalIndent(g, " ", "  ")
	if err == nil {
		err = ioutil.WriteFile(f, bytes, 0644)
	}
	return
}

func load() (err error) {
	d, err := os.UserHomeDir()
	if err != nil {
		return
	}
	f := path.Join(d, "adventure.json")
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &game)
	if err != nil {
		return
	}
	if game.Prop[BIRD] == BIRD_CAGED {
		objects[CAGE].pic = "cage_bird"
	}
	if game.Prop[LAMP] == LAMP_BRIGHT {
		objects[LAMP].pic = "lamp_on"
	}
	switch game.Prop[BEAR] {
	case SITTING_BEAR:
		objects[BEAR].pic = "bear_sitting"
	case CONTENTED_BEAR:
		objects[BEAR].pic = "bear_wandering"
	}
	return
}
