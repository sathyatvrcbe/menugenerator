package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	kootuCategory := Category{
	    name: "Kootu",
	    items: []Item{
	    	Item{"Kootu1"},
			Item{"Kootu2"},
			Item{"Kootu3"},
	    },
	}
	kootuSelector := OneItemSelector{kootuCategory}

	poriyalCategory := Category{
	    name: "Porial",
	    items: []Item{
	    	Item{"porial1"},
			Item{"Porial2"},
			Item{"Porial3"},
	    },
	}
	poriyalSelector := OneItemSelector{poriyalCategory}

	varuvalCategory := Category{
	    name: "Varuval",
	    items: []Item{
	    	Item{"varuval1"},
			Item{"varuval2"},
			Item{"varuval3"},
	    },
	}

	varuvalSelector := OneItemSelector{varuvalCategory}

	sideDishSelector := CategorySelector{
		categoriesToSelect: 2,
		childSelectors: []Selector{
			kootuSelector,
			poriyalSelector,
			varuvalSelector,
		},
	}

	gravyCategory := Category{
		name: "Gravy",
		items: []Item{
			Item{"GravyItem1"},
			Item{"GravyItem2"},
			Item{"GravyItem3"},
		},
	}

	gravySelector := OneItemSelector{gravyCategory}

	lunchSelector := CategorySelector{
		categoriesToSelect: 2,
		childSelectors: []Selector{
			gravySelector,
			sideDishSelector,
		},
	}

	breakfastItems := Category{
	    name: "Breakfast",
	    items: []Item{
	    	Item{"breakfast1"},
			Item{"breakfast2"},
			Item{"breakfast3"},
	    },
	}
	breakfastSelector := OneItemSelector{breakfastItems}

	chutneyItems := Category{
		name: "Chutneys",
		items: []Item{
			Item{"chutney1"},
			Item{"chutney2"},
			Item{"chutney3"},
		},
	}
	chutneySelector := OneItemSelector{chutneyItems}

	breakfastDinnerSelector := CategorySelector{
		categoriesToSelect: 2,
		childSelectors: []Selector{
			breakfastSelector,
			chutneySelector,
		},
	}

	startDate := "1999-12-31"
	t, _ := time.Parse("2006-01-02", startDate)
	
	fmt.Println("## Menu")
	 for i := 1; i <= 7; i++ {
		fmt.Println("###", t.Format("January 2, 2006"))
		fmt.Println()
		fmt.Println("Breakfast | Lunch | Dinner |")
		fmt.Println("|---|---|---|")
		b := getItemsString(breakfastDinnerSelector.Select())
		l := getItemsString(lunchSelector.Select())
		d := getItemsString(breakfastDinnerSelector.Select())
		fmt.Println(b+" | "+l+" | "+d+" |")
		fmt.Println()
		t = t.AddDate(0,0,1)
	}
	
}

type Item struct{
	name string
}

func (i Item) String() string{
	return i.name
}

func getItemsString(items []Item) string{
	ret := ""
	for _, item := range items {
		ret = ret + item.name + ", "
	}
	return ret[0:len(ret)-2]
}

type Category struct{
	name string
	items []Item
}

type Selector interface{
	Select() []Item
}

type OneItemSelector struct{
	category Category
}

func (s OneItemSelector) Select() []Item {
    n := rand.Intn(len(s.category.items))
	return []Item{s.category.items[n]}
}

type CategorySelector struct{
	categoriesToSelect int
	childSelectors []Selector
}

func (s CategorySelector) Select() []Item {
	selections := make(map[int]int)
	for len(selections)<s.categoriesToSelect {
	    n := rand.Intn(len(s.childSelectors))
	    _, ok := selections[n]
	    if !ok {
	       selections[n]=0
	    }
	}
	
	var itemsToReturn []Item
	for key,_ := range selections {
	    itemsToReturn = append(itemsToReturn, s.childSelectors[key].Select()...)
	}
	return itemsToReturn
}

