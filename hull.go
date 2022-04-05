package main

import (
	"fmt"

	"fyne.io/fyne/widget"
)

type hullProperties struct {
	code        string
	hullTons    int
	tons        int
	price       int
	maxHP       int
	armored     bool
	fireControl bool
}

var hull = hullProperties{
	code:        defaultHullCode,
	hullTons:    defaultTons,
	tons:        defaultTons + defaultMaxhardpts,
	price:       defaultHullPrice,
	maxHP:       defaultMaxhardpts,
	armored:     false,
	fireControl: false,
}

var (
	hullSelect          *widget.Select
	fireControlCheck    *widget.Check
	detailHull          *widget.Label = widget.NewLabel("Hull")
	detailMaxHardPoints *widget.Label = widget.NewLabel("Hard Points")
	hullDetails         *widget.Box   = widget.NewVBox(detailHull, detailMaxHardPoints)
)

func (h *hullProperties) init(form *widget.Form, box *widget.Box) {
	hullSelect = widget.NewSelect(hullSelectionCode, nothing)
	hullSelect.Selected = hullSelectionCode[2]
	form.AppendItem(widget.NewFormItem("hull", hullSelect))

	fireControlCheck = widget.NewCheck("Fire Control in all hard points", h.fireControlChanged)
	fireControlCheck.Checked = false
	form.AppendItem(widget.NewFormItem("Fire Control", fireControlCheck))

	box.Children = append(box.Children, hullDetails)

	h.updateHull()
	h.updateHardPoints()
	hullSelect.OnChanged = h.hullChanged
}

func (h *hullProperties) armorChanged(armored bool) {
	hull.armored = armored
}

func (h *hullProperties) hullChanged(value string) {
	index := getIndexFromHull(value)
	if index > -1 {
		hull.code = value
		hull.hullTons = hullSelections[index].tons
		if h.fireControl {
			hull.tons = hull.hullTons + hull.maxHP
		} else {
			hull.tons = hull.hullTons
		}
		hull.price = hullSelections[index].price
		j, m, p := drives.minDrives(value)
		drives.jumpChanged(TrvIndex[j])
		drives.maneuverChanged(TrvIndex[m])
		drives.powerChanged(TrvIndex[p])
	}
	h.updateHull()
}

func (h *hullProperties) fireControlChanged(fc bool) {
	h.fireControl = fc
	if h.fireControl {
		hull.tons = hull.hullTons + hull.maxHP
	} else {
		hull.tons = hull.hullTons
	}
	h.updateHull()
	detailHull.Refresh()
	summary.update()
}

func (h *hullProperties) updateHull() {
	detailHull.SetText(fmt.Sprintf("Hull %s tons: %d, cost: %d MCr", hull.code, hull.hullTons, hull.price))
	detailHull.Refresh()
	summary.update()
}

func (h *hullProperties) updateHardPoints() {
	index := getIndexFromHull(hull.code)
	if index > -1 {
		hull.hullTons = hullSelections[index].tons
		hull.maxHP = hull.tons / 100
		if h.fireControl {
			hull.tons = hull.hullTons + hull.maxHP
			detailMaxHardPoints.SetText(fmt.Sprintf("Hardpoints w/Fire Control x %d, %d tons, %d MCr", hull.maxHP, hull.maxHP, hull.maxHP))
		} else {
			hull.tons = hull.hullTons
			detailMaxHardPoints.SetText(fmt.Sprintf("Maximum Hardpoints (no fire control): %d", hull.maxHP))
		}
		detailMaxHardPoints.Refresh()
		detailMaxHardPoints.Show()
		summary.update()
	}
}

func (h *hullProperties) getTons() int {
	return hull.tons
}

func (h *hullProperties) getMCr() int {
	return hull.price
}
