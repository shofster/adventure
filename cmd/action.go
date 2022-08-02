package cmd

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

var actions = [...]Action{
	/* 0 */ {words: nil, noaction: false,
		message: ""},
	/* 1 */ {words: []string{"g", "carry", "take", "keep", "catch", "steal", "captu", "get", "tote", "snarf"}, noaction: false,
		message: "You are already carrying it!"},
	/* 2 */ {words: []string{"drop", "relea", "free", "disca", "dump"}, noaction: false,
		message: "You aren't carrying it!"},
	/* 3 */ {words: []string{"say", "chant", "sing", "utter", "mumbl"}, noaction: false,
		message: "NO_MESSAGE"},
	/* 4 */ {words: []string{"unloc", "open"}, noaction: false,
		message: "I don't know how to lock or unlock such a thing."},
	/* 5 */ {words: []string{"z", "nothi"}, noaction: false,
		message: "NO_MESSAGE"},
	/* 6 */ {words: []string{"lock", "close"}, noaction: false,
		message: "I don't know how to lock or unlock such a thing."},
	/* 7 */ {words: []string{"light", "on"}, noaction: false,
		message: "I'm afraid I don't understand."},
	/* 8 */ {words: []string{"extin", "off"}, noaction: false,
		message: "I'm afraid I don't understand."},
	/* 9 */ {words: []string{"wave", "shake", "swing"}, noaction: false,
		message: "Nothing happens."},
	/* 10 */ {words: []string{"calm", "placa", "tame"}, noaction: false,
		message: "I'm game.  Would you care to explain how?"},
	/* 11 */ {words: []string{"walk", "run", "trave", "go", "proce", "conti", "explo", "follo", "turn"}, noaction: false,
		message: "Where?"},
	/* 12 */ {words: []string{"attac", "kill", "fight", "hit", "strik", "slay"}, noaction: false,
		message: "Don't be ridiculous!"},
	/* 13 */ {words: []string{"pour"}, noaction: false,
		message: "You aren't carrying it!"},
	/* 14 */ {words: []string{"eat", "devou"}, noaction: false,
		message: "Don't be ridiculous!"},
	/* 15 */ {words: []string{"drink"}, noaction: false,
		message: "You have taken a drink from the stream.  The water tastes strongly of\nminerals, but is not unpleasant.  It is extremely cold."},
	/* 16 */ {words: []string{"rub"}, noaction: false,
		message: "Rubbing the electric lamp is not particularly rewarding.  Anyway,\nnothing exciting happens."},
	/* 17 */ {words: []string{"throw", "toss"}, noaction: false,
		message: "You aren't carrying it!"},
	/* 18 */ {words: []string{"quit"}, noaction: false,
		message: "Huh?"},
	/* 19 */ {words: []string{"find", "where"}, noaction: false,
		message: "I can only tell you what you see as you move about and manipulate\nthings.  I cannot tell you where remote things are."},
	/* 20 */ {words: []string{"i", "inven"}, noaction: false,
		message: "I can only tell you what you see as you move about and manipulate\nthings.  I cannot tell you where remote things are."},
	/* 21 */ {words: []string{"feed"}, noaction: false,
		message: "There is nothing here to eat."},
	/* 22 */ {words: []string{"fill"}, noaction: false,
		message: "You can't fill that."},
	/* 23 */ {words: []string{"blast", "deton", "ignit", "blowu"}, noaction: false,
		message: "Blasting requires dynamite."},
	/* 24 */ {words: []string{"score"}, noaction: false,
		message: "Huh?"},
	/* 25 */ {words: []string{"fee"}, noaction: false,
		message: "I don't know how."},
	/* 26 */ {words: []string{"fie"}, noaction: false,
		message: "I don't know how."},
	/* 27 */ {words: []string{"foe"}, noaction: false,
		message: "I don't know how."},
	/* 28 */ {words: []string{"foo"}, noaction: false,
		message: "I don't know how."},
	/* 29 */ {words: []string{"fum"}, noaction: false,
		message: "I don't know how."},
	/* 30 */ {words: []string{"brief"}, noaction: false,
		message: "On what?"},
	/* 31 */ {words: []string{"read", "perus"}, noaction: false,
		message: "I'm afraid I don't understand."},
	/* 32 */ {words: []string{"break", "shatt", "smash"}, noaction: false,
		message: "It is beyond your power to do that."},
	/* 33 */ {words: []string{"wake", "distu"}, noaction: false,
		message: "Don't be ridiculous!"},
	/* 34 */ {words: []string{"suspe", "pause", "save"}, noaction: false,
		message: "Huh?"},
	/* 35 */ {words: []string{"resum", "resta"}, noaction: false,
		message: "Huh?"},
	/* 36 */ {words: []string{"fly"}, noaction: false,
		message: "I'm game.  Would you care to explain how?"},
	/* 37 */ {words: []string{"liste"}, noaction: false,
		message: "I'm afraid I don't understand."},
	/* 38 NOTE: escaped the \ */ {words: []string{"z\\'zzz"}, noaction: false,
		message: "Nothing happens."},
	/* 39 */ {words: []string{"seed"}, noaction: false,
		message: "Seed set to %d"},
	/* 40 */ {words: []string{"waste"}, noaction: false,
		message: "Game Limit is now %d"},
	/* 41 */ {words: nil, noaction: false,
		message: "Huh?"},
	/* 42 */ {words: []string{"thank"}, noaction: true,
		message: "You're quite welcome."},
	/* 43 */ {words: []string{"sesam", "opens", "abra", "abrac", "shaza", "hocus", "pocus"}, noaction: true,
		message: "Good try, but that is an old worn-out magic word."},
	/* 44 */ {words: []string{"help", "?"}, noaction: true,
		message: "I know of places, actions, and things.  Most of my vocabulary\ndescribes places and is used to move you there.  To move, try words\nlike forest, building, downstream, enter, east, west, north, south,\nup, or down.  I know about a few special objects, like a black rod\nhidden in the cave.  These objects can be manipulated using some of\nthe action words that I know.  Usually you will need to give both the\nobject and action words (in either order), but sometimes I can infer\nthe object from the verb alone.  Some objects also imply verbs; in\nparticular, \"inventory\" implies \"take inventory\", which causes me to\ngive you a list of what you're carrying.  Some objects have unexpected\neffects; the effects are not always desirable!  Usually people having\ntrouble moving just need to try a few more words.  Usually people\ntrying unsuccessfully to manipulate an object are attempting something\nbeyond their (or my!) capabilities and should try a completely\ndifferent tack.  One point often confusing to beginners is that, when\nthere are several ways to go in a certain direction (e.g., if there\nare several holes in a wall), choosing that direction in effect\nchooses one of the ways at random; often, though, by specifying the\nPlace you want to reach you can guarantee choosing the right path.\nAlso, to speed the game you can sometimes move long distances with a\nsingle word.  For example, \"building\" usually gets you to the building\nfrom anywhere above ground except when lost in the forest.  Also, note\nthat cave passages and forest paths turn a lot, so leaving one Place\nheading north doesn't guarantee entering the next from the south.\nHowever (another important point), except when you've used a \"long\ndistance\" word such as \"building\", there is always a way to go back\nwhere you just came from unless I warn you to the contrary, even\nthough the direction that takes you back might not be the reverse of\nwhat got you here.  Good luck, and have fun!"},
	/* 45 */ {words: []string{"no"}, noaction: true,
		message: "OK"},
	/* 46 */ {words: []string{"tree", "trees"}, noaction: true,
		message: "The trees of the forest are large hardwood oak and maple, with an\noccasional grove of pine or spruce.  There is quite a bit of under-\ngrowth, largely birch and ash saplings plus nondescript bushes of\nvarious sorts.  This time of year visibility is quite restricted by\nall the leaves, but travel is quite easy if you detour around the\nspruce and berry bushes."},
	/* 47 */ {words: []string{"dig", "excav"}, noaction: true,
		message: "Digging without a shovel is quite impractical.  Even with a shovel\nprogress is unlikely."},
	/* 48 */ {words: []string{"lost"}, noaction: true,
		message: "I'm as confused as you are."},
	/* 49 */ {words: []string{"mist"}, noaction: true,
		message: "Mist is a white vapor, usually water, seen from time to time in\ncaverns.  It can be found anywhere but is frequently a sign of a deep\npit leading down to water.'"},
	/* 50 */ {words: []string{"fuck", "shit", "crap", "damn"}, noaction: true,
		message: "Watch it!"},
	/* 51 */ {words: []string{"stop"}, noaction: true,
		message: "I don't know the word \"stop\".  Use \"quit\" if you want to give up."},
	/* 52 */ {words: []string{"info", "infor"}, noaction: true,
		message: "For a summary of the most recent states to the game, say \"news\".\nIf you want to end your adventure early, say \"quit\".  To suspend your\nadventure such that you can continue later, say \"suspend\" (or \"pause\"\nor \"save\").  To see how well you're doing, say \"score\".  To get full\ncredit for a treasure, you must have left it safely in the building,\nthough you get partial credit just for locating it.  You lose points\nfor getting killed, or for quitting, though the former costs you more.\nThere are also points based on how much (if any) of the cave you've\nmanaged to explore; in particular, there is a large Bonus just for\ngetting in (to distinguish the beginners from the rest of the pack),\nand there are other ways to determine whether you've been through some\nof the more harrowing sections.  If you think you've found all the\ntreasures, just keep exploring for a while.  If nothing interesting\nhappens, you haven't found them all yet.  If something interesting\n*DOES* happen (incidentally, there *ARE* ways to hasten things along),\nit means you're getting a Bonus and have an opportunity to garner many\nmore points in the Master's section.  I may occasionally offer hints\nif you seem to be having trouble.  If I do, I'll warn you in advance\nhow much it will affect your score to accept the hints.  Finally, to\nsave time, you may specify \"brief\", which tells me never to repeat the\nfull description of a Place unless you explicitly ask me to."},
	/* 53 */ {words: []string{"swim"}, noaction: true,
		message: "I don't know how."},
	/* 54 */ {words: []string{"wizar"}, noaction: true,
		message: "Wizards are not to be disturbed by such as you."},
	/* 55 */ {words: []string{"yes"}, noaction: true,
		message: "Guess again."},
	/* 56 */ {words: []string{"news"}, noaction: true,
		message: "Open Adventure is an author-approved open-source release of\nVersion 2.5 with, as yet, no gameplay states.\nVersion 2.5 was essentially the same as Version II; the cave and the\nhazards therein are unchanged, and top score is still 430 points.\nThere are a few more hints, especially for some of the more obscure\npuzzles.  There are a few minor bugfixes and cosmetic states.  You\ncan now save a game and resume it at once (formerly you had to wait a\nwhile first), but it now costs you a few points each time you save the\ngame.  Saved games are now stored in much smaller files than before."},
	/* 57 */ {words: []string{"versi"}, noaction: true,
		message: "There is a puff of orange smoke; within it, fiery runes spell out:\n\n\tOpen Adventure %V - http://www.catb.org/esr/open-adventure/"},
	/* 58 */ {words: nil, noaction: false,
		message: ""},
}

