package cmd

import (
	display "adventure/images"
	"fmt"
	"fyne.io/fyne/v2"
)

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

type CommandState int
type CondType int
type DestType int
type LocationType int
type ObjectType int
type PhaseCodeType int
type ScoreType int
type SpeakType int
type SpeechPart int
type StringGroup []string
type TerminationType int
type TurnType int
type VerbType int
type VocabType int
type WordType int

var settings Settings

type Settings struct {
	App       fyne.App
	Window    fyne.Window
	TTY       *Console
	CRT       *display.Display
	AllowSave bool
	Logger    bool
}
type Action struct {
	words    StringGroup
	message  string
	noaction bool
}
type Class struct {
	threshold int
	message   string
}
type Command struct {
	part  SpeechPart
	word  [2]CommandWord
	verb  VerbType
	obj   ObjectType
	state CommandState
}

// "toString"
func (c *Command) String() string {
	return fmt.Sprintf("[%s] part %s, verb %s, object %s\n 0(%s) 1(%s)", CommandStateText(c.state),
		SpeechPartText(c.part), ActionTypeText(c.verb), ObjectText(c.obj), c.word[0], c.word[1])
}

type CommandWord struct {
	raw string // char raw[LINESIZE]
	id  VocabType
	typ WordType
}

// "toString"
func (cw CommandWord) String() string {
	return fmt.Sprintf("\"%s\",%s,id:%d", cw.raw, WordTypeText(cw.typ), cw.id)
}

var EmptyCommandWord = CommandWord{raw: "", id: WORD_EMPTY, typ: NO_WORD_TYPE}

type Descriptions struct {
	small string
	big   string // note: lines := strings.Split(big, "\\n")
}
type Hint struct {
	number   int
	turns    int
	penalty  int
	question string
	hint     string
}
type Location struct {
	descriptions Descriptions
	sound        VocabType
	loud         bool
	pic          string
}

func LocationText(l LocationType) string {
	switch {
	case l < 0:
		return fmt.Sprintf("(Location.%d)unknown", l)
	case l == 0:
		return "(Location.0)DESTROYED"
	}
	return fmt.Sprintf("(Location.%d)%s", l, locations[l].descriptions.small)
}

type Motion struct {
	words StringGroup
}
type Obituary struct {
	query    string
	response string
}
type Object struct {
	words        StringGroup
	inventory    string
	plac         LocationType // 0 = LOC_NOWHERE
	fixd         LocationType // 0 = FREE, -1 = FIXED
	treasure     bool
	descriptions []string // note: lines := strings.Split(X, "\\n")
	sounds       StringGroup
	pic          string
	texts        StringGroup
	states       StringGroup
}

func ObjectText(o ObjectType) string {
	if o > NOBJECTS {
		o -= NOBJECTS
	}
	words := objects[o].words
	if words != nil && len(words) > 0 {
		return fmt.Sprintf("(Object.%d)%s", o, words[0])
	}
	return fmt.Sprintf("(Object.%d)unknown", o)
}

type Threshold struct {
	threshold TurnType
	loss      int
	message   string
}
type Travel struct {
	motion    int
	condtype  CondType
	condarg1  int
	condarg2  int
	desttype  DestType
	destval   LocationType
	nodwarves bool
	stop      bool
}

// "toString"
func (t *Travel) String() string {
	return fmt.Sprintf("motion %d, type %s c1 %d c2 %d, destType %s dest %s, stop %t",
		t.motion, CondTypeText(t.condtype), t.condarg1, t.condarg2,
		DestTypeText(t.desttype), LocationText(t.destval), t.stop)
}

const SILENT = -1 // no sound
//const ignore = "LXGZI"

func (t Travel) Terminate() bool {
	return t.motion == 1
}

const ( // ScoreType (Bonus)
	bonus_none = iota
	bonus_splatter
	bonus_defeat
	bonus_victory
)
const ( // CondType
	cond_goto = iota
	cond_pct
	cond_carry
	cond_with
	cond_not
)

func CondTypeText(c CondType) string {
	switch c {
	case cond_goto:
		return "cond_goto"
	case cond_pct:
		return "cond_pct"
	case cond_carry:
		return "cond_carry"
	case cond_with:
		return "cond_with"
	case cond_not:
		return "cond_not"
	}
	return fmt.Sprintf("unknown(%d)", c)
}

const ( // DestType
	dest_goto = iota
	dest_special
	dest_speak
)

