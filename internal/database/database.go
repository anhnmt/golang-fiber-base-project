package database

import (
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent/migrate"
)

var (
	once   *sync.Once
	client *ent.Client
)

func Connection(conf *config.DefaultConfig) *ent.Client {
	if client == nil {
		once = &sync.Once{}

		once.Do(func() {
			if client == nil {
				client = Client(conf)
			}
		})
	}

	return client
}

func Client(conf *config.DefaultConfig) *ent.Client {
	log.Printf("Connect to [%s] %s", conf.DBDriver, conf.DBUrl)

	conn, err := ent.Open(conf.DBDriver, conf.DBUrl)

	if err != nil {
		log.Fatal(err)
	}

	// Run migration.
	Migration(conf, conn)

	return conn
}

func Migration(conf *config.DefaultConfig, client *ent.Client) {
	log.Println("Migrating...")
	// Run migration.
	if err := client.Schema.Create(
		conf.Ctx,
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("Migrated")
}

func Close() error {
	return client.Close()
}
