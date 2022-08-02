package cmd

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */
var locations = [...]Location{
	/* 0: LOC_NOWHERE */ {descriptions: Descriptions{small: "",
		big: ""},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 1: LOC_START */ {descriptions: Descriptions{small: "You're in front of building.",
		big: "You are standing at the end of a road before a small brick building.\nAround you is a forest.  A small stream flows out of the building and\ndown a gully."},
		sound: STREAM_GURGLES, loud: false, pic: "building"},
	/* 2: LOC_HILL */ {descriptions: Descriptions{small: "You're at hill in road.",
		big: "You have walked up a hill, still in the forest.  The road slopes back\ndown the other side of the hill.  There is a building in the distance."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 3: LOC_BUILDING */ {descriptions: Descriptions{small: "You're inside building.",
		big: "You are inside a building, a well house for a large spring."},
		sound: STREAM_GURGLES, loud: false, pic: "building"},
	/* 4: LOC_VALLEY */ {descriptions: Descriptions{small: "You're in valley.",
		big: "You are in a valley in the forest beside a stream tumbling along a\nrocky bed."},
		sound: STREAM_GURGLES, loud: false, pic: "valley"},
	/* 5: LOC_ROADEND */ {descriptions: Descriptions{small: "You're at end of road.",
		big: "The road, which approaches from the east, ends here amid the trees."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 6: LOC_CLIFF */ {descriptions: Descriptions{small: "You're at cliff.",
		big: "The forest thins out here to reveal a steep cliff.  There is no way\ndown, but a small ledge can be seen to the west across the chasm."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 7: LOC_SLIT */ {descriptions: Descriptions{small: "You're at slit in streambed.",
		big: "At your feet all the water of the stream splashes into a 2-inch slit\nin the rock.  Downstream the streambed is bare rock."},
		sound: STREAM_GURGLES, loud: false, pic: ""},
	/* 8: LOC_GRATE */ {descriptions: Descriptions{small: "You're outside grate.",
		big: "You are in a 20-foot depression floored with bare dirt.  Set into the\ndirt is a strong steel grate mounted in concrete.  A dry streambed\nleads into the depression."},
		sound: SILENT, loud: false, pic: "grate"},
	/* 9: LOC_BELOWGRATE */ {descriptions: Descriptions{small: "You're below the grate.",
		big: "You are in a small chamber beneath a 3x3 steel grate to the surface.\nA low crawl over cobbles leads inward to the west."},
		sound: SILENT, loud: false}, // pic: "chamber_grate"},
	/* 10: LOC_COBBLE */ {descriptions: Descriptions{small: "You're in cobble crawl.",
		big: "You are crawling over cobbles in a low passage.  There is a dim light\nat the east end of the passage."},
		sound: SILENT, loud: false, pic: "xyzzy"},
	/* 11: LOC_DEBRIS */ {descriptions: Descriptions{small: "You're in debris room.",
		big: "You are in a debris room filled with stuff washed in from the surface.\nA low wide passage with cobbles becomes plugged with mud and debris\nhere, but an awkward canyon leads upward and west.  In the mud someone\nhas scrawled, 'MAGIC WORD XYZZY'."},
		sound: SILENT, loud: false, pic: "xyzzy"},
	/* 12: LOC_AWKWARD */ {descriptions: Descriptions{small: "",
		big: "You are in an awkward sloping east/west canyon."},
		sound: SILENT, loud: false, pic: "xyzzy"},
	/* 13: LOC_BIRDCHAMBER */ {descriptions: Descriptions{small: "You're in bird chamber.",
		big: "You are in a splendid chamber thirty feet high.  The walls are frozen\nrivers of orange stone.  An awkward canyon and a good passage exit\nfrom east and west sides of the chamber."},
		sound: SILENT, loud: false, pic: "bird_chamber"},
	/* 14: LOC_PITTOP */ {descriptions: Descriptions{small: "You're at top of small pit.",
		big: "At your feet is a small pit breathing traces of white mist.  An east\npassage ends here except for a small crack leading on."},
		sound: SILENT, loud: false, pic: "top_pit"},
	/* 15: LOC_MISTHALL */ {descriptions: Descriptions{small: "You're in Hall of Mists.",
		big: "You are at one end of a vast hall stretching forward out of sight to\nthe west.  There are openings to either side.  Nearby, a wide stone\nstaircase leads downward.  The hall is filled with wisps of white mist\nswaying to and fro almost as if alive.  A cold wind blows up the\nstaircase.  There is a passage at the top of a dome behind you."},
		sound: WIND_WHISTLES, loud: false, pic: "hall_mists_east"},
	/* 16: LOC_CRACK */ {descriptions: Descriptions{small: "",
		big: "The crack is far too small for you to follow.  At its widest it is\nbarely wide enough to admit your foot."},
		sound: SILENT, loud: false, pic: ""},
	/* 17: LOC_EASTBANK */ {descriptions: Descriptions{small: "You're on east bank of fissure.",
		big: "You are on the east bank of a fissure slicing clear across the hall.\nThe mist is quite thick here, and the fissure is too wide to jump."},
		sound: SILENT, loud: false, pic: "hall_mists_east"},
	/* 18: LOC_NUGGET */ {descriptions: Descriptions{small: "You're in nugget-of-gold room.",
		big: "This is a low room with a crude note on the wall.  The note says,\n'You won't get it up the steps'."},
		sound: SILENT, loud: false, pic: "hall_mists_east"},
	/* 19: LOC_KINGHALL */ {descriptions: Descriptions{small: "You're in Hall of Mt King.",
		big: "You are in the Hall of the Mountain King, with passages off in all\ndirections."},
		sound: SILENT, loud: false, pic: "hall_mountain_king"},
	/* 20: LOC_NECKBROKE */ {descriptions: Descriptions{small: "",
		big: "You are at the bottom of the pit with a broken neck."},
		sound: SILENT, loud: false, pic: ""},
	/* 21: LOC_NOMAKE */ {descriptions: Descriptions{small: "",
		big: "You didn't make it."},
		sound: SILENT, loud: false, pic: ""},
	/* 22: LOC_DOME */ {descriptions: Descriptions{small: "",
		big: "The dome is unclimbable."},
		sound: SILENT, loud: false, pic: ""},
	/* 23: LOC_WESTEND */ {descriptions: Descriptions{small: "You're at west end of Twopit Room.",
		big: "You are at the west end of the Twopit Room.  There is a large hole in\nthe wall above the pit at this end of the room."},
		sound: SILENT, loud: false, pic: "two_pit"},
	/* 24: LOC_EASTPIT */ {descriptions: Descriptions{small: "You're in east pit.",
		big: "You are at the bottom of the eastern pit in the Twopit Room.  There is\na small pool of oil in one corner of the pit."},
		sound: SILENT, loud: false, pic: "two_pit"},
	/* 25: LOC_WESTPIT */ {descriptions: Descriptions{small: "You're in west pit.",
		big: "You are at the bottom of the western pit in the Twopit Room.  There is\na large hole in the wall about 25 feet above you."},
		sound: SILENT, loud: false, pic: "two_pit"},
	/* 26: LOC_CLIMBSTALK */ {descriptions: Descriptions{small: "",
		big: "You clamber up the plant and scurry through the hole at the top."},
		sound: SILENT, loud: false, pic: "two_pit"},
	/* 27: LOC_WESTBANK */ {descriptions: Descriptions{small: "You're on west bank of fissure.",
		big: "You are on the west side of the fissure in the Hall of Mists."},
		sound: SILENT, loud: false, pic: "hall_mist_west"},
	/* 28: LOC_FLOORHOLE */ {descriptions: Descriptions{small: "You're in n/s passage above e/w passage.",
		big: "You are in a low n/s passage at a hole in the floor.  The hole goes\ndown to an e/w passage."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 29: LOC_SOUTHSIDE */ {descriptions: Descriptions{small: "",
		big: "You are in the south side chamber."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 30: LOC_WESTSIDE */ {descriptions: Descriptions{small: "You're in the west side chamber.",
		big: "You are in the west side chamber of the Hall of the Mountain King.\nA passage continues west and up here."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 31: LOC_BUILDING1 */ {descriptions: Descriptions{small: "",
		big: ""},
		sound: SILENT, loud: false, pic: "building"},
	/* 32: LOC_SNAKEBLOCK */ {descriptions: Descriptions{small: "",
		big: "You can't get by the snake."},
		sound: SILENT, loud: false, pic: "hall_mountain_king"},
	/* 33: LOC_Y2 */ {descriptions: Descriptions{small: "You're at 'Y2'.",
		big: "You are in a large room, with a passage to the south, a passage to the\nwest, and a wall of broken rock to the east.  There is a large 'Y2' on\na rock in the room's center."},
		sound: SILENT, loud: false, pic: "y2"},
	/* 34: LOC_JUMBLE */ {descriptions: Descriptions{small: "",
		big: "You are in a jumble of rock, with cracks everywhere."},
		sound: SILENT, loud: false, pic: "y2"},
	/* 35: LOC_WINDOW1 */ {descriptions: Descriptions{small: "You're at window on pit.",
		big: "You're at a low window overlooking a huge pit, which extends up out of\nsight.  A floor is indistinctly visible over 50 feet below.  Traces of\nwhite mist cover the floor of the pit, becoming thicker to the right.\nMarks in the dust around the window would seem to indicate that\nsomeone has been here recently.  Directly across the pit from you and\n25 feet away there is a similar window looking into a lighted room.  A\nshadowy figure can be seen there peering back at you."},
		sound: SILENT, loud: false, pic: "y2"},
	/* 36: LOC_BROKEN */ {descriptions: Descriptions{small: "You're in dirty passage.",
		big: "You are in a dirty broken passage.  To the east is a crawl.  To the\nwest is a large passage.  Above you is a hole to another passage."},
		sound: SILENT, loud: false, pic: "passage_broken"},
	/* 37: LOC_SMALLPITBRINK */ {descriptions: Descriptions{small: "You're at brink of small pit.",
		big: "You are on the brink of a small clean climbable pit.  A crawl leads\nwest."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 38: LOC_SMALLPIT */ {descriptions: Descriptions{small: "You're at bottom of pit with stream.",
		big: "You are in the bottom of a small pit with a little stream, which\nenters and exits through tiny slits."},
		sound: STREAM_GURGLES, loud: false, pic: "nsewud"},
	/* 39: LOC_DUSTY */ {descriptions: Descriptions{small: "You're in dusty rock room.",
		big: "You are in a large room full of dusty rocks.  There is a big hole in\nthe floor.  There are cracks everywhere, and a passage leading east."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 40: LOC_PARALLEL1 */ {descriptions: Descriptions{small: "",
		big: "You have crawled through a very low wide passage parallel to and north\nof the Hall of Mists."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 41: LOC_MISTWEST */ {descriptions: Descriptions{small: "You're at west end of Hall of Mists.",
		big: "You are at the west end of the Hall of Mists.  A low wide crawl\ncontinues west and another goes north.  To the south is a little\npassage 6 feet off the floor."},
		sound: SILENT, loud: false, pic: "hall_mists_west"},
	/* 42: LOC_ALIKE1 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 43: LOC_ALIKE2 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 44: LOC_ALIKE3 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 45: LOC_ALIKE4 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 46: LOC_MAZEEND1 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 47: LOC_MAZEEND2 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 48: LOC_MAZEEND3 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 49: LOC_ALIKE5 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 50: LOC_ALIKE6 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 51: LOC_ALIKE7 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 52: LOC_ALIKE8 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 53: LOC_ALIKE9 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 54: LOC_MAZEEND4 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 55: LOC_ALIKE10 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 56: LOC_MAZEEND5 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 57: LOC_PITBRINK */ {descriptions: Descriptions{small: "You're at brink of pit.",
		big: "You are on the brink of a thirty foot pit with a massive orange column\ndown one wall.  You could climb down here but you could not get back\nup.  The maze continues at this level."},
		sound: SILENT, loud: false, pic: "brink_pit"},
	/* 58: LOC_MAZEEND6 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 59: LOC_PARALLEL2 */ {descriptions: Descriptions{small: "",
		big: "You have crawled through a very low wide passage parallel to and north\nof the Hall of Mists."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 60: LOC_LONGEAST */ {descriptions: Descriptions{small: "You're at east end of long hall.",
		big: "You are at the east end of a very long hall apparently without side\nchambers.  To the east a low wide crawl slants up.  To the north a\nround two foot hole slants down."},
		sound: SILENT, loud: false, pic: "hall_long"},
	/* 61: LOC_LONGWEST */ {descriptions: Descriptions{small: "You're at west end of long hall.",
		big: "You are at the west end of a very long featureless hall.  The hall\njoins up with a narrow north/south passage."},
		sound: SILENT, loud: false, pic: "hall_long"},
	/* 62: LOC_CROSSOVER */ {descriptions: Descriptions{small: "",
		big: "You are at a crossover of a high n/s passage and a low e/w one."},
		sound: SILENT, loud: false, pic: "hall_long"},
	/* 63: LOC_DEADEND7 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 64: LOC_COMPLEX */ {descriptions: Descriptions{small: "You're at complex junction.",
		big: "You are at a complex junction.  A low hands and knees passage from the\nnorth joins a higher crawl from the east to make a walking passage\ngoing west.  There is also a large room above.  The air is damp here."},
		sound: WIND_WHISTLES, loud: false, pic: "bedquilt"},
	/* 65: LOC_BEDQUILT */ {descriptions: Descriptions{small: "You're in Bedquilt.",
		big: "You are in Bedquilt, a long east/west passage with holes everywhere.\nTo explore at random select north, south, up, or down."},
		sound: SILENT, loud: false, pic: "bedquilt"},
	/* 66: LOC_SWISSCHEESE */ {descriptions: Descriptions{small: "You're in Swiss Cheese Room.",
		big: "You are in a room whose walls resemble Swiss cheese.  Obvious passages\ngo west, east, ne, and nw.  Part of the room is occupied by a large\nbedrock block."},
		sound: SILENT, loud: false, pic: "room_swiss"},
	/* 67: LOC_EASTEND */ {descriptions: Descriptions{small: "You're at east end of Twopit Room.",
		big: "You are at the east end of the Twopit Room.  The floor here is\nlittered with thin rock slabs, which make it easy to descend the pits.\nThere is a path here bypassing the pits to connect passages from east\nand west.  There are holes all over, but the only big one is on the\nwall directly over the west pit where you can't get to it."},
		sound: SILENT, loud: false, pic: "two_pit"},
	/* 68: LOC_SLAB */ {descriptions: Descriptions{small: "You're in Slab Room.",
		big: "You are in a large low circular chamber whose floor is an immense slab\nfallen from the ceiling (Slab Room).  East and west there once were\nlarge passages, but they are now filled with boulders.  Low small\npassages go north and south, and the south one quickly bends west\naround the boulders."},
		sound: SILENT, loud: false, pic: "room_slab"},
	/* 69: LOC_SECRET1 */ {descriptions: Descriptions{small: "",
		big: "You are in a secret n/s canyon above a large room."},
		sound: SILENT, loud: false, pic: "canyon_mirror_secret"},
	/* 70: LOC_SECRET2 */ {descriptions: Descriptions{small: "",
		big: "You are in a secret n/s canyon above a sizable passage."},
		sound: SILENT, loud: false, pic: "canyon_junction"},
	/* 71: LOC_THREEJUNCTION */ {descriptions: Descriptions{small: "You're at junction of three secret canyons.",
		big: "You are in a secret canyon at a junction of three canyons, bearing\nnorth, south, and se.  The north one is as tall as the other two\ncombined."},
		sound: SILENT, loud: false},
	/* 72: LOC_LOWROOM */ {descriptions: Descriptions{small: "You're in large low room.",
		big: "You are in a large low room.  Crawls lead north, se, and sw."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 73: LOC_DEADCRAWL */ {descriptions: Descriptions{small: "",
		big: "Dead end crawl."},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 74: LOC_SECRET3 */ {descriptions: Descriptions{small: "You're in secret e/w canyon above tight canyon.",
		big: "You are in a secret canyon which here runs e/w.  It crosses over a\nvery tight canyon 15 feet below.  If you go down you may not be able\nto get back up."},
		sound: SILENT, loud: false, pic: "canyon_mirror_secret"},
	/* 75: LOC_WIDEPLACE */ {descriptions: Descriptions{small: "",
		big: "You are at a wide Place in a very tight n/s canyon."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 76: LOC_TIGHTPLACE */ {descriptions: Descriptions{small: "",
		big: "The canyon here becomes too tight to go further south."},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 77: LOC_TALL */ {descriptions: Descriptions{small: "",
		big: "You are in a tall e/w canyon.  A low tight crawl goes 3 feet north and\nseems to open up."},
		sound: SILENT, loud: false, pic: "canyon_tall"},
	/* 78: LOC_BOULDERS1 */ {descriptions: Descriptions{small: "",
		big: "The canyon runs into a mass of boulders -- dead end."},
		sound: SILENT, loud: false, pic: "volcano_boulders"},
	/* 79: LOC_SEWER */ {descriptions: Descriptions{small: "",
		big: "The stream flows out through a pair of 1 foot diameter sewer pipes.\nIt would be advisable to use the exit."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 80: LOC_ALIKE11 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 81: LOC_MAZEEND8 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 82: LOC_MAZEEND9 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 83: LOC_ALIKE12 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 84: LOC_ALIKE13 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 85: LOC_MAZEEND10 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 86: LOC_MAZEEND11 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 87: LOC_ALIKE14 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all alike."},
		sound: SILENT, loud: false, pic: "maze_alike"},
	/* 88: LOC_NARROW */ {descriptions: Descriptions{small: "You're in narrow corridor.",
		big: "You are in a long, narrow corridor stretching out of sight to the\nwest.  At the eastern end is a hole through which you can see a\nprofusion of leaves."},
		sound: SILENT, loud: false, pic: "sapling"},
	/* 89: LOC_NOCLIMB */ {descriptions: Descriptions{small: "",
		big: "There is nothing here to climb.  Use 'up' or 'out' to leave the pit."},
		sound: SILENT, loud: false, pic: "two_pit"},
	/* 90: LOC_PLANTTOP */ {descriptions: Descriptions{small: "",
		big: "You have climbed up the plant and out of the pit."},
		sound: SILENT, loud: false, pic: "two_pit"},
	/* 91: LOC_INCLINE */ {descriptions: Descriptions{small: "You're at steep incline above large room.",
		big: "You are at the top of a steep incline above a large room.  You could\nclimb down here, but you would not be able to climb up.  There is a\npassage leading back to the north."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 92: LOC_GIANTROOM */ {descriptions: Descriptions{small: "You're in Giant Room.",
		big: "You are in the Giant Room.  The ceiling here is too high up for your\nlamp to show it.  Cavernous passages lead east, north, and south.  On\nthe west wall is scrawled the inscription, 'FEE FIE FOE FOO' [sic]."},
		sound: SILENT, loud: false, pic: "room_giant"},
	/* 93: LOC_CAVEIN */ {descriptions: Descriptions{small: "",
		big: "The passage here is blocked by a recent cave-in."},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 94: LOC_IMMENSE */ {descriptions: Descriptions{small: "",
		big: "You are at one end of an immense north/south passage."},
		sound: WIND_WHISTLES, loud: false, pic: "nsewud"},
	/* 95: LOC_WATERFALL */ {descriptions: Descriptions{small: "You're in cavern with waterfall.",
		big: "You are in a magnificent cavern with a rushing stream, which cascades\nover a sparkling waterfall into a roaring whirlpool which disappears\nthrough a hole in the floor.  Passages exit to the south and west."},
		sound: STREAM_SPLASHES, loud: false, pic: "cavern_magnificent"},
	/* 96: LOC_SOFTROOM */ {descriptions: Descriptions{small: "You're in Soft Room.",
		big: "You are in the Soft Room.  The walls are covered with heavy curtains,\nthe floor with a thick pile carpet.  Moss covers the ceiling."},
		sound: SILENT, loud: false, pic: "room_soft"},
	/* 97: LOC_ORIENTAL */ {descriptions: Descriptions{small: "You're in Oriental Room.",
		big: "This is the Oriental Room.  Ancient oriental cave drawings cover the\nwalls.  A gently sloping passage leads upward to the north, another\npassage leads se, and a hands and knees crawl leads west."},
		sound: SILENT, loud: false, pic: "room_oriental"},
	/* 98: LOC_MISTY */ {descriptions: Descriptions{small: "You're in misty cavern.",
		big: "You are following a wide path around the outer edge of a large cavern.\nFar below, through a heavy white mist, strange splashing noises can be\nheard.  The mist rises up through a fissure in the ceiling.  The path\nexits to the south and west."},
		sound: NO_MEANING, loud: false, pic: "plover"},
	/* 99: LOC_ALCOVE */ {descriptions: Descriptions{small: "You're in alcove.",
		big: "You are in an alcove.  A small nw path seems to widen after a short\ndistance.  An extremely tight tunnel leads east.  It looks like a very\ntight squeeze.  An eerie light can be seen at the other end."},
		sound: SILENT, loud: false, pic: "plover"},
	/* 100: LOC_PLOVER */ {descriptions: Descriptions{small: "You're in Plover Room.",
		big: "You're in a small chamber lit by an eerie green light.  An extremely\nnarrow tunnel exits to the west.  A dark corridor leads ne."},
		sound: SILENT, loud: false, pic: "plover"},
	/* 101: LOC_DARKROOM */ {descriptions: Descriptions{small: "You're in dark-room.",
		big: "You're in the dark-room.  A corridor leading south is the only exit."},
		sound: SILENT, loud: false, pic: "room_dark"},
	/* 102: LOC_ARCHED */ {descriptions: Descriptions{small: "You're in arched hall.",
		big: "You are in an arched hall.  A coral passage once continued up and east\nfrom here, but is now blocked by debris.  The air smells of sea water."},
		sound: SILENT, loud: false, pic: "room_shell"},
	/* 103: LOC_SHELLROOM */ {descriptions: Descriptions{small: "You're in Shell Room.",
		big: "You're in a large room carved out of sedimentary rock.  The floor and\nwalls are littered with bits of shells embedded in the stone.  A\nshallow passage proceeds downward, and a somewhat steeper one leads\nup.  A low hands and knees passage enters from the south."},
		sound: SILENT, loud: false, pic: "room_shell"},
	/* 104: LOC_SLOPING1 */ {descriptions: Descriptions{small: "",
		big: "You are in a long sloping corridor with ragged sharp walls."},
		sound: SILENT, loud: false, pic: "corridor_sloping"},
	/* 105: LOC_CULDESAC */ {descriptions: Descriptions{small: "",
		big: "You are in a cul-de-sac about eight feet across."},
		sound: SILENT, loud: false, pic: "corridor_sloping"},
	/* 106: LOC_ANTEROOM */ {descriptions: Descriptions{small: "You're in anteroom.",
		big: "You are in an anteroom leading to a large passage to the east.  Small\npassages go west and up.  The remnants of recent digging are evident.\nA sign in midair here says 'Cave under construction beyond this point.\nProceed at own risk.  [Witt Construction Company]'"},
		sound: SILENT, loud: false, pic: "room_ante"},
	/* 107: LOC_DIFFERENT1 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisty little passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 108: LOC_WITTSEND */ {descriptions: Descriptions{small: "You're at Witt's End.",
		big: "You are at Witt's End.  Passages lead off in *ALL* directions."},
		sound: SILENT, loud: false, pic: "room_ante"},
	/* 109: LOC_MIRRORCANYON */ {descriptions: Descriptions{small: "You're in Mirror Canyon.",
		big: "You are in a north/south canyon about 25 feet across.  The floor is\ncovered by white mist seeping in from the north.  The walls extend\nupward for well over 100 feet.  Suspended from some unseen point far\nabove you, an enormous two-sided mirror is hanging parallel to and\nmidway between the canyon walls.  (The mirror is obviously provided\nfor the use of the dwarves who, as you know, are extremely vain.)  A\nsmall window can be seen in either wall, some fifty feet up."},
		sound: WIND_WHISTLES, loud: false, pic: "canyon_mirror_secret"},
	/* 110: LOC_WINDOW2 */ {descriptions: Descriptions{small: "You're at window on pit.",
		big: "You're at a low window overlooking a huge pit, which extends up out of\nsight.  A floor is indistinctly visible over 50 feet below.  Traces of\nwhite mist cover the floor of the pit, becoming thicker to the left.\nMarks in the dust around the window would seem to indicate that\nsomeone has been here recently.  Directly across the pit from you and\n25 feet away there is a similar window looking into a lighted room.  A\nshadowy figure can be seen there peering back at you."},
		sound: SILENT, loud: false, pic: "canyon_junction"},
	/* 111: LOC_TOPSTALACTITE */ {descriptions: Descriptions{small: "You're at top of stalactite.",
		big: "A large stalactite extends from the roof and almost reaches the floor\nbelow.  You could climb down it, and jump from it to the floor, but\nhaving done so you would be unable to reach it to climb back up."},
		sound: SILENT, loud: false, pic: "canyon_junction"},
	/* 112: LOC_DIFFERENT2 */ {descriptions: Descriptions{small: "",
		big: "You are in a little maze of twisting passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 113: LOC_RESERVOIR */ {descriptions: Descriptions{small: "You're at reservoir.",
		big: "You are at the edge of a large underground reservoir.  An opaque cloud\nof white mist fills the room and rises rapidly upward.  The lake is\nfed by a stream, which tumbles out of a hole in the wall about 10 feet\noverhead and splashes noisily into the water somewhere within the\nmist.  There is a passage going back toward the south."},
		sound: STREAM_SPLASHES, loud: false, pic: "reservoir"},
	/* 114: LOC_MAZEEND12 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 115: LOC_NE */ {descriptions: Descriptions{small: "You're at ne end.",
		big: "You are at the northeast end of an immense room, even larger than the\nGiant Room.  It appears to be a repository for the 'Adventure'\nprogram.  Massive torches far overhead bathe the room with smoky\nyellow light.  Scattered about you can be seen a pile of bottles (all\nof them empty), a nursery of young beanstalks murmuring quietly, a bed\nof oysters, a bundle of black rods with rusty stars on their ends, and\na collection of brass lanterns.  Off to one side a great many dwarves\nare sleeping on the floor, snoring loudly.  A notice nearby reads: 'Do\nnot disturb the dwarves!'  An immense mirror is hanging against one\nwall, and stretches to the other end of the room, where various other\nsundry objects can be glimpsed dimly in the distance."},
		sound: MURMURING_SNORING, loud: false, pic: "canyon_mirror_secret"},
	/* 116: LOC_SW */ {descriptions: Descriptions{small: "You're at sw end.",
		big: "You are at the southwest end of the repository.  To one side is a pit\nfull of fierce green snakes.  On the other side is a row of small\nwicker cages, each of which contains a little sulking bird.  In one\ncorner is a bundle of black rods with rusty marks on their ends.  A\nlarge number of velvet pillows are scattered about on the floor.  A\nvast mirror stretches off to the northeast.  At your feet is a large\nsteel grate, next to which is a sign that reads, 'Treasure Vault.\nKeys in main office.'"},
		sound: SNAKES_HISSING, loud: false, pic: "nsewud"},
	/* 117: LOC_SWCHASM */ {descriptions: Descriptions{small: "You're on sw side of chasm.",
		big: "You are on one side of a large, deep chasm.  A heavy white mist rising\nup from below obscures all view of the far side.  A sw path leads away\nfrom the chasm into a winding corridor."},
		sound: SILENT, loud: false, pic: "bridge_troll"},
	/* 118: LOC_WINDING */ {descriptions: Descriptions{small: "You're in sloping corridor.",
		big: "You are in a long winding corridor sloping out of sight in both\ndirections."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 119: LOC_SECRET4 */ {descriptions: Descriptions{small: "",
		big: "You are in a secret canyon which exits to the north and east."},
		sound: SILENT, loud: false, pic: "canyon_mirror_secret"},
	/* 120: LOC_SECRET5 */ {descriptions: Descriptions{small: "",
		big: "You are in a secret canyon which exits to the north and east."},
		sound: SILENT, loud: false, pic: "canyon_mirror_secret"},
	/* 121: LOC_SECRET6 */ {descriptions: Descriptions{small: "",
		big: "You are in a secret canyon which exits to the north and east."},
		sound: SILENT, loud: false, pic: "canyon_mirror_secret"},
	/* 122: LOC_NECHASM */ {descriptions: Descriptions{small: "You're on ne side of chasm.",
		big: "You are on the far side of the chasm.  A ne path leads away from the\nchasm on this side."},
		sound: SILENT, loud: false, pic: "bridge_troll"},
	/* 123: LOC_CORRIDOR */ {descriptions: Descriptions{small: "You're in corridor.",
		big: "You're in a long east/west corridor.  A faint rumbling noise can be\nheard in the distance."},
		sound: DULL_RUMBLING, loud: false, pic: "path_fork"},
	/* 124: LOC_FORK */ {descriptions: Descriptions{small: "You're at fork in path.",
		big: "The path forks here.  The left fork leads northeast.  A dull rumbling\nseems to get louder in that direction.  The right fork leads southeast\ndown a gentle slope.  The main corridor enters from the west."},
		sound: DULL_RUMBLING, loud: false, pic: "path_fork"},
	/* 125: LOC_WARMWALLS */ {descriptions: Descriptions{small: "You're at junction with warm walls.",
		big: "The walls are quite warm here.  From the north can be heard a steady\nroar, so loud that the entire cave seems to be trembling.  Another\npassage leads south, and a low crawl goes east."},
		sound: LOUD_ROAR, loud: false, pic: "volcano_boulders"},
	/* 126: LOC_BREATHTAKING */ {descriptions: Descriptions{small: "You're at breath-taking view.",
		big: "You are on the edge of a breath-taking view.  Far below you is an\nactive volcano, from which great gouts of molten lava come surging\nout, cascading back down into the depths.  The glowing rock fills the\nfarthest reaches of the cavern with a blood-red glare, giving every-\nthing an eerie, macabre appearance.  The air is filled with flickering\nsparks of ash and a heavy smell of brimstone.  The walls are hot to\nthe touch, and the thundering of the volcano drowns out all other\nsounds.  Embedded in the jagged roof far overhead are myriad twisted\nformations composed of pure white alabaster, which scatter the murky\nlight into sinister apparitions upon the walls.  To one side is a deep\ngorge, filled with a bizarre chaos of tortured rock which seems to\nhave been crafted by the devil himself.  An immense river of fire\ncrashes out from the depths of the volcano, burns its way through the\ngorge, and plummets into a bottomless pit far off to your left.  To\nthe right, an immense geyser of blistering steam erupts continuously\nfrom a barren island in the center of a sulfurous lake, which bubbles\nominously.  The far right wall is aflame with an incandescence of its\nown, which lends an additional infernal splendor to the already\nhellish scene.  A dark, foreboding passage exits to the south."},
		sound: TOTAL_ROAR, loud: true, pic: "volcano"},
	/* 127: LOC_BOULDERS2 */ {descriptions: Descriptions{small: "You're in Chamber of Boulders.",
		big: "You are in a small chamber filled with large boulders.  The walls are\nvery warm, causing the air in the room to be almost stifling from the\nheat.  The only exit is a crawl heading west, through which is coming\na low rumbling."},
		sound: DULL_RUMBLING, loud: false, pic: "volcano_boulders"},
	/* 128: LOC_LIMESTONE */ {descriptions: Descriptions{small: "You're in limestone passage.",
		big: "You are walking along a gently sloping north/south passage\nlined with oddly shaped limestone formations."},
		sound: SILENT, loud: false, pic: "path_fork"},
	/* 129: LOC_BARRENFRONT */ {descriptions: Descriptions{small: "You're in front of Barren Room.",
		big: "You are standing at the entrance to a large, barren room.  A notice\nabove the entrance reads:  'Caution!  Bear in room!'"},
		sound: SILENT, loud: false, pic: "room_baren"},
	/* 130: LOC_BARRENROOM */ {descriptions: Descriptions{small: "You're in Barren Room.",
		big: "You are inside a barren room.  The center of the room is completely\nempty except for some dust.  Marks in the dust lead away toward the\nfar end of the room.  The only exit is the way you came in."},
		sound: SILENT, loud: false, pic: "room_baren"},
	/* 131: LOC_DIFFERENT3 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of twisting little passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 132: LOC_DIFFERENT4 */ {descriptions: Descriptions{small: "",
		big: "You are in a little maze of twisty passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 133: LOC_DIFFERENT5 */ {descriptions: Descriptions{small: "",
		big: "You are in a twisting maze of little passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 134: LOC_DIFFERENT6 */ {descriptions: Descriptions{small: "",
		big: "You are in a twisting little maze of passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 135: LOC_DIFFERENT7 */ {descriptions: Descriptions{small: "",
		big: "You are in a twisty little maze of passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 136: LOC_DIFFERENT8 */ {descriptions: Descriptions{small: "",
		big: "You are in a twisty maze of little passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 137: LOC_DIFFERENT9 */ {descriptions: Descriptions{small: "",
		big: "You are in a little twisty maze of passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 138: LOC_DIFFERENT10 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of little twisting passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 139: LOC_DIFFERENT11 */ {descriptions: Descriptions{small: "",
		big: "You are in a maze of little twisty passages, all different."},
		sound: SILENT, loud: false, pic: "maze_different"},
	/* 140: LOC_DEADEND13 */ {descriptions: Descriptions{small: "",
		big: "Dead end"},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 141: LOC_ROUGHHEWN */ {descriptions: Descriptions{small: "",
		big: "You are in a long, rough-hewn, north/south corridor."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 142: LOC_BADDIRECTION */ {descriptions: Descriptions{small: "",
		big: "There is no way to go that direction."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 143: LOC_LARGE */ {descriptions: Descriptions{small: "",
		big: "You are in a large chamber with passages to the west and north."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 144: LOC_STOREROOM */ {descriptions: Descriptions{small: "",
		big: "You are in the ogre's storeroom.  The only exit is to the south."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 145: LOC_FOREST1 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 146: LOC_FOREST2 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 147: LOC_FOREST3 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 148: LOC_FOREST4 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 149: LOC_FOREST5 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 150: LOC_FOREST6 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 151: LOC_FOREST7 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 152: LOC_FOREST8 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 153: LOC_FOREST9 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 154: LOC_FOREST10 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 155: LOC_FOREST11 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 156: LOC_FOREST12 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 157: LOC_FOREST13 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 158: LOC_FOREST14 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 159: LOC_FOREST15 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 160: LOC_FOREST16 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 161: LOC_FOREST17 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 162: LOC_FOREST18 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 163: LOC_FOREST19 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 164: LOC_FOREST20 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 165: LOC_FOREST21 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 166: LOC_FOREST22 */ {descriptions: Descriptions{small: "",
		big: "You are wandering aimlessly through the forest."},
		sound: SILENT, loud: false, pic: "forest"},
	/* 167: LOC_LEDGE */ {descriptions: Descriptions{small: "You're on ledge.",
		big: "You are on a small ledge on one face of a sheer cliff.  There are no\npaths away from the ledge.  Across the chasm is a small clearing\nsurrounded by forest."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 168: LOC_RESBOTTOM */ {descriptions: Descriptions{small: "You're at bottom of reservoir.",
		big: "You are walking across the bottom of the reservoir.  Walls of water\nrear up on either side.  The roar of the water cascading past is\nnearly deafening, and the mist is so thick you can barely see."},
		sound: TOTAL_ROAR, loud: true, pic: "nsewud"},
	/* 169: LOC_RESNORTH */ {descriptions: Descriptions{small: "You're north of reservoir.",
		big: "You are at the northern edge of the reservoir.  A northwest passage\nleads sharply up from here."},
		sound: WATERS_CRASHING, loud: false, pic: "nsewud"},
	/* 170: LOC_TREACHEROUS */ {descriptions: Descriptions{small: "",
		big: "You are scrambling along a treacherously steep, rocky passage."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 171: LOC_STEEP */ {descriptions: Descriptions{small: "",
		big: "You are on a very steep incline, which widens at it goes upward."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 172: LOC_CLIFFBASE */ {descriptions: Descriptions{small: "You're at base of cliff.",
		big: "You are at the base of a nearly vertical cliff.  There are some\nslim footholds which would enable you to climb up, but it looks\nextremely dangerous.  Here at the base of the cliff lie the remains\nof several earlier adventurers who apparently failed to make it."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 173: LOC_CLIFFACE */ {descriptions: Descriptions{small: "",
		big: "You are climbing along a nearly vertical cliff."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 174: LOC_FOOTSLIP */ {descriptions: Descriptions{small: "",
		big: "Just as you reach the top, your foot slips on a loose rock and you\ntumble several hundred feet to join the other unlucky adventurers."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 175: LOC_CLIFFTOP */ {descriptions: Descriptions{small: "",
		big: "Just as you reach the top, your foot slips on a loose rock and you\nmake one last desperate grab.  Your luck holds, as does your grip.\nWith an enormous heave, you lift yourself to the ledge above."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 176: LOC_CLIFFLEDGE */ {descriptions: Descriptions{small: "You're at top of cliff.",
		big: "You are on a small ledge at the top of a nearly vertical cliff.\nThere is a low crawl leading off to the northeast."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 177: LOC_REACHDEAD */ {descriptions: Descriptions{small: "",
		big: "You have reached a dead end."},
		sound: SILENT, loud: false, pic: "dead_end"},
	/* 178: LOC_GRUESOME */ {descriptions: Descriptions{small: "",
		big: "There is now one more gruesome aspect to the spectacular vista."},
		sound: SILENT, loud: false, pic: "nsewud"},
	/* 179: LOC_FOOF1 */ {descriptions: Descriptions{small: "",
		big: ">>Foof!<<"},
		sound: SILENT, loud: false, pic: ""},
	/* 180: LOC_FOOF2 */ {descriptions: Descriptions{small: "",
		big: ">>Foof!<<"},
		sound: SILENT, loud: false, pic: ""},
	/* 181: LOC_FOOF3 */ {descriptions: Descriptions{small: "",
		big: ">>Foof!<<"},
		sound: SILENT, loud: false, pic: ""},
	/* 182: LOC_FOOF4 */ {descriptions: Descriptions{small: "",
		big: ">>Foof!<<"},
		sound: SILENT, loud: false, pic: ""},
	/* 183: LOC_FOOF5 */ {descriptions: Descriptions{small: "",
		big: ">>Foof!<<"},
		sound: SILENT, loud: false, pic: ""},
	/* 184: LOC_FOOF6 */ {descriptions: Descriptions{small: "",
		big: ">>Foof!<<"},
		sound: SILENT, loud: false, pic: ""},
}

const (
	LOC_NOWHERE = iota
	LOC_START
	LOC_HILL
	LOC_BUILDING
	LOC_VALLEY
	LOC_ROADEND
	LOC_CLIFF
	LOC_SLIT
	LOC_GRATE
	LOC_BELOWGRATE
	LOC_COBBLE
	LOC_DEBRIS
	LOC_AWKWARD
	LOC_BIRDCHAMBER
	LOC_PITTOP
	LOC_MISTHALL
	LOC_CRACK
	LOC_EASTBANK
	LOC_NUGGET
	LOC_KINGHALL
	LOC_NECKBROKE
	LOC_NOMAKE
	LOC_DOME
	LOC_WESTEND
	LOC_EASTPIT
	LOC_WESTPIT
	LOC_CLIMBSTALK
	LOC_WESTBANK
	LOC_FLOORHOLE
	LOC_SOUTHSIDE
	LOC_WESTSIDE
	LOC_BUILDING1
	LOC_SNAKEBLOCK
	LOC_Y2
	LOC_JUMBLE
	LOC_WINDOW1
	LOC_BROKEN
	LOC_SMALLPITBRINK
	LOC_SMALLPIT
	LOC_DUSTY
	LOC_PARALLEL1
	LOC_MISTWEST
	LOC_ALIKE1
	LOC_ALIKE2
	LOC_ALIKE3
	LOC_ALIKE4
	LOC_MAZEEND1
	LOC_MAZEEND2
	LOC_MAZEEND3
	LOC_ALIKE5
	LOC_ALIKE6
	LOC_ALIKE7
	LOC_ALIKE8
	LOC_ALIKE9
	LOC_MAZEEND4
	LOC_ALIKE10
	LOC_MAZEEND5
	LOC_PITBRINK
	LOC_MAZEEND6
	LOC_PARALLEL2
	LOC_LONGEAST
	LOC_LONGWEST
	LOC_CROSSOVER
	LOC_DEADEND7
	LOC_COMPLEX
	LOC_BEDQUILT
	LOC_SWISSCHEESE
	LOC_EASTEND
	LOC_SLAB
	LOC_SECRET1
	LOC_SECRET2
	LOC_THREEJUNCTION
	LOC_LOWROOM
	LOC_DEADCRAWL
	LOC_SECRET3
	LOC_WIDEPLACE
	LOC_TIGHTPLACE
	LOC_TALL
	LOC_BOULDERS1
	LOC_SEWER
	LOC_ALIKE11
	LOC_MAZEEND8
	LOC_MAZEEND9
	LOC_ALIKE12
	LOC_ALIKE13
	LOC_MAZEEND10
	LOC_MAZEEND11
	LOC_ALIKE14
	LOC_NARROW
	LOC_NOCLIMB
	LOC_PLANTTOP
	LOC_INCLINE
	LOC_GIANTROOM
	LOC_CAVEIN
	LOC_IMMENSE
	LOC_WATERFALL
	LOC_SOFTROOM
	LOC_ORIENTAL
	LOC_MISTY
	LOC_ALCOVE
	LOC_PLOVER
	LOC_DARKROOM
	LOC_ARCHED
	LOC_SHELLROOM
	LOC_SLOPING1
	LOC_CULDESAC
	LOC_ANTEROOM
	LOC_DIFFERENT1
	LOC_WITTSEND
	LOC_MIRRORCANYON
	LOC_WINDOW2
	LOC_TOPSTALACTITE
	LOC_DIFFERENT2
	LOC_RESERVOIR
	LOC_MAZEEND12
	LOC_NE
	LOC_SW
	LOC_SWCHASM
	LOC_WINDING
	LOC_SECRET4
	LOC_SECRET5
	LOC_SECRET6
	LOC_NECHASM
	LOC_CORRIDOR
	LOC_FORK
	LOC_WARMWALLS
	LOC_BREATHTAKING
	LOC_BOULDERS2
	LOC_LIMESTONE
	LOC_BARRENFRONT
	LOC_BARRENROOM
	LOC_DIFFERENT3
	LOC_DIFFERENT4
	LOC_DIFFERENT5
	LOC_DIFFERENT6
	LOC_DIFFERENT7
	LOC_DIFFERENT8
	LOC_DIFFERENT9
	LOC_DIFFERENT10
	LOC_DIFFERENT11
	LOC_DEADEND13
	LOC_ROUGHHEWN
	LOC_BADDIRECTION
	LOC_LARGE
	LOC_STOREROOM
	LOC_FOREST1
	LOC_FOREST2
	LOC_FOREST3
	LOC_FOREST4
	LOC_FOREST5
	LOC_FOREST6
	LOC_FOREST7
	LOC_FOREST8
	LOC_FOREST9
	LOC_FOREST10
	LOC_FOREST11
	LOC_FOREST12
	LOC_FOREST13
	LOC_FOREST14
	LOC_FOREST15
	LOC_FOREST16
	LOC_FOREST17
	LOC_FOREST18
	LOC_FOREST19
	LOC_FOREST20
	LOC_FOREST21
	LOC_FOREST22
	LOC_LEDGE
	LOC_RESBOTTOM
	LOC_RESNORTH
	LOC_TREACHEROUS
	LOC_STEEP
	LOC_CLIFFBASE
	LOC_CLIFFACE
	LOC_FOOTSLIP
	LOC_CLIFFTOP
	LOC_CLIFFLEDGE
	LOC_REACHDEAD
	LOC_GRUESOME
	LOC_FOOF1
	LOC_FOOF2
	LOC_FOOF3
	LOC_FOOF4
	LOC_FOOF5
	LOC_FOOF6
)