func DestTypeText(d DestType) string {
	switch d {
	case dest_goto:
		return "dest_goto"
	case dest_special:
		return "dest_special"
	case dest_speak:
		return "dest_speak"
	}
	return fmt.Sprintf("unknown(%d)", d)
}

const ( // SpeechPart
	unknown = iota
	intransitive
	transitive
)

func SpeechPartText(t SpeechPart) string {
	switch t {
	case intransitive:
		return "intransitive"
	case transitive:
		return "transitive"
	}
	return "empty"
}

const ( // SpeakType
	touch = iota
	look
	hear
	study
	change
)

func SpeechTypeText(t SpeakType) string {
	switch t {
	case touch:
		return "touch"
	case look:
		return "look"
	case hear:
		return "hear"
	case study:
		return "study"
	case change:
		return "change"
	}
	return fmt.Sprintf("%d -> SILENT", t)
}

const ( // TerminationType for score:  (endgame, quitgame, scoregame)
	term_end = iota
	term_quit
	term_score
)

func TerminationTypeText(t TerminationType) string {
	switch t {
	case term_end:
		return "End Game"
	case term_quit:
		return "Quit Game"
	case term_score:
		return "Score Game"
	}
	return fmt.Sprintf("Unknown %d", t)
}

const ( // WordType
	NO_WORD_TYPE = iota
	MOTION
	OBJECT
	ACTION
	NUMERIC
)

func WordTypeText(t WordType) string {
	switch t {
	case MOTION:
		return "Motion"
	case OBJECT:
		return "Object"
	case ACTION:
		return "Action"
	case NUMERIC:
		return "Numeric"
	}
	return fmt.Sprintf("Unknown %d", t)
}

/* closure potentially can cause side effects within its scope. */
func ifElse(cond bool, yes, no func()) {
	if cond {
		yes()
	} else {
		no()
	}
}

const ( // CommandState
	EMPTY = iota
	RAW
	TOKENIZED
	GIVEN
	PREPROCESSED
	PROCESSING
	EXECUTED
)

func CommandStateText(t CommandState) string {
	switch t {
	case EMPTY:
		return "Empty"
	case RAW:
		return "Raw"
	case TOKENIZED:
		return "Tokenized"
	case GIVEN:
		return "Given"
	case PREPROCESSED:
		return "PreProcessed"
	case PROCESSING:
		return "Processing"
	case EXECUTED:
		return "Executed"
	}
	return fmt.Sprintf("Unknown %d", t)
}

/* Phase codes for action returns.
 * These were at one time FORTRAN line numbers.
 * The values don't matter, but perturb their order at your peril.
 */
const ( // PhaseCodeType
	GO_TERMINATE = iota
	GO_MOVE
	GO_TOP
	GO_CLEAROBJ
	GO_CHECKHINT
	GO_WORD2
	GO_UNKNOWN
	GO_DWARFWAKE
)

func PhaseCodeText(t PhaseCodeType) string {
	switch t {
	case GO_TERMINATE:
		return "Terminate"
	case GO_MOVE:
		return "Move"
	case GO_TOP:
		return "Top"
	case GO_CLEAROBJ:
		return "Clear Object"
	case GO_WORD2:
		return "2nd Word"
	case GO_UNKNOWN:
		return "Unkown"
	case GO_DWARFWAKE:
		return "Dwarf Wake"
	}
	return fmt.Sprintf("Unknown %d", t)
}

// NLOCATIONS  et all... Game array sizes
const NLOCATIONS = 184
const NOBJECTS = 69
const NHINTS = 10
const NCLASSES = 10
const NDEATHS = 3
const NTHRESHOLDS = 4
const NMOTIONS = 76
const NACTIONS = 58

// other constants

const NDWARVES = 6       // number of dwarves
const GAMELIMIT = 330    // base Limit of Turns
const NOVICELIMIT = 1000 // Limit of Turns for Novice
const WARNTIME = 30      // late game starts at game.Limit-this
const PANICTIME = 15     // time left after closing
const FLASHTIME = 50     // Turns from first warning till blinding flash
const WORD_EMPTY = 0     // "Word empty" flag value for the vocab hash functions
const CARRIED = -1       // Player is toting it
// LCG_A /* LCG PRNG parameters tested against
const LCG_A int32 = 1093
const LCG_C int32 = 221587
const LCG_M int32 = 1048576
const TOKLEN = 5 // # sigificant characters in a token
// STATE_NOTFOUND Special object-state values - integers > 0 are object-specific
const STATE_NOTFOUND = -1 // "Not found" state of treasures
const STATE_FOUND = 0     // After discovered, before messed with
const STATE_IN_CAVITY = 1 // State value common to all gemstones
const WORD_NOT_FOUND = -1 // "Word not found" flag value for the vocab hash functions.
const PIRATE = NDWARVES   // must be NDWARVES-1 when zero-origin
const DALTLC = LOC_NUGGET // alternate dwarf location
// FIXED Special Fixed object-state values - integers > 0 are location
const FIXED = -1
const FREE = 0
const BATTERYLIFE = 2500 // turn Limit increment from batteries
const INTRANSITIVE = -1  // illegal object number
const INVLIMIT = 7       // inventory Limit (# of objects)

