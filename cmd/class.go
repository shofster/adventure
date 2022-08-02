package cmd

/*
 * Copyright (c) 1977, 2005 by Will Crowther and Don Woods
 * Copyright (c) 2017 by Eric S. Raymond
 * Translate to GO by Robert Shofner (c) 2022
 * BSD License - see file "COPYING"
 */

var classes = [...]Class{
	{threshold: 0, message: ""},
	{threshold: 45, message: "You are obviously a rank amateur.  Better luck next time."},
	{threshold: 120, message: "Your score qualifies you as a Novice class adventurer."},
	{threshold: 170, message: "You have achieved the rating: \"Experienced Adventurer\"."},
	{threshold: 250, message: "You may now consider yourself a \"Seasoned Adventurer\"."},
	{threshold: 320, message: "You have reached \"Junior Master\" status."},
	{threshold: 375, message: "Your score puts you in Master Adventurer Class C."},
	{threshold: 410, message: "Your score puts you in Master Adventurer Class B."},
	{threshold: 426, message: "Your score puts you in Master Adventurer Class A."},
	{threshold: 429, message: "All of Adventuredom gives tribute to you, Adventurer Grandmaster!"},
	{threshold: 9999, message: "'Adventuredom stands in awe -- you have now joined the ranks of the\n       W O R L D   C H A M P I O N   A D V E N T U R E R S !\nIt may interest you to know that the Dungeon-Master himself has, to\nmy knowledge, never achieved this threshold in fewer than 330 Turns.'"},
}
var hints = [...]Hint{
	{number: 1, turns: 2, penalty: 4, question: "Are you trying to get into the cave?", hint: "The grate is very solid and has a hardened steel lock.  You cannot\nenter without a key, and there are no keys nearby.  I would recommend\nlooking elsewhere for the keys."},
	{number: 2, turns: 2, penalty: 5, question: "Are you trying to catch the bird?", hint: "Something about you seems to be frightening the bird.  Perhaps you\nmight figure out what it is."},
	{number: 3, turns: 2, penalty: 8, question: "Are you trying to somehow deal with the snake?", hint: "You can't kill the snake, or drive it away, or avoid it, or anything\nlike that.  There is a way to get by, but you don't have the necessary\nresources right now."},
	{number: 4, turns: 4, penalty: 75, question: "Do you need help getting out of the maze?", hint: "You can make the passages look less alike by dropping things."},
	{number: 5, turns: 5, penalty: 25, question: "Are you trying to explore beyond the plover room?", hint: "here is a way to explore that region without having to worry about\nfalling into a pit.  None of the objects available is immediately\nuseful in discovering the secret."},
	{number: 6, turns: 3, penalty: 20, question: "Do you need help getting out of here?", hint: "Don't go west."},
	{number: 7, turns: 2, penalty: 8, question: "Are you wondering what to do here?", hint: "This section is quite advanced.  Find the cave first."},
	{number: 8, turns: 2, penalty: 25, question: "Would you like to be shown out of the forest?", hint: "Go east ten times.  If that doesn't get you out, then go south, then\nwest twice, then south."},
	{number: 9, turns: 4, penalty: 10, question: "Do you need help dealing with the ogre?", hint: "There is nothing the presence of which will prevent you from defeating\nhim; thus it can't hurt to fetch everything you possibly can."},
	{number: 10, turns: 4, penalty: 1, question: "You're missing only one other treasure.  Do you need help finding it?", hint: "Once you've found all the other treasures, it is no longer possible to\nlocate the one you're now missing."},
}
var thresholds = [...]Threshold{
	{threshold: 350, loss: 2, message: "Tsk!  A wizard wouldn't have to take 350 Turns.  This is going to cost\nyou a couple of points."},
	{threshold: 500, loss: 3, message: "500 Turns?  That's another few points you've lost."},
	{threshold: 1000, loss: 5, message: "Are you still at it?  Five points off for exceeding 1000 Turns!"},
	{threshold: 2500, loss: 10, message: "Good grief, don't you *EVER* give up?  Do you realize you've spent\nover 2500 Turns at this?  That's another ten points off, a total of\ntwenty points lost for taking so long."},
}
var obituaries = []Obituary{
	{query: "Oh dear, you seem to have gotten yourself killed.  I might be able to\nhelp you out, but I've never really done this before.  Do you want me\nto try to reincarnate you?",
		response: "All right.  But don't blame me if something goes wr......\n                    --- POOF!! ---\nYou are engulfed in a cloud of orange smoke.  Coughing and gasping,\nyou emerge from the smoke and find...."},
	{query: "You clumsy oaf, you\\'ve done it again!  I don't know how long I can\\nkeep this up.  Do you want me to try reincarnating you again?",
		response: "Okay, now where did I put my orange smoke?....  >POOF!<\nEverything disappears in a dense cloud of orange smoke."},
	{query: "Now you've really done it!  I'm out of orange smoke!  You don't expect\nme to do a decent reincarnation without any orange smoke, do you?",
		response: "Okay, if you're so smart, do it yourself!  I'm leaving!"},
}
var keys = []int{
	0, 1, 16, 23, 30, 41, 49, 53, 68, 80,
	89, 97, 110, 119, 127, 138, 154, 155, 164, 167,
	180, 181, 182, 183, 190, 192, 196, 197, 206, 213,
	216, 221, 223, 224, 232, 235, 238, 244, 249, 259,
	265, 266, 274, 279, 282, 286, 292, 294, 296, 298,
	300, 304, 308, 314, 317, 319, 323, 325, 331, 333,
	334, 341, 344, 348, 350, 358, 370, 377, 382, 386,
	392, 396, 399, 404, 407, 411, 413, 414, 418, 419,
	420, 424, 426, 428, 431, 434, 436, 438, 440, 446,
	447, 448, 453, 456, 459, 466, 470, 472, 478, 481,
	486, 494, 497, 500, 506, 509, 512, 515, 525, 536,
	539, 541, 547, 557, 563, 564, 565, 567, 576, 578,
	582, 584, 588, 598, 603, 611, 617, 623, 628, 635,
	643, 647, 657, 667, 677, 687, 697, 707, 717, 727,
	737, 741, 743, 744, 747, 749, 753, 757, 761, 765,
	769, 773, 777, 781, 785, 789, 793, 797, 801, 805,
	809, 813, 817, 821, 825, 829, 833, 837, 838, 840,
	847, 851, 855, 859, 862, 863, 864, 868, 871, 872,
	873, 874, 875, 876, 877,
}
