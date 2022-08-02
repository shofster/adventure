package cmd

import "strings"

/*
 * Copyright (c 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

// COND_HBASE /* Bits past 10 indicate areas of interest to "hint" routines */
const COND_HBASE = 10 /* Base for location hint bits */

const (
	COND_LIT    = 1
	COND_OILY   = 1 << 1
	COND_FLUID  = 1 << 2
	COND_NOARRR = 1 << 3
	COND_NOBACK = 1 << 4
	COND_ABOVE  = 1 << 5
	COND_DEEP   = 1 << 6
	COND_FOREST = 1 << 7
	COND_FORCED = 1 << 8
	COND_CAVE   = 1 << 11
	COND_BIRD   = 1 << 12
	COND_SNAKE  = 1 << 13
	COND_MAZE   = 1 << 14
	COND_DARK   = 1 << 15
	COND_WITT   = 1 << 16
	COND_CLIFF  = 1 << 17
	COND_WOODS  = 1 << 18
	COND_OGRE   = 1 << 19
	COND_JADE   = 1 << 20
)

func ConditionText(c uint32) (s string) {
	if c&COND_LIT != 0 {
		s = "Lit,"
	}
	if c&COND_OILY != 0 {
		s = "Oily,"
	}
	if c&COND_FLUID != 0 {
		s = "Fluid,"
	}
	//if c&COND_NOARRR != 0 {
	//	s = "Noarr,"
	//}
	//if c&COND_NOBACK != 0 {
	//	s = "Noback,"
	//}
	if c&COND_ABOVE != 0 {
		s = "Above,"
	}
	if c&COND_DEEP != 0 {
		s = "Deep,"
	}
	if c&COND_FOREST != 0 {
		s = "Forest,"
	}
	if c&COND_FORCED != 0 {
		s = "Forced,"
	}
	if c&COND_CAVE != 0 {
		s = "Cave,"
	}
	if c&COND_BIRD != 0 {
		s = "Bird,"
	}
	if c&COND_SNAKE != 0 {
		s = "Snake,"
	}
	if c&COND_MAZE != 0 {
		s = "Maze,"
	}
	if c&COND_DARK != 0 {
		s = "Dark,"
	}
	if c&COND_WITT != 0 {
		s = "Witt,"
	}
	if c&COND_CLIFF != 0 {
		s = "Cliff,"
	}
	if c&COND_WOODS != 0 {
		s = "Woods,"
	}
	if c&COND_OGRE != 0 {
		s = "Ogre,"
	}
	if c&COND_JADE != 0 {
		s = "Jade,"
	}
	if strings.HasSuffix(s, ",") {
		s = s[:len(s)-1]
		s += "."
	}
	return
}

