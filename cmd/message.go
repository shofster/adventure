package cmd

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */
func plural(n int) string {
	if n != 1 {
		return "s"
	}
	return ""
}

var messages = [...]string{
	"NULL",
	"Somewhere nearby is Colossal Cave, where others have found fortunes in\ntreasure and gold, though it is rumored that some who enter are never\nseen again.  Magic is said to work in the cave.  I will be your eyes\nand hands.  Direct me with commands of 1 or 2 words.  I should warn\nyou that I look at only the first five letters of each word, so you'll\nhave to enter \"northeast\" as \"ne\" to distinguish it from \"north\".\nYou can type \"help\" for some general hints.  For information on how\nto end your adventure, scoring, etc., type \"info\".\n\tThis program was originally developed by Willie Crowther.  Most of the\nfeatures of the current program were added by Don Woods.\nConverted to GO & Visuals added by Bob Shofner.",
	"A little dwarf with a big knife blocks your way.",
	"A little dwarf just walked around a corner, saw you, threw a little\naxe at you which missed, cursed, and ran away.",
	"There are %d threatening little dwarves in the room with you.",
	"There is a threatening little dwarf in the room with you!",
	"One sharp nasty knife is thrown at you!",
	"A hollow voice says \"PLUGH\".",
	"It gets you!",
	"It misses!",
	"I am unsure how you are facing.  Use compass points or nearby objects.",
	"I don't know in from out here.  Use compass points or name something\nin the general direction you want to go.",
	"I don't know how to apply that word here.",
	"I'm game.  Would you care to explain how?",
	"Sorry, but I am not allowed to give more Detail.  I will repeat the\nlong description of your location.",
	"It is now pitch dark.  If you proceed you will likely fall into a pit.",
	"If you prefer, simply type w rather than west.",
	"Do you really want to quit now?",
	"You fell into a pit and broke every bone in your body!",
	"You are already carrying it!",
	"You can't be serious!",
	"The bird seemed unafraid at first, but as you approach it becomes\ndisturbed and you cannot catch it.",
	"You can catch the bird, but you cannot carry it.",
	"There is nothing here with a lock!",
	"You aren't carrying it!",
	"The little bird attacks the green snake, and in an astounding flurry\ndrives the snake away.",
	"You have no keys!",
	"It has no lock.",
	"I don't know how to lock or unlock such a thing.",
	"It was already locked.",
	"It was already unlocked.",
	"There is no way to get past the bear to unlock the chain, which is\nprobably just as well.",
	"Nothing happens.",
	"Where?",
	"There is nothing here to attack.",
	"The little bird is now dead.  Its body disappears.",
	"Attacking the snake both doesn't work and is very dangerous.",
	"You killed a little dwarf.",
	"You attack a little dwarf, but he dodges out of the way.",
	"With what?  Your bare hands?",
	"There is no way to go that direction.",
	"Please stick to 1 or 2 word commands.",
	"OK",
	"You can't unlock the keys.",
	"You have crawled around in some little holes and wound up back in the\nmain passage.",
	"I don't know where the cave is, but hereabouts no stream can run on\nthe surface for long.  I would try the stream.",
	"I need more detailed instructions to do that.",
	"I can only tell you what you see as you move about and manipulate\nthings.  I cannot tell you where remote things are.",
	"The ogre snarls and shoves you back.",
	"Huh?",
	"\nWelcome to Adventure!!  Would you like instructions?",
	"Blasting requires dynamite.",
	"Your feet are now wet.",
	"I think I just lost my appetite.",
	"Thank you, it was delicious!",
	"Peculiar.  Nothing unexpected happens.",
	"Your bottle is empty and the ground is wet.",
	"You can't pour that.",
	"Which way?",
	"Sorry, but I no longer seem to remember how it was you got here.",
	"You can't carry anything more.  You'll have to drop something first.",
	"You can't go through a locked steel grate!",
	"I believe what you want is right here with you.",
	"You don't fit through a two-inch slit!",
	"I respectfully suggest you go across the bridge instead of jumping.",
	"There is no way across the fissure.",
	"You're not carrying anything.",
	"You are currently holding the following:",
	"It's not hungry (it's merely pinin' for the fjords).  Besides, you\nhave no bird seed.",
	"The snake has now devoured your bird.",
	"There's nothing here it wants to eat (except perhaps you).",
	"You fool, dwarves eat only coal!  Now you've made him *REALLY* mad!!",
	"You have nothing in which to carry it.",
	"Your bottle is already full.",
	"There is nothing here with which to fill the bottle.",
	"Don't be ridiculous!",
	"The door is extremely rusty and refuses to open.",
	"The plant indignantly shakes the oil off its leaves and asks, \"Water?\"",
	"The plant has exceptionally deep roots and cannot be pulled free.",
	"The dwarves' knives vanish as they strike the walls of the cave.",
	"Something you're carrying won't fit through the tunnel with you.\nYou'd best take inventory and drop something.",
	"You can't fit this five-foot clam through that little passage!",
	"You can't fit this five-foot oyster through that little passage!",
	"I advise you to put down the clam before opening it.  >STRAIN!<",
	"I advise you to put down the oyster before opening it.  >WRENCH!<",
	"You don't have anything strong enough to open the clam.",
	"You don't have anything strong enough to open the oyster.",
	"A glistening pearl falls out of the clam and rolls away.  Goodness,\nthis must really be an oyster.  (I never was very good at identifying\nbivalves.)  Whatever it is, it has now snapped shut again.",
	"The oyster creaks open, revealing nothing but oyster inside.  It\npromptly snaps shut again.",
	"You have crawled around in some little holes and found your way\nblocked by a recent cave-in.  You are now back in the main passage.",
	"There are faint rustling noises from the darkness behind you.",
	"Out from the shadows behind you pounces a bearded pirate!  \"Har, har,\"\nhe chortles, \"I'll just take all this booty and hide it away with me\nchest deep in the maze!\"  He snatches your treasure and vanishes into\nthe gloom.",
	"A sepulchral voice reverberating through the cave, says, \"Cave closing\nsoon.  All adventurers exit immediately through main office.\"",
	"A mysterious recorded voice groans into life and announces:\n   \"This exit is Closed.  Please leave via main office.\"",
	"It looks as though you're dead.  Well, seeing as how it's so close to\nclosing time anyway, I think we'll just call it a day.",
	"The sepulchral voice intones, \"The cave is now Closed.\"  As the echoes\nfade, there is a blinding flash of light (and a small puff of orange\nsmoke). . . .    As your eyes refocus, you look around and find...",
	"There is a loud explosion, and a twenty-foot hole appears in the far\nwall, burying the dwarves in the rubble.  You march through the hole\nand find yourself in the main office, where a cheering band of\nfriendly elves carry the conquering adventurer off into the sunset.",
	"There is a loud explosion, and a twenty-foot hole appears in the far\nwall, burying the snakes in the rubble.  A river of molten lava pours\nin through the hole, destroying everything in its path, including you!",
	"There is a loud explosion, and you are suddenly splashed across the\nwalls of the room.",
	"The resulting ruckus has awakened the dwarves.  There are now several\nthreatening little dwarves in the room with you!  Most of them throw\nknives at you!  All of them get you!",
	"Oh, leave the poor unhappy bird alone.",
	"I daresay whatever you want is around here somewhere.",
	"You can't get there from here.",
	"You are being followed by a very large, tame bear.",
	"Now let's see you do it without suspending in mid-Adventure.",
	"There is nothing here with which to fill it.",
	"The sudden change in temperature has delicately shattered the vase.",
	"It is beyond your power to do that.",
	"I don't know how.",
	"It is too far up for you to reach.",
	"You killed a little dwarf.  The body vanishes in a cloud of greasy\nblack smoke.",
	"The shell is very strong and is impervious to attack.",
	"What's the matter, can't you read?  Now you'd best start over.",
	"The axe bounces harmlessly off the dragon's thick scales.",
	"The dragon looks rather nasty.  You'd best not try to get by.",
	"The little bird attacks the green dragon, and in an astounding flurry\ngets burnt to a cinder.  The ashes blow away.",
	"Okay, from now on I'll only describe a Place in full the first time\nyou come to it.  To get the full description, say \"look\".",
	"Trolls are close relatives with the rocks and have skin as tough as\nthat of a rhinoceros.  The troll fends off your blows effortlessly.",
	"The troll deftly catches the axe, examines it carefully, and tosses it\nback, declaring, \"Good workmanship, but it's not valuable enough.\"",
	"The troll catches your treasure and scurries away out of sight.",
	"The troll refuses to let you cross.",
	"There is no longer any way across the chasm.",
	"With what?  Your bare hands?  Against *HIS* bear hands??",
	"The bear is confused; he only wants to be your friend.",
	"For crying out loud, the poor thing is already dead!",
	"The bear is still chained to the wall.",
	"The chain is still locked.",
	"The chain is now unlocked.",
	"The chain is now locked.",
	"There is nothing here to which the chain can be locked.",
	"Do you want the hint?",
	"Gluttony is not one of the troll's vices.  Avarice, however, is.",
	"Your lamp is getting dim.  You'd best start wrapping this up, unless\nyou can find some fresh batteries.  I seem to recall there's a vending\nmachine in the maze.  Bring some coins with you.",
	"Your lamp has run out of power.",
	"Please answer the question.",
	"There are faint rustling noises from the darkness behind you.  As you\nturn toward them, the beam of your lamp falls across a bearded pirate.\nHe is carrying a large chest.  \"Shiver me timbers!\" he cries, \"I've\nbeen spotted!  I'd best hie meself off to the maze to hide me chest!\"\nWith that, he vanishes into the gloom.",
	"Your lamp is getting dim.  You'd best go back for those batteries.",
	"Your lamp is getting dim.  I'm taking the liberty of replacing the\nbatteries.",
	"Your lamp is getting dim, and you're out of spare batteries.  You'd\nbest start wrapping this up.",
	"You sift your fingers through the dust, but succeed only in\nobliterating the cryptic message.",
	"Hmmm, this looks like a clue, which means it'll cost you 10 points to\nread it.  Should I go ahead and read it anyway?",
	"It says, \"There is a way out of this Place.  Do you need any more\ninformation to escape?  Sorry, but this initial hint is all you get.\"",
	"I'm afraid I don't understand.",
	"Your hand passes through it as though it weren't there.",
	"You prod the nearest dwarf, who wakes up grumpily, takes one look at\nyou, curses, and grabs for his axe.",
	"Is this acceptable?",
	"The ogre doesn't appear to be hungry.",
	"The ogre, who despite his bulk is quite agile, easily dodges your\nattack.  He seems almost amused by your puny effort.",
	"The ogre, distracted by your rush, is struck by the knife.  With a\nblood-curdling yell he Turns and bounds after the dwarves, who flee\nin Panic.  You are left alone in the room.",
	"The ogre, distracted by your rush, is struck by the knife.  With a\nblood-curdling yell he Turns and bounds after the dwarf, who flees\nin Panic.  You are left alone in the room.",
	"The bird flies about agitatedly for a moment.",
	"The bird flies agitatedly about the cage.",
	"The bird flies about agitatedly for a moment, then disappears through\nthe crack.  It reappears shortly, carrying in its beak a jade\nnecklace, which it drops at your feet.",
	"You empty the bottle into the urn, which promptly ejects the water\nwith uncanny accuracy, squirting you directly between the eyes.",
	"Your bottle is now empty and the urn is full of oil.",
	"The urn is already full of oil.",
	"There's no way to get the oil out of the urn.",
	"The urn is far too firmly embedded for your puny strength to budge it.",
	"As you rub the urn, there is a flash of light and a genie appears.\nHis aspect is stern as he advises: \"One who wouldst traffic in\nprecious stones must first learn to recognize the signals thereof.\"\nHe wrests the urn from the stone, leaving a small cavity.  Turning to\nface you again, he fixes you with a steely eye and intones: \"Caution!\"\nGenie and urn vanish in a cloud of amber smoke.  The smoke condenses\nto form a rare amber gemstone, resting in the cavity in the rock.",
	"I suppose you collect doughnut holes, too?",
	"The gem fits easily into the cavity.",
	"The Persian rug stiffens and rises a foot or so off the ground.",
	"The Persian rug draped over your shoulder seems to wriggle for a\nmoment, but then subsides.",
	"The Persian rug settles gently to the ground.",
	"The rug hovers stubbornly where it is.",
	"The rug does not appear inclined to cooperate.",
	"If you mean to use the Persian rug, it does not appear inclined to\ncooperate.",
	"Though you flap your arms furiously, it is to no avail.",
	"You board the Persian rug, which promptly whisks you across the chasm.\nYou have time for a fleeting glimpse of a two thousand foot drop to a\nmighty river; then you find yourself on the other side.",
	"The rug ferries you back across the chasm.",
	"All is silent.",
	"The stream is gurgling placidly.",
	"The wind whistles coldly past your ears.",
	"The stream splashes loudly into the pool.",
	"You are unable to make anything of the splashing noise.",
	"You can hear the murmuring of the beanstalks and the snoring of the\ndwarves.",
	"A loud hissing emanates from the snake pit.",
	"The air is filled with a dull rumbling sound.",
	"The roar is quite loud here.",
	"The roaring is so loud that it drowns out all other sound.",
	"The bird eyes you suspiciously and flutters away.  A moment later you\nfeel something wet land on your head, but upon looking up you can see\nno sign of the culprit.",
	"There are only a few drops--not enough to carry.",
	"(Uh, y'know, that wasn't very bright.)",
	"It's a pity you took so long about it.",
	"Upstream or downstream?",
	"NULL",
	"The waters are crashing loudly against the shore.",
	"%d of them throw knives at you!",
	"%d of them get you!",
	"One of them gets you!",
	"None of them hits you!",
	"Sorry, I don't know the word '%s'.",
	"What do you want to do with the '%s'?",
	"I see no '%s' here.",
	"'%s' what?",
	"Okay, \"%s\".",
	"You have garnered %d out of a possible %d points, using %d turn%s.",
	"I can save your Adventure for you so that you can resume later, but\nit will cost you 5 points.",
	"I am prepared to give you a hint, but it will cost you %d point%s.",
	"You scored %d out of a possible %d, using %d turn%s.",
	"To achieve the next higher rating, you need %d more point%s.",
	"To achieve the next higher rating would be a neat trick!\nCongratulations!!",
	"You just went off my scale!!",
	"To resume your Adventure, start a new game and it will RESUME.",
	"To resume an earlier Adventure, you must abandon the current one.",
	"I'm sorry, but that Adventure was begun using Version %d.%d of the\nsave file format, and this program uses Version %d.%d.  You must find an instance\nusing that other version in order to resume that Adventure.",
	"Sorry, but the path twisted and turned so much that I can't figure\nout which way to go to get back.",
	"You don't have to say \"go\" every time; just specify a direction or, if\nit's nearby, name the Place to which you wish to move.",
	"This command requires a numeric argument.",
}