/* State definitions - used as "skip" value in pspeak */

// LAMP_DARK /* States for LAMP */
const LAMP_DARK = 0
const LAMP_BRIGHT = 1

// GRATE_CLOSED /* States for GRATE */
const GRATE_CLOSED = 0
const GRATE_OPEN = 1

// STEPS_DOWN /* States for STEPS */
const STEPS_DOWN = 0
const STEPS_UP = 1

// BIRD_UNCAGED /* States for BIRD */
const BIRD_UNCAGED = 0
const BIRD_CAGED = 1
const BIRD_FOREST_UNCAGED = 2
const BIRD_ENDSTATE = 5

// DOOR_RUSTED /* States for DOOR */
const DOOR_RUSTED = 0
const DOOR_UNRUSTED = 1

// SNAKE_BLOCKS /* States for SNAKE */
const SNAKE_BLOCKS = 0
const SNAKE_CHASED = 1

// UNBRIDGED /* States for FISSURE */
const UNBRIDGED = 0
const BRIDGED = 1

// WATER_BOTTLE /* States for BOTTLE */
const WATER_BOTTLE = 0
const EMPTY_BOTTLE = 1
const OIL_BOTTLE = 2

// MIRROR_UNBROKEN /* States for MIRROR */
const MIRROR_UNBROKEN = 0
const MIRROR_BROKEN = 1

// PLANT_THIRSTY /* States for PLANT */
const PLANT_THIRSTY = 0
const PLANT_BELLOWING = 1
const PLANT_GROWN = 2

// AXE_HERE /* States for AXE */
const AXE_HERE = 0
const AXE_LOST = 1

// DRAGON_BARS /* States for DRAGON */
const DRAGON_BARS = 0
const DRAGON_DEAD = 1
const DRAGON_BLOODLESS = 2

// TROLL_BRIDGE /* States for CHASM */
const TROLL_BRIDGE = 0
const BRIDGE_WRECKED = 1

// TROLL_UNPAID /* States for TROLL */
const TROLL_UNPAID = 0
const TROLL_PAIDONCE = 1
const TROLL_GONE = 2

// UNTAMED_BEAR /* States for BEAR */
const UNTAMED_BEAR = 0
const SITTING_BEAR = 1
const CONTENTED_BEAR = 2
const BEAR_DEAD = 3

// VEND_BLOCKS /* States for VEND */
const VEND_BLOCKS = 0
const VEND_UNBLOCKS = 1

// FRESH_BATTERIES /* States for BATTERY */
const FRESH_BATTERIES = 0
const DEAD_BATTERIES = 1

// URN_EMPTY /* States for URN */
const URN_EMPTY = 0
const URN_DARK = 1
const URN_LIT = 2

// CAVITY_FULL /* States for CAVITY */
const CAVITY_FULL = 0
const CAVITY_EMPTY = 1

// WATERS_UNPARTED /* States for RESER */
const WATERS_UNPARTED = 0
const WATERS_PARTED = 1

// INGAME_SIGN /* States for SIGN */
const INGAME_SIGN = 0
const ENDGAME_SIGN = 1

// EGGS_HERE /* States for EGGS */
const EGGS_HERE = 0
const EGGS_VANISHED = 1
const EGGS_DONE = 2

// VASE_WHOLE /* States for VASE */
const VASE_WHOLE = 0
const VASE_DROPPED = 1
const VASE_BROKEN = 2

// RUG_FLOOR /* States for RUG */
const RUG_FLOOR = 0
const RUG_DRAGON = 1
const RUG_HOVER = 2

// CHAIN_HEAP /* States for CHAIN */
const CHAIN_HEAP = 0
const CHAINING_BEAR = 1
const CHAIN_FIXED = 2

// AMBER_IN_URN /* States for AMBER */
const AMBER_IN_URN = 0
const AMBER_IN_ROCK = 1
