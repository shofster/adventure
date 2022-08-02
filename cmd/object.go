package cmd

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */
var objects = [...]Object{
	/* 0: NO_OBJECT */ {words: nil, inventory: "", plac: LOC_NOWHERE, fixd: LOC_NOWHERE, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil,
		states:       nil},
	/* 1: KEYS */ {words: []string{"keys", "key"}, inventory: "Set of keys", plac: LOC_BUILDING, fixd: 0, treasure: false,
		descriptions: []string{
			"There are some keys on the ground here."},
		sounds: nil, pic: "keys",
		texts:  nil,
		states: nil},
	/* 2: LAMP */ {words: []string{"lamp", "lante"}, inventory: "Brass lantern", plac: LOC_BUILDING, fixd: 0, treasure: false,
		descriptions: []string{
			"There is a shiny brass lamp nearby.",
			"There is a lamp shining nearby."},
		sounds: nil, pic: "lamp_off",
		texts: nil,
		states: []string{
			"Your lamp is now off.",
			"Your lamp is now on."}},
	/* 3: GRATE */ {words: []string{"grate"}, inventory: "*grate", plac: LOC_GRATE, fixd: LOC_BELOWGRATE, treasure: false,
		descriptions: []string{
			"The grate is locked.",
			"The grate is open."},
		sounds: nil,
		texts:  nil,
		states: []string{
			"The grate is now locked.",
			"The grate is now unlocked."}},
	/* 4: CAGE */ {words: []string{"cage"}, inventory: "Wicker cage", plac: LOC_COBBLE, fixd: 0, treasure: false,
		descriptions: []string{
			"There is a small wicker cage discarded nearby."},
		sounds: nil, pic: "cage",
		texts:  nil,
		states: nil},
	/* 5: ROD */ {words: []string{"rod"}, inventory: "Black rod", plac: LOC_DEBRIS, fixd: 0, treasure: false,
		descriptions: []string{
			"A three foot black rod with a rusty star on an end lies nearby."},
		sounds: nil, pic: "rod_star",
		texts:  nil,
		states: []string{""}},
	/* 6: ROD2 */ {words: []string{"rod"}, inventory: "Black rod", plac: LOC_NOWHERE, fixd: 0, treasure: false,
		descriptions: []string{
			"A three foot black rod with a rusty mark on an end lies nearby."},
		sounds: nil, pic: "rod_mark",
		texts:  nil,
		states: []string{""}},
	/* 7: STEPS */ {words: []string{"steps"}, inventory: "*steps", plac: LOC_PITTOP, fixd: LOC_MISTHALL, treasure: false,
		descriptions: []string{
			"Rough stone steps lead down the pit.",
			"Rough stone stepsf lead up the dome."},
		sounds: nil, pic: "steps",
		texts:  nil,
		states: nil},
	/* 8: BIRD */ {words: []string{"bird"}, inventory: "Little bird in cage", plac: LOC_BIRDCHAMBER, fixd: 0, treasure: false,
		descriptions: []string{
			"A cheerful little bird is sitting here singing.",
			"There is a little bird in the cage.",
			"A cheerful little bird is sitting here singing."},
		sounds: []string{
			"The bird's singing is quite melodious.",
			"The bird does not seem inclined to sing while in the cage.",
			"It almost seems as though the bird is trying to tell you something.",
			"To your surprise, you can understand the bird''s chirping; it is\nsinging about the joys of its forest home.",
			"The bird does not seem inclined to sing while in the cage.",
			"The bird is singing to you in gratitude for your having returned it to\nits home.  In return, it informs you of a magic word which it thinks\nyou may find useful somewhere near the Hall of Mists.  The magic word\nchanges frequently, but for now the bird believes it is '%s'.  You\nthank the bird for this information, and it flies off into the forest."},
		pic: "bird", texts: nil,
		states: nil},
	/* 9: DOOR */ {words: []string{"door"}, inventory: "*rusty door", plac: LOC_IMMENSE, fixd: -1, treasure: false,
		descriptions: []string{
			"The way north is barred by a massive, rusty, iron door.",
			"The way north leads through a massive, rusty, iron door."},
		sounds: nil, pic: "door",
		texts: nil,
		states: []string{
			"The hinges are quite thoroughly rusted now and won't budge.",
			"The oil has freed up the hinges so that the door will now move,\nalthough it requires some effort."}},
	/* 10: PILLOW */ {words: []string{"pillo", "velve"}, inventory: "Velvet pillow", plac: LOC_SOFTROOM, fixd: 0, treasure: false,
		descriptions: []string{
			"A small velvet pillow lies on the floor."},
		sounds: nil, pic: "pillow",
		texts:  nil,
		states: nil},
	/* 11: SNAKE */ {words: []string{"snake"}, inventory: "*snake", plac: LOC_KINGHALL, fixd: -1, treasure: false,
		descriptions: []string{
			"A huge green fierce snake bars the way!",
			""},
		sounds: []string{
			"The snake is hissing venomously.",
			""},
		pic: "snake", texts: nil,
		states: nil},
	/* 12: FISSURE */ {words: []string{"fissu"}, inventory: "*fissure", plac: LOC_EASTBANK, fixd: LOC_WESTBANK, treasure: false,
		descriptions: []string{
			"",
			"A crystal bridge spans the fissure."},
		sounds: nil, pic: "",
		texts: nil,
		states: []string{
			"The crystal bridge has vanished!",
			"A crystal bridge now spans the fissure."}},
	/* 13: OBJ_13 */ {words: []string{"table"}, inventory: "*stone tablet", plac: LOC_DARKROOM, fixd: -1, treasure: false,
		descriptions: []string{
			"A massive stone tablet embedded in the wall reads:\n'Congratulations on bringing light into the dark-room!'"},
		sounds: nil,
		texts: []string{
			"'Congratulations on bringing light into the dark-room!'"},
		states: nil},
	/* 14: CLAM */ {words: []string{"clam"}, inventory: "Giant clam  >GRUNT!<", plac: LOC_SHELLROOM, fixd: 0, treasure: false,
		descriptions: []string{
			"There is an enormous clam here with its shell tightly Closed."},
		sounds: []string{
			"The clam is as tight-mouthed as a, er, clam."},
		pic: "clam", texts: nil,
		states: nil},
	/* 15: OYSTER */ {words: []string{"oyste"}, inventory: "Giant oyster  >GROAN!<", plac: LOC_NOWHERE, fixd: 0, treasure: false,
		descriptions: []string{
			"There is an enormous oyster here with its shell tightly Closed.",
			"Interesting.  There seems to be something written on the underside of\nthe oyster."},
		sounds: []string{
			"Even though it's an oyster, the critter's as tight-mouthed as a clam.",
			"It says the same thing it did before.  Hm, maybe it's a pun?"},
		pic: "oyster", texts: nil,
		states: nil},
	/* 16: MAGAZINE */ {words: []string{"magaz", "issue", "spelu", "\"spel"}, inventory: "\"Spelunker Today\"", plac: LOC_ANTEROOM, fixd: 0, treasure: false,
		descriptions: []string{
			"There are a few recent issues of 'Spelunker Today' magazine here."},
		sounds: nil,
		texts: []string{
			"I'm afraid the magazine is written in dwarvish.  But pencilled on one\ncover you see, 'Please leave the magazines at the construction site.'"},
		states: nil},
	/* 17: DWARF */ {words: []string{"dwarf", "dwarv"}, inventory: "", plac: LOC_NOWHERE, fixd: -1, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil, pic: "dwarf",
		states: []string{""}},
	/* 18: KNIFE */ {words: []string{"knife", "knive"}, inventory: "", plac: LOC_NOWHERE, fixd: 0, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil, pic: "knife",
		states: []string{""}},
	/* 19: FOOD */ {words: []string{"food", "ratio"}, inventory: "Tasty food", plac: LOC_BUILDING, fixd: 0, treasure: false,
		descriptions: []string{
			"There is food here."},
		sounds: nil, pic: "food",
		texts:  nil,
		states: nil},
	/* 20: BOTTLE */ {words: []string{"bottl", "jar"}, inventory: "Small bottle", plac: LOC_BUILDING, fixd: 0, treasure: false,
		descriptions: []string{
			"There is a bottle of water here.",
			"There is an empty bottle here.",
			"There is a bottle of oil here."},
		sounds: nil, pic: "bottle",
		texts: nil,
		states: []string{
			"Your bottle is now full of water.",
			"The bottle of water is now empty.",
			"Your bottle is now full of oil."}},
	/* 21: WATER */ {words: []string{"water", "h2o"}, inventory: "Water in the bottle", plac: LOC_NOWHERE, fixd: 0, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil,
		states:       nil},
	/* 22: OIL */ {words: []string{"oil"}, inventory: "Oil in the bottle", plac: LOC_NOWHERE, fixd: 0, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil,
		states:       nil},
	/* 23: MIRROR */ {words: []string{"mirro"}, inventory: "*mirror", plac: LOC_MIRRORCANYON, fixd: -1, treasure: false,
		descriptions: []string{
			"",
			""},
		sounds: nil,
		texts:  nil,
		states: []string{
			"",
			"You strike the mirror a resounding blow, whereupon it shatters into a\nmyriad tiny fragments."}},
	/* 24: PLANT */ {words: []string{"plant", "beans"}, inventory: "*plant", plac: LOC_WESTPIT, fixd: -1, treasure: false,
		descriptions: []string{
			"There is a tiny little plant in the pit, murmuring 'water, water, ...'",
			"There is a 12-foot-tall beanstalk stretching up out of the pit,\nbellowing 'WATER!! WATER!!'",
			"There is a gigantic beanstalk stretching all the way up to the hole."},
		sounds: []string{
			"The plant continues to ask plaintively for water.",
			"The plant continues to demand water.",
			"The plant now maintains a contented silence."},
		texts: nil, pic: "sapling",
		states: []string{
			"You've over-watered the plant!  It's shriveling up!  And now . . .",
			"The plant spurts into furious growth for a few seconds.",
			"The plant grows explosively, almost filling the bottom of the pit."}},
	/* 25: PLANT2 */ {words: []string{"plant"}, inventory: "*phony plant", plac: LOC_WESTEND, fixd: LOC_EASTEND, treasure: false,
		descriptions: []string{
			"",
			"The top of a 12-foot-tall beanstalk is poking out of the west pit.",
			"There is a huge beanstalk growing out of the west pit up to the hole."},
		sounds: nil,
		texts:  nil, pic: "sapling",
		states: nil},
	/* 26: OBJ_26 */ {words: []string{"stala"}, inventory: "*stalactite", plac: LOC_TOPSTALACTITE, fixd: -1, treasure: false,
		descriptions: []string{
			""},
		sounds: nil,
		texts:  nil,
		states: nil},
	/* 27: OBJ_27 */ {words: []string{"shado", "figur", "windo"}, inventory: "*shadowy figure and/or window", plac: LOC_WINDOW1, fixd: LOC_WINDOW2, treasure: false,
		descriptions: []string{
			"A shadowy figure seems to be trying to attract your attention."},
		sounds: nil, pic: "gnome",
		texts:  nil,
		states: nil},
	/* 28: AXE */ {words: []string{"axe"}, inventory: "Dwarf's axe", plac: LOC_NOWHERE, fixd: 0, treasure: false,
		descriptions: []string{
			"There is a little axe here.",
			"There is a little axe lying beside the bear."},
		sounds: nil, pic: "axe",
		texts: nil,
		states: []string{
			"",
			"The axe misses and lands near the bear where you can't get at it."}},
	/* 29: OBJ_29 */ {words: []string{"drawi"}, inventory: "*cave drawings", plac: LOC_ORIENTAL, fixd: -1, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil,
		states:       nil},
	/* 30: OBJ_30 */ {words: []string{"pirat", "genie", "djinn"}, inventory: "*pirate/genie", plac: LOC_NOWHERE, fixd: -1, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil,
		states:       nil},
	/* 31: DRAGON */ {words: []string{"drago"}, inventory: "*dragon", plac: LOC_SECRET4, fixd: LOC_SECRET6, treasure: false,
		descriptions: []string{
			"A huge green fierce dragon bars the way!",
			"The blood-specked body of a huge green dead dragon lies to one side.",
			"The body of a huge green dead dragon is lying off to one side."},
		sounds: []string{
			"The dragon's ominous hissing does not bode well for you.",
			"The dragon is, not surprisingly, silent.",
			"The dragon is, not surprisingly, silent."},
		pic: "dragon", texts: nil,
		states: []string{
			"",
			"Congratulations!  You have just vanquished a dragon with your bare\nhands!  (Unbelievable, isn't it?)",
			"Your head buzzes strangely for a moment."}},
	/* 32: CHASM */ {words: []string{"chasm"}, inventory: "*chasm", plac: LOC_SWCHASM, fixd: LOC_NECHASM, treasure: false,
		descriptions: []string{
			"A rickety wooden bridge extends across the chasm, vanishing into the\nmist.  A notice posted on the bridge reads, 'Stop! Pay troll!'",
			"The wreckage of a bridge (and a dead bear) can be seen at the bottom\nof the chasm."},
		sounds: nil,
		texts:  nil,
		states: []string{
			"",
			"Just as you reach the other side, the bridge buckles beneath the\nweight of the bear, which was still following you around.  You\nscrabble desperately for support, but as the bridge collapses you\nstumble back and fall into the chasm."}},
	/* 33: TROLL */ {words: []string{"troll"}, inventory: "*troll", plac: LOC_SWCHASM, fixd: LOC_NECHASM, treasure: false,
		descriptions: []string{
			"A burly troll stands by the bridge and insists you throw him a\ntreasure before you may cross.",
			"The troll steps out from beneath the bridge and blocks your way.",
			""},
		sounds: []string{
			"The troll sounds quite adamant in his demand for a treasure.",
			"The troll sounds quite adamant in his demand for a treasure.",
			""},
		pic: "troll", texts: nil,
		states: []string{
			"",
			"",
			"The bear lumbers toward the troll, who lets out a startled shriek and\nscurries away.  The bear soon gives up the pursuit and wanders back."}},
	/* 34: TROLL2 */ {words: []string{"troll"}, inventory: "*phony troll", plac: LOC_NOWHERE, fixd: LOC_NOWHERE, treasure: false,
		descriptions: []string{
			"The troll is nowhere to be seen."},
		sounds: nil,
		texts:  nil,
		states: nil},
	/* 35: BEAR */ {words: []string{"bear"}, inventory: "", plac: LOC_BARRENROOM, fixd: -1, treasure: false,
		descriptions: []string{
			"There is a ferocious cave bear eyeing you from the far end of the room!",
			"There is a gentle cave bear sitting placidly in one corner.",
			"There is a contented-looking bear wandering about nearby.",
			""},
		sounds: nil, pic: "bear",
		texts: nil,
		states: []string{
			"",
			"The bear eagerly wolfs down your food, after which he seems to calm\ndown considerably and even becomes rather friendly.",
			"",
			""}},
	/* 36: MESSAG */ {words: []string{"messa"}, inventory: "*message in second maze", plac: LOC_NOWHERE, fixd: -1, treasure: false,
		descriptions: []string{
			"There is a message scrawled in the dust in a flowery script, reading:\n'This is not the maze where the pirate leaves his treasure chest.'"},
		sounds: nil,
		texts: []string{
			"'This is not the maze where the pirate leaves his treasure chest.'"},
		states: nil},
	/* 37: VOLCANO */ {words: []string{"volca", "geyse"}, inventory: "*volcano and/or geyser", plac: LOC_BREATHTAKING, fixd: -1, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil,
		states:       nil},
	/* 38: VEND */ {words: []string{"machi", "vendi"}, inventory: "*vending machine", plac: LOC_DEADEND13, fixd: -1, treasure: false,
		descriptions: []string{
			"There is a massive and somewhat battered vending machine here.  The\ninstructions on it read: 'Drop coins here to receive fresh batteries.'",
			"There is a massive vending machine here, swung back to reveal a\nsouthward passage."},
		sounds: nil,
		texts: []string{
			"'Drop coins here to receive fresh batteries.'",
			"'Drop coins here to receive fresh batteries.'"},
		states: []string{
			"The vending machine swings back to block the passage.",
			"As you strike the vending machine, it pivots backward along with a\nsection of wall, revealing a dark passage leading south."}},
	/* 39: BATTERY */ {words: []string{"batte"}, inventory: "Batteries", plac: LOC_NOWHERE, fixd: 0, treasure: false,
		descriptions: []string{
			"There are fresh batteries here.",
			"Some worn-out batteries have been discarded nearby."},
		sounds: nil, pic: "battery",
		texts:  nil,
		states: nil},
	/* 40: OBJ_40 */ {words: []string{"carpe", "moss"}, inventory: "*carpet and/or moss and/or curtains", plac: LOC_SOFTROOM, fixd: -1, treasure: false,
		descriptions: nil,
		sounds:       nil,
		texts:        nil,
		states:       nil},
	/* 41: OGRE */ {words: []string{"ogre"}, inventory: "*ogre", plac: LOC_LARGE, fixd: -1, treasure: false,
		descriptions: []string{
			"A formidable ogre bars the northern exit."},
		sounds: []string{
			"The ogre is apparently the strong, silent type."},
		texts: nil, pic: "ogre",
		states: nil},
	/* 42: URN */ {words: []string{"urn"}, inventory: "*urn", plac: LOC_CLIFF, fixd: -1, treasure: false,
		descriptions: []string{
			"A small urn is embedded in the rock.",
			"A small urn full of oil is embedded in the rock.",
			"A small oil flame extrudes from an urn embedded in the rock."},
		sounds: nil,
		texts:  nil, pic: "urn",
		states: []string{
			"The urn is empty and will not light.",
			"The urn is now dark.",
			"The urn is now lit."}},
	/* 43: CAVITY */ {words: []string{"cavit"}, inventory: "*cavity", plac: LOC_NOWHERE, fixd: -1, treasure: false,
		descriptions: []string{
			"",
			"There is a small urn-shaped cavity in the rock."},
		sounds: nil,
		texts:  nil,
		states: nil},
	/* 44: BLOOD */ {words: []string{"blood"}, inventory: "*blood", plac: LOC_NOWHERE, fixd: -1, treasure: false,
		descriptions: []string{
			""},
		sounds: nil, pic: "blood",
		texts:  nil,
		states: nil},
	/* 45: RESER */ {words: []string{"reser"}, inventory: "*reservoir", plac: LOC_RESERVOIR, fixd: LOC_RESNORTH, treasure: false,
		descriptions: []string{
			"",
			"The waters have parted to form a narrow path across the reservoir."},
		sounds: nil, pic: "reservoir",
		texts: nil,
		states: []string{
			"The waters crash together again.",
			"The waters have parted to form a narrow path across the reservoir."}},
	/* 46: RABBITFOOT */ {words: []string{"appen", "lepor"}, inventory: "Leporine appendage", plac: LOC_FOREST22, fixd: 0, treasure: false,
		descriptions: []string{
			"Your keen eye spots a severed leporine appendage lying on the ground."},
		sounds: nil,
		texts:  nil,
		states: nil},
	/* 47: OBJ_47 */ {words: []string{"mud"}, inventory: "*mud", plac: LOC_DEBRIS, fixd: -1, treasure: false,
		descriptions: []string{
			""},
		sounds: nil, pic: "word_xyzzy",
		texts: []string{
			"'MAGIC WORD XYZZY'"},
		states: nil},
	/* 48: OBJ_48 */ {words: []string{"note"}, inventory: "*note", plac: LOC_NUGGET, fixd: -1, treasure: false,
		descriptions: []string{
			""},
		sounds: nil,
		texts: []string{
			"'You won't get it up the steps'"},
		states: nil},
	/* 49: SIGN */ {words: []string{"sign"}, inventory: "*sign", plac: LOC_ANTEROOM, fixd: -1, treasure: false,
		descriptions: []string{
			"",
			""},
		sounds: nil,
		texts: []string{
			"Cave under construction beyond this point.\n           Proceed at own risk.\n       [Witt Construction Company]",
			"'Treasure Vault.  Keys in main office.'"},
		states: nil},
	/* 50: NUGGET */ {words: []string{"gold", "nugge"}, inventory: "Large gold nugget", plac: LOC_NUGGET, fixd: 0, treasure: true,
		descriptions: []string{
			"There is a large sparkling nugget of gold here!"},
		sounds: nil, pic: "gold",
		texts:  nil,
		states: nil},
	/* 51: OBJ_51 */ {words: []string{"diamo"}, inventory: "Several diamonds", plac: LOC_WESTBANK, fixd: 0, treasure: true,
		descriptions: []string{
			"There are diamonds here!"},
		sounds: nil, pic: "diamonds",
		texts:  nil,
		states: nil},
	/* 52: OBJ_52 */ {words: []string{"silve", "bars"}, inventory: "Bars of silver", plac: LOC_FLOORHOLE, fixd: 0, treasure: true,
		descriptions: []string{
			"There are bars of silver here!"},
		sounds: nil, pic: "silver_bars",
		texts:  nil,
		states: nil},
	/* 53: OBJ_53 */ {words: []string{"jewel"}, inventory: "Precious jewelry", plac: LOC_SOUTHSIDE, fixd: 0, treasure: true,
		descriptions: []string{
			"There is precious jewelry here!"},
		sounds: nil, pic: "jewels",
		texts:  nil,
		states: nil},
	/* 54: COINS */ {words: []string{"coins"}, inventory: "Rare coins", plac: LOC_WESTSIDE, fixd: 0, treasure: true,
		descriptions: []string{
			"There are many coins here!"},
		sounds: nil, pic: "coins",
		texts:  nil,
		states: nil},
	/* 55: CHEST */ {words: []string{"chest", "box", "treas"}, inventory: "Treasure chest", plac: LOC_NOWHERE, fixd: 0, treasure: true,
		descriptions: []string{
			"The pirate's treasure chest is here!"},
		sounds: nil, pic: "chest",
		texts:  nil,
		states: nil},
	/* 56: EGGS */ {words: []string{"eggs", "egg", "nest"}, inventory: "Golden eggs", plac: LOC_GIANTROOM, fixd: 0, treasure: true,
		descriptions: []string{
			"There is a large nest here, full of golden eggs!",
			"The nest of golden eggs has vanished!",
			"Done!"},
		sounds: nil, pic: "eggs",
		texts:  nil,
		states: nil},
	/* 57: TRIDENT */ {words: []string{"tride"}, inventory: "Jeweled trident", plac: LOC_WATERFALL, fixd: 0, treasure: true,
		descriptions: []string{
			"There is a jewel-encrusted trident here!"},
		sounds: nil, pic: "trident",
		texts:  nil,
		states: nil},
	/* 58: VASE */ {words: []string{"vase", "ming", "shard", "potte"}, inventory: "Ming vase", plac: LOC_ORIENTAL, fixd: 0, treasure: true,
		descriptions: []string{
			"There is a delicate, precious, ming vase here!",
			"The floor is littered with worthless shards of pottery.",
			"The floor is littered with worthless shards of pottery."},
		sounds: nil, pic: "vase",
		texts: nil,
		states: []string{
			"The vase is now resting, delicately, on a velvet pillow.",
			"The ming vase drops with a delicate crash.",
			"You have taken the vase and hurled it delicately to the ground."}},
	/* 59: EMERALD */ {words: []string{"emera"}, inventory: "Egg-sized emerald", plac: LOC_PLOVER, fixd: 0, treasure: true,
		descriptions: []string{
			"There is an emerald here the size of a plover's egg!",
			"There is an emerald resting in a small cavity in the rock!"},
		sounds: nil, pic: "emerald",
		texts:  nil,
		states: nil},
	/* 60: PYRAMID */ {words: []string{"plati", "pyram"}, inventory: "Platinum pyramid", plac: LOC_DARKROOM, fixd: 0, treasure: true,
		descriptions: []string{
			"There is a platinum pyramid here, 8 inches on a side!"},
		sounds: nil, pic: "pyramid",
		texts:  nil,
		states: nil},
	/* 61: PEARL */ {words: []string{"pearl"}, inventory: "Glistening pearl", plac: LOC_NOWHERE, fixd: 0, treasure: true,
		descriptions: []string{
			"Off to one side lies a glistening pearl!"},
		sounds: nil, pic: "pearl",
		texts:  nil,
		states: nil},
	/* 62: RUG */ {words: []string{"rug", "persi"}, inventory: "Persian rug", plac: LOC_SECRET4, fixd: LOC_SECRET6, treasure: true,
		descriptions: []string{
			"There is a Persian rug spread out on the floor!",
			"The dragon is sprawled out on a Persian rug!!",
			"There is a Persian rug here, hovering in mid-air!"},
		sounds: nil, pic: "rug",
		texts:  nil,
		states: nil},
	/* 63: OBJ_63 */ {words: []string{"spice"}, inventory: "Rare spices", plac: LOC_BOULDERS2, fixd: 0, treasure: true,
		descriptions: []string{
			"There are rare spices here!"},
		sounds: nil, pic: "spices",
		texts:  nil,
		states: nil},
	/* 64: CHAIN */ {words: []string{"chain"}, inventory: "Golden chain", plac: LOC_BARRENROOM, fixd: -1, treasure: true,
		descriptions: []string{
			"There is a golden chain lying in a heap on the floor!",
			"The bear is locked to the wall with a golden chain!",
			"There is a golden chain locked to the wall!"},
		sounds: nil, pic: "chain",
		texts:  nil,
		states: nil},
	/* 65: RUBY */ {words: []string{"ruby"}, inventory: "Giant ruby", plac: LOC_STOREROOM, fixd: 0, treasure: true,
		descriptions: []string{
			"There is an enormous ruby here!",
			"There is a ruby resting in a small cavity in the rock!"},
		sounds: nil, pic: "ruby",
		texts:  nil,
		states: nil},
	/* 66: COND_JADE */ {words: []string{"jade", "neckl"}, inventory: "Jade necklace", plac: LOC_NOWHERE, fixd: 0, treasure: true,
		descriptions: []string{
			"A precious jade necklace has been dropped here!"},
		sounds: nil, pic: "jade",
		texts:  nil,
		states: nil},
	/* 67: AMBER */ {words: []string{"amber", "gemst"}, inventory: "Amber gemstone", plac: LOC_NOWHERE, fixd: 0, treasure: true,
		descriptions: []string{
			"There is a rare amber gemstone here!",
			"There is an amber gemstone resting in a small cavity in the rock!"},
		sounds: nil, pic: "amber",
		texts:  nil,
		states: nil},
	/* 68: SAPPH */ {words: []string{"sapph"}, inventory: "Star sapphire", plac: LOC_LEDGE, fixd: 0, treasure: true,
		descriptions: []string{
			"A brilliant blue star sapphire is here!",
			"There is a star sapphire resting in a small cavity in the rock!"},
		sounds: nil, pic: "sapphire",
		texts:  nil,
		states: nil},
	/* 69: OBJ_69 */ {words: []string{"ebony", "statu"}, inventory: "Ebony statuette", plac: LOC_REACHDEAD, fixd: 0, treasure: true,
		descriptions: []string{
			"There is a richly-carved ebony statuette here!"},
		sounds: nil, pic: "statue",
		texts:  nil,
		states: nil},
}

const (
	NO_OBJECT = iota
	KEYS
	LAMP
	GRATE
	CAGE
	ROD
	ROD2
	STEPS
	BIRD
	DOOR
	PILLOW
	SNAKE
	FISSURE
	OBJ_13
	CLAM
	OYSTER
	MAGAZINE
	DWARF
	KNIFE
	FOOD
	BOTTLE
	WATER
	OIL
	MIRROR
	PLANT
	PLANT2
	OBJ_26
	OBJ_27
	AXE
	OBJ_29
	OBJ_30
	DRAGON
	CHASM
	TROLL
	TROLL2
	BEAR
	MESSAG
	VOLCANO
	VEND
	BATTERY
	OBJ_40
	OGRE
	URN
	CAVITY
	BLOOD
	RESER
	RABBITFOOT
	OBJ_47
	OBJ_48
	SIGN
	NUGGET
	OBJ_51
	OBJ_52
	OBJ_53
	COINS
	CHEST
	EGGS
	TRIDENT
	VASE
	EMERALD
	PYRAMID
	PEARL
	RUG
	OBJ_63
	CHAIN
	RUBY
	JADE
	AMBER
	SAPPH
	OBJ_69
)
