package main

import "fmt"

func main() {
	var (
		name, lastname string
		age            int

		name2, lastname2 string
		age2             int
	)

	name, lastname, age = "pablo", "picasso", 91
	name2, lastname2, age2 = "sigmund", "freud", 83

	fmt.Println("Picasso:", name, lastname, age)
	fmt.Println("Freud:", name2, lastname2, age2)

	type person struct {
		name, lastname string
		age            int
	}
	// picasso := person{name: "Pablo", lastname: "Picasso", age: 91}

	picasso := person{
		name:     "Pablo",
		lastname: "Picasso",
		age:      91,
	}
	var freud person

	freud.name = "Sigmund"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("\n%s's age is %d\n",
		picasso.name, picasso.age)

	fmt.Printf("\nPicasso:  %#v\n", picasso)
	fmt.Printf("\nFreud:  %#v\n", freud)

	//Struct comparisons and assignments
	type song struct {
		title, artist string
	}
	type playlst struct {
		genre string
		songs []song
	}

	songs := []song{
		{title: "wonderwall", artist: "oasis"},
		{title: "super sonic", artist: "oasis"},
	}

	rock := playlst{genre: "indie rock", songs: songs}

	rock.songs[0].title = "live forever"

	fmt.Printf("\n%-20s %20s\n", "title", "artist")
	for _, s := range rock.songs {

		//s : rock.songs[i]
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)

	}

	fmt.Printf("\n%-20s %20s\n", "title", "artist")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)

	}

	//struct object oriented programming
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn  string
		title string
	}

	moby := book{
		text: text{title: "moby dick", words: 206052},
		isbn: "102030"}

	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.title, moby.words, moby.isbn)

}