const (
	NO_MESSAGE = iota
	CAVE_NEARBY
	DWARF_BLOCK
	DWARF_RAN
	DWARF_PACK
	DWARF_SINGLE
	KNIFE_THROWN
	SAYS_PLUGH
	GETS_YOU
	MISSES_YOU
	UNSURE_FACING
	NO_INOUT_HERE
	CANT_APPLY
	AM_GAME
	NO_MORE_DETAIL
	PITCH_DARK
	W_IS_WEST
	REALLY_QUIT
	PIT_FALL
	ALREADY_CARRYING
	YOU_JOKING
	BIRD_EVADES
	CANNOT_CARRY
	NOTHING_LOCKED
	ARENT_CARRYING
	BIRD_ATTACKS
	NO_KEYS
	NO_LOCK
	NOT_LOCKABLE
	ALREADY_LOCKED
	ALREADY_UNLOCKED
	BEAR_BLOCKS
	NOTHING_HAPPENS
	WHERE_QUERY
	NO_TARGET
	BIRD_DEAD
	SNAKE_WARNING
	KILLED_DWARF
	DWARF_DODGES
	BARE_HANDS_QUERY
	BAD_DIRECTION
	TWO_WORDS
	OK_MAN
	CANNOT_UNLOCK
	FUTILE_CRAWL
	FOLLOW_STREAM
	NEED_DETAIL
	NEARBY
	OGRE_SNARL
	HUH_MAN
	WELCOME_YOU
	REQUIRES_DYNAMITE
	FEET_WET
	LOST_APPETITE
	THANKS_DELICIOUS
	PECULIAR_NOTHING
	GROUND_WET
	CANT_POUR
	WHICH_WAY
	FORGOT_PATH
	CARRY_LIMIT
	GRATE_NOWAY
	YOU_HAVEIT
	DONT_FIT
	CROSS_BRIDGE
	NO_CROSS
	NO_CARRY
	NOW_HOLDING
	BIRD_PINING
	BIRD_DEVOURED
	NOTHING_EDIBLE
	REALLY_MAD
	NO_CONTAINER
	BOTTLE_FULL
	NO_LIQUID
	RIDICULOUS_ATTEMPT
	RUSTY_DOOR
	SHAKING_LEAVES
	DEEP_ROOTS
	KNIVES_VANISH
	MUST_DROP
	CLAM_BLOCKER
	OYSTER_BLOCKER
	DROP_CLAM
	DROP_OYSTER
	CLAM_OPENER
	OYSTER_OPENER
	PEARL_FALLS
	OYSTER_OPENS
	WAY_BLOCKED
	PIRATE_RUSTLES
	PIRATE_POUNCES
	CAVE_CLOSING
	EXIT_CLOSED
	DEATH_CLOSING
	CAVE_CLOSED
	VICTORY_MESSAGE
	DEFEAT_MESSAGE
	SPLATTER_MESSAGE
	DWARVES_AWAKEN
	UNHAPPY_BIRD
	NEEDED_NEARBY
	NOT_CONNECTED
	TAME_BEAR
	WITHOUT_SUSPENDS
	FILL_INVALID
	SHATTER_VASE
	BEYOND_POWER
	NOT_KNOWHOW
	TOO_FAR
	DWARF_SMOKE
	SHELL_IMPERVIOUS
	START_OVER
	DRAGON_SCALES
	NASTY_DRAGON
	BIRD_BURNT
	BRIEF_CONFIRM
	ROCKY_TROLL
	TROLL_RETURNS
	TROLL_SATISFIED
	TROLL_BLOCKS
	BRIDGE_GONE
	BEAR_HANDS
	BEAR_CONFUSED
	ALREADY_DEAD
	BEAR_CHAINED
	STILL_LOCKED
	CHAIN_UNLOCKED
	CHAIN_LOCKED
	NO_LOCKSITE
	WANT_HINT
	TROLL_VICES
	LAMP_DIM
	LAMP_OUT
	PLEASE_ANSWER
	PIRATE_SPOTTED
	GET_BATTERIES
	REPLACE_BATTERIES
	MISSING_BATTERIES
	REMOVE_MESSAGE
	CLUE_QUERY
	WAYOUT_CLUE
	DONT_UNDERSTAND
	HAND_PASSTHROUGH
	PROD_DWARF
	THIS_ACCEPTABLE
	OGRE_FULL
	OGRE_DODGE
	OGRE_PANIC1
	OGRE_PANIC2
	FREE_FLY
	CAGE_FLY
	NECKLACE_FLY
	WATER_URN
	OIL_URN
	FULL_URN
	URN_NOPOUR
	URN_NOBUDGE
	URN_GENIES
	DOUGHNUT_HOLES
	GEM_FITS
	RUG_RISES
	RUG_WIGGLES
	RUG_SETTLES
	RUG_HOVERS
	RUG_NOTHING1
	RUG_NOTHING2
	FLAP_ARMS
	RUG_GOES
	RUG_RETURNS
	ALL_SILENT
	STREAM_GURGLES
	WIND_WHISTLES
	STREAM_SPLASHES
	NO_MEANING
	MURMURING_SNORING
	SNAKES_HISSING
	DULL_RUMBLING
	LOUD_ROAR
	TOTAL_ROAR
	BIRD_CRAP
	FEW_DROPS
	NOT_BRIGHT
	TOOK_LONG
	UPSTREAM_DOWNSTREAM
	FOREST_QUERY
	WATERS_CRASHING
	THROWN_KNIVES
	MULTIPLE_HITS
	ONE_HIT
	NONE_HIT
	DONT_KNOW
	WHAT_DO
	NO_SEE
	DO_WHAT
	OKEY_DOKEY
	GARNERED_POINTS
	SUSPEND_WARNING
	HINT_COST
	TOTAL_SCORE
	NEXT_HIGHER
	NO_HIGHER
	OFF_SCALE
	RESUME_HELP
	RESUME_ABANDON
	VERSION_SKEW
	TWIST_TURN
	GO_UNNEEDED
	NUMERIC_REQUIRED
)

// message text row # - 16 + 228 is the const
