package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/JavascriptDev347/social.git/internal/store"
)

var usernames = []string{
	"maya", "elliot", "nora", "caleb", "ivy", "felix", "luna", "theo", "zara", "miles",
	"aria", "jasper", "ruby", "silas", "clara", "dylan", "freya", "hugo", "iris", "kai",
	"lena", "mason", "nina", "oscar", "piper", "quinn", "rhea", "sean", "tessa", "victor",
	"gopher_tom", "nil_pointer", "chan_master", "goroutine_gal", "defer_dan", "slice_sam",
	"mutex_mike", "panic_at_dawn", "byte_baker", "stack_trace", "heap_hero", "null_ninja",
	"debug_duck", "commit_carl", "merge_mia", "rebase_rex", "binary_bob", "kernel_kate",
	"socket_sid", "cache_cat",
}

var titles = []string{
	"What Really Happens When You Press a Key",
	"Inside the Machine That Keeps Time",
	"The Hidden Physics of Flight",
	"How Your Phone Knows Where You Are",
	"Why Bridges Don't Fall Down",
	"The Strange Math Behind Compression",
	"How a Microwave Heats Food",
	"What Happens When You Type a URL",
	"The Engineering of a Skyscraper",
	"How Noise-Cancelling Headphones Work",
	"The Secret Life of a Hard Drive",
	"Why Batteries Slowly Die",
	"How Touch Screens Feel Your Finger",
	"The Truth About Cloud Storage",
	"How Elevators Stay Safe",
	"What Makes a Lightbulb Glow",
	"How Submarines Dive and Surface",
	"The Science of Fingerprint Scanners",
	"How Airplanes Navigate Over the Ocean",
	"Why Your Fridge Hums at Night",
}

var contents = []string{
	"One small detail explains the whole thing.",
	"The real mechanism is far simpler than you would guess.",
	"Engineers spent decades getting this right.",
	"Most people walk past this every single day.",
	"Behind the scenes, a clever trick does all the work.",
	"It seems like magic until you see the parts.",
	"A single design choice made all the difference.",
	"This idea started as a happy accident.",
	"The trick is hiding where nobody looks.",
	"Break it down, and it makes perfect sense.",
	"What looks complicated is just simple steps stacked together.",
	"The history of this invention is wilder than expected.",
	"Without this, modern life would grind to a halt.",
	"It works because of one elegant principle.",
	"Here is the part that textbooks usually skip.",
	"A tiny component carries most of the load.",
	"The solution was obvious only after someone found it.",
	"Think of it as a chain reaction with a purpose.",
	"This is the quiet technology you rely on constantly.",
	"Once you see it, you cannot unsee it.",
}

var tags = []string{
	"tech", "sciencefacts", "explainer", "engineeringlife", "mechanics", "stem",
	"funfacts", "behindthescenes", "deepdive", "curious", "inventions", "physics",
	"gadgets", "futuretech", "learneveryday", "mindblown", "howthingswork",
	"smartliving", "discovery", "shortsvideo",
}

var comments = []string{
	"Great breakdown, this finally clicked for me.",
	"I have wondered about this for years.",
	"Honestly one of the clearest explanations I have seen.",
	"Never thought about it this way before.",
	"Please do a follow-up on this topic!",
	"Wow, engineers are on another level.",
	"I am going to send this to my whole class.",
	"This channel keeps getting better.",
	"So that is why it behaves like that!",
	"Simple, clear, and actually interesting.",
	"My brain just did a little flip 🤯",
	"I wish my teachers explained things like this.",
	"Can you cover how the next step works too?",
	"Learned more in one minute than in a whole lecture.",
	"This changes how I look at everyday stuff.",
	"Underrated content, more people need to see this.",
	"Okay, that was genuinely clever.",
	"I rewatched this twice, so good.",
	"The visuals made it so easy to follow.",
	"Subscribed just for explanations like this.",
}

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
