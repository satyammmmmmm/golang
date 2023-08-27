package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	app.InfoLog.Println("starting web server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	//connect to db
	db := initDB()
	//db.Ping()

	session := initSession()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create channels

	//create waitgroups
	wg := sync.WaitGroup()

	//setup app config
	app := Config{
		Session:  session,
		DB:       db,
		infoLog:  infoLog,
		ErrorLog: errorLog,
		Wait:     &wg,
	}
	app.Mailer=app.createMail()
go app.listenForMail()
	go app.ListenForShutdown()

	//listen for web connection
	app.serve()
}

func (app *Config) createMail() {
	errorChan :=make(chan error)
	mailerChan:=make(chan Message,100) ,

	mailerDoneChan :=make(Chan bool)
	m:=Mail{
		Domain:"localhost",
		Host:"localhost".
		Port:1025,
		Encryption:"none",
		FromName:"info",
		FromAddress : "info@my.com"
		Wait:app.Wait ,
		ErrorChan:errorChan,
		Wait:app.Wait,
		MailerChan:mailerChan,
		DoneChan:mailerDoneChan, 

	}

}
func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("cannt connect to db")
	}
	return conn
}
func connectToDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("posttgres not ready yet")
		} else {
			log.Println("connected to db")
			return connection
		}
		if counts > 10 {
			return nil
		}
		log.Println("backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func initSession() *scs.SessionManager {
	gob.Register(data.User{})
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true
	return

}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	return redisPool
}

func (App *Config) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)

}
func (app *Config) shutdown() {
	//perform cleeanuo
	app.InfoLog.Println("would run cleanup")
	//block until waitgroup is empty
	app.Wait.Wait()
	app.InfoLog.Println("closing channel and shutting doen application")
}