const ( // Action / Verb Type
	ACT_NULL = iota
	CARRY
	DROP
	SAY
	UNLOCK
	NOTHING
	LOCK
	LIGHT
	EXTINGUISH
	WAVE
	TAME
	GO
	ATTACK
	POUR
	EAT
	DRINK
	RUB
	THROW
	QUIT
	FIND
	INVENTORY
	FEED
	FILL
	BLAST
	SCORE
	FEE
	FIE
	FOE
	FOO
	FUM
	BRIEF
	READ
	BREAK
	WAKE
	SAVE
	RESUME
	FLY
	LISTEN
	PART
	SEED
	WASTE
	/* UNUSED ??
	ACT_UNKNOWN
	THANKYOU
	INVALIDMAGIC
	HELP
	False
	TREE
	DIG
	LOST
	MIST
	FBOMB
	STOP
	INFO
	SWIM
	WIZARD
	YES
	NEWS
	ACT_VERSION
	*/
)

func ActionTypeText(t VerbType) string {
	switch t {
	case ACT_NULL:
		return "none"
	case CARRY:
		return "carry"
	case DROP:
		return "drop"
	case SAY:
		return "say"
	case UNLOCK:
		return "unlock"
	case NOTHING:
		return "nothing"
	case LOCK:
		return "lock"
	case LIGHT:
		return "light"
	case EXTINGUISH:
		return "extinguish"
	case WAVE:
		return "wave"
	case TAME:
		return "tame"
	case GO:
		return "go"
	case ATTACK:
		return "attack"
	case POUR:
		return "pour"
	case EAT:
		return "eat"
	case DRINK:
		return "drink"
	case RUB:
		return "rub"
	case THROW:
		return "throw"
	case QUIT:
		return "quit"
	case FIND:
		return "find"
	case INVENTORY:
		return "inventory"
	case FEED:
		return "feed"
	case FILL:
		return "fill"
	case BLAST:
		return "blast"
	case SCORE:
		return "score"
	case FEE:
		return "fee"
	case FIE:
		return "fie"
	case FOE:
		return "foe"
	case FOO:
		return "foo"
	case FUM:
		return "fum"
	case BRIEF:
		return "brief"
	case READ:
		return "read"
	case BREAK:
		return "break"
	case WAKE:
		return "wake"
	case SAVE:
		return "save"
	case RESUME:
		return "resume"
	case FLY:
		return "fly"
	case LISTEN:
		return "listen"
	case PART:
		return "part"
	case SEED:
		return "seed"
	case WASTE:
		return "waste"
	}
	return "Unknown Action"
}