var conditions = [...]uint32{
	0,                                  // LOC_NOWHERE
	COND_FLUID | COND_ABOVE | COND_LIT, // LOC_START
	COND_ABOVE | COND_LIT,              // LOC_HILL
	COND_FLUID | COND_ABOVE | COND_LIT, // LOC_BUILDING
	COND_FLUID | COND_ABOVE | COND_LIT, // LOC_VALLEY
	COND_ABOVE | COND_LIT,              // LOC_ROADEND
	COND_ABOVE | COND_NOBACK | COND_LIT | COND_CLIFF, // LOC_CLIFF
	COND_FLUID | COND_ABOVE | COND_LIT,               // LOC_SLIT
	COND_ABOVE | COND_LIT | COND_CAVE | COND_JADE,    // LOC_GRATE
	COND_LIT,                            // LOC_BELOWGRATE
	COND_LIT,                            // LOC_COBBLE
	0,                                   // LOC_DEBRIS
	0,                                   // LOC_AWKWARD
	COND_BIRD,                           // LOC_BIRDCHAMBER
	0,                                   // LOC_PITTOP
	COND_DEEP | COND_JADE,               // LOC_MISTHALL
	COND_DEEP,                           // LOC_CRACK
	COND_DEEP,                           // LOC_EASTBANK
	COND_DEEP,                           // LOC_NUGGET
	COND_DEEP | COND_SNAKE,              // LOC_KINGHALL
	COND_DEEP,                           // LOC_NECKBROKE
	COND_DEEP,                           // LOC_NOMAKE
	COND_DEEP,                           // LOC_DOME
	COND_DEEP,                           // LOC_WESTEND
	COND_FLUID | COND_DEEP | COND_OILY,  // LOC_EASTPIT
	COND_DEEP,                           // LOC_WESTPIT
	COND_DEEP,                           // LOC_CLIMBSTALK
	COND_DEEP,                           // LOC_WESTBANK
	COND_DEEP,                           // LOC_FLOORHOLE
	COND_DEEP,                           // LOC_SOUTHSIDE
	COND_DEEP,                           // LOC_WESTSIDE
	COND_DEEP,                           // LOC_BUILDING1
	COND_DEEP,                           // LOC_SNAKEBLOCK
	COND_DEEP,                           // LOC_Y2
	COND_DEEP,                           // LOC_JUMBLE
	COND_DEEP,                           // LOC_WINDOW1
	COND_DEEP,                           // LOC_BROKEN
	COND_DEEP,                           // LOC_SMALLPITBRINK
	COND_FLUID | COND_DEEP,              // LOC_SMALLPIT
	COND_DEEP,                           // LOC_DUSTY
	COND_DEEP,                           // LOC_PARALLEL1
	COND_DEEP,                           // LOC_MISTWEST
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE1
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE2
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE3
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE4
	COND_DEEP | COND_NOARRR | COND_MAZE, // LOC_MAZEEND1
	COND_DEEP | COND_NOARRR | COND_MAZE, // LOC_MAZEEND2
	COND_DEEP | COND_NOARRR | COND_MAZE, // LOC_MAZEEND3
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE5
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE6
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE7
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE8
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE9
	COND_DEEP | COND_NOARRR | COND_MAZE, // LOC_MAZEEND4
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE10
	COND_DEEP | COND_NOARRR | COND_MAZE, // LOC_MAZEEND5
	COND_DEEP | COND_NOBACK,             // LOC_PITBRINK
	COND_NOARRR | COND_DEEP,             // LOC_MAZEEND6
	COND_DEEP,                           // LOC_PARALLEL2
	COND_DEEP,                           // LOC_LONGEAST
	COND_DEEP,                           // LOC_LONGWEST
	COND_DEEP,                           // LOC_CROSSOVER
	COND_DEEP,                           // LOC_DEADEND7
	COND_DEEP | COND_JADE,               // LOC_COMPLEX
	COND_DEEP,                           // LOC_BEDQUILT
	COND_DEEP,                           // LOC_SWISSCHEESE
	COND_DEEP,                           // LOC_EASTEND
	COND_DEEP,                           // LOC_SLAB
	COND_DEEP,                           // LOC_SECRET1
	COND_DEEP,                           // LOC_SECRET2
	COND_DEEP,                           // LOC_THREEJUNCTION
	COND_DEEP,                           // LOC_LOWROOM
	COND_DEEP,                           // LOC_DEADCRAWL
	COND_DEEP,                           // LOC_SECRET3
	COND_DEEP,                           // LOC_WIDEPLACE
	COND_DEEP,                           // LOC_TIGHTPLACE
	COND_DEEP,                           // LOC_TALL
	COND_DEEP,                           // LOC_BOULDERS1
	COND_DEEP,                           // LOC_SEWER
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE11
	COND_DEEP | COND_MAZE,               // LOC_MAZEEND8
	COND_DEEP | COND_NOARRR | COND_MAZE, // LOC_MAZEEND9
	COND_DEEP | COND_NOBACK,             // LOC_ALIKE12
	COND_DEEP | COND_NOBACK,             // LOC_ALIKE13
	COND_NOARRR | COND_DEEP,             // LOC_MAZEEND10
	COND_DEEP | COND_NOARRR | COND_MAZE, // LOC_MAZEEND11
	COND_DEEP | COND_NOBACK | COND_MAZE, // LOC_ALIKE14
	COND_DEEP,                           // LOC_NARROW
	COND_DEEP,                           // LOC_NOCLIMB
	COND_DEEP,                           // LOC_PLANTTOP
	COND_DEEP,                           // LOC_INCLINE
	COND_DEEP,                           // LOC_GIANTROOM
	COND_DEEP,                           // LOC_CAVEIN
	COND_DEEP,                           // LOC_IMMENSE
	COND_FLUID | COND_DEEP,              // LOC_WATERFALL
	COND_DEEP,                           // LOC_SOFTROOM
	COND_DEEP,                           // LOC_ORIENTAL
	COND_DEEP,                           // LOC_MISTY
	COND_DEEP | COND_DARK,               // LOC_ALCOVE
	COND_DEEP | COND_LIT | COND_DARK,    // LOC_PLOVER
	COND_DEEP | COND_DARK,               // LOC_DARKROOM
	COND_DEEP,                           // LOC_ARCHED
	COND_DEEP,                           // LOC_SHELLROOM
	COND_DEEP,                           // LOC_SLOPING1
	COND_DEEP,                           // LOC_CULDESAC
	COND_DEEP,                           // LOC_ANTEROOM
	COND_DEEP | COND_NOBACK,             // LOC_DIFFERENT1
	COND_DEEP | COND_NOBACK | COND_WITT, // LOC_WITTSEND
	COND_DEEP | COND_JADE,               // LOC_MIRRORCANYON
	COND_DEEP,                           // LOC_WINDOW2
	COND_DEEP,                           // LOC_TOPSTALACTITE
	COND_DEEP | COND_NOBACK,             // LOC_DIFFERENT2
	COND_FLUID | COND_DEEP,              // LOC_RESERVOIR
	COND_DEEP,                           // LOC_MAZEEND12
	COND_DEEP | COND_LIT,                // LOC_NE
	COND_DEEP | COND_LIT,                // LOC_SW
	COND_DEEP,                           // LOC_SWCHASM
	COND_DEEP,                           // LOC_WINDING
	COND_DEEP,                           // LOC_SECRET4
	COND_DEEP,                           // LOC_SECRET5
	COND_DEEP,                           // LOC_SECRET6
	COND_NOARRR | COND_DEEP,             // LOC_NECHASM
	COND_NOARRR | COND_DEEP,             // LOC_CORRIDOR
	COND_NOARRR | COND_DEEP,             // LOC_FORK
	COND_NOARRR | COND_DEEP,             // LOC_WARMWALLS
	COND_NOARRR | COND_LIT | COND_DEEP | COND_JADE, // LOC_BREATHTAKING
	COND_NOARRR | COND_DEEP,                        // LOC_BOULDERS2
	COND_NOARRR | COND_DEEP,                        // LOC_LIMESTONE
	COND_NOARRR | COND_DEEP,                        // LOC_BARRENFRONT
	COND_NOARRR | COND_DEEP,                        // LOC_BARRENROOM
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT3
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT4
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT5
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT6
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT7
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT8
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT9
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT10
	COND_DEEP | COND_NOBACK,                        // LOC_DIFFERENT11
	COND_DEEP,                                      // LOC_DEADEND13
	COND_DEEP,                                      // LOC_ROUGHHEWN
	COND_DEEP,                                      // LOC_BADDIRECTION
	COND_DEEP | COND_OGRE,                          // LOC_LARGE
	COND_DEEP,                                      // LOC_STOREROOM
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST1
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST2
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST3
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST4
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST5
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST6
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST7
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST8
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST9
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST10
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST11
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST12
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST13
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST14
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST15
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST16
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST17
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST18
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST19
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST20
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST21
	COND_FOREST | COND_NOBACK | COND_LIT | COND_WOODS, // LOC_FOREST22
	COND_ABOVE | COND_LIT,                             // LOC_LEDGE
	COND_FLUID | COND_DEEP,                            // LOC_RESBOTTOM
	COND_FLUID | COND_DEEP,                            // LOC_RESNORTH
	COND_DEEP,                                         // LOC_TREACHEROUS
	COND_DEEP,                                         // LOC_STEEP
	COND_DEEP,                                         // LOC_CLIFFBASE
	COND_DEEP,                                         // LOC_CLIFFACE
	COND_DEEP,                                         // LOC_FOOTSLIP
	COND_DEEP,                                         // LOC_CLIFFTOP
	COND_DEEP,                                         // LOC_CLIFFLEDGE
	COND_DEEP,                                         // LOC_REACHDEAD
	COND_DEEP,                                         // LOC_GRUESOME
	0,                                                 // LOC_FOOF1
	COND_ABOVE,                                        // LOC_FOOF2
	COND_DEEP,                                         // LOC_FOOF3
	COND_ABOVE,                                        // LOC_FOOF4
	COND_DEEP,                                         // LOC_FOOF5
	COND_DEEP,                                         // LOC_FOOF6
}
