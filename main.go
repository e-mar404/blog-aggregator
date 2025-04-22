package main

import (
	"database/sql"
	"e-mar404/gator/internal/config"
	"e-mar404/gator/internal/database"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", c.DBURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	s := &state {
		config: c,
		db: dbQueries,
	}

	cmds := commands {
		available: newHandlerList(),	
	}
	
	cmds.register("login", handelerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handerlAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("there should be atleast 2: gator <arg>\n")
		os.Exit(1)
	}

	command := command {
		name: args[1],
		arguments: args[2:],
	}
	if err := cmds.run(s, command); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
