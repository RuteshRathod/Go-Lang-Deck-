package main
import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)
type deck [] string//deck of cards of type string

func newDeck() deck{ // for particular set of cards
	cards:= deck{}

	cardsSuits:=[]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsvalues:=[]string{"Jack","King","Queen","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten"}

	for _,  suits:=range cardsSuits{
		for _, values:= range cardsvalues{
			cards = append(cards, suits +" of " + values)
		}

	}
	return cards
}

func (d deck) print(){ // for printing cards selected
	fmt.Println("I am Inside the Deck")
	for i, card:=range d {
		fmt.Println(i,card)
	}
}

func  deal (d deck, handSize int)(deck, deck){//To get handful of cards from number of set of Cards
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{//conversion of set of carsd of type "Deck" into string
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error{// to store the array of deck into the file , but content is coverted into bytes for write purpose
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck{//to read from the file
	bs, err:= ioutil.ReadFile(filename)
	if err !=nil{
		fmt.Println("Error:-	", err)
		os.Exit(1)
	}
	s:= strings.Split(string(bs),",")
	return deck(s)
}


func (d deck) shuffle(){//to shuffle the number of cards with custom seed number.
source:= rand.NewSource(time.Now().UnixNano())//source of random numbers by giving current time and random it
r:= rand.New(source)//Seed Generation by randomizing it

	for i:= range d {
		newPosition:= r.Intn(len(d)-1)

		d[i], d[newPosition]= d[newPosition],d[i]
	}

}