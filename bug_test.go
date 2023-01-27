package bug

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
)

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"14": 5432} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=root sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func fakeData(client *ent.Client) {
	ctx := context.Background()
	for i := 1; i < 3; i++ {
		catname := "cat" + strconv.Itoa(i)
		cat, _ := client.Category.Create().SetName(catname).Save(ctx)
		for i := 1; i < 5; i++ {
			post, _ := client.Post.Create().SetName(catname + "_post" + strconv.Itoa(i)).Save(ctx)
			cat.Update().AddPosts(post).Save(ctx)
		}
	}
}

// Convert any struct to JSON String with Pretty Print
func ToJSONIndent(val interface{}) (string, error) {
	b, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func PrintToJSON(val interface{}) {
	b, err := ToJSONIndent(val)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(b)
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	client.Category.Delete().ExecX(ctx)
	client.Post.Delete().ExecX(ctx)

	fakeData(client)

	categories, _ := client.Category.Query().WithPosts(func(pq *ent.PostQuery) {
		pq.Limit(3)
	}).All(ctx)

	for _, category := range categories {
		fmt.Println(len(category.Edges.Posts))
		if n := len(category.Edges.Posts); n != 3 {
			t.Errorf("expected 3 posts but got: %d", n)
		}
	}
	// PrintToJSON(categories)
}
