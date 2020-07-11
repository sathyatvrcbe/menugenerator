package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	kootuCategory := Category{
	    name: "Kootu",
	    items: []Item{
	    	Item{"கீரை கூட்டு"},
			Item{"கத்தரிக்காய் கூட்டு"},
			Item{"வாழைக்காய் கூட்டு "},
			Item{"சுரைக்காய் கூட்டு"},
			Item{"முட்டைகோஸ் கூட்டு"},
			Item{"பீர்க்கங்காய்  கூட்டு"},
			Item{"வெள்ளரிக்காய் கூட்டு"},
			Item{"புடலங்காய் கூட்டு"},
			Item{"பாகற்காய் கூட்டு"},
			Item{"பரங்கிக்காய் கூட்டு"},
	    },
	}
	kootuSelector := OneItemSelector{kootuCategory}

	poriyalCategory := Category{
	    name: "Porial",
	    items: []Item{
	    	Item{"வாழைக்காய் பொரியல்"},
			Item{"கேரட் பீன்ஸ் பொரியல்"},
			Item{"முட்டைகோஸ் பொரியல்"},
			Item{"புடலங்காய் பொரியல்"},
			Item{"பயத்தங்காய் பொரியல்"},
	    },
	}
	poriyalSelector := OneItemSelector{poriyalCategory}

	varuvalCategory := Category{
	    name: "Varuval",
	    items: []Item{
	    	Item{"கத்தரிக்காய் வறுவல்"},
			Item{"வாழைக்காய் வறுவல்"},
			Item{"கேரட் வறுவல்"},
			Item{"பீன்ஸ் வறுவல்"},
			Item{"முட்டைகோஸ் வறுவல்"},
			Item{"கோவைக்காய் வறுவல்"},
			Item{"கொத்தரங்காய் வறுவல்"},
			Item{"பாகற்காய் வறுவல்"},
			Item{"சேப்பங்கிழங்கு வறுவல்"},
			Item{"பரங்கிக்காய் வறுவல்"},
			Item{"ஆட்டுக்கறி வறுவல்"},
			Item{"கோழி வறுவல்"},
			Item{"முட்டை வறுவல்"},

			
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
			Item{"மோர் குழம்பு"},
			Item{"பூண்டு குழம்பு "},
			Item{"வெந்திய குழம்பு"},
			Item{"வெங்காய குழம்பு"},
			Item{"வத்தல் குழம்பு"},
			Item{"சாம்பார்"},
			Item{"ரசம்"},
			Item{"உருண்டை குழம்பு"},
			Item{"பொரிச்ச குழம்பு"},
			Item{"முட்டை குழம்பு"},
			Item{"ஆட்டுக்கறி குழம்பு"},
			Item{"கோழி குழம்பு"},
			Item{"தொக்கு"},
			Item{"புளி குழம்பு"},
			Item{"பிரியாணி"},
			Item{"புலாவ்"},
			Item{"புதினா சாதம்"},
			Item{"கொத்தமல்லி சாதம்"},
			Item{"தேங்காய் சாதம்"},
			Item{"தக்காளி சாதம்"},
			Item{"மிளகு சாதம்"},
			Item{"தொக்கு"},

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
	    	Item{"வரகு உப்புமா"},
			Item{"சாமை உப்புமா"},
			Item{"குதிரைவாலி உப்புமா"},
			Item{"ரவா உப்புமா"},
			Item{"கோதுமை ரவா உப்புமா"},
			Item{"சேமியா உப்புமா"},
			Item{"சிகப்பு அவல் உப்புமா "},
			Item{"கம்பு அடை"},
			Item{"பருப்பு அடை"},
			Item{"கீரை அடை"},
			Item{"அரிசி அடை"},
			Item{"கேப்பை அடை"},
			Item{"ஆப்பம்"},
			Item{"பாசிப்பயறு தோசை"},
			Item{"கம்பு தோசை"},
			Item{"கேழ்வரகு தோசை"},
			Item{"கீரை அடைதோசை"},
			Item{"முடக்கத்தான் தோசை"},
			Item{"ரவா இட்லி"},
			Item{"மாவு இட்லி"},
			Item{"பணியாரம்"},
			Item{"கேழ்வரகு புட்டு"},
			Item{"சிவப்பரிசி புட்டு"},
			Item{"கம்பு புட்டு"},
			Item{"குதிரைவாலி பொங்கல்"},
			Item{"வரகு பொங்கல்"},
			Item{"சாமை பொங்கல்"},
			Item{"திணை பொங்கல்"},
			Item{"அரிசி பொங்கல்"},
			Item{"குதிரைவாலி கஞ்சி"},
			Item{"வரகு கஞ்சி"},
			Item{"சாமை கஞ்சி"},
			Item{"stuffed சப்பாத்தி"},
			Item{"கீரை சப்பாத்தி"},
			Item{"இடியப்பம்"},
			Item{"ரவா கிச்சடி"},
			Item{"சேமியா கிச்சடி"},
			Item{"கேழ்வரகு களி"},
			Item{"உளுந்து களி"},
			
	    },
	}
	breakfastSelector := OneItemSelector{breakfastItems}

	chutneyItems := Category{
		name: "Chutneys",
		items: []Item{
			Item{"உளுந்து சட்னி"},
			Item{"பூண்டு சட்னி"},
			Item{"தேங்காய் சட்னி"},
			Item{"தக்காளி சட்னி"},
			Item{"கொத்தமல்லி சட்னி"},
			Item{"புதினா சட்னி"},
			Item{"பீர்க்கங்காய் சட்னி"},
			Item{"கலவை சட்னி"},
			Item{"முட்டைகோஸ் சட்னி"},
			Item{"கத்தரிக்காய் சட்னி"},
			Item{"வெங்காய சட்னி"},
			Item{"பிரண்டை சட்னி"},
			Item{"வெள்ளை குருமா"},
			Item{"சப்பாத்தி குர்மா"},
			Item{"இட்லி சாம்பார்"},
			
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

	startDate := "2020-07-14"
	t, _ := time.Parse("2006-01-02", startDate)
	
	f, err := os.Create("out.md")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	 for i := 1; i <= 7; i++ {
		f.WriteString("#### "+ t.Format("January 2, 2006"))
		f.WriteString("\n")
		f.WriteString("காலை | மதியம் | இரவு |\n")
		f.WriteString("|---|---|---|\n")
		b := getItemsString(breakfastDinnerSelector.Select())
		l := getItemsString(lunchSelector.Select())
		d := getItemsString(breakfastDinnerSelector.Select())
		f.WriteString(b+" | "+l+" | "+d+" |\n\n")
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

