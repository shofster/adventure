package display

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"time"
)

/*
*  Copyright (c) 2022 by Robert (Bob) Shofner
*  SPDX-License-Identifier: BSD-2-clause

 */
/*
  Description: Set of functions to simulate a 2 pane display device.
*/
type Display struct {
	Content  *fyne.Container
	upper    *fyne.Container
	lower    *fyne.Container
	location Pane
	picture  Pane
}
type Pane struct {
	container *fyne.Container
	channel   chan *canvas.Image
	duration  time.Duration
}

func NewDisplay() *Display {
	display := Display{
		upper: container.NewMax(widget.NewLabel("LOCATION")),
		lower: container.NewMax(widget.NewLabel("PICTURE"))}
	display.Content = container.NewMax(container.NewVSplit(display.upper, display.lower))
	display.location = NewPane(display.upper, time.Duration(0))
	display.picture = NewPane(display.lower, time.Duration(500*time.Millisecond))
	return &display
}
func NewPane(container *fyne.Container, duration time.Duration) Pane {
	p := Pane{container: container, duration: duration}
	p.channel = make(chan *canvas.Image, 8)
	go func() {
		for object := range p.channel {
			container.Objects = make([]fyne.CanvasObject, 1)
			container.Objects[0] = object
			container.Refresh()
			if p.duration != 0 {
				time.Sleep(p.duration)
			}
		}
	}()
	return p
}
func (p *Pane) Show(object *canvas.Image) {
	p.channel <- object
}

// unusedPicFile. if a file is needed
func unusedPicFile() *canvas.Image {
	return canvas.NewImageFromFile("path_to_png or jpg")
}
func (d *Display) Show(top bool, pic string) {
	picure, ok := pictures[pic]
	if !ok {
		return
	}
	switch pictures[pic].picType {
	case picFile:
		if _, err := os.Stat(""); err == nil {
			image := picure.picFunc()
			image.FillMode = canvas.ImageFillContain
			if top {
				d.location.Show(image)
				d.Show(false, "nsewud")
			} else {
				d.picture.Show(image)
			}
		}
	case picResource:
		image := picure.picFunc()
		image.FillMode = canvas.ImageFillContain
		if top {
			d.location.Show(image)
			d.Show(false, "nsewud")
		} else {
			d.picture.Show(image)
		}
	}
}

const ( // picture type
	picFile = iota
	picResource
)

type picture struct {
	picType int
	picFunc func() *canvas.Image
}

var pictures = map[string]picture{
	// locations - "nsewud" is the default
	"nsewud": {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceNsewudJpg) }},
	"map":    {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceMapJpg) }},
	"poof":   {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePoofJpg) }},
	// Object sounds
	"bird":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBirdJpg) }},
	"clam":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceClamJpg) }},
	"dragon":     {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceDragonJpg) }},
	"oyster":     {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceOysterJpg) }},
	"pirate":     {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePirateJpg) }},
	"snake":      {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceSnakeJpg) }},
	"troll":      {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceTrollJpg) }},
	"word_xyzzy": {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceWordxyzzyJpg) }},
	"word_plugh": {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceWordplughJpg) }},
	// Object descriptions
	"amber":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceAmberJpg) }},
	"axe":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceAxeJpg) }},
	"battery":        {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBatteryJpg) }},
	"bear":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBearJpg) }},
	"bear_sitting":   {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBearsittingJpg) }},
	"bear_wandering": {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBearwanderingJpg) }},
	"blood":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBloodJpg) }},
	"bottle":         {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBottleJpg) }},
	"cage":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCageJpg) }},
	"cage_bird":      {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCagebirdJpg) }},
	"chain":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceChainJpg) }},
	"coins":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCoinsJpg) }},
	"chest":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceChestJpg) }},
	"diamonds":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceDiamondsJpg) }},
	"door":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceDoorJpg) }},
	"dwarf":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceDwarfJpg) }},
	"eggs":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceEggsJpg) }},
	"emerald":        {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceEmeraldJpg) }},
	"food":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceFoodJpg) }},
	"gold":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceGoldJpg) }},
	"gnome":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceGnomeJpg) }},
	"jade":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceJadeJpg) }},
	"jewels":         {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceJewelsJpg) }},
	"keys":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceKeysJpg) }},
	"knife":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceKnifeJpg) }},
	"lamp_off":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceLampoffJpg) }},
	"lamp_on":        {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceLamponJpg) }},
	"pearl":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePearlJpg) }},
	"pillow":         {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePillowJpg) }},
	"pyramid":        {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePyramidJpg) }},
	"rod_star":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRodstarJpg) }},
	"rod_mark":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRodmarkJpg) }},
	"ruby":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRubyJpg) }},
	"rug":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRugJpg) }},
	"sapphire":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceSapphireJpg) }},
	"silver_bars":    {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceSilverbarsJpg) }},
	"spices":         {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceSpicesJpg) }},
	"statue":         {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceStatueJpg) }},
	"steps":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceStepsJpg) }},
	"trident":        {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceTridentJpg) }},
	"urn":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceUrnJpg) }},
	"vase":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceVaseJpg) }},
	// locations
	"bedquilt":             {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBedquiltJpg) }},
	"building":             {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBuildingJpg) }},
	"dead_end":             {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceDeadendJpg) }},
	"forest":               {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceForestJpg) }},
	"grate":                {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceGrateJpg) }},
	"plover":               {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePloverJpg) }},
	"reservoir":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceReservoirJpg) }},
	"valley":               {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceValleyJpg) }},
	"xyzzy":                {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceXyzzyJpg) }},
	"y2":                   {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceY2Jpg) }},
	"bird_chamber":         {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceChamberbirdJpg) }},
	"brink_pit":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBrinkpitJpg) }},
	"bridge_crystal":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBridgecrystalJpg) }},
	"bridge_troll":         {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceBridgetrollJpg) }},
	"canyon_junction":      {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCanyonjunctionJpg) }},
	"canyon_mirror_secret": {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCanyonmirrorsecretJpg) }},
	"canyon_tall":          {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCanyontallJpg) }},
	"cavern_magnificent":   {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCavernmagnificentJpg) }},
	"chamber_grate":        {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceChambergrateJpg) }},
	"corridor_sloping":     {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceCorridorslopingJpg) }},
	"hall_long":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceHalllongJpg) }},
	"hall_mists_east":      {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceHallmistseastJpg) }},
	"hall_mists_west":      {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceHallmistswestJpg) }},
	"hall_mountain_king":   {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceHallmountainkingJpg) }},
	"maze_alike":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceMazealikeJpg) }},
	"maze_different":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceMazedifferentJpg) }},
	"passage_broken":       {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePassagebrokenJpg) }},
	"path_fork":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourcePathforkJpg) }},
	"room_ante":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomanteJpg) }},
	"room_baren":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoombarenJpg) }},
	"room_dark":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomdarkJpg) }},
	"room_giant":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomgiantJpg) }},
	"room_oriental":        {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomorientalJpg) }},
	"room_shell":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomshellJpg) }},
	"room_slab":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomslabJpg) }},
	"room_soft":            {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomsoftJpg) }},
	"room_swiss":           {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceRoomswissJpg) }},
	"top_pit":              {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceToppitJpg) }},
	"two_pit":              {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceTwopitJpg) }},
	"volcano":              {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceVolcanoJpg) }},
	"volcano_boulders":     {picType: picResource, picFunc: func() *canvas.Image { return canvas.NewImageFromResource(resourceVolcanoboldersJpg) }},
}
