package main

import (
	"fmt"
	"golang.conradwood.net/apis/auth"
)

// given a set of people, will optimize minimum numbers of groups required to match all

type GroupOptimizer struct {
	groupAll  *auth.Group
	allGroups []*auth.Group
	Users     []*auth.User
}

type result struct {
}

func (gm *GroupOptimizer) Groups() []*auth.Group {
	var res []*auth.Group
	var all []*auth.Group
	// build up sum of all groups of all users:
	for _, c := range gm.Users {
		for _, g := range c.Groups {
			if g == nil {
				continue
			}
			if g.ID == "all" {
				gm.groupAll = g
				continue
			}
			found := false
			for _, a := range all {
				if a.ID == g.ID {
					found = true
					break
				}
			}
			if !found {
				all = append(all, g)
			}
		}
	}
	gm.allGroups = all
	// now find the group(s) which cover the highest amount of users in the least amount of groups

	// surely, there's a better algorithm out there...
	tg := gm.Group("34") // dev
	if gm.Covered(tg) {
		return tg
	}
	tg = gm.Group("36") // guru staff
	if gm.Covered(tg) {
		return tg
	}
	for i := 0; i < len(all); i++ {
		g := gm.GetGroups(i)
		if !gm.Covered(g) {
			continue
		}
		res = g
		break
	}
	if len(res) == 0 {
		fmt.Printf("Found no groups to match users:\n")
		for _, u := range gm.Users {
			fmt.Printf("   %s %s\n", u.ID, u.Email)
		}
		return []*auth.Group{&auth.Group{ID: "all"}}
	}
	fmt.Printf("Users groups:\n")
	for _, r := range res {
		fmt.Printf("   Group: %s\n", r)
	}
	return res
}
func (gm *GroupOptimizer) Covered(gr []*auth.Group) bool {
	for _, u := range gm.Users {
		covered := false
		for _, g := range gr {
			for _, ug := range u.Groups {
				if ug.ID == g.ID {
					covered = true
				}
			}
		}
		if !covered {
			return false
		}
	}
	return true
}
func (gm *GroupOptimizer) GetGroups(x int) []*auth.Group {
	return gm.allGroups[:x]
}
func (gm *GroupOptimizer) Group(id string) []*auth.Group {
	for _, g := range gm.allGroups {
		if g.ID == id {
			return []*auth.Group{g}
		}
	}
	return nil
}
