package main

import (
	"context"
	"log"

	"github.com/fosshostorg/teardrop/api"
	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/migrate"
)

func main() {
	client, err := ent.Open("sqlite3", "file:lmao?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// //testing stuffs
	// proj := client.Project.Create().SetName("teardrop-test").SetGit("https://github.com/JcdeA/website.git").SetDefaultBranch("main").SaveX(context.TODO())
	// client.User.Create().SetName("babo").SetEmail("io@fosshost.org").SetImage("https://avatars.githubusercontent.com/u/31413538?v=4").AddProjects(proj).SaveX(context.TODO())

	// domain := client.Domain.Create().SetDomain("localhost:1323").SetDomain("http://localhost:3000").SetProject(proj).SaveX(context.TODO())
	// client.Deployment.Create().SetAddress("127.0.0.1:3000").SetBranch("main").AddDomains(domain).SaveX(context.TODO())

	api.StartAPI()
}
