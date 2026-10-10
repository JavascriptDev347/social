package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/JavascriptDev347/social.git/internal/store"
)

var usernames = []string{
	"alex", "jordan", "michael", "david", "daniel", "james", "robert", "william", "chris", "matthew", "andrew", "ryan", "kevin", "brian", "jason", "ethan", "noah", "liam", "oliver", "lucas", "jack", "henry", "benjamin", "samuel", "charlie", "leo", "max", "owen", "gopher", "go_dev", "golangdev", "coder123", "devpro", "hacklab", "techwolf", "codefox", "devmaster", "programmer", "code_master", "techguy", "dev_king", "bytehunter", "codewizard", "ninja_dev", "cyberfox", "pixelmaster", "darkcoder", "webmaster", "programming", "developer"}

var titles = []string{"The Future of AI", "How the Internet Works", "Why Planes Can Fly", "The Secret of Black Holes", "How GPS Finds You", "Why We Dream", "The Science of Sleep", "How Credit Cards Work", "Inside a Nuclear Reactor", "How Rockets Land", "The Mystery of Time", "How Wi-Fi Works", "Why the Sky Is Blue", "How Electric Cars Work", "The Secret of QR Codes", "How Your Brain Learns", "Why Ships Don't Sink", "How Refrigerators Work", "The Science Behind Rain", "How Satellites Stay in Orbit"}
var contents = []string{"A strange discovery changed everything.", "Nobody expected this to happen.", "Here is what actually happens behind the scenes.", "This simple idea is more powerful than it looks.", "The science behind this is surprisingly simple.", "You use this every day without knowing how it works.", "This invention completely changed the world.", "There is a hidden system working inside it.", "Most people have never thought about this.", "Here is the secret nobody tells you.", "It looks simple, but the technology is incredible.", "One tiny mistake can cause a huge problem.", "This is how engineers solved the impossible.", "The answer is hiding in plain sight.", "What happens next is surprisingly clever.", "This everyday object has a fascinating story.", "Scientists discovered something unexpected.", "The technology behind this is everywhere.", "It took years to figure this out.", "Now you finally know how it works."}

var tags = []string{"technology", "science", "howitworks", "engineering", "facts", "education", "interestingfacts", "didyouknow", "explained", "technologyexplained", "scienceexplained", "engineeringexplained", "curiosity", "knowledge", "innovation", "inventions", "learning", "educational", "viral", "shorts"}
var comments = []string{"This is actually fascinating!", "I never knew this worked like that.", "The technology behind this is crazy.", "Why did nobody explain this before?", "That makes so much sense now.", "I learned something new today.", "The more you know!", "This is way more complicated than I thought.", "Mind blown 🤯", "I need to see how this works in real life.", "Science is amazing.", "That explanation was actually really clear.", "I had no idea about this.", "Now I can't stop thinking about it.", "This deserves way more views.", "The engineering behind this is incredible.", "Wait... so THAT'S how it works?", "I use this every day and never knew this.", "Okay, that's actually pretty cool.", "Part 2 please!"}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.User.Create(ctx, user); err != nil {
			log.Println("Error creating user")
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Post.Create(ctx, post); err != nil {
			log.Println("Error creating post")
		}
	}

	comments := generateComments(500, posts, users)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment")
		}
	}

	log.Print("Seeding completed")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123321",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Content: contents[rand.Intn(len(contents))],
			Title:   titles[rand.Intn(len(titles))],
			Tags:    []string{tags[rand.Intn(len(tags))], tags[rand.Intn(len(tags))]},
		}
	}
	return posts
}

func generateComments(num int, posts []*store.Post, users []*store.User) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
